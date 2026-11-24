// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// The notification destinations that receive access requests for a Unity
// Catalog securable. A securable without destinations of its own uses its
// owner's email when default owner email destinations are on; otherwise it
// inherits the destinations of its nearest parent that has any.
type AccessRequestDestinations struct {
	// Indicates whether any destinations are hidden from the caller due to a
	// lack of permissions. This value is true if the caller does not have
	// permission to see all destinations.
	AreAnyDestinationsHidden bool `json:"are_any_destinations_hidden,omitempty"`
	// The source securable from which the destinations are inherited. Either
	// the same value as securable (if destination is set directly on the
	// securable) or the nearest parent securable with destinations set.
	DestinationSourceSecurable *Securable `json:"destination_source_securable,omitempty"`
	// The access request destinations for the securable.
	Destinations []NotificationDestination `json:"destinations,omitempty"`
	// The full name of the securable, repeated from **securable**.
	FullName string `json:"full_name,omitempty"`
	// The securable for which the access request destinations are being
	// modified or read.
	Securable Securable `json:"securable"`
	// The type of the securable, repeated from **securable** as a plain string.
	SecurableType string `json:"securable_type,omitempty"`
	// The access request URLs shared by every workspace in the metastore. Only
	// set when the securable is a metastore. These URLs are used when neither
	// the requested securable nor its parents have an explicit destination or a
	// default owner email destination, and they take precedence over metastore
	// destinations.
	//
	// A **dynamic_url** and default owner email destinations cannot both be on:
	// saving one turns the other off. Saving a **dynamic_url** while they are
	// on requires a metastore admin.
	//
	// Setting this field removes the metastore destinations in this workspace.
	// The first save also removes the metastore URL destinations of other
	// workspaces. Once set, the configuration can be cleared but not removed.
	UrlConfiguration *UrlConfiguration `json:"url_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AccessRequestDestinations) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AccessRequestDestinations) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Properties of the new metastore.
type AccountsCreateMetastore struct {
	MetastoreInfo *CreateAccountsMetastore `json:"metastore_info,omitempty"`
}

func (s *AccountsCreateMetastore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The mapping from workspace to metastore.
type AccountsCreateMetastoreAssignment struct {
	MetastoreAssignment *CreateMetastoreAssignment `json:"metastore_assignment,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *AccountsCreateMetastoreAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The metastore assignment was successfully created.
type AccountsCreateMetastoreAssignmentResponse struct {
}

func (s *AccountsCreateMetastoreAssignmentResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AccountsCreateMetastoreResponse struct {
	MetastoreInfo *MetastoreInfo `json:"metastore_info,omitempty"`
}

func (s *AccountsCreateMetastoreResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AccountsCreateStorageCredential struct {
	CredentialInfo *CreateAccountsStorageCredential `json:"credential_info,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Optional, default false. Supplying true to this argument skips validation
	// of the created set of credentials.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AccountsCreateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AccountsCreateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AccountsCreateStorageCredentialInfo struct {
	CredentialInfo *StorageCredentialInfo `json:"credential_info,omitempty"`
}

func (s *AccountsCreateStorageCredentialInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The metastore assignment was successfully deleted.
type AccountsDeleteMetastoreAssignmentResponse struct {
}

func (s *AccountsDeleteMetastoreAssignmentResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The metastore was successfully deleted.
type AccountsDeleteMetastoreResponse struct {
}

func (s *AccountsDeleteMetastoreResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The storage credential was successfully deleted.
type AccountsDeleteStorageCredentialResponse struct {
}

func (s *AccountsDeleteStorageCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The metastore was successfully returned.
type AccountsGetMetastoreResponse struct {
	MetastoreInfo *MetastoreInfo `json:"metastore_info,omitempty"`
}

func (s *AccountsGetMetastoreResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Metastores were returned successfully.
type AccountsListMetastoresResponse struct {
	// An array of metastore information objects.
	Metastores []MetastoreInfo `json:"metastores,omitempty"`
}

func (s *AccountsListMetastoresResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The workspace metastore assignment was successfully returned.
type AccountsMetastoreAssignment struct {
	MetastoreAssignment *MetastoreAssignment `json:"metastore_assignment,omitempty"`
}

func (s *AccountsMetastoreAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The storage credential was successfully retrieved.
type AccountsStorageCredentialInfo struct {
	CredentialInfo *StorageCredentialInfo `json:"credential_info,omitempty"`
}

func (s *AccountsStorageCredentialInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Properties of the metastore to change.
type AccountsUpdateMetastore struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Properties of the metastore to change.
	MetastoreInfo *UpdateAccountsMetastore `json:"metastore_info,omitempty"`
}

func (s *AccountsUpdateMetastore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The metastore assignment to update.
type AccountsUpdateMetastoreAssignment struct {
	MetastoreAssignment *UpdateMetastoreAssignment `json:"metastore_assignment,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *AccountsUpdateMetastoreAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The metastore assignment was successfully updated.
type AccountsUpdateMetastoreAssignmentResponse struct {
}

func (s *AccountsUpdateMetastoreAssignmentResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The metastore update request succeeded.
type AccountsUpdateMetastoreResponse struct {
	MetastoreInfo *MetastoreInfo `json:"metastore_info,omitempty"`
}

func (s *AccountsUpdateMetastoreResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The storage credential to update.
type AccountsUpdateStorageCredential struct {
	CredentialInfo *UpdateAccountsStorageCredential `json:"credential_info,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Optional. Supplying true to this argument skips validation of the updated
	// set of credentials.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// Name of the storage credential.
	StorageCredentialName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AccountsUpdateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AccountsUpdateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The storage credential was successfully updated.
type AccountsUpdateStorageCredentialResponse struct {
	CredentialInfo *StorageCredentialInfo `json:"credential_info,omitempty"`
}

func (s *AccountsUpdateStorageCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ArtifactAllowlistInfo struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers []ArtifactMatcher `json:"artifact_matchers,omitempty"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of the user who set the artifact allowlist.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ArtifactAllowlistInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ArtifactAllowlistInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ArtifactMatcher struct {
	// The artifact path or maven coordinate
	Artifact string `json:"artifact"`
	// The pattern matching type of the artifact
	MatchType MatchType `json:"match_type"`
}

func (s *ArtifactMatcher) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
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

// Values returns all possible values for ArtifactType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ArtifactType) Values() []ArtifactType {
	return []ArtifactType{
		ArtifactTypeInitScript,
		ArtifactTypeLibraryJar,
		ArtifactTypeLibraryMaven,
	}
}

// Type always returns ArtifactType to satisfy [pflag.Value] interface
func (f *ArtifactType) Type() string {
	return "ArtifactType"
}

// AWS temporary credentials for API authentication. Read more at
// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
type AwsCredentials struct {
	// The access key ID that identifies the temporary credentials.
	AccessKeyId string `json:"access_key_id,omitempty"`
	// The Amazon Resource Name (ARN) of the S3 access point for temporary
	// credentials related the external location.
	AccessPoint string `json:"access_point,omitempty"`
	// The secret access key that can be used to sign AWS API requests.
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	// The token that users must pass to AWS API to use the temporary
	// credentials.
	SessionToken string `json:"session_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsCredentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsCredentials) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The AWS IAM role configuration
type AwsIamRole struct {
	// The external ID used in role assumption to prevent the confused deputy
	// problem.
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	RoleArn string `json:"role_arn,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsIamRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsIamRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The AWS IAM role configuration
type AwsIamRoleRequest struct {
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	RoleArn string `json:"role_arn"`
}

func (s *AwsIamRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The AWS IAM role configuration
type AwsIamRoleResponse struct {
	// The external ID used in role assumption to prevent the confused deputy
	// problem.
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	RoleArn string `json:"role_arn"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsIamRoleResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsIamRoleResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AwsSqsQueue struct {
	// Unique identifier included in the name of file events managed cloud
	// resources.
	ManagedResourceId string `json:"managed_resource_id,omitempty"`
	// The SQS queue URL in the format
	// `https://sqs.{region}.amazonaws.com/{account id}/{queue name}`. Only
	// required for provided_sqs.
	QueueUrl string `json:"queue_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsSqsQueue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsSqsQueue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Azure Active Directory token, essentially the Oauth token for Azure Service
// Principal or Managed Identity. Read more at
// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
type AzureActiveDirectoryToken struct {
	// Opaque token that contains claims that you can use in Azure Active
	// Directory to access cloud services.
	AadToken string `json:"aad_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureActiveDirectoryToken) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureActiveDirectoryToken) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AzureEncryptionSettings struct {
	AzureCmkAccessConnectorId string `json:"azure_cmk_access_connector_id,omitempty"`

	AzureCmkManagedIdentityId string `json:"azure_cmk_managed_identity_id,omitempty"`

	AzureTenantId string `json:"azure_tenant_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureEncryptionSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureEncryptionSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The Azure managed identity configuration.
type AzureManagedIdentity struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	AccessConnectorId string `json:"access_connector_id"`
	// The Databricks internal ID that represents this managed identity.
	CredentialId string `json:"credential_id,omitempty"`
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureManagedIdentity) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureManagedIdentity) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The Azure managed identity configuration.
type AzureManagedIdentityRequest struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	AccessConnectorId string `json:"access_connector_id"`
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureManagedIdentityRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureManagedIdentityRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The Azure managed identity configuration.
type AzureManagedIdentityResponse struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	AccessConnectorId string `json:"access_connector_id"`
	// The Databricks internal ID that represents this managed identity.
	CredentialId string `json:"credential_id,omitempty"`
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureManagedIdentityResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureManagedIdentityResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AzureQueueStorage struct {
	// Unique identifier included in the name of file events managed cloud
	// resources.
	ManagedResourceId string `json:"managed_resource_id,omitempty"`
	// The AQS queue URL in the format `https://{storage
	// account}.queue.core.windows.net/{queue name}`. Only required for
	// provided_aqs.
	QueueUrl string `json:"queue_url,omitempty"`
	// Optional resource group for the queue, event grid subscription, and
	// external location storage account. Only required for locations with a
	// service principal storage credential
	ResourceGroup string `json:"resource_group,omitempty"`
	// Optional subscription id for the queue, event grid subscription, and
	// external location storage account. Required for locations with a service
	// principal storage credential
	SubscriptionId string `json:"subscription_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureQueueStorage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureQueueStorage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The Azure service principal configuration. Only applicable when purpose is
// **STORAGE**.
type AzureServicePrincipal struct {
	// The application ID of the application registration within the referenced
	// AAD tenant.
	ApplicationId string `json:"application_id"`
	// The client secret generated for the above app ID in AAD.
	ClientSecret string `json:"client_secret"`
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	DirectoryId string `json:"directory_id"`
}

func (s *AzureServicePrincipal) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Azure temporary credentials for API authentication. Read more at
// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
type AzureUserDelegationSas struct {
	// The signed URI (SAS Token) used to access blob services for a given path
	SasToken string `json:"sas_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureUserDelegationSas) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureUserDelegationSas) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BatchCreateAccessRequestsRequest struct {
	// A list of individual access requests, where each request corresponds to a
	// set of permissions being requested on a list of securables for a
	// specified principal.
	//
	// At most 30 requests per API call.
	Requests []CreateAccessRequest `json:"requests,omitempty"`
}

func (s *BatchCreateAccessRequestsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The access request destinations resolved for each principal and securable in
// the batch.
type BatchCreateAccessRequestsResponse struct {
	// The access request destinations for each securable object the principal
	// requested.
	Responses []CreateAccessRequestResponse `json:"responses,omitempty"`
}

func (s *BatchCreateAccessRequestsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CancelRefreshRequest struct {
	RefreshId int64 `json:"-" url:"-"`
	// UC table name in format `catalog.schema.table_name`. table_name is case
	// insensitive and spaces are disallowed.
	TableName string `json:"-" url:"-"`
}

func (s *CancelRefreshRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A catalog is the first layer of Unity Catalog's three-level namespace and
// groups the schemas, tables, volumes, and other data assets under it. Catalogs
// created through this API are standard catalogs, Delta Sharing catalogs
// created from a share, or foreign catalogs that mirror a database behind a
// connection to an external data source.
type CatalogInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The server sets the catalog type from the create request. Catalogs
	// created through this API are MANAGED_CATALOG (standard catalogs),
	// DELTASHARING_CATALOG when created from a share with provider_name and
	// share_name, or FOREIGN_CATALOG when created from a connection with
	// connection_name. The remaining types are provisioned by Databricks.
	CatalogType CatalogType `json:"catalog_type,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Naming a connection to an external data source creates a foreign catalog
	// that mirrors a database in that source, which requires ownership of or
	// the CREATE FOREIGN CATALOG privilege on the connection.
	// Connection-specific settings, such as the database to mirror, go in
	// options. The connection cannot be changed after creation.
	ConnectionName string `json:"connection_name,omitempty"`
	// Time at which this catalog was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of catalog creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Dropped objects in the catalog remain recoverable for this number of
	// hours before they are permanently purged. The value applies to standard
	// catalogs only, must be 0 or within the range allowed for the metastore,
	// and is inherited by schemas and tables that do not set their own value.
	CustomMaxRetentionHours int64 `json:"custom_max_retention_hours,omitempty"`
	// Indicates the predictive optimization setting that is in effect for the
	// catalog once inherited values are resolved.
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `json:"effective_predictive_optimization_flag,omitempty"`
	// Whether predictive optimization is enabled for this catalog and the
	// objects under it. Valid values are ENABLE, DISABLE, and INHERIT, where
	// INHERIT applies the account-level setting. Standard catalogs start as
	// INHERIT and other catalog types as DISABLE. It cannot be set when
	// creating a catalog; use the update method instead.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// The full name of the catalog. Corresponds with the name field.
	FullName string `json:"full_name,omitempty"`
	// Whether the catalog is accessible from all workspaces attached to the
	// metastore (OPEN) or only from the workspaces bound to it (ISOLATED). It
	// cannot be set when creating a catalog; change it with the update method
	// and manage the bound workspaces with the workspace bindings API. See
	// [Workspace-catalog binding].
	//
	// [Workspace-catalog binding]: https://docs.databricks.com/aws/en/catalogs/binding
	IsolationMode CatalogIsolationMode `json:"isolation_mode,omitempty"`
	// Customer-managed key (CMK) settings control the encryption at rest of the
	// catalog's managed data.
	ManagedEncryptionSettings *EncryptionSettings `json:"managed_encryption_settings,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the catalog is stored in lowercase, can be at most 255
	// characters long, and cannot contain periods, spaces, forward slashes, or
	// control characters.
	Name string `json:"name,omitempty"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options,omitempty"`
	// Username of current owner of catalog.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// The Delta Sharing provider that owns the share must be named here,
	// together with share_name, to create a Delta Sharing catalog that exposes
	// the data of the share on a remote sharing server. The provider cannot be
	// changed after creation.
	ProviderName string `json:"provider_name,omitempty"`
	// Indicates the provisioning status of the catalog. It is present only for
	// catalogs that are provisioned asynchronously.
	ProvisioningInfo *ProvisioningInfo `json:"provisioning_info,omitempty"`
	// The securable type is always CATALOG for a catalog.
	SecurableType SecurableType `json:"securable_type,omitempty"`
	// The share under the Delta Sharing provider must be named here, together
	// with provider_name, to create a Delta Sharing catalog. The share cannot
	// be changed after creation.
	ShareName string `json:"share_name,omitempty"`
	// The server derives this full storage location URL for managed tables and
	// volumes within the catalog from storage_root and the catalog ID.
	StorageLocation string `json:"storage_location,omitempty"`
	// Managed tables and volumes in the catalog are stored under this storage
	// root URL. The path must be inside an external location on which the
	// caller has the CREATE MANAGED STORAGE privilege. If omitted, managed data
	// is stored under the metastore's root storage location, which must then be
	// configured. Delta Sharing catalogs cannot have a storage root.
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this catalog was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified catalog.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CatalogInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CatalogInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CatalogIsolationMode string

const CatalogIsolationModeIsolated CatalogIsolationMode = `ISOLATED`

const CatalogIsolationModeOpen CatalogIsolationMode = `OPEN`

// String representation for [fmt.Print]
func (f *CatalogIsolationMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CatalogIsolationMode) Set(v string) error {
	switch v {
	case `ISOLATED`, `OPEN`:
		*f = CatalogIsolationMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ISOLATED", "OPEN"`, v)
	}
}

// Values returns all possible values for CatalogIsolationMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *CatalogIsolationMode) Values() []CatalogIsolationMode {
	return []CatalogIsolationMode{
		CatalogIsolationModeIsolated,
		CatalogIsolationModeOpen,
	}
}

// Type always returns CatalogIsolationMode to satisfy [pflag.Value] interface
func (f *CatalogIsolationMode) Type() string {
	return "CatalogIsolationMode"
}

// The type of the catalog.
type CatalogType string

const CatalogTypeDeltasharingCatalog CatalogType = `DELTASHARING_CATALOG`

const CatalogTypeForeignCatalog CatalogType = `FOREIGN_CATALOG`

const CatalogTypeInternalCatalog CatalogType = `INTERNAL_CATALOG`

const CatalogTypeManagedCatalog CatalogType = `MANAGED_CATALOG`

const CatalogTypeManagedOnlineCatalog CatalogType = `MANAGED_ONLINE_CATALOG`

const CatalogTypeSystemCatalog CatalogType = `SYSTEM_CATALOG`

// String representation for [fmt.Print]
func (f *CatalogType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CatalogType) Set(v string) error {
	switch v {
	case `DELTASHARING_CATALOG`, `FOREIGN_CATALOG`, `INTERNAL_CATALOG`, `MANAGED_CATALOG`, `MANAGED_ONLINE_CATALOG`, `SYSTEM_CATALOG`:
		*f = CatalogType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTASHARING_CATALOG", "FOREIGN_CATALOG", "INTERNAL_CATALOG", "MANAGED_CATALOG", "MANAGED_ONLINE_CATALOG", "SYSTEM_CATALOG"`, v)
	}
}

// Values returns all possible values for CatalogType.
//
// There is no guarantee on the order of the values in the slice.
func (f *CatalogType) Values() []CatalogType {
	return []CatalogType{
		CatalogTypeDeltasharingCatalog,
		CatalogTypeForeignCatalog,
		CatalogTypeInternalCatalog,
		CatalogTypeManagedCatalog,
		CatalogTypeManagedOnlineCatalog,
		CatalogTypeSystemCatalog,
	}
}

// Type always returns CatalogType to satisfy [pflag.Value] interface
func (f *CatalogType) Type() string {
	return "CatalogType"
}

// The Cloudflare API token configuration. Read more at
// https://developers.cloudflare.com/r2/api/s3/tokens/
type CloudflareApiToken struct {
	// The access key ID associated with the API token.
	AccessKeyId string `json:"access_key_id"`
	// The ID of the account associated with the API token.
	AccountId string `json:"account_id"`
	// The secret access token generated for the above access key ID.
	SecretAccessKey string `json:"secret_access_key"`
}

func (s *CloudflareApiToken) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ColumnInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// The column mask function applied to this column, if any. A column mask is
	// a SQL user-defined function that transforms the values a query returns
	// for the column. See
	// https://docs.databricks.com/aws/en/tables/row-and-column-filters for
	// details.
	Mask *ColumnMask `json:"mask,omitempty"`
	// Name of Column.
	Name string `json:"name,omitempty"`
	// Whether field may be Null (default: true).
	Nullable bool `json:"nullable,omitempty"`
	// Partition index for column.
	PartitionIndex int `json:"partition_index,omitempty"`
	// Ordinal position of column (starting at position 0).
	Position int `json:"position,omitempty"`
	// Format of IntervalType.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// Full data type specification, JSON-serialized.
	TypeJson string `json:"type_json,omitempty"`
	// The name of the column's data type, for example STRING, DECIMAL or
	// STRUCT. The complete type, including parameters such as the precision of
	// a decimal or the element type of an array, is in type_text.
	TypeName ColumnTypeName `json:"type_name,omitempty"`
	// Digits of precision; required for DecimalTypes.
	TypePrecision int `json:"type_precision,omitempty"`
	// Digits to right of decimal; Required for DecimalTypes.
	TypeScale int `json:"type_scale,omitempty"`
	// Full data type specification as SQL/catalogString text.
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ColumnMask struct {
	// The full name of the column mask SQL UDF.
	FunctionName string `json:"function_name,omitempty"`
	// The list of table columns or literals to be passed as additional
	// arguments to a column mask function, carrying the type (column reference
	// vs constant literal) of each argument. Deprecated: use using_column_names
	// instead.
	UsingArguments []PolicyFunctionArgument `json:"using_arguments,omitempty"`
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	UsingColumnNames []string `json:"using_column_names,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ColumnMask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnMask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ColumnMaskOptions struct {
	// The fully qualified name of the column mask function. The function is
	// called on each row of the target table. The function's first argument and
	// its return type should match the type of the masked column. Required on
	// create and update.
	FunctionName string `json:"function_name"`
	// The alias of the column to be masked. The alias must refer to one of
	// matched columns. The values of the column is passed to the column mask
	// function as the first argument. Required on create and update.
	OnColumn string `json:"on_column"`
	// Optional list of column aliases or constant literals to be passed as
	// additional arguments to the column mask function. The type of each column
	// should match the positional argument of the column mask function.
	Using []FunctionArgument `json:"using,omitempty"`
}

func (s *ColumnMaskOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Mapping from a column of the source object to a column of the target object.
type ColumnRelationship struct {
	// Name of the column in the source object.
	Source string `json:"source,omitempty"`
	// Name of the column in the target object.
	Target string `json:"target,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ColumnRelationship) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnRelationship) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Extracts the value of a column-level tag: get_column_tag_value(col,
// "tagKey").
type ColumnTagValueExtraction struct {
	// The alias from MATCH COLUMNS that identifies the column.
	ColumnAlias string `json:"column_alias"`
	// 1024 matches the max_length on FunctionArgument.constant above.
	TagKey string `json:"tag_key"`
}

func (s *ColumnTagValueExtraction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

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

const ColumnTypeNameGeography ColumnTypeName = `GEOGRAPHY`

const ColumnTypeNameGeometry ColumnTypeName = `GEOMETRY`

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

const ColumnTypeNameVariant ColumnTypeName = `VARIANT`

// String representation for [fmt.Print]
func (f *ColumnTypeName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ColumnTypeName) Set(v string) error {
	switch v {
	case `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `GEOGRAPHY`, `GEOMETRY`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TABLE_TYPE`, `TIMESTAMP`, `TIMESTAMP_NTZ`, `USER_DEFINED_TYPE`, `VARIANT`:
		*f = ColumnTypeName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARRAY", "BINARY", "BOOLEAN", "BYTE", "CHAR", "DATE", "DECIMAL", "DOUBLE", "FLOAT", "GEOGRAPHY", "GEOMETRY", "INT", "INTERVAL", "LONG", "MAP", "NULL", "SHORT", "STRING", "STRUCT", "TABLE_TYPE", "TIMESTAMP", "TIMESTAMP_NTZ", "USER_DEFINED_TYPE", "VARIANT"`, v)
	}
}

// Values returns all possible values for ColumnTypeName.
//
// There is no guarantee on the order of the values in the slice.
func (f *ColumnTypeName) Values() []ColumnTypeName {
	return []ColumnTypeName{
		ColumnTypeNameArray,
		ColumnTypeNameBinary,
		ColumnTypeNameBoolean,
		ColumnTypeNameByte,
		ColumnTypeNameChar,
		ColumnTypeNameDate,
		ColumnTypeNameDecimal,
		ColumnTypeNameDouble,
		ColumnTypeNameFloat,
		ColumnTypeNameGeography,
		ColumnTypeNameGeometry,
		ColumnTypeNameInt,
		ColumnTypeNameInterval,
		ColumnTypeNameLong,
		ColumnTypeNameMap,
		ColumnTypeNameNull,
		ColumnTypeNameShort,
		ColumnTypeNameString,
		ColumnTypeNameStruct,
		ColumnTypeNameTableType,
		ColumnTypeNameTimestamp,
		ColumnTypeNameTimestampNtz,
		ColumnTypeNameUserDefinedType,
		ColumnTypeNameVariant,
	}
}

// Type always returns ColumnTypeName to satisfy [pflag.Value] interface
func (f *ColumnTypeName) Type() string {
	return "ColumnTypeName"
}

// A connection that is dependent on a SQL object.
type ConnectionDependency struct {
	// Full name of the dependent connection, in the form of
	// __connection_name__.
	ConnectionName string `json:"connection_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ConnectionDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ConnectionDependency) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ConnectionInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Unique identifier of the Connection.
	ConnectionId string `json:"connection_id,omitempty"`
	// The type of external data source or service that the connection reaches.
	// The connection type determines which keys are required in options and
	// cannot be changed after creation.
	ConnectionType ConnectionType `json:"connection_type,omitempty"`
	// Time at which this connection was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of connection creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The type of credential that is used to authenticate to the external data
	// source. The server derives it from the connection type and the keys
	// present in options.
	CredentialType CredentialType `json:"credential_type,omitempty"`
	// Environment settings for the connection, such as the Java libraries that
	// it needs and the environment version.
	EnvironmentSettings *EnvironmentSettings `json:"environment_settings,omitempty"`
	// Full name of connection.
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of the connection.
	Name string `json:"name,omitempty"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options,omitempty"`
	// Username of the current owner of the connection. The creator becomes the
	// owner at creation, and ownership can be transferred on update.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// Provisioning status of the connection.
	ProvisioningInfo *ProvisioningInfo `json:"provisioning_info,omitempty"`
	// Whether the connection is read only. All connections are currently read
	// only, so a value set on create or update is ignored.
	ReadOnly bool `json:"read_only,omitempty"`
	// A map of option names to UC Secret references. Keys are connection option
	// names (same as in OptionsKVPairs) and values are UC Secret fully
	// qualified names.
	Secrets map[string]string `json:"secrets,omitempty"`
	// The type of Unity Catalog securable. This is always CONNECTION.
	SecurableType SecurableType `json:"securable_type,omitempty"`
	// Time at which this connection was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified connection.
	UpdatedBy string `json:"updated_by,omitempty"`
	// URL of the remote data source. The server derives it from the host, port
	// and path options.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ConnectionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ConnectionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ConnectionType string

const ConnectionTypeAwsSecretsManager ConnectionType = `AWS_SECRETS_MANAGER`

const ConnectionTypeAzureKeyVault ConnectionType = `AZURE_KEY_VAULT`

const ConnectionTypeBigquery ConnectionType = `BIGQUERY`

const ConnectionTypeConfluence ConnectionType = `CONFLUENCE`

const ConnectionTypeDatabricks ConnectionType = `DATABRICKS`

const ConnectionTypeDynamics365 ConnectionType = `DYNAMICS365`

const ConnectionTypeGa4RawData ConnectionType = `GA4_RAW_DATA`

const ConnectionTypeGithub ConnectionType = `GITHUB`

const ConnectionTypeGlue ConnectionType = `GLUE`

const ConnectionTypeGoogleAds ConnectionType = `GOOGLE_ADS`

const ConnectionTypeHiveMetastore ConnectionType = `HIVE_METASTORE`

const ConnectionTypeHttp ConnectionType = `HTTP`

const ConnectionTypeHubspot ConnectionType = `HUBSPOT`

const ConnectionTypeJdbc ConnectionType = `JDBC`

const ConnectionTypeMetaMarketing ConnectionType = `META_MARKETING`

const ConnectionTypeMysql ConnectionType = `MYSQL`

const ConnectionTypeNetsuite ConnectionType = `NETSUITE`

const ConnectionTypeOracle ConnectionType = `ORACLE`

const ConnectionTypeOutlook ConnectionType = `OUTLOOK`

const ConnectionTypePostgresql ConnectionType = `POSTGRESQL`

const ConnectionTypePowerBi ConnectionType = `POWER_BI`

const ConnectionTypeRedshift ConnectionType = `REDSHIFT`

const ConnectionTypeSalesforce ConnectionType = `SALESFORCE`

const ConnectionTypeSalesforceDataCloud ConnectionType = `SALESFORCE_DATA_CLOUD`

const ConnectionTypeServicenow ConnectionType = `SERVICENOW`

const ConnectionTypeSmartsheet ConnectionType = `SMARTSHEET`

const ConnectionTypeSnowflake ConnectionType = `SNOWFLAKE`

const ConnectionTypeSqldw ConnectionType = `SQLDW`

const ConnectionTypeSqlserver ConnectionType = `SQLSERVER`

const ConnectionTypeTeradata ConnectionType = `TERADATA`

const ConnectionTypeTiktokAds ConnectionType = `TIKTOK_ADS`

const ConnectionTypeUnknownConnectionType ConnectionType = `UNKNOWN_CONNECTION_TYPE`

const ConnectionTypeWorkdayRaas ConnectionType = `WORKDAY_RAAS`

const ConnectionTypeZendesk ConnectionType = `ZENDESK`

// String representation for [fmt.Print]
func (f *ConnectionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ConnectionType) Set(v string) error {
	switch v {
	case `AWS_SECRETS_MANAGER`, `AZURE_KEY_VAULT`, `BIGQUERY`, `CONFLUENCE`, `DATABRICKS`, `DYNAMICS365`, `GA4_RAW_DATA`, `GITHUB`, `GLUE`, `GOOGLE_ADS`, `HIVE_METASTORE`, `HTTP`, `HUBSPOT`, `JDBC`, `META_MARKETING`, `MYSQL`, `NETSUITE`, `ORACLE`, `OUTLOOK`, `POSTGRESQL`, `POWER_BI`, `REDSHIFT`, `SALESFORCE`, `SALESFORCE_DATA_CLOUD`, `SERVICENOW`, `SMARTSHEET`, `SNOWFLAKE`, `SQLDW`, `SQLSERVER`, `TERADATA`, `TIKTOK_ADS`, `UNKNOWN_CONNECTION_TYPE`, `WORKDAY_RAAS`, `ZENDESK`:
		*f = ConnectionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AWS_SECRETS_MANAGER", "AZURE_KEY_VAULT", "BIGQUERY", "CONFLUENCE", "DATABRICKS", "DYNAMICS365", "GA4_RAW_DATA", "GITHUB", "GLUE", "GOOGLE_ADS", "HIVE_METASTORE", "HTTP", "HUBSPOT", "JDBC", "META_MARKETING", "MYSQL", "NETSUITE", "ORACLE", "OUTLOOK", "POSTGRESQL", "POWER_BI", "REDSHIFT", "SALESFORCE", "SALESFORCE_DATA_CLOUD", "SERVICENOW", "SMARTSHEET", "SNOWFLAKE", "SQLDW", "SQLSERVER", "TERADATA", "TIKTOK_ADS", "UNKNOWN_CONNECTION_TYPE", "WORKDAY_RAAS", "ZENDESK"`, v)
	}
}

// Values returns all possible values for ConnectionType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ConnectionType) Values() []ConnectionType {
	return []ConnectionType{
		ConnectionTypeAwsSecretsManager,
		ConnectionTypeAzureKeyVault,
		ConnectionTypeBigquery,
		ConnectionTypeConfluence,
		ConnectionTypeDatabricks,
		ConnectionTypeDynamics365,
		ConnectionTypeGa4RawData,
		ConnectionTypeGithub,
		ConnectionTypeGlue,
		ConnectionTypeGoogleAds,
		ConnectionTypeHiveMetastore,
		ConnectionTypeHttp,
		ConnectionTypeHubspot,
		ConnectionTypeJdbc,
		ConnectionTypeMetaMarketing,
		ConnectionTypeMysql,
		ConnectionTypeNetsuite,
		ConnectionTypeOracle,
		ConnectionTypeOutlook,
		ConnectionTypePostgresql,
		ConnectionTypePowerBi,
		ConnectionTypeRedshift,
		ConnectionTypeSalesforce,
		ConnectionTypeSalesforceDataCloud,
		ConnectionTypeServicenow,
		ConnectionTypeSmartsheet,
		ConnectionTypeSnowflake,
		ConnectionTypeSqldw,
		ConnectionTypeSqlserver,
		ConnectionTypeTeradata,
		ConnectionTypeTiktokAds,
		ConnectionTypeUnknownConnectionType,
		ConnectionTypeWorkdayRaas,
		ConnectionTypeZendesk,
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
	InitialPipelineSyncProgress *PipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ContinuousUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ContinuousUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateAccessRequest struct {
	// Optional. The principal this request is for. Empty `behalf_of` defaults
	// to the requester's identity.
	//
	// Principals must be unique across the API call.
	BehalfOf *Principal `json:"behalf_of,omitempty"`
	// Optional. Comment associated with the request.
	//
	// At most 200 characters, can only contain lowercase/uppercase letters
	// (a-z, A-Z), numbers (0-9), punctuation, and spaces.
	Comment string `json:"comment,omitempty"`
	// List of securables and their corresponding requested UC privileges.
	//
	// At most 30 securables can be requested for a principal per batched call.
	// Each securable can only be requested once per principal.
	SecurablePermissions []SecurablePermissions `json:"securable_permissions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateAccessRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAccessRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateAccessRequestResponse struct {
	// The principal the request was made on behalf of.
	BehalfOf *Principal `json:"behalf_of,omitempty"`
	// The access request destinations for all the securables the principal
	// requested.
	RequestDestinations []AccessRequestDestinations `json:"request_destinations,omitempty"`
}

func (s *CreateAccessRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateAccountsMetastore struct {
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled bool `json:"external_access_enabled,omitempty"`
	// The user-specified name of the metastore.
	Name string `json:"name"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateAccountsMetastore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAccountsMetastore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A storage credential represents an authentication and authorization mechanism
// for accessing data stored on your cloud tenant. It holds exactly one cloud
// credential, for example an AWS IAM role, an Azure managed identity or service
// principal, or a GCP service account. Only the field for the selected
// credential type is set in requests and responses.
type CreateAccountsStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleRequest `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityRequest `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name string `json:"name"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateAccountsStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAccountsStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCatalog struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Naming a connection to an external data source creates a foreign catalog
	// that mirrors a database in that source, which requires ownership of or
	// the CREATE FOREIGN CATALOG privilege on the connection.
	// Connection-specific settings, such as the database to mirror, go in
	// options. The connection cannot be changed after creation.
	ConnectionName string `json:"connection_name,omitempty"`
	// Dropped objects in the catalog remain recoverable for this number of
	// hours before they are permanently purged. The value applies to standard
	// catalogs only, must be 0 or within the range allowed for the metastore,
	// and is inherited by schemas and tables that do not set their own value.
	CustomMaxRetentionHours int64 `json:"custom_max_retention_hours,omitempty"`
	// Customer-managed key (CMK) settings control the encryption at rest of the
	// catalog's managed data.
	ManagedEncryptionSettings *EncryptionSettings `json:"managed_encryption_settings,omitempty"`
	// The name of the catalog is stored in lowercase, can be at most 255
	// characters long, and cannot contain periods, spaces, forward slashes, or
	// control characters.
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// The Delta Sharing provider that owns the share must be named here,
	// together with share_name, to create a Delta Sharing catalog that exposes
	// the data of the share on a remote sharing server. The provider cannot be
	// changed after creation.
	ProviderName string `json:"provider_name,omitempty"`
	// The share under the Delta Sharing provider must be named here, together
	// with provider_name, to create a Delta Sharing catalog. The share cannot
	// be changed after creation.
	ShareName string `json:"share_name,omitempty"`
	// Managed tables and volumes in the catalog are stored under this storage
	// root URL. The path must be inside an external location on which the
	// caller has the CREATE MANAGED STORAGE privilege. If omitted, managed data
	// is stored under the metastore's root storage location, which must then be
	// configured. Delta Sharing catalogs cannot have a storage root.
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateConnection struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// The type of external data source or service that the connection reaches.
	// The connection type determines which keys are required in options and
	// cannot be changed after creation.
	ConnectionType ConnectionType `json:"connection_type"`
	// Environment settings for the connection, such as the Java libraries that
	// it needs and the environment version.
	EnvironmentSettings *EnvironmentSettings `json:"environment_settings,omitempty"`
	// Name of the connection.
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options"`
	// Parent schema for schema-level connections, in format
	// "schemas/{catalog}.{schema}". Absent for metastore-level (L1)
	// connections.
	Parent string `json:"parent,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// Whether the connection is read only. All connections are currently read
	// only, so a value set on create or update is ignored.
	ReadOnly bool `json:"read_only,omitempty"`
	// A map of option names to UC Secret references. Keys are connection option
	// names (same as in OptionsKVPairs) and values are UC Secret fully
	// qualified names.
	Secrets map[string]string `json:"secrets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateConnection) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCredentialRequest struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentity `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name string `json:"name"`
	// Indicates the purpose of the credential.
	Purpose CredentialPurpose `json:"purpose,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`
	// Optional. Supplying true to this argument skips validation of the created
	// set of credentials.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateEntityTagAssignmentRequest struct {
	TagAssignment EntityTagAssignment `json:"tag_assignment"`
}

func (s *CreateEntityTagAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateExternalLineageRelationshipRequest struct {
	ExternalLineageRelationship CreateRequestExternalLineage `json:"external_lineage_relationship"`
}

func (s *CreateExternalLineageRelationshipRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateExternalLocation struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the storage credential used with this location.
	CredentialName string `json:"credential_name"`
	// The effective value of `enable_file_events` after applying server-side
	// defaults.
	EffectiveEnableFileEvents bool `json:"effective_enable_file_events,omitempty"`
	// The effective file event queue configuration after applying server-side
	// defaults. Always populated when a queue is provisioned, regardless of
	// whether the user explicitly set `enable_file_events`. Use this field
	// instead of `file_event_queue` for reading the actual queue state.
	EffectiveFileEventQueue *FileEventQueue `json:"effective_file_event_queue,omitempty"`
	// Whether to enable file events on this external location. Defaults to
	// `true`. Set to `false` to disable file events. The actual applied value
	// may differ due to server-side defaults; check
	// `effective_enable_file_events` for the effective state.
	EnableFileEvents bool `json:"enable_file_events,omitempty"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback bool `json:"fallback,omitempty"`
	// File event queue settings. When `enable_file_events` is `true`, exactly
	// one of the queue options must be set. Must be omitted when
	// `enable_file_events` is `false`.
	FileEventQueue *FileEventQueue `json:"file_event_queue,omitempty"`
	// Name of the external location.
	Name string `json:"name"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// URL of the cloud storage path that the external location points to,
	// including the scheme (for example `s3://`, `abfss://`, or `gs://`).
	Url string `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateExternalLocation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExternalLocation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateExternalMetadataRequest struct {
	ExternalMetadata ExternalMetadata `json:"external_metadata"`
}

func (s *CreateExternalMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A user-defined function (UDF) registered in Unity Catalog. A function belongs
// to a schema and is referenced as catalog_name.schema_name.function_name. Its
// body is either a SQL expression or query, or source code in an external
// language such as Python. See
// https://docs.databricks.com/aws/en/udf/unity-catalog for an overview.
type CreateFunction struct {
	// Name of parent Catalog.
	CatalogName string `json:"catalog_name"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Data type of the value the function returns. A table function uses
	// TABLE_TYPE and describes its result columns in return_params, except for
	// a polymorphic Python UDTF, which omits them.
	DataType ColumnTypeName `json:"data_type"`
	// Language of the function body for an EXTERNAL function. It is required
	// when routine_body is EXTERNAL.
	ExternalLanguage string `json:"external_language,omitempty"`
	// Name of the function in its external language. It is reserved for future
	// use.
	ExternalName string `json:"external_name,omitempty"`
	// Return data type of the function in SQL syntax, including any type
	// parameters such as precision and scale.
	FullDataType string `json:"full_data_type"`
	// Input parameters of the function. When it is set, parameters must contain
	// at least one entry. A function that takes no arguments omits
	// input_params, even though the Create request schema marks it as required.
	InputParams FunctionParameterInfos `json:"input_params"`
	// Whether the function returns the same result for the same arguments. It
	// corresponds to the DETERMINISTIC characteristic of CREATE FUNCTION.
	IsDeterministic bool `json:"is_deterministic"`
	// Whether the function is a null-call function in the sense of the SQL
	// standard, which returns NULL whenever any argument is NULL.
	IsNullCall bool `json:"is_null_call"`
	// Name of function, relative to parent schema.
	Name string `json:"name"`
	// How arguments are passed to the function body. S means that each argument
	// is passed as a single scalar value.
	ParameterStyle CreateFunctionParameterStyle `json:"parameter_style"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties string `json:"properties,omitempty"`
	// Result columns of a table function. Set it when data_type is TABLE_TYPE,
	// except for a polymorphic Python UDTF (an EXTERNAL Python function with
	// full_data_type "()"), which must not set it.
	ReturnParams *FunctionParameterInfos `json:"return_params,omitempty"`
	// Language of the function body. SQL means that routine_definition holds a
	// SQL expression or query. EXTERNAL means that it holds source code in the
	// language named by external_language, and sql_data_access must then be
	// NO_SQL.
	RoutineBody CreateFunctionRoutineBody `json:"routine_body"`
	// Body of the function. For a SQL function it is the expression or query
	// that computes the result; for an EXTERNAL function it is the source code
	// of the function. It is omitted when the caller can only browse the
	// function.
	RoutineDefinition string `json:"routine_definition"`
	// Function dependencies. For external UDFs, dependencies may contain only
	// credential, secret, or volume objects.
	RoutineDependencies *DependencyList `json:"routine_dependencies,omitempty"`
	// Name of parent Schema relative to its parent Catalog.
	SchemaName string `json:"schema_name"`
	// Whose privileges the function body runs with. DEFINER runs it with the
	// privileges of the function owner.
	SecurityType CreateFunctionSecurityType `json:"security_type"`
	// Specific name of the function. Set it to the same value as name.
	SpecificName string `json:"specific_name"`
	// How the function body interacts with SQL data. CONTAINS_SQL means that
	// the body contains SQL but reads no table or view, READS_SQL_DATA means
	// that it reads a table or view, and NO_SQL is required for EXTERNAL
	// functions.
	SqlDataAccess CreateFunctionSqlDataAccess `json:"sql_data_access"`
	// List of schemas whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Values returns all possible values for CreateFunctionParameterStyle.
//
// There is no guarantee on the order of the values in the slice.
func (f *CreateFunctionParameterStyle) Values() []CreateFunctionParameterStyle {
	return []CreateFunctionParameterStyle{
		CreateFunctionParameterStyleS,
	}
}

// Type always returns CreateFunctionParameterStyle to satisfy [pflag.Value] interface
func (f *CreateFunctionParameterStyle) Type() string {
	return "CreateFunctionParameterStyle"
}

type CreateFunctionRequest struct {
	// Partial __FunctionInfo__ specifying the function to be created.
	FunctionInfo CreateFunction `json:"function_info"`
}

func (s *CreateFunctionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

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

// Values returns all possible values for CreateFunctionRoutineBody.
//
// There is no guarantee on the order of the values in the slice.
func (f *CreateFunctionRoutineBody) Values() []CreateFunctionRoutineBody {
	return []CreateFunctionRoutineBody{
		CreateFunctionRoutineBodyExternal,
		CreateFunctionRoutineBodySql,
	}
}

// Type always returns CreateFunctionRoutineBody to satisfy [pflag.Value] interface
func (f *CreateFunctionRoutineBody) Type() string {
	return "CreateFunctionRoutineBody"
}

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

// Values returns all possible values for CreateFunctionSecurityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *CreateFunctionSecurityType) Values() []CreateFunctionSecurityType {
	return []CreateFunctionSecurityType{
		CreateFunctionSecurityTypeDefiner,
	}
}

// Type always returns CreateFunctionSecurityType to satisfy [pflag.Value] interface
func (f *CreateFunctionSecurityType) Type() string {
	return "CreateFunctionSecurityType"
}

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

// Values returns all possible values for CreateFunctionSqlDataAccess.
//
// There is no guarantee on the order of the values in the slice.
func (f *CreateFunctionSqlDataAccess) Values() []CreateFunctionSqlDataAccess {
	return []CreateFunctionSqlDataAccess{
		CreateFunctionSqlDataAccessContainsSql,
		CreateFunctionSqlDataAccessNoSql,
		CreateFunctionSqlDataAccessReadsSqlData,
	}
}

// Type always returns CreateFunctionSqlDataAccess to satisfy [pflag.Value] interface
func (f *CreateFunctionSqlDataAccess) Type() string {
	return "CreateFunctionSqlDataAccess"
}

type CreateMcpServiceRequest struct {
	// The MCP service to create. Do not set `name`; the server derives it from
	// `parent` and `mcp_service_id`. `source_connection` is required.
	McpService McpService `json:"mcp_service"`
	// Name of the MCP service within the parent schema. The name must not
	// contain periods, spaces, forward slashes, percent signs, or control
	// characters, and it is stored in lowercase.
	McpServiceId string `json:"-" url:"mcp_service_id"`
	// Name of the parent schema. Format: `schemas/{catalog}.{schema}`. Each
	// `{...}` component is capped at 255 characters individually.
	Parent string `json:"-" url:"parent"`
}

func (s *CreateMcpServiceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateMcpServiceUserMappedCredentialRequest struct {
	Login McpServiceUserMappedCredentialLogin `json:"login"`
	// Resource name of the MCP service. Format:
	// `mcp-services/{catalog}.{schema}.{mcp_service}`.
	Name string `json:"-" url:"-"`
}

func (s *CreateMcpServiceUserMappedCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateMetastore struct {
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled bool `json:"external_access_enabled,omitempty"`
	// The user-specified name of the metastore.
	Name string `json:"name"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateMetastore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateMetastore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// deprecated. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName string `json:"default_catalog_name"`
	// The unique ID of the metastore.
	MetastoreId string `json:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *CreateMetastoreAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateModelProviderServiceRequest struct {
	// The model provider service to create. Do not set `name`; the server
	// derives it from `parent` and `model_provider_service_id`.
	ModelProviderService ModelProviderService `json:"model_provider_service"`
	// Name of the model provider service within the parent schema. The name
	// must not contain periods, spaces, forward slashes, percent signs, or
	// control characters, must not end in `_DO_NOT_MODIFY`, and must not exceed
	// 241 characters. It is stored in lowercase.
	ModelProviderServiceId string `json:"-" url:"model_provider_service_id"`
	// Name of the parent schema. Format: `schemas/{catalog}.{schema}`. Each
	// `{...}` component is capped at 255 characters individually.
	Parent string `json:"-" url:"parent"`
}

func (s *CreateModelProviderServiceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateModelServiceRequest struct {
	// The model service to create. Do not set `name`; the server derives it
	// from `parent` and `model_service_id`.
	ModelService ModelService `json:"model_service"`
	// Name of the model service within the parent schema. The name must not
	// contain periods, spaces, forward slashes, percent signs, or control
	// characters, and it is stored in lowercase.
	ModelServiceId string `json:"-" url:"model_service_id"`
	// Name of the parent schema. Format: `schemas/{catalog}.{schema}`. Each
	// `{...}` component is capped at 255 characters individually.
	Parent string `json:"-" url:"parent"`
}

func (s *CreateModelServiceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateMonitor struct {
	// [Create:REQ Update:IGN] Field for specifying the absolute path to a
	// custom directory to store data-monitoring assets. Normally prepopulated
	// to a default user location via UI and Python APIs.
	AssetsDir string `json:"assets_dir"`
	// [Create:OPT Update:OPT] Baseline table name. Baseline data is used to
	// compute drift from the data in the monitored `table_name`. The baseline
	// table and the monitored table shall have the same schema.
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// [Create:OPT Update:OPT] Custom metrics.
	CustomMetrics []MonitorMetric `json:"custom_metrics,omitempty"`
	// [Create:OPT Update:OPT] Data classification related config.
	DataClassificationConfig *MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`

	InferenceLog *MonitorInferenceLog `json:"inference_log,omitempty"`
	// [Create:ERR Update:IGN] The latest error message for a monitor failure.
	LatestMonitorFailureMsg string `json:"latest_monitor_failure_msg,omitempty"`
	// [Create:OPT Update:OPT] Field for specifying notification settings.
	Notifications *MonitorNotifications `json:"notifications,omitempty"`
	// [Create:REQ Update:REQ] Schema where output tables are created. Needs to
	// be in 2-level format {catalog}.{schema}
	OutputSchemaName string `json:"output_schema_name"`
	// [Create:OPT Update:OPT] The monitor schedule.
	Schedule *MonitorCronSchedule `json:"schedule,omitempty"`
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	SkipBuiltinDashboard bool `json:"skip_builtin_dashboard,omitempty"`
	// [Create:OPT Update:OPT] List of column expressions to slice data with for
	// targeted analysis. The data is grouped by each expression independently,
	// resulting in a separate slice for each predicate and its complements. For
	// example `slicing_exprs=[“col_1”, “col_2 > 10”]` will generate the
	// following slices: two slices for `col_2 > 10` (True and False), and one
	// slice per unique value in `col1`. For high-cardinality columns, only the
	// top 100 unique values by frequency will generate slices.
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	Snapshot *MonitorSnapshot `json:"snapshot,omitempty"`
	// UC table name in format `catalog.schema.table_name`. This field
	// corresponds to the {full_table_name_arg} arg in the endpoint path.
	TableName string `json:"-" url:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries *MonitorTimeSeries `json:"time_series,omitempty"`
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateMonitor) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateMonitor) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateOnlineTableRequest struct {
	// Specification of the online table to be created.
	Table OnlineTable `json:"table"`
}

func (s *CreateOnlineTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreatePolicyRequest struct {
	// Required. The policy to create.
	PolicyInfo PolicyInfo `json:"policy_info"`
}

func (s *CreatePolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateRegisteredModelRequest struct {
	// List of aliases associated with the registered model
	Aliases []RegisteredModelAlias `json:"aliases,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog where the schema and the registered model reside
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the registered model
	Comment string `json:"comment,omitempty"`
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the registered model
	CreatedBy string `json:"created_by,omitempty"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the metastore
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the registered model
	Name string `json:"name,omitempty"`
	// The identifier of the user who owns the registered model
	Owner string `json:"owner,omitempty"`
	// The name of the schema where the registered model resides
	SchemaName string `json:"schema_name,omitempty"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `json:"storage_location,omitempty"`
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the registered model last time
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A lineage relationship between a source object and a target object, where at
// least one of the two is an external metadata object. The other object can be
// a Unity Catalog table, a cloud storage path, a model version, or another
// external metadata object. A relationship is identified by its source and
// target; optional column mappings and key-value properties describe the data
// flow between them.
type CreateRequestExternalLineage struct {
	// List of column relationships between source and target objects. The
	// combined size of all source and target column names is limited.
	Columns []ColumnRelationship `json:"columns,omitempty"`
	// Unique identifier of the external lineage relationship.
	Id string `json:"id,omitempty"`
	// Key-value properties associated with the external lineage relationship.
	// The combined length of all keys and values must not exceed 5000
	// characters.
	Properties map[string]string `json:"properties,omitempty"`
	// Source object of the external lineage relationship.
	Source ExternalLineageObject `json:"source"`
	// Target object of the external lineage relationship.
	Target ExternalLineageObject `json:"target"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateRequestExternalLineage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRequestExternalLineage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateSchema struct {
	// Name of parent catalog.
	CatalogName string `json:"catalog_name"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Dropped objects in the schema remain recoverable for this number of hours
	// before they are permanently purged. The value must be 0 or within the
	// range allowed for the metastore. Schemas that do not set their own value
	// inherit the parent catalog's value.
	CustomMaxRetentionHours int64 `json:"custom_max_retention_hours,omitempty"`
	// Name of schema, relative to parent catalog.
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// Managed tables and volumes in the schema are stored under this storage
	// root URL. The path must be inside an external location on which the
	// caller has the CREATE MANAGED STORAGE privilege. If omitted, managed data
	// is stored under the parent catalog's storage root, or under the
	// metastore's root storage location when the catalog has none. Schemas in
	// Delta Sharing catalogs cannot have a storage root, and schemas in a
	// federated Hive metastore catalog must set one.
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateSchema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateSchema) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateSecretRequest struct {
	// The secret object to create. The **name**, **catalog_name**,
	// **schema_name**, and **value** fields are required.
	Secret Secret `json:"secret"`
}

func (s *CreateSecretRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateSkillRequest struct {
	// Name of the parent schema. Format: `schemas/{catalog}.{schema}`. Each
	// `{...}` component is capped at 255 characters individually.
	Parent string `json:"-" url:"parent"`
	// The skill to create. `comment` is the only accepted client input and may
	// be omitted. Do not set `name`; the server derives it from `parent` and
	// `skill_id`.
	Skill Skill `json:"skill"`
	// Name of the skill within the parent schema. The name must not contain
	// periods, spaces, forward slashes, percent signs, or control characters,
	// and it is stored in lowercase. It is independent of the bundle name read
	// from SKILL.md.
	SkillId string `json:"-" url:"skill_id"`
}

func (s *CreateSkillRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleRequest `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityRequest `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name string `json:"name"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`
	// Supplying true to this argument skips validation of the created
	// credential.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateTableConstraint struct {
	Constraint TableConstraint `json:"constraint"`
	// The full name of the table referenced by the constraint.
	FullNameArg string `json:"full_name_arg"`
}

func (s *CreateTableConstraint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateTableRequest struct {
	// Name of parent catalog. Required for creation using catalog and schema
	// names.
	CatalogName string `json:"catalog_name"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns []ColumnInfo `json:"columns,omitempty"`
	// The format of the data that backs the table, such as DELTA, PARQUET, CSV
	// or JSON. Foreign tables report the external system that holds them, for
	// example SNOWFLAKE_FORMAT or MYSQL_FORMAT.
	DataSourceFormat DataSourceFormat `json:"data_source_format"`
	// Name of table, relative to parent schema.
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// Name of parent schema relative to its parent catalog. Required for
	// creation using catalog and schema names.
	SchemaName string `json:"schema_name"`
	// The cloud storage URL under which the data files of the table are stored.
	// For an EXTERNAL table, it must be set on creation to a path within an
	// existing external location. For a MANAGED table, the server assigns a
	// path under the managed storage location of the parent schema, catalog or
	// metastore and returns it in this field. Views have no storage location.
	StorageLocation string `json:"storage_location"`
	// The type of the table. MANAGED tables store their data in storage that
	// Unity Catalog manages, EXTERNAL tables store it in a registered external
	// location, FOREIGN tables mirror objects of an external database through a
	// connection, and the shallow clone types reference the data files of a
	// base table. VIEW, MATERIALIZED_VIEW, STREAMING_TABLE and METRIC_VIEW are
	// defined by a view definition. See
	// https://docs.databricks.com/aws/en/tables/ for an overview of the table
	// types.
	TableType TableType `json:"table_type"`
}

func (s *CreateTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateVolumeRequestContent struct {
	// The name of the catalog where the schema and the volume are
	CatalogName string `json:"catalog_name"`
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`
	// The name of the volume
	Name string `json:"name"`
	// The name of the schema where the volume is
	SchemaName string `json:"schema_name"`
	// The cloud storage URL where the files of the volume are stored. For an
	// external volume, it must be set on creation to a path within an existing
	// external location. For a managed volume, it must be omitted on creation;
	// the volume is placed under the managed storage location of the parent
	// schema, catalog, or metastore, and the assigned path is returned in this
	// field.
	StorageLocation string `json:"storage_location,omitempty"`
	// The type of the volume. An external volume is located in the specified
	// external location. A managed volume is located in the default location
	// which is specified by the parent schema, or the parent catalog, or the
	// Metastore. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/aws/en/volumes/managed-vs-external
	VolumeType VolumeType `json:"volume_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateVolumeRequestContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateVolumeRequestContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A credential that is dependent on a SQL object.
type CredentialDependency struct {
	// Full name of the dependent credential, in the form of
	// __credential_name__.
	CredentialName string `json:"credential_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CredentialDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CredentialDependency) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A storage credential represents an authentication and authorization mechanism
// for accessing data stored on your cloud tenant. It holds exactly one cloud
// credential, for example an AWS IAM role, an Azure managed identity or service
// principal, or a GCP service account. Only the field for the selected
// credential type is set in requests and responses.
type CredentialInfo struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentity `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// Time at which this credential was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	// The full name of the credential.
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the credential.
	Id string `json:"id,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of the parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name string `json:"name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Indicates the purpose of the credential.
	Purpose CredentialPurpose `json:"purpose,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Whether this credential is the current metastore's root storage
	// credential. Only applicable when purpose is **STORAGE**.
	UsedForManagedStorage bool `json:"used_for_managed_storage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CredentialInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CredentialInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CredentialPurpose string

const CredentialPurposeService CredentialPurpose = `SERVICE`

const CredentialPurposeStorage CredentialPurpose = `STORAGE`

// String representation for [fmt.Print]
func (f *CredentialPurpose) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CredentialPurpose) Set(v string) error {
	switch v {
	case `SERVICE`, `STORAGE`:
		*f = CredentialPurpose(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SERVICE", "STORAGE"`, v)
	}
}

// Values returns all possible values for CredentialPurpose.
//
// There is no guarantee on the order of the values in the slice.
func (f *CredentialPurpose) Values() []CredentialPurpose {
	return []CredentialPurpose{
		CredentialPurposeService,
		CredentialPurposeStorage,
	}
}

// Type always returns CredentialPurpose to satisfy [pflag.Value] interface
func (f *CredentialPurpose) Type() string {
	return "CredentialPurpose"
}

type CredentialType string

const CredentialTypeAnyStaticCredential CredentialType = `ANY_STATIC_CREDENTIAL`

const CredentialTypeBearerToken CredentialType = `BEARER_TOKEN`

const CredentialTypeEdgegridAkamai CredentialType = `EDGEGRID_AKAMAI`

const CredentialTypeOauthAccessToken CredentialType = `OAUTH_ACCESS_TOKEN`

const CredentialTypeOauthM2m CredentialType = `OAUTH_M2M`

const CredentialTypeOauthMtls CredentialType = `OAUTH_MTLS`

const CredentialTypeOauthRefreshToken CredentialType = `OAUTH_REFRESH_TOKEN`

const CredentialTypeOauthResourceOwnerPassword CredentialType = `OAUTH_RESOURCE_OWNER_PASSWORD`

const CredentialTypeOauthU2m CredentialType = `OAUTH_U2M`

const CredentialTypeOauthU2mMapping CredentialType = `OAUTH_U2M_MAPPING`

const CredentialTypeOidcToken CredentialType = `OIDC_TOKEN`

const CredentialTypePemPrivateKey CredentialType = `PEM_PRIVATE_KEY`

const CredentialTypeServiceCredential CredentialType = `SERVICE_CREDENTIAL`

const CredentialTypeSswsToken CredentialType = `SSWS_TOKEN`

const CredentialTypeUnknownCredentialType CredentialType = `UNKNOWN_CREDENTIAL_TYPE`

const CredentialTypeUsernamePassword CredentialType = `USERNAME_PASSWORD`

// String representation for [fmt.Print]
func (f *CredentialType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CredentialType) Set(v string) error {
	switch v {
	case `ANY_STATIC_CREDENTIAL`, `BEARER_TOKEN`, `EDGEGRID_AKAMAI`, `OAUTH_ACCESS_TOKEN`, `OAUTH_M2M`, `OAUTH_MTLS`, `OAUTH_REFRESH_TOKEN`, `OAUTH_RESOURCE_OWNER_PASSWORD`, `OAUTH_U2M`, `OAUTH_U2M_MAPPING`, `OIDC_TOKEN`, `PEM_PRIVATE_KEY`, `SERVICE_CREDENTIAL`, `SSWS_TOKEN`, `UNKNOWN_CREDENTIAL_TYPE`, `USERNAME_PASSWORD`:
		*f = CredentialType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ANY_STATIC_CREDENTIAL", "BEARER_TOKEN", "EDGEGRID_AKAMAI", "OAUTH_ACCESS_TOKEN", "OAUTH_M2M", "OAUTH_MTLS", "OAUTH_REFRESH_TOKEN", "OAUTH_RESOURCE_OWNER_PASSWORD", "OAUTH_U2M", "OAUTH_U2M_MAPPING", "OIDC_TOKEN", "PEM_PRIVATE_KEY", "SERVICE_CREDENTIAL", "SSWS_TOKEN", "UNKNOWN_CREDENTIAL_TYPE", "USERNAME_PASSWORD"`, v)
	}
}

// Values returns all possible values for CredentialType.
//
// There is no guarantee on the order of the values in the slice.
func (f *CredentialType) Values() []CredentialType {
	return []CredentialType{
		CredentialTypeAnyStaticCredential,
		CredentialTypeBearerToken,
		CredentialTypeEdgegridAkamai,
		CredentialTypeOauthAccessToken,
		CredentialTypeOauthM2m,
		CredentialTypeOauthMtls,
		CredentialTypeOauthRefreshToken,
		CredentialTypeOauthResourceOwnerPassword,
		CredentialTypeOauthU2m,
		CredentialTypeOauthU2mMapping,
		CredentialTypeOidcToken,
		CredentialTypePemPrivateKey,
		CredentialTypeServiceCredential,
		CredentialTypeSswsToken,
		CredentialTypeUnknownCredentialType,
		CredentialTypeUsernamePassword,
	}
}

// Type always returns CredentialType to satisfy [pflag.Value] interface
func (f *CredentialType) Type() string {
	return "CredentialType"
}

type CredentialValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message string `json:"message,omitempty"`
	// The results of the tested operation.
	Result ValidateCredentialResult `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CredentialValidationResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CredentialValidationResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Data source format
type DataSourceFormat string

const DataSourceFormatAvro DataSourceFormat = `AVRO`

const DataSourceFormatBigqueryFormat DataSourceFormat = `BIGQUERY_FORMAT`

const DataSourceFormatCsv DataSourceFormat = `CSV`

const DataSourceFormatDatabricksFormat DataSourceFormat = `DATABRICKS_FORMAT`

const DataSourceFormatDatabricksRowStoreFormat DataSourceFormat = `DATABRICKS_ROW_STORE_FORMAT`

const DataSourceFormatDelta DataSourceFormat = `DELTA`

const DataSourceFormatDeltasharing DataSourceFormat = `DELTASHARING`

const DataSourceFormatDeltaUniformHudi DataSourceFormat = `DELTA_UNIFORM_HUDI`

const DataSourceFormatDeltaUniformIceberg DataSourceFormat = `DELTA_UNIFORM_ICEBERG`

const DataSourceFormatHive DataSourceFormat = `HIVE`

const DataSourceFormatIceberg DataSourceFormat = `ICEBERG`

const DataSourceFormatJson DataSourceFormat = `JSON`

const DataSourceFormatMongodbFormat DataSourceFormat = `MONGODB_FORMAT`

const DataSourceFormatMysqlFormat DataSourceFormat = `MYSQL_FORMAT`

const DataSourceFormatNetsuiteFormat DataSourceFormat = `NETSUITE_FORMAT`

const DataSourceFormatOracleFormat DataSourceFormat = `ORACLE_FORMAT`

const DataSourceFormatOrc DataSourceFormat = `ORC`

const DataSourceFormatParquet DataSourceFormat = `PARQUET`

const DataSourceFormatPostgresqlFormat DataSourceFormat = `POSTGRESQL_FORMAT`

const DataSourceFormatRedshiftFormat DataSourceFormat = `REDSHIFT_FORMAT`

const DataSourceFormatSalesforceDataCloudFormat DataSourceFormat = `SALESFORCE_DATA_CLOUD_FORMAT`

const DataSourceFormatSalesforceFormat DataSourceFormat = `SALESFORCE_FORMAT`

const DataSourceFormatSnowflakeFormat DataSourceFormat = `SNOWFLAKE_FORMAT`

const DataSourceFormatSqldwFormat DataSourceFormat = `SQLDW_FORMAT`

const DataSourceFormatSqlserverFormat DataSourceFormat = `SQLSERVER_FORMAT`

const DataSourceFormatTeradataFormat DataSourceFormat = `TERADATA_FORMAT`

const DataSourceFormatText DataSourceFormat = `TEXT`

const DataSourceFormatUnityCatalog DataSourceFormat = `UNITY_CATALOG`

const DataSourceFormatVectorIndexFormat DataSourceFormat = `VECTOR_INDEX_FORMAT`

const DataSourceFormatWorkdayRaasFormat DataSourceFormat = `WORKDAY_RAAS_FORMAT`

// String representation for [fmt.Print]
func (f *DataSourceFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataSourceFormat) Set(v string) error {
	switch v {
	case `AVRO`, `BIGQUERY_FORMAT`, `CSV`, `DATABRICKS_FORMAT`, `DATABRICKS_ROW_STORE_FORMAT`, `DELTA`, `DELTASHARING`, `DELTA_UNIFORM_HUDI`, `DELTA_UNIFORM_ICEBERG`, `HIVE`, `ICEBERG`, `JSON`, `MONGODB_FORMAT`, `MYSQL_FORMAT`, `NETSUITE_FORMAT`, `ORACLE_FORMAT`, `ORC`, `PARQUET`, `POSTGRESQL_FORMAT`, `REDSHIFT_FORMAT`, `SALESFORCE_DATA_CLOUD_FORMAT`, `SALESFORCE_FORMAT`, `SNOWFLAKE_FORMAT`, `SQLDW_FORMAT`, `SQLSERVER_FORMAT`, `TERADATA_FORMAT`, `TEXT`, `UNITY_CATALOG`, `VECTOR_INDEX_FORMAT`, `WORKDAY_RAAS_FORMAT`:
		*f = DataSourceFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVRO", "BIGQUERY_FORMAT", "CSV", "DATABRICKS_FORMAT", "DATABRICKS_ROW_STORE_FORMAT", "DELTA", "DELTASHARING", "DELTA_UNIFORM_HUDI", "DELTA_UNIFORM_ICEBERG", "HIVE", "ICEBERG", "JSON", "MONGODB_FORMAT", "MYSQL_FORMAT", "NETSUITE_FORMAT", "ORACLE_FORMAT", "ORC", "PARQUET", "POSTGRESQL_FORMAT", "REDSHIFT_FORMAT", "SALESFORCE_DATA_CLOUD_FORMAT", "SALESFORCE_FORMAT", "SNOWFLAKE_FORMAT", "SQLDW_FORMAT", "SQLSERVER_FORMAT", "TERADATA_FORMAT", "TEXT", "UNITY_CATALOG", "VECTOR_INDEX_FORMAT", "WORKDAY_RAAS_FORMAT"`, v)
	}
}

// Values returns all possible values for DataSourceFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *DataSourceFormat) Values() []DataSourceFormat {
	return []DataSourceFormat{
		DataSourceFormatAvro,
		DataSourceFormatBigqueryFormat,
		DataSourceFormatCsv,
		DataSourceFormatDatabricksFormat,
		DataSourceFormatDatabricksRowStoreFormat,
		DataSourceFormatDelta,
		DataSourceFormatDeltasharing,
		DataSourceFormatDeltaUniformHudi,
		DataSourceFormatDeltaUniformIceberg,
		DataSourceFormatHive,
		DataSourceFormatIceberg,
		DataSourceFormatJson,
		DataSourceFormatMongodbFormat,
		DataSourceFormatMysqlFormat,
		DataSourceFormatNetsuiteFormat,
		DataSourceFormatOracleFormat,
		DataSourceFormatOrc,
		DataSourceFormatParquet,
		DataSourceFormatPostgresqlFormat,
		DataSourceFormatRedshiftFormat,
		DataSourceFormatSalesforceDataCloudFormat,
		DataSourceFormatSalesforceFormat,
		DataSourceFormatSnowflakeFormat,
		DataSourceFormatSqldwFormat,
		DataSourceFormatSqlserverFormat,
		DataSourceFormatTeradataFormat,
		DataSourceFormatText,
		DataSourceFormatUnityCatalog,
		DataSourceFormatVectorIndexFormat,
		DataSourceFormatWorkdayRaasFormat,
	}
}

// Type always returns DataSourceFormat to satisfy [pflag.Value] interface
func (f *DataSourceFormat) Type() string {
	return "DataSourceFormat"
}

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccount struct {
	// The Databricks internal ID that represents this managed identity.
	CredentialId string `json:"credential_id,omitempty"`
	// The email of the service account.
	Email string `json:"email,omitempty"`
	// The ID that represents the private key for this Service Account
	PrivateKeyId string `json:"private_key_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksGcpServiceAccount) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksGcpServiceAccount) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccountRequest struct {
}

func (s *DatabricksGcpServiceAccountRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccountResponse struct {
	// The Databricks internal ID that represents this managed identity.
	CredentialId string `json:"credential_id,omitempty"`
	// The email of the service account.
	Email string `json:"email,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksGcpServiceAccountResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksGcpServiceAccountResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteAccountMetastoreAssignmentRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *DeleteAccountMetastoreAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteAccountMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool `json:"-" url:"force,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteAccountMetastoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAccountMetastoreRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteAccountStorageCredentialRequest struct {
	// Force deletion even if the Storage Credential is not empty. Default is
	// false.
	Force bool `json:"-" url:"force,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Name of the storage credential.
	StorageCredentialName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteAccountStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAccountStorageCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteAliasRequest struct {
	// The name of the alias
	Alias string `json:"-" url:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
}

func (s *DeleteAliasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteCatalogRequest struct {
	// Force deletion even if the catalog is not empty.
	Force bool `json:"-" url:"force,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteCatalogRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteConnectionRequest struct {
	// The name of the connection to be deleted.
	Name string `json:"-" url:"-"`
}

func (s *DeleteConnectionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteCredentialRequest struct {
	// Force deletion even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the credential.
	NameArg string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteEntityTagAssignmentRequest struct {
	// The fully qualified name of the entity to which the tag is assigned
	EntityName string `json:"-" url:"-"`
	// The type of the entity to which the tag is assigned.
	EntityType string `json:"-" url:"-"`
	// Required. The key of the tag to delete
	TagKey string `json:"-" url:"-"`
}

func (s *DeleteEntityTagAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteExternalLineageRelationshipRequest struct {
	// The relationship to delete, identified by its `source` and `target`
	// objects. Since this field is a query parameter, flatten the nested
	// fields, for example
	// `external_lineage_relationship.source.table.name=main.sales.orders`.
	ExternalLineageRelationship DeleteRequestExternalLineage `json:"-" url:"external_lineage_relationship"`
}

func (s *DeleteExternalLineageRelationshipRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteExternalLocationRequest struct {
	// Force deletion even if there are dependent external tables or mounts.
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the external location.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteExternalLocationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteExternalLocationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteExternalMetadataRequest struct {
	// Name of the external metadata object to delete.
	Name string `json:"-" url:"-"`
}

func (s *DeleteExternalMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteFunctionRequest struct {
	// Force deletion even if the function is not empty.
	Force bool `json:"-" url:"force,omitempty"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function_name__).
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteFunctionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteFunctionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteMcpServiceRequest struct {
	// Optimistic concurrency token from the most recent read. When set, the
	// delete succeeds only if the resource has not changed. Leave unset for an
	// unconditional delete. For REST requests, URL-encode the base64 string
	// returned by the API when setting the `etag` query parameter.
	Etag string `json:"-" url:"etag,omitempty"`
	// Resource name of the MCP service. Format:
	// `mcp-services/{catalog}.{schema}.{mcp_service}`. Each `{...}` component
	// is capped at 255 characters individually.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteMcpServiceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteMcpServiceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteMcpServiceUserMappedCredentialRequest struct {
	// Resource name of the MCP service. Format:
	// `mcp-services/{catalog}.{schema}.{mcp_service}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteMcpServiceUserMappedCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Delete returns no resource; a dedicated (empty) response keeps the revoke
// RPC's shape owned here rather than google.protobuf.Empty.
type DeleteMcpServiceUserMappedCredentialResponse struct {
}

func (s *DeleteMcpServiceUserMappedCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool `json:"-" url:"force,omitempty"`
	// Unique ID of the metastore.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteMetastoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteMetastoreRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteModelProviderServiceRequest struct {
	// Optimistic concurrency token from the most recent read. When set, the
	// delete succeeds only if the resource has not changed. Leave unset for an
	// unconditional delete. For REST requests, URL-encode the base64 string
	// returned by the API when setting the `etag` query parameter.
	Etag string `json:"-" url:"etag,omitempty"`
	// Resource name of the model provider service. Format:
	// `model-provider-services/{catalog}.{schema}.{model_provider_service}`.
	// Each `{...}` component is capped at 255 characters individually.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteModelProviderServiceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteModelProviderServiceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteModelServiceRequest struct {
	// Optimistic concurrency token from the most recent read. When set, the
	// delete succeeds only if the resource has not changed. Leave unset for an
	// unconditional delete. For REST requests, URL-encode the base64 string
	// returned by the API when setting the `etag` query parameter.
	Etag string `json:"-" url:"etag,omitempty"`
	// Resource name of the model service. Format:
	// `model-services/{catalog}.{schema}.{model_service}`. Each `{...}`
	// component is capped at 255 characters individually.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteModelServiceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteModelServiceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteModelVersionRequest struct {
	// The three-level (fully qualified) name of the parent registered model, in
	// the form `catalog_name.schema_name.model_name`.
	FullName string `json:"-" url:"-"`
	// The integer version number of the model version
	Version int `json:"-" url:"-"`
}

func (s *DeleteModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteMonitorResponse struct {
}

func (s *DeleteMonitorResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"-" url:"-"`
}

func (s *DeleteOnlineTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeletePolicyRequest struct {
	// Required. The name of the policy to delete
	Name string `json:"-" url:"-"`
	// Required. The fully qualified name of the securable to delete the policy
	// from.
	OnSecurableFullname string `json:"-" url:"-"`
	// Required. The type of the securable to delete the policy from.
	OnSecurableType string `json:"-" url:"-"`
}

func (s *DeletePolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeletePolicyResponse struct {
}

func (s *DeletePolicyResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteQualityMonitorRequest struct {
	// UC table name in format `catalog.schema.table_name`. This field
	// corresponds to the {full_table_name_arg} arg in the endpoint path.
	TableName string `json:"-" url:"-"`
}

func (s *DeleteQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
}

func (s *DeleteRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A lineage relationship between a source object and a target object, where at
// least one of the two is an external metadata object. The other object can be
// a Unity Catalog table, a cloud storage path, a model version, or another
// external metadata object. A relationship is identified by its source and
// target; optional column mappings and key-value properties describe the data
// flow between them.
type DeleteRequestExternalLineage struct {
	// Unique identifier of the external lineage relationship.
	Id string `json:"id,omitempty" url:"id,omitempty"`
	// Source object of the external lineage relationship.
	Source ExternalLineageObject `json:"source" url:"source"`
	// Target object of the external lineage relationship.
	Target ExternalLineageObject `json:"target" url:"target"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteRequestExternalLineage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteRequestExternalLineage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteSchemaRequest struct {
	// Force deletion even if the schema is not empty.
	Force bool `json:"-" url:"force,omitempty"`
	// Full name of the schema, in the form catalog_name.schema_name.
	FullName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteSchemaRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteSchemaRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteSecretRequest struct {
	// The three-level (fully qualified) name of the secret (for example,
	// **catalog_name.schema_name.secret_name**).
	FullName string `json:"-" url:"-"`
}

func (s *DeleteSecretRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteSkillRequest struct {
	// Optimistic concurrency token from the most recent read. When set, the
	// delete succeeds only if the resource has not changed. Leave unset for an
	// unconditional delete. For REST requests, URL-encode the base64 string
	// returned by the API when setting the `etag` query parameter.
	Etag string `json:"-" url:"etag,omitempty"`
	// Full resource name of the skill. Format:
	// `skills/{catalog}.{schema}.{skill}`. Each `{...}` component is capped at
	// 255 characters individually.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteSkillRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteSkillRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteStorageCredentialRequest struct {
	// Force deletion even if there are dependent external locations or external
	// tables.
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the storage credential.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteStorageCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteTableConstraintRequest struct {
	// If true, try deleting all child constraints of the current constraint. If
	// false, reject this operation if the current constraint has any child
	// constraints.
	Cascade bool `json:"-" url:"cascade"`
	// The name of the constraint to delete.
	ConstraintName string `json:"-" url:"constraint_name"`
	// Full name of the table referenced by the constraint.
	FullName string `json:"-" url:"-"`
}

func (s *DeleteTableConstraintRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteTableRequest struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`
}

func (s *DeleteTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteVolumeRequest struct {
	// The three-level (fully qualified) name of the volume
	Name string `json:"-" url:"-"`
}

func (s *DeleteVolumeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Properties pertaining to the current state of the delta table as given by the
// commit server. This does not contain **delta.*** (input) properties in
// __TableInfo.properties__.
type DeltaRuntimePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	DeltaRuntimeProperties map[string]string `json:"delta_runtime_properties"`
}

func (s *DeltaRuntimePropertiesKvPairs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeltaSharingScopeEnum string

const DeltaSharingScopeEnumInternal DeltaSharingScopeEnum = `INTERNAL`

const DeltaSharingScopeEnumInternalAndExternal DeltaSharingScopeEnum = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (f *DeltaSharingScopeEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeltaSharingScopeEnum) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*f = DeltaSharingScopeEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Values returns all possible values for DeltaSharingScopeEnum.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeltaSharingScopeEnum) Values() []DeltaSharingScopeEnum {
	return []DeltaSharingScopeEnum{
		DeltaSharingScopeEnumInternal,
		DeltaSharingScopeEnumInternalAndExternal,
	}
}

// Type always returns DeltaSharingScopeEnum to satisfy [pflag.Value] interface
func (f *DeltaSharingScopeEnum) Type() string {
	return "DeltaSharingScopeEnum"
}

type DenyOptions struct {
	// List of privileges to deny. When any of these privileges are requested,
	// the policy will deny access if the principal and condition match.
	// Required on create and update.
	Privileges []string `json:"privileges"`
}

func (s *DenyOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A dependency of a SQL object. One of the following fields must be defined:
// __table__, __function__, __connection__, __credential__, __volume__, or
// __secret__.
type Dependency struct {
	// A dependency on a Unity Catalog connection.
	Connection *ConnectionDependency `json:"connection,omitempty"`
	// A dependency on a Unity Catalog credential.
	Credential *CredentialDependency `json:"credential,omitempty"`
	// A dependency on a Unity Catalog function.
	Function *FunctionDependency `json:"function,omitempty"`
	// A dependency on a Unity Catalog table.
	Table *TableDependency `json:"table,omitempty"`
}

func (s *Dependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A list of dependencies.
type DependencyList struct {
	// Array of dependencies.
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

func (s *DependencyList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DestinationType string

const DestinationTypeEmail DestinationType = `EMAIL`

const DestinationTypeGenericWebhook DestinationType = `GENERIC_WEBHOOK`

const DestinationTypeMicrosoftTeams DestinationType = `MICROSOFT_TEAMS`

const DestinationTypeSlack DestinationType = `SLACK`

const DestinationTypeUrl DestinationType = `URL`

// String representation for [fmt.Print]
func (f *DestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DestinationType) Set(v string) error {
	switch v {
	case `EMAIL`, `GENERIC_WEBHOOK`, `MICROSOFT_TEAMS`, `SLACK`, `URL`:
		*f = DestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EMAIL", "GENERIC_WEBHOOK", "MICROSOFT_TEAMS", "SLACK", "URL"`, v)
	}
}

// Values returns all possible values for DestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DestinationType) Values() []DestinationType {
	return []DestinationType{
		DestinationTypeEmail,
		DestinationTypeGenericWebhook,
		DestinationTypeMicrosoftTeams,
		DestinationTypeSlack,
		DestinationTypeUrl,
	}
}

// Type always returns DestinationType to satisfy [pflag.Value] interface
func (f *DestinationType) Type() string {
	return "DestinationType"
}

type DisableRequest struct {
	// The metastore ID under which the system schema lives.
	MetastoreId string `json:"-" url:"-"`
	// Name of the system schema, without the catalog prefix.
	SchemaName string `json:"-" url:"-"`
}

func (s *DisableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EffectivePermissionsList struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// The privileges conveyed to each principal (either directly or via
	// inheritance)
	PrivilegeAssignments []EffectivePrivilegeAssignment `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EffectivePermissionsList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePermissionsList) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EffectivePredictiveOptimizationFlag struct {
	// The name of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromName string `json:"inherited_from_name,omitempty"`
	// The type of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromType EffectivePredictiveOptimizationFlagInheritedFromType `json:"inherited_from_type,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	Value EnablePredictiveOptimization `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EffectivePredictiveOptimizationFlag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePredictiveOptimizationFlag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Values returns all possible values for EffectivePredictiveOptimizationFlagInheritedFromType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EffectivePredictiveOptimizationFlagInheritedFromType) Values() []EffectivePredictiveOptimizationFlagInheritedFromType {
	return []EffectivePredictiveOptimizationFlagInheritedFromType{
		EffectivePredictiveOptimizationFlagInheritedFromTypeCatalog,
		EffectivePredictiveOptimizationFlagInheritedFromTypeSchema,
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
	InheritedFromName string `json:"inherited_from_name,omitempty"`
	// The type of the object that conveys this privilege via inheritance. This
	// field is omitted when privilege is not inherited (it's assigned to the
	// securable itself).
	InheritedFromType SecurableType `json:"inherited_from_type,omitempty"`
	// The privilege assigned to the principal.
	Privilege Privilege `json:"privilege,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EffectivePrivilege) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePrivilege) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EffectivePrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal string `json:"principal,omitempty"`
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	Privileges []EffectivePrivilege `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EffectivePrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Values returns all possible values for EnablePredictiveOptimization.
//
// There is no guarantee on the order of the values in the slice.
func (f *EnablePredictiveOptimization) Values() []EnablePredictiveOptimization {
	return []EnablePredictiveOptimization{
		EnablePredictiveOptimizationDisable,
		EnablePredictiveOptimizationEnable,
		EnablePredictiveOptimizationInherit,
	}
}

// Type always returns EnablePredictiveOptimization to satisfy [pflag.Value] interface
func (f *EnablePredictiveOptimization) Type() string {
	return "EnablePredictiveOptimization"
}

type EnableRequest struct {
	// Name of the catalog in which to enable the system schema. The supported
	// values are `system` and `samples`. If not set, the system schema is
	// enabled in the `system` catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// The metastore ID under which the system schema lives.
	MetastoreId string `json:"-" url:"-"`
	// Name of the system schema, without the catalog prefix.
	SchemaName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Encryption options that apply to clients connecting to cloud storage.
type EncryptionDetails struct {
	// Server-Side Encryption properties for clients communicating with AWS s3.
	SseEncryptionDetails *SseEncryptionDetails `json:"sse_encryption_details,omitempty"`
}

func (s *EncryptionDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Encryption Settings are used to carry metadata for securable encryption at
// rest. Currently used for catalogs, we can use the information supplied here
// to interact with a CMK.
type EncryptionSettings struct {
	// optional Azure settings - only required if an Azure CMK is used.
	AzureEncryptionSettings *AzureEncryptionSettings `json:"azure_encryption_settings,omitempty"`
	// the AKV URL in Azure, null otherwise.
	AzureKeyVaultKeyId string `json:"azure_key_vault_key_id,omitempty"`
	// the CMK uuid in AWS and GCP, null otherwise.
	CustomerManagedKeyId string `json:"customer_managed_key_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EncryptionSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EncryptionSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Represents a tag assignment to an entity
type EntityTagAssignment struct {
	// The fully qualified name of the entity to which the tag is assigned
	EntityName string `json:"entity_name"`
	// The type of the entity to which the tag is assigned.
	EntityType string `json:"entity_type"`
	// The source type of the tag assignment, e.g., user-assigned or
	// system-assigned
	SourceType TagAssignmentSourceType `json:"source_type,omitempty"`
	// The key of the tag
	TagKey string `json:"tag_key"`
	// The value of the tag
	TagValue string `json:"tag_value,omitempty"`
	// The timestamp when the tag assignment was last updated
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// The user or principal who updated the tag assignment
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EntityTagAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EntityTagAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EnvironmentSettings struct {
	EnvironmentVersion string `json:"environment_version,omitempty"`

	JavaDependencies []string `json:"java_dependencies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnvironmentSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnvironmentSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExistsRequest struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`
}

func (s *ExistsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An external metadata object at one end of an external lineage relationship.
type ExternalLineageExternalMetadata struct {
	// Name of the external metadata object.
	Name string `json:"name,omitempty" url:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLineageExternalMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageExternalMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Represents the external metadata object in the lineage event.
type ExternalLineageExternalMetadataInfo struct {
	// Type of entity represented by the external metadata object.
	EntityType string `json:"entity_type,omitempty"`
	// Timestamp of the lineage event.
	EventTime string `json:"event_time,omitempty"`
	// Name of the external metadata object.
	Name string `json:"name,omitempty"`
	// Type of external system.
	SystemType SystemType `json:"system_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLineageExternalMetadataInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageExternalMetadataInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Represents the path information in the lineage event.
type ExternalLineageFileInfo struct {
	// Timestamp of the lineage event.
	EventTime string `json:"event_time,omitempty"`
	// URL of the path.
	Path string `json:"path,omitempty"`
	// The full name of the securable on the path.
	SecurableName string `json:"securable_name,omitempty"`
	// The securable type of the securable on the path.
	SecurableType string `json:"securable_type,omitempty"`
	// The storage location associated with securable on the path.
	StorageLocation string `json:"storage_location,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLineageFileInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageFileInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Lineage response containing lineage information of a data asset.
type ExternalLineageInfo struct {
	// Information about the edge metadata of the external lineage relationship.
	ExternalLineageInfo *ExternalLineageRelationshipInfo `json:"external_lineage_info,omitempty"`
	// Information about external metadata involved in the lineage relationship.
	ExternalMetadataInfo *ExternalLineageExternalMetadataInfo `json:"external_metadata_info,omitempty"`
	// Information about the file involved in the lineage relationship.
	FileInfo *ExternalLineageFileInfo `json:"file_info,omitempty"`
	// Information about the model version involved in the lineage relationship.
	ModelInfo *ExternalLineageModelVersionInfo `json:"model_info,omitempty"`
	// Information about the table involved in the lineage relationship.
	TableInfo *ExternalLineageTableInfo `json:"table_info,omitempty"`
}

func (s *ExternalLineageInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A version of a Unity Catalog registered model at one end of an external
// lineage relationship.
type ExternalLineageModelVersion struct {
	// Full name of the registered model, in the form `catalog.schema.model`.
	Name string `json:"name,omitempty" url:"name,omitempty"`
	// Version number of the model, as a string.
	Version string `json:"version,omitempty" url:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLineageModelVersion) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageModelVersion) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Represents the model version information in the lineage event.
type ExternalLineageModelVersionInfo struct {
	// Timestamp of the lineage event.
	EventTime string `json:"event_time,omitempty"`
	// Name of the model.
	ModelName string `json:"model_name,omitempty"`
	// Version number of the model.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLineageModelVersionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageModelVersionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An object at one end of an external lineage relationship. Exactly one of the
// fields is set.
type ExternalLineageObject struct {
	// An external metadata object registered in Unity Catalog.
	ExternalMetadata *ExternalLineageExternalMetadata `json:"external_metadata,omitempty" url:"external_metadata,omitempty"`
	// A version of a Unity Catalog registered model.
	ModelVersion *ExternalLineageModelVersion `json:"model_version,omitempty" url:"model_version,omitempty"`
	// A cloud storage path registered in Unity Catalog, i.e. under a table,
	// volume, or external location.
	Path *ExternalLineagePath `json:"path,omitempty" url:"path,omitempty"`
	// A Unity Catalog table.
	Table *ExternalLineageTable `json:"table,omitempty" url:"table,omitempty"`
}

func (s *ExternalLineageObject) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A cloud storage path at one end of an external lineage relationship.
type ExternalLineagePath struct {
	// URL of the cloud storage path. The path must be registered in Unity
	// Catalog, i.e. under a table, volume, or external location.
	Url string `json:"url,omitempty" url:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLineagePath) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineagePath) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A lineage relationship between a source object and a target object, where at
// least one of the two is an external metadata object. The other object can be
// a Unity Catalog table, a cloud storage path, a model version, or another
// external metadata object. A relationship is identified by its source and
// target; optional column mappings and key-value properties describe the data
// flow between them.
type ExternalLineageRelationship struct {
	// List of column relationships between source and target objects. The
	// combined size of all source and target column names is limited.
	Columns []ColumnRelationship `json:"columns,omitempty"`
	// Unique identifier of the external lineage relationship.
	Id string `json:"id,omitempty"`
	// Key-value properties associated with the external lineage relationship.
	// The combined length of all keys and values must not exceed 5000
	// characters.
	Properties map[string]string `json:"properties,omitempty"`
	// Source object of the external lineage relationship.
	Source ExternalLineageObject `json:"source"`
	// Target object of the external lineage relationship.
	Target ExternalLineageObject `json:"target"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLineageRelationship) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageRelationship) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A lineage relationship between a source object and a target object, where at
// least one of the two is an external metadata object. The other object can be
// a Unity Catalog table, a cloud storage path, a model version, or another
// external metadata object. A relationship is identified by its source and
// target; optional column mappings and key-value properties describe the data
// flow between them.
type ExternalLineageRelationshipInfo struct {
	// List of column relationships between source and target objects. The
	// combined size of all source and target column names is limited.
	Columns []ColumnRelationship `json:"columns,omitempty"`
	// Unique identifier of the external lineage relationship.
	Id string `json:"id,omitempty"`
	// Key-value properties associated with the external lineage relationship.
	// The combined length of all keys and values must not exceed 5000
	// characters.
	Properties map[string]string `json:"properties,omitempty"`
	// Source object of the external lineage relationship.
	Source ExternalLineageObject `json:"source"`
	// Target object of the external lineage relationship.
	Target ExternalLineageObject `json:"target"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLineageRelationshipInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageRelationshipInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Unity Catalog table at one end of an external lineage relationship.
type ExternalLineageTable struct {
	// Full name of the table, in the form `catalog.schema.table`.
	Name string `json:"name,omitempty" url:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLineageTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Represents the table information in the lineage event.
type ExternalLineageTableInfo struct {
	// Name of Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// Timestamp of the lineage event.
	EventTime string `json:"event_time,omitempty"`
	// Name of Table.
	Name string `json:"name,omitempty"`
	// Name of Schema.
	SchemaName string `json:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLineageTableInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageTableInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An external location combines a cloud storage path with a storage credential
// that authorizes access to that path. Privileges granted on the external
// location control which users and groups can access the cloud storage path,
// for example to create external tables, external volumes, or managed storage
// locations under it.
type ExternalLocationInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this external location was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of external location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique ID of the location's storage credential.
	CredentialId string `json:"credential_id,omitempty"`
	// Name of the storage credential used with this location.
	CredentialName string `json:"credential_name,omitempty"`
	// The effective value of `enable_file_events` after applying server-side
	// defaults.
	EffectiveEnableFileEvents bool `json:"effective_enable_file_events,omitempty"`
	// The effective file event queue configuration after applying server-side
	// defaults. Always populated when a queue is provisioned, regardless of
	// whether the user explicitly set `enable_file_events`. Use this field
	// instead of `file_event_queue` for reading the actual queue state.
	EffectiveFileEventQueue *FileEventQueue `json:"effective_file_event_queue,omitempty"`
	// Whether to enable file events on this external location. Defaults to
	// `true`. Set to `false` to disable file events. The actual applied value
	// may differ due to server-side defaults; check
	// `effective_enable_file_events` for the effective state.
	EnableFileEvents bool `json:"enable_file_events,omitempty"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback bool `json:"fallback,omitempty"`
	// File event queue settings. When `enable_file_events` is `true`, exactly
	// one of the queue options must be set. Must be omitted when
	// `enable_file_events` is `false`.
	FileEventQueue *FileEventQueue `json:"file_event_queue,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces. Cannot be set when creating an external
	// location.
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of metastore hosting the external location.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of the external location.
	Name string `json:"name,omitempty"`
	// The owner of the external location.
	Owner string `json:"owner,omitempty"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which this external location was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the external location.
	UpdatedBy string `json:"updated_by,omitempty"`
	// URL of the cloud storage path that the external location points to,
	// including the scheme (for example `s3://`, `abfss://`, or `gs://`).
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLocationInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLocationInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata about an object that lives outside of Databricks, such as a Kafka
// topic, a Tableau dashboard, or a Salesforce table. An external metadata
// object is a Unity Catalog securable that can be the source or the target of
// external lineage relationships, so that data flows into and out of Databricks
// appear in the Unity Catalog lineage graph. See [External lineage].
//
// [External lineage]: https://docs.databricks.com/aws/en/data-governance/unity-catalog/external-lineage
type ExternalMetadata struct {
	// Names of the columns of the object, used for column-level mappings in
	// external lineage relationships. Column names must be unique within the
	// list, and the combined size of all column names is limited.
	Columns []string `json:"columns,omitempty"`
	// Time at which this external metadata object was created.
	CreateTime string `json:"create_time,omitempty"`
	// Username of external metadata object creator.
	CreatedBy string `json:"created_by,omitempty"`
	// User-provided free-form text description of the object. The description
	// must not exceed 5000 characters.
	Description string `json:"description,omitempty"`
	// Type of the object within the external system, such as a table, a report,
	// or a dashboard. The value is free-form text of at most 255 characters.
	EntityType string `json:"entity_type"`
	// Unique identifier of the external metadata object. The ID is a UUID that
	// the system assigns on creation.
	Id string `json:"id,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of the external metadata object. The name is unique within the
	// metastore, must follow the [Unity Catalog naming requirements], and is
	// stored in lowercase. It cannot be changed after creation, and external
	// lineage relationships refer to the object by this name.
	//
	// [Unity Catalog naming requirements]: https://docs.databricks.com/aws/en/data-governance/unity-catalog/requirements#naming-requirements
	Name string `json:"name"`
	// Owner of the external metadata object. Defaults to the creator. To change
	// the owner, send an update whose update mask includes `owner`.
	Owner string `json:"owner,omitempty"`
	// Key-value properties attached to the object. Keys and values must be
	// non-empty, and the combined length of all keys and values must not exceed
	// 5000 characters.
	Properties map[string]string `json:"properties,omitempty"`
	// Type of the external system that holds the object. Use `OTHER` for a
	// system that is not in the list.
	SystemType SystemType `json:"system_type"`
	// Time at which this external metadata object was last modified.
	UpdateTime string `json:"update_time,omitempty"`
	// Username of user who last modified external metadata object.
	UpdatedBy string `json:"updated_by,omitempty"`
	// URL of the object in the external system. When set, lineage graph viewers
	// can follow it to the external asset. The URL must not exceed 2048
	// characters.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of an online table. Shown if the online table is in the
// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
type FailedStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may only be partially synced to the online
	// table. Only populated if the table is still online and available for
	// serving.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table. Only populated if the table is still online
	// and available for serving.
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FailedStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FailedStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileEventQueue struct {
	ManagedAqs *AzureQueueStorage `json:"managed_aqs,omitempty"`

	ManagedPubsub *GcpPubsub `json:"managed_pubsub,omitempty"`

	ManagedSqs *AwsSqsQueue `json:"managed_sqs,omitempty"`

	ProvidedAqs *AzureQueueStorage `json:"provided_aqs,omitempty"`

	ProvidedPubsub *GcpPubsub `json:"provided_pubsub,omitempty"`

	ProvidedSqs *AwsSqsQueue `json:"provided_sqs,omitempty"`
}

func (s *FileEventQueue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type FinalizeSkillRequest struct {
	// Full resource name of the skill. Format:
	// `skills/{catalog}.{schema}.{skill}`. Each `{...}` component is capped at
	// 255 characters individually.
	Name string `json:"-" url:"-"`
}

func (s *FinalizeSkillRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ForeignKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns []string `json:"child_columns"`
	// The name of the constraint.
	Name string `json:"name"`
	// Column names for this constraint.
	ParentColumns []string `json:"parent_columns"`
	// The full name of the parent constraint.
	ParentTable string `json:"parent_table"`
	// True if the constraint is RELY, false or unset if NORELY.
	Rely bool `json:"rely,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ForeignKeyConstraint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ForeignKeyConstraint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An expression that is evaluated at query time against per-request context.
// New variants (e.g., identity attributes) are added as additional oneof cases.
type FunctionArgExpression struct {
	// An expression that introspects tags at query time.
	TagIntrospection *TagIntrospectionExpression `json:"tag_introspection,omitempty"`
}

func (s *FunctionArgExpression) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type FunctionArgument struct {
	// The alias of a matched column.
	Alias string `json:"alias,omitempty"`
	// A constant literal.
	Constant string `json:"constant,omitempty"`
	// An expression evaluated at query time. Wraps per-request expression
	// variants (e.g., tag introspection) so new variants can be added without
	// extending the FunctionArgument oneof.
	FunctionArgExpression *FunctionArgExpression `json:"function_arg_expression,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FunctionArgument) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FunctionArgument) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A function that is dependent on a SQL object.
type FunctionDependency struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	FunctionFullName string `json:"function_full_name"`
}

func (s *FunctionDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A user-defined function (UDF) registered in Unity Catalog. A function belongs
// to a schema and is referenced as catalog_name.schema_name.function_name. Its
// body is either a SQL expression or query, or source code in an external
// language such as Python. See
// https://docs.databricks.com/aws/en/udf/unity-catalog for an overview.
type FunctionInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Name of parent Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this function was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of function creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Data type of the value the function returns. A table function uses
	// TABLE_TYPE and describes its result columns in return_params, except for
	// a polymorphic Python UDTF, which omits them.
	DataType ColumnTypeName `json:"data_type,omitempty"`
	// Language of the function body for an EXTERNAL function. It is required
	// when routine_body is EXTERNAL.
	ExternalLanguage string `json:"external_language,omitempty"`
	// Name of the function in its external language. It is reserved for future
	// use.
	ExternalName string `json:"external_name,omitempty"`
	// Return data type of the function in SQL syntax, including any type
	// parameters such as precision and scale.
	FullDataType string `json:"full_data_type,omitempty"`
	// Full name of Function, in form of
	// **catalog_name**.**schema_name**.**function_name**
	FullName string `json:"full_name,omitempty"`
	// Id of Function, relative to parent schema.
	FunctionId string `json:"function_id,omitempty"`
	// Input parameters of the function. When it is set, parameters must contain
	// at least one entry. A function that takes no arguments omits
	// input_params, even though the Create request schema marks it as required.
	InputParams *FunctionParameterInfos `json:"input_params,omitempty"`
	// Whether the function returns the same result for the same arguments. It
	// corresponds to the DETERMINISTIC characteristic of CREATE FUNCTION.
	IsDeterministic bool `json:"is_deterministic,omitempty"`
	// Whether the function is a null-call function in the sense of the SQL
	// standard, which returns NULL whenever any argument is NULL.
	IsNullCall bool `json:"is_null_call,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of function, relative to parent schema.
	Name string `json:"name,omitempty"`
	// Username of current owner of the function.
	Owner string `json:"owner,omitempty"`
	// How arguments are passed to the function body. S means that each argument
	// is passed as a single scalar value.
	ParameterStyle FunctionInfoParameterStyle `json:"parameter_style,omitempty"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties string `json:"properties,omitempty"`
	// Result columns of a table function. Set it when data_type is TABLE_TYPE,
	// except for a polymorphic Python UDTF (an EXTERNAL Python function with
	// full_data_type "()"), which must not set it.
	ReturnParams *FunctionParameterInfos `json:"return_params,omitempty"`
	// Language of the function body. SQL means that routine_definition holds a
	// SQL expression or query. EXTERNAL means that it holds source code in the
	// language named by external_language, and sql_data_access must then be
	// NO_SQL.
	RoutineBody FunctionInfoRoutineBody `json:"routine_body,omitempty"`
	// Body of the function. For a SQL function it is the expression or query
	// that computes the result; for an EXTERNAL function it is the source code
	// of the function. It is omitted when the caller can only browse the
	// function.
	RoutineDefinition string `json:"routine_definition,omitempty"`
	// Function dependencies. For external UDFs, dependencies may contain only
	// credential, secret, or volume objects.
	RoutineDependencies *DependencyList `json:"routine_dependencies,omitempty"`
	// Name of parent Schema relative to its parent Catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// Whose privileges the function body runs with. DEFINER runs it with the
	// privileges of the function owner.
	SecurityType FunctionInfoSecurityType `json:"security_type,omitempty"`
	// Specific name of the function. Set it to the same value as name.
	SpecificName string `json:"specific_name,omitempty"`
	// How the function body interacts with SQL data. CONTAINS_SQL means that
	// the body contains SQL but reads no table or view, READS_SQL_DATA means
	// that it reads a table or view, and NO_SQL is required for EXTERNAL
	// functions.
	SqlDataAccess FunctionInfoSqlDataAccess `json:"sql_data_access,omitempty"`
	// List of schemas whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// Time at which this function was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the function.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FunctionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FunctionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Values returns all possible values for FunctionInfoParameterStyle.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionInfoParameterStyle) Values() []FunctionInfoParameterStyle {
	return []FunctionInfoParameterStyle{
		FunctionInfoParameterStyleS,
	}
}

// Type always returns FunctionInfoParameterStyle to satisfy [pflag.Value] interface
func (f *FunctionInfoParameterStyle) Type() string {
	return "FunctionInfoParameterStyle"
}

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

// Values returns all possible values for FunctionInfoRoutineBody.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionInfoRoutineBody) Values() []FunctionInfoRoutineBody {
	return []FunctionInfoRoutineBody{
		FunctionInfoRoutineBodyExternal,
		FunctionInfoRoutineBodySql,
	}
}

// Type always returns FunctionInfoRoutineBody to satisfy [pflag.Value] interface
func (f *FunctionInfoRoutineBody) Type() string {
	return "FunctionInfoRoutineBody"
}

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

// Values returns all possible values for FunctionInfoSecurityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionInfoSecurityType) Values() []FunctionInfoSecurityType {
	return []FunctionInfoSecurityType{
		FunctionInfoSecurityTypeDefiner,
	}
}

// Type always returns FunctionInfoSecurityType to satisfy [pflag.Value] interface
func (f *FunctionInfoSecurityType) Type() string {
	return "FunctionInfoSecurityType"
}

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

// Values returns all possible values for FunctionInfoSqlDataAccess.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionInfoSqlDataAccess) Values() []FunctionInfoSqlDataAccess {
	return []FunctionInfoSqlDataAccess{
		FunctionInfoSqlDataAccessContainsSql,
		FunctionInfoSqlDataAccessNoSql,
		FunctionInfoSqlDataAccessReadsSqlData,
	}
}

// Type always returns FunctionInfoSqlDataAccess to satisfy [pflag.Value] interface
func (f *FunctionInfoSqlDataAccess) Type() string {
	return "FunctionInfoSqlDataAccess"
}

type FunctionParameterInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of Parameter.
	Name string `json:"name"`
	// Default value of the parameter as a SQL expression, used when a call
	// omits the argument.
	ParameterDefault string `json:"parameter_default,omitempty"`
	// Direction of the parameter. Only IN parameters, which pass a value into
	// the function, are supported.
	ParameterMode FunctionParameterMode `json:"parameter_mode,omitempty"`
	// Whether the entry is an input parameter (PARAM) or a result column of a
	// table function (COLUMN).
	ParameterType FunctionParameterType `json:"parameter_type,omitempty"`
	// Zero-based position of the parameter in the parameter list. Positions
	// must be unique and consecutive.
	Position int `json:"position"`
	// Unit range of the data type for INTERVAL parameters. It is unset for
	// other types.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// Data type of the parameter as the JSON representation of a Spark SQL
	// struct field, with name, type, nullable and metadata keys.
	TypeJson string `json:"type_json,omitempty"`
	// Name of type (INT, STRUCT, MAP, etc.)
	TypeName ColumnTypeName `json:"type_name"`
	// Digits of precision; required on Create for DecimalTypes.
	TypePrecision int `json:"type_precision,omitempty"`
	// Digits to right of decimal; Required on Create for DecimalTypes.
	TypeScale int `json:"type_scale,omitempty"`
	// Data type of the parameter in SQL syntax, including any type parameters
	// such as precision and scale.
	TypeText string `json:"type_text"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FunctionParameterInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FunctionParameterInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FunctionParameterInfos struct {
	// The parameters that make up the list, ordered by position.
	Parameters []FunctionParameterInfo `json:"parameters,omitempty"`
}

func (s *FunctionParameterInfos) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

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

// Values returns all possible values for FunctionParameterMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionParameterMode) Values() []FunctionParameterMode {
	return []FunctionParameterMode{
		FunctionParameterModeIn,
	}
}

// Type always returns FunctionParameterMode to satisfy [pflag.Value] interface
func (f *FunctionParameterMode) Type() string {
	return "FunctionParameterMode"
}

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

// Values returns all possible values for FunctionParameterType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionParameterType) Values() []FunctionParameterType {
	return []FunctionParameterType{
		FunctionParameterTypeColumn,
		FunctionParameterTypeParam,
	}
}

// Type always returns FunctionParameterType to satisfy [pflag.Value] interface
func (f *FunctionParameterType) Type() string {
	return "FunctionParameterType"
}

// GCP temporary credentials for API authentication. Read more at
// https://developers.google.com/identity/protocols/oauth2/service-account
type GcpOauthToken struct {
	// The OAuth 2.0 access token to pass as a bearer token to Google Cloud
	// APIs.
	OauthToken string `json:"oauth_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpOauthToken) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpOauthToken) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GcpPubsub struct {
	// Unique identifier included in the name of file events managed cloud
	// resources.
	ManagedResourceId string `json:"managed_resource_id,omitempty"`
	// The Pub/Sub subscription name in the format
	// `projects/{project}/subscriptions/{subscription name}`. Only required for
	// provided_pubsub.
	SubscriptionName string `json:"subscription_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpPubsub) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpPubsub) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenerateTemporaryPathCredentialRequest struct {
	// Optional. When set to true, the service will not validate that the
	// generated credentials can perform write operations, therefore no new
	// paths will be created and the response will not contain valid
	// credentials. Defaults to false.
	DryRun bool `json:"dry_run,omitempty"`
	// The operation being performed on the path.
	Operation PathOperation `json:"operation"`
	// URL for path-based access.
	Url string `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenerateTemporaryPathCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateTemporaryPathCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenerateTemporaryPathCredentialResponse struct {
	// Temporary AWS credentials that can be used to sign AWS API requests.
	AwsTempCredentials *AwsCredentials `json:"aws_temp_credentials,omitempty"`
	// A Microsoft Entra ID (formerly Azure Active Directory) token that is
	// issued for the credential's identity.
	AzureAad *AzureActiveDirectoryToken `json:"azure_aad,omitempty"`
	// An Azure user delegation shared access signature (SAS) that grants access
	// to the requested storage path.
	AzureUserDelegationSas *AzureUserDelegationSas `json:"azure_user_delegation_sas,omitempty"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// A GCP OAuth 2.0 access token that is issued for the credential's service
	// account.
	GcpOauthToken *GcpOauthToken `json:"gcp_oauth_token,omitempty"`
	// Temporary Cloudflare R2 credentials that can be used with the R2
	// S3-compatible API.
	R2TempCredentials *R2Credentials `json:"r2_temp_credentials,omitempty"`
	// The URL of the storage path accessible by the temporary credential.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenerateTemporaryPathCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateTemporaryPathCredentialResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The Azure cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialAzureOptions struct {
	// The resources to which the temporary Azure credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://learn.microsoft.com/python/api/azure-core/azure.core.credentials.tokencredential?view=azure-python).
	// Exactly one resource is supported, and its URL must end with the
	// `.default` scope.
	Resources []string `json:"resources,omitempty"`
}

func (s *GenerateTemporaryServiceCredentialAzureOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The GCP cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialGcpOptions struct {
	// The scopes to which the temporary GCP credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://google-auth.readthedocs.io/en/latest/reference/google.auth.html#google.auth.credentials.Credentials).
	// Each scope must be a Google OAuth 2.0 scope URL on the `googleapis.com`
	// or `www.googleapis.com` host.
	Scopes []string `json:"scopes,omitempty"`
}

func (s *GenerateTemporaryServiceCredentialGcpOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenerateTemporaryServiceCredentialRequest struct {
	// Options for an Azure service credential. They are required when the
	// service credential uses an Azure managed identity.
	AzureOptions *GenerateTemporaryServiceCredentialAzureOptions `json:"azure_options,omitempty"`
	// The name of the service credential used to generate a temporary
	// credential
	CredentialName string `json:"credential_name"`
	// Options for a GCP service credential. They are required when the service
	// credential uses a Databricks-managed GCP service account.
	GcpOptions *GenerateTemporaryServiceCredentialGcpOptions `json:"gcp_options,omitempty"`
}

func (s *GenerateTemporaryServiceCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenerateTemporaryTableCredentialRequest struct {
	// The operation performed against the table data, either READ or
	// READ_WRITE. If READ_WRITE is specified, the credentials returned will
	// have write permissions, otherwise, it will be read only.
	Operation TableOperation `json:"operation,omitempty"`
	// UUID of the table to read or write.
	TableId string `json:"table_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenerateTemporaryTableCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateTemporaryTableCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenerateTemporaryTableCredentialResponse struct {
	// Temporary AWS credentials that can be used to sign AWS API requests.
	AwsTempCredentials *AwsCredentials `json:"aws_temp_credentials,omitempty"`
	// A Microsoft Entra ID (formerly Azure Active Directory) token that is
	// issued for the credential's identity.
	AzureAad *AzureActiveDirectoryToken `json:"azure_aad,omitempty"`
	// An Azure user delegation shared access signature (SAS) that grants access
	// to the requested storage path.
	AzureUserDelegationSas *AzureUserDelegationSas `json:"azure_user_delegation_sas,omitempty"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// A GCP OAuth 2.0 access token that is issued for the credential's service
	// account.
	GcpOauthToken *GcpOauthToken `json:"gcp_oauth_token,omitempty"`
	// Temporary Cloudflare R2 credentials that can be used with the R2
	// S3-compatible API.
	R2TempCredentials *R2Credentials `json:"r2_temp_credentials,omitempty"`
	// The URL of the storage path accessible by the temporary credential.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenerateTemporaryTableCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateTemporaryTableCredentialResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Generate volume credentials RPC
type GenerateTemporaryVolumeCredentialRequest struct {
	// The operation performed against the volume data, either READ_VOLUME or
	// WRITE_VOLUME. If WRITE_VOLUME is specified, the credentials returned will
	// have write permissions, otherwise, it will be read only.
	Operation VolumeOperation `json:"operation,omitempty"`
	// Id of the volume to read or write.
	VolumeId string `json:"volume_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenerateTemporaryVolumeCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateTemporaryVolumeCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenerateTemporaryVolumeCredentialResponse struct {
	// Temporary AWS credentials that can be used to sign AWS API requests.
	AwsTempCredentials *AwsCredentials `json:"aws_temp_credentials,omitempty"`
	// A Microsoft Entra ID (formerly Azure Active Directory) token that is
	// issued for the credential's identity.
	AzureAad *AzureActiveDirectoryToken `json:"azure_aad,omitempty"`
	// An Azure user delegation shared access signature (SAS) that grants access
	// to the requested storage path.
	AzureUserDelegationSas *AzureUserDelegationSas `json:"azure_user_delegation_sas,omitempty"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// A GCP OAuth 2.0 access token that is issued for the credential's service
	// account.
	GcpOauthToken *GcpOauthToken `json:"gcp_oauth_token,omitempty"`
	// Temporary Cloudflare R2 credentials that can be used with the R2
	// S3-compatible API.
	R2TempCredentials *R2Credentials `json:"r2_temp_credentials,omitempty"`
	// The URL of the storage path accessible by the temporary credential.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenerateTemporaryVolumeCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateTemporaryVolumeCredentialResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetAccessRequestDestinationsRequest struct {
	// The full name of the securable: the names of its parents followed by its
	// own name, separated by dots, for example `main.sales.orders` for a table.
	FullName string `json:"-" url:"-"`
	// The type of the securable. Matching is case-insensitive, and hyphens may
	// replace underscores in multi-word types.
	SecurableType string `json:"-" url:"-"`
}

func (s *GetAccessRequestDestinationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetAccountMetastoreAssignmentRequest struct {
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *GetAccountMetastoreAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetAccountMetastoreRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

func (s *GetAccountMetastoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetAccountStorageCredentialRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Required. Name of the storage credential.
	StorageCredentialName string `json:"-" url:"-"`
}

func (s *GetAccountStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetArtifactAllowlistRequest struct {
	// The artifact type of the allowlist.
	ArtifactType ArtifactType `json:"-" url:"-"`
}

func (s *GetArtifactAllowlistRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetBindingsRequest struct {
	// Maximum number of workspace bindings to return. - When set to 0, the page
	// length is set to a server configured value (recommended); - When set to a
	// value greater than 0, the page length is the minimum of this value and a
	// server configured value; - When set to a value less than 0, an invalid
	// parameter error is returned; - If not set, all the workspace bindings are
	// returned (not recommended).
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The name of the securable.
	SecurableName string `json:"-" url:"-"`
	// The type of the securable to bind to a workspace (catalog,
	// storage_credential, credential, or external_location).
	SecurableType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetBindingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetBindingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetByAliasRequest struct {
	// The name of the alias
	Alias string `json:"-" url:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases bool `json:"-" url:"include_aliases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetByAliasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetByAliasRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetCatalogRequest struct {
	// Whether to return the catalog when the principal can only access
	// selective metadata for it through the BROWSE privilege. The response then
	// has browse_only set to true.
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCatalogRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetCatalogWorkspaceBindingsResponse struct {
	// The IDs of the workspaces that are bound to the catalog.
	Workspaces []int64 `json:"workspaces,omitempty"`
}

func (s *GetCatalogWorkspaceBindingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetConnectionRequest struct {
	// Name of the connection.
	Name string `json:"-" url:"-"`
}

func (s *GetConnectionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetCredentialRequest struct {
	// Name of the credential.
	NameArg string `json:"-" url:"-"`
}

func (s *GetCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetEffectiveRequest struct {
	// Full name of securable.
	FullName string `json:"-" url:"-"`
	// Specifies the maximum number of privileges to return (page length). Every
	// EffectivePrivilegeAssignment present in a single page response is
	// guaranteed to contain all the effective privileges granted on (or
	// inherited by) the requested Securable for the respective principal.
	//
	// If not set, all the effective permissions are returned. If set to -
	// lesser than 0: invalid parameter error - 0: page length is set to a
	// server configured value - lesser than 150 but greater than 0: invalid
	// parameter error (this is to ensure that server is able to return at least
	// one complete EffectivePrivilegeAssignment in a single page response) -
	// greater than (or equal to) 150: page length is the minimum of this value
	// and a server configured value
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token for the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	Principal string `json:"-" url:"principal,omitempty"`
	// Type of securable.
	SecurableType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetEffectiveRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetEffectiveRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetEntityTagAssignmentRequest struct {
	// The fully qualified name of the entity to which the tag is assigned
	EntityName string `json:"-" url:"-"`
	// The type of the entity to which the tag is assigned.
	EntityType string `json:"-" url:"-"`
	// Required. The key of the tag
	TagKey string `json:"-" url:"-"`
}

func (s *GetEntityTagAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExternalLocationRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata.
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Name of the external location.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetExternalLocationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetExternalLocationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetExternalMetadataRequest struct {
	// Name of the external metadata object to get.
	Name string `json:"-" url:"-"`
}

func (s *GetExternalMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetFunctionRequest struct {
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function_name__).
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetFunctionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetFunctionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetGrantRequest struct {
	// Full name of the securable. Tables, views, volumes and functions use the
	// three-level name `catalog.schema.object`; a metastore uses its metastore
	// ID.
	FullName string `json:"-" url:"-"`
	// Specifies the maximum number of privileges to return (page length). Every
	// PrivilegeAssignment present in a single page response is guaranteed to
	// contain all the privileges granted on the requested Securable for the
	// respective principal.
	//
	// If not set, all the permissions are returned. If set to - lesser than 0:
	// invalid parameter error - 0: page length is set to a server configured
	// value - lesser than 150 but greater than 0: invalid parameter error (this
	// is to ensure that server is able to return at least one complete
	// PrivilegeAssignment in a single page response) - greater than (or equal
	// to) 150: page length is the minimum of this value and a server configured
	// value
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	Principal string `json:"-" url:"principal,omitempty"`
	// Type of the securable, such as `catalog`, `schema`, `table` or `volume`.
	// The value is case-insensitive.
	SecurableType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetGrantRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetGrantRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetMcpServiceRequest struct {
	// Resource name of the MCP service. Format:
	// `mcp-services/{catalog}.{schema}.{mcp_service}`. Each `{...}` component
	// is capped at 255 characters individually.
	Name string `json:"-" url:"-"`
}

func (s *GetMcpServiceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetMcpServiceUserMappedCredentialRequest struct {
	// Resource name of the MCP service. Format:
	// `mcp-services/{catalog}.{schema}.{mcp_service}`.
	Name string `json:"-" url:"-"`
}

func (s *GetMcpServiceUserMappedCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetMetastoreRequest struct {
	// Unique ID of the metastore.
	Id string `json:"-" url:"-"`
}

func (s *GetMetastoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetMetastoreSummaryResponse struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud string `json:"cloud,omitempty"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope DeltaSharingScopeEnum `json:"delta_sharing_scope,omitempty"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled bool `json:"external_access_enabled,omitempty"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// Unique identifier of metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the metastore.
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string `json:"storage_root_credential_name,omitempty"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the metastore.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetMetastoreSummaryResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetMetastoreSummaryResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetModelProviderServiceRequest struct {
	// Resource name of the model provider service. Format:
	// `model-provider-services/{catalog}.{schema}.{model_provider_service}`.
	// Each `{...}` component is capped at 255 characters individually.
	Name string `json:"-" url:"-"`
}

func (s *GetModelProviderServiceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetModelServiceRequest struct {
	// Resource name of the model service. Format:
	// `model-services/{catalog}.{schema}.{model_service}`. Each `{...}`
	// component is capped at 255 characters individually.
	Name string `json:"-" url:"-"`
}

func (s *GetModelServiceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetModelVersionRequest struct {
	// The three-level (fully qualified) name of the parent registered model, in
	// the form `catalog_name.schema_name.model_name`.
	FullName string `json:"-" url:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases bool `json:"-" url:"include_aliases,omitempty"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// The integer version number of the model version
	Version int `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetModelVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"-" url:"-"`
}

func (s *GetOnlineTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// The privileges assigned to each principal
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPermissionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPermissionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPolicyRequest struct {
	// Required. The name of the policy to retrieve.
	Name string `json:"-" url:"-"`
	// Required. The fully qualified name of securable to retrieve policy for.
	OnSecurableFullname string `json:"-" url:"-"`
	// Required. The type of the securable to retrieve the policy for.
	OnSecurableType string `json:"-" url:"-"`
}

func (s *GetPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetQualityMonitorRequest struct {
	// UC table name in format `catalog.schema.table_name`. This field
	// corresponds to the {full_table_name_arg} arg in the endpoint path.
	TableName string `json:"-" url:"-"`
}

func (s *GetQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetQuotaRequest struct {
	// Full name of the parent resource. Provide the metastore ID if the parent
	// is a metastore.
	ParentFullName string `json:"-" url:"-"`
	// Securable type of the quota parent.
	ParentSecurableType string `json:"-" url:"-"`
	// Name of the quota. Follows the pattern of the quota type, with "-quota"
	// added as a suffix.
	QuotaName string `json:"-" url:"-"`
}

func (s *GetQuotaRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetQuotaResponse struct {
	// The returned QuotaInfo.
	QuotaInfo *QuotaInfo `json:"quota_info,omitempty"`
}

func (s *GetQuotaResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRefreshRequest struct {
	// ID of the refresh.
	RefreshId int64 `json:"-" url:"-"`
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

func (s *GetRefreshRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
	// Whether to include registered model aliases in the response
	IncludeAliases bool `json:"-" url:"include_aliases,omitempty"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetSchemaRequest struct {
	// Full name of the schema, in the form catalog_name.schema_name.
	FullName string `json:"-" url:"-"`
	// Whether to return the schema when the principal can only access selective
	// metadata for it through the BROWSE privilege. The response then has
	// browse_only set to true.
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetSchemaRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetSchemaRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetSecretRequest struct {
	// The three-level (fully qualified) name of the secret (for example,
	// **catalog_name.schema_name.secret_name**).
	FullName string `json:"-" url:"-"`
	// Whether to include the secret value in the response. Defaults to false.
	// Requires the **READ_SECRET** privilege.
	IncludeValue bool `json:"-" url:"include_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetSecretRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetSecretRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetSkillRequest struct {
	// Full resource name of the skill. Format:
	// `skills/{catalog}.{schema}.{skill}`. Each `{...}` component is capped at
	// 255 characters individually.
	Name string `json:"-" url:"-"`
}

func (s *GetSkillRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetStorageCredentialRequest struct {
	// Name of the storage credential.
	Name string `json:"-" url:"-"`
}

func (s *GetStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetTableRequest struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for.
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata bool `json:"-" url:"include_delta_metadata,omitempty"`
	// Whether to include a manifest containing table capabilities in the
	// response.
	IncludeManifestCapabilities bool `json:"-" url:"include_manifest_capabilities,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetTableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetWorkspaceBindingRequest struct {
	// The name of the catalog.
	Name string `json:"-" url:"-"`
}

func (s *GetWorkspaceBindingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceBindingsResponse struct {
	// List of workspace bindings
	Bindings []WorkspaceBinding `json:"bindings,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetWorkspaceBindingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetWorkspaceBindingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A governed tag referenced by a policy condition.
type GovernedTagReference struct {
	// The governed tag key.
	TagKey string `json:"tag_key"`
	// The governed tag value, when the policy references a specific value.
	TagValue string `json:"tag_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GovernedTagReference) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GovernedTagReference) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GrantOptions struct {
	// List of privileges to grant. When any of these privileges are requested,
	// the policy will grant access if the principal and condition match.
	// Required on create and update.
	Privileges []string `json:"privileges"`
}

func (s *GrantOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Configuration for logging request and response payloads to a Unity Catalog
// inference table. When this configuration is present, payload logging is
// enabled by default.
type InferenceTableConfig struct {
	// Whether the referenced inference table has been deleted. The
	// configuration remains visible so you can identify the broken dependency.
	// Payload logging cannot continue until the table is restored or the
	// configuration is updated.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// Parent Unity Catalog schema where the inference table is created, in the
	// form `schemas/{catalog}.{schema}`. Required when configuring an inference
	// table. After the inference table is created, this field cannot be
	// changed.
	Parent string `json:"parent"`
	// Resource name of the Unity Catalog table that receives the payload logs.
	// Format: `tables/{catalog}.{schema}.{table}`.
	Table string `json:"table,omitempty"`
	// Prefix used to form the inference table's registered name. AI Gateway
	// appends `_payload`; for example, `table_name_prefix = "orders"` creates
	// `orders_payload`. If unset, the prefix defaults to the service name. Read
	// `table` from the response for the resulting resource name. After the
	// inference table is created, this field cannot be changed.
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InferenceTableConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InferenceTableConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type IsolationMode string

const IsolationModeIsolationModeIsolated IsolationMode = `ISOLATION_MODE_ISOLATED`

const IsolationModeIsolationModeOpen IsolationMode = `ISOLATION_MODE_OPEN`

// String representation for [fmt.Print]
func (f *IsolationMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IsolationMode) Set(v string) error {
	switch v {
	case `ISOLATION_MODE_ISOLATED`, `ISOLATION_MODE_OPEN`:
		*f = IsolationMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ISOLATION_MODE_ISOLATED", "ISOLATION_MODE_OPEN"`, v)
	}
}

// Values returns all possible values for IsolationMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *IsolationMode) Values() []IsolationMode {
	return []IsolationMode{
		IsolationModeIsolationModeIsolated,
		IsolationModeIsolationModeOpen,
	}
}

// Type always returns IsolationMode to satisfy [pflag.Value] interface
func (f *IsolationMode) Type() string {
	return "IsolationMode"
}

type LineageDirection string

const LineageDirectionDownstream LineageDirection = `DOWNSTREAM`

const LineageDirectionUpstream LineageDirection = `UPSTREAM`

// String representation for [fmt.Print]
func (f *LineageDirection) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LineageDirection) Set(v string) error {
	switch v {
	case `DOWNSTREAM`, `UPSTREAM`:
		*f = LineageDirection(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DOWNSTREAM", "UPSTREAM"`, v)
	}
}

// Values returns all possible values for LineageDirection.
//
// There is no guarantee on the order of the values in the slice.
func (f *LineageDirection) Values() []LineageDirection {
	return []LineageDirection{
		LineageDirectionDownstream,
		LineageDirectionUpstream,
	}
}

// Type always returns LineageDirection to satisfy [pflag.Value] interface
func (f *LineageDirection) Type() string {
	return "LineageDirection"
}

type ListAccountMetastoreAssignmentsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

func (s *ListAccountMetastoreAssignmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The metastore assignments were successfully returned.
type ListAccountMetastoreAssignmentsResponse struct {
	WorkspaceIds []int64 `json:"workspace_ids,omitempty"`
}

func (s *ListAccountMetastoreAssignmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListAccountStorageCredentialsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

func (s *ListAccountStorageCredentialsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The metastore storage credentials were successfully returned.
type ListAccountStorageCredentialsResponse struct {
	// An array of metastore storage credentials.
	StorageCredentials []StorageCredentialInfo `json:"storage_credentials,omitempty"`
}

func (s *ListAccountStorageCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListCatalogsRequest struct {
	// Whether to include catalogs for which the principal can only access
	// selective metadata through the BROWSE privilege. Such catalogs are
	// returned with browse_only set to true.
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Whether to include catalogs not bound to the workspace. Effective only if
	// the user has permission to update the catalog–workspace binding.
	IncludeUnbound bool `json:"-" url:"include_unbound,omitempty"`
	// Maximum number of catalogs to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid catalogs are returned (not
	// recommended). - Note: The number of returned catalogs might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further catalogs can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCatalogsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCatalogsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCatalogsResponse struct {
	// An array of catalog information objects.
	Catalogs []CatalogInfo `json:"catalogs,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCatalogsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCatalogsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListConnectionsRequest struct {
	// Maximum number of connections to return. - If not set, all connections
	// are returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Optional. Parent schema filter for listing schema-level connections, in
	// format "schemas/{catalog}.{schema}".
	Parent string `json:"-" url:"parent,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListConnectionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListConnectionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListConnectionsResponse struct {
	// An array of connection information objects.
	Connections []ConnectionInfo `json:"connections,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListConnectionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListConnectionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCredentialsRequest struct {
	// Whether to include credentials not bound to the workspace. Effective only
	// if the user has permission to update the credential–workspace binding.
	IncludeUnbound bool `json:"-" url:"include_unbound,omitempty"`
	// Maximum number of credentials to return. - If not set, the default max
	// page size is used. - When set to a value greater than 0, the page length
	// is the minimum of this value and a server-configured value. - When set to
	// 0, the page length is set to a server-configured value (recommended). -
	// When set to a value less than 0, an invalid parameter error is returned.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token to retrieve the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Return only credentials for the specified purpose.
	Purpose CredentialPurpose `json:"-" url:"purpose,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCredentialsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCredentialsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCredentialsResponse struct {
	Credentials []CredentialInfo `json:"credentials,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCredentialsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEffectivePrivilegeAssignmentsRequest struct {
	// Full name of securable.
	FullName string `json:"-" url:"-"`
	// Specifies the maximum number of privilege assignments to return (page
	// length). Every EffectivePrivilegeAssignment present in a single page
	// response is guaranteed to contain all the effective privileges granted on
	// (or inherited by) the requested Securable for the respective principal.
	//
	// If not set, a server-configured default is used. If set to - lesser than
	// 0: invalid parameter error - 0: page length is set to a server configured
	// value - lesser than 150 but greater than 0: invalid parameter error (this
	// is to ensure that server is able to return at least one complete
	// EffectivePrivilegeAssignment in a single page response) - greater than
	// (or equal to) 150: page length is the minimum of this value and a server
	// configured value
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	Principal string `json:"-" url:"principal,omitempty"`
	// Type of securable.
	SecurableType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEffectivePrivilegeAssignmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEffectivePrivilegeAssignmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEffectivePrivilegeAssignmentsResponse struct {
	// The effective privilege assignments for the securable (and optional
	// principal).
	EffectivePrivilegeAssignments []EffectivePrivilegeAssignment `json:"effective_privilege_assignments,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEffectivePrivilegeAssignmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEffectivePrivilegeAssignmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEntityTagAssignmentsRequest struct {
	// The fully qualified name of the entity to which the tag is assigned
	EntityName string `json:"-" url:"-"`
	// The type of the entity to which the tag is assigned.
	EntityType string `json:"-" url:"-"`
	// Optional. Maximum number of tag assignments to return in a single page
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Optional. Pagination token to retrieve the next page of results
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEntityTagAssignmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEntityTagAssignmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEntityTagAssignmentsResponse struct {
	// Optional. Pagination token for retrieving the next page of results
	NextPageToken string `json:"next_page_token,omitempty"`
	// The list of tag assignments
	TagAssignments []EntityTagAssignment `json:"tag_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEntityTagAssignmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEntityTagAssignmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExternalLineageRelationshipsRequest struct {
	// Direction of the relationships to return, relative to `object_info`.
	// `UPSTREAM` returns the relationships in which the object is the target,
	// so each result describes a source; `DOWNSTREAM` returns the relationships
	// in which the object is the source, so each result describes a target.
	LineageDirection LineageDirection `json:"-" url:"lineage_direction"`
	// The object to query external lineage relationships for. Since this field
	// is a query parameter, please flatten the nested fields. For example, if
	// the object is a table, the query parameter should look like:
	// `object_info.table.name=main.sales.customers`
	ObjectInfo ExternalLineageObject `json:"-" url:"object_info"`
	// Specifies the maximum number of external lineage relationships to return
	// in a single response. The value must be less than or equal to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExternalLineageRelationshipsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalLineageRelationshipsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExternalLineageRelationshipsResponse struct {
	// Relationships that match the query. Each entry describes the object at
	// the other end of the relationship in one of `table_info`, `file_info`,
	// `model_info`, or `external_metadata_info`, together with the relationship
	// itself in `external_lineage_info`.
	ExternalLineageRelationships []ExternalLineageInfo `json:"external_lineage_relationships,omitempty"`
	// Opaque token to retrieve the next page of results. Omitted when there are
	// no more results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExternalLineageRelationshipsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalLineageRelationshipsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExternalLocationsRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata.
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Whether to include external locations not bound to the workspace.
	// Effective only if the user has permission to update the
	// location–workspace binding.
	IncludeUnbound bool `json:"-" url:"include_unbound,omitempty"`
	// Maximum number of external locations to return. If not set, all the
	// external locations are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExternalLocationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalLocationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExternalLocationsResponse struct {
	// An array of external locations.
	ExternalLocations []ExternalLocationInfo `json:"external_locations,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExternalLocationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalLocationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExternalMetadataRequest struct {
	// Specifies the maximum number of external metadata objects to return in a
	// single response. The value must be less than or equal to 1000. Defaults
	// to 1000 when omitted.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExternalMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExternalMetadataResponse struct {
	ExternalMetadata []ExternalMetadata `json:"external_metadata,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExternalMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFunctionsRequest struct {
	// Name of parent catalog for functions of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of functions to return. If not set, all the functions are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent schema of functions.
	SchemaName string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFunctionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFunctionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFunctionsResponse struct {
	// An array of function information objects.
	Functions []FunctionInfo `json:"functions,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFunctionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFunctionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListMcpServicesRequest struct {
	// Maximum number of MCP services to return. Defaults to 200 when unset or
	// 0; the maximum is 1000 for `view = BASIC` and 200 for `view = FULL`. Use
	// `page_token` to retrieve additional pages.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token from the previous response.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent schema to list within, in the form `schemas/{catalog}.{schema}`.
	// Required. Each `{...}` component is capped at 255 characters
	// individually.
	Parent string `json:"-" url:"parent,omitempty"`
	// Fields to return for each service. `FULL` includes source-connection
	// details and rate-limit principal names. `BASIC` omits the source
	// connection and omits principal names from rate limits. Defaults to
	// `BASIC` when unset.
	View ListMcpServicesRequestView `json:"-" url:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMcpServicesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMcpServicesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Controls which fields are populated on each McpService in the response. The
// server uses `BASIC` when `view` is unset. Callers needing the full
// configuration must request it explicitly with `view = FULL`.
type ListMcpServicesRequestView string

const ListMcpServicesRequestViewBasic ListMcpServicesRequestView = `BASIC`

const ListMcpServicesRequestViewFull ListMcpServicesRequestView = `FULL`

// String representation for [fmt.Print]
func (f *ListMcpServicesRequestView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListMcpServicesRequestView) Set(v string) error {
	switch v {
	case `BASIC`, `FULL`:
		*f = ListMcpServicesRequestView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BASIC", "FULL"`, v)
	}
}

// Values returns all possible values for ListMcpServicesRequestView.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListMcpServicesRequestView) Values() []ListMcpServicesRequestView {
	return []ListMcpServicesRequestView{
		ListMcpServicesRequestViewBasic,
		ListMcpServicesRequestViewFull,
	}
}

// Type always returns ListMcpServicesRequestView to satisfy [pflag.Value] interface
func (f *ListMcpServicesRequestView) Type() string {
	return "ListMcpServicesRequestView"
}

// Response for listing MCP services.
type ListMcpServicesResponse struct {
	// The list of MCP services.
	McpServices []McpService `json:"mcp_services,omitempty"`
	// Pagination token for retrieving the next page. Empty when there are no
	// more results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMcpServicesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMcpServicesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListMetastoresRequest struct {
	// Maximum number of metastores to return. - when set to a value greater
	// than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned; - If not set, all the metastores are
	// returned (not recommended). - Note: The number of returned metastores
	// might be less than the specified max_results size, even zero. The only
	// definitive indication that no further metastores can be fetched is when
	// the next_page_token is unset from the response.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMetastoresRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMetastoresRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListMetastoresResponse struct {
	// An array of metastore information objects.
	Metastores []MetastoreInfo `json:"metastores,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMetastoresResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMetastoresResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListModelProviderServicesRequest struct {
	// Maximum number of provider services to return. Defaults to 200 when unset
	// or 0; the maximum is 1000 for `view = BASIC` and 200 for `view = FULL`.
	// Use `page_token` to retrieve additional pages.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token from the previous response.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent schema to list within, in the form `schemas/{catalog}.{schema}`.
	// Required. Each `{...}` component is capped at 255 characters
	// individually.
	Parent string `json:"-" url:"parent,omitempty"`
	// Fields to return for each service. `FULL` includes resolved
	// service-credential and inference-table details and rate-limit principal
	// names. `BASIC` omits those details and principal names from rate limits.
	// Defaults to `BASIC` when unset.
	View ListModelProviderServicesRequestView `json:"-" url:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListModelProviderServicesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelProviderServicesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Controls which fields are populated on each ModelProviderService in the
// response. The server uses `BASIC` when `view` is unset. Callers needing the
// full configuration must request it explicitly with `view = FULL`.
type ListModelProviderServicesRequestView string

const ListModelProviderServicesRequestViewBasic ListModelProviderServicesRequestView = `BASIC`

const ListModelProviderServicesRequestViewFull ListModelProviderServicesRequestView = `FULL`

// String representation for [fmt.Print]
func (f *ListModelProviderServicesRequestView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListModelProviderServicesRequestView) Set(v string) error {
	switch v {
	case `BASIC`, `FULL`:
		*f = ListModelProviderServicesRequestView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BASIC", "FULL"`, v)
	}
}

// Values returns all possible values for ListModelProviderServicesRequestView.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListModelProviderServicesRequestView) Values() []ListModelProviderServicesRequestView {
	return []ListModelProviderServicesRequestView{
		ListModelProviderServicesRequestViewBasic,
		ListModelProviderServicesRequestViewFull,
	}
}

// Type always returns ListModelProviderServicesRequestView to satisfy [pflag.Value] interface
func (f *ListModelProviderServicesRequestView) Type() string {
	return "ListModelProviderServicesRequestView"
}

// Response for listing model provider services.
type ListModelProviderServicesResponse struct {
	// The list of model provider services.
	ModelProviderServices []ModelProviderService `json:"model_provider_services,omitempty"`
	// Pagination token for retrieving the next page. Empty when there are no
	// more results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListModelProviderServicesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelProviderServicesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListModelServicesRequest struct {
	// Maximum number of model services to return. Defaults to 200 when unset or
	// 0; the maximum is 1000 for `view = BASIC` and 200 for `view = FULL`. Use
	// `page_token` to retrieve additional pages.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token from the previous response.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent schema to list within, in the form `schemas/{catalog}.{schema}`.
	// Required. Each `{...}` component is capped at 255 characters
	// individually.
	Parent string `json:"-" url:"parent,omitempty"`
	// Fields to return for each service. `FULL` includes destinations,
	// inference-table details, and rate-limit principal names. `BASIC` omits
	// destinations and inference-table details and omits principal names from
	// rate limits. Defaults to `BASIC` when unset.
	View ListModelServicesRequestView `json:"-" url:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListModelServicesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelServicesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Controls which fields are populated on each ModelService in the response. The
// server uses `BASIC` when `view` is unset. Callers needing the full
// configuration must request it explicitly with `view = FULL`.
type ListModelServicesRequestView string

const ListModelServicesRequestViewBasic ListModelServicesRequestView = `BASIC`

const ListModelServicesRequestViewFull ListModelServicesRequestView = `FULL`

// String representation for [fmt.Print]
func (f *ListModelServicesRequestView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListModelServicesRequestView) Set(v string) error {
	switch v {
	case `BASIC`, `FULL`:
		*f = ListModelServicesRequestView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BASIC", "FULL"`, v)
	}
}

// Values returns all possible values for ListModelServicesRequestView.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListModelServicesRequestView) Values() []ListModelServicesRequestView {
	return []ListModelServicesRequestView{
		ListModelServicesRequestViewBasic,
		ListModelServicesRequestViewFull,
	}
}

// Type always returns ListModelServicesRequestView to satisfy [pflag.Value] interface
func (f *ListModelServicesRequestView) Type() string {
	return "ListModelServicesRequestView"
}

// Response for listing model services.
type ListModelServicesResponse struct {
	// The list of model services.
	ModelServices []ModelService `json:"model_services,omitempty"`
	// Pagination token for retrieving the next page. Empty when there are no
	// more results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListModelServicesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelServicesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListModelVersionsRequest struct {
	// The full three-level name of the registered model under which to list
	// model versions
	FullName string `json:"-" url:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of model versions to return. If not set, the page length
	// is set to a server configured value (100, as of 1/3/2024). - when set to
	// a value greater than 0, the page length is the minimum of this value and
	// a server configured value(1000, as of 1/3/2024); - when set to 0, the
	// page length is set to a server configured value (100, as of 1/3/2024)
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListModelVersionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelVersionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListModelVersionsResponse struct {
	ModelVersions []ModelVersionInfo `json:"model_versions,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListModelVersionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelVersionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPoliciesRequest struct {
	// Optional. Whether to include policies defined on parent securables. By
	// default, the inherited policies are not included.
	IncludeInherited bool `json:"-" url:"include_inherited,omitempty"`
	// Optional. Maximum number of policies to return on a single page (page
	// length). - When not set or set to 0, the page length is set to a server
	// configured value (recommended); - When set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Required. The fully qualified name of securable to list policies for.
	OnSecurableFullname string `json:"-" url:"-"`
	// Required. The type of the securable to list policies for.
	OnSecurableType string `json:"-" url:"-"`
	// Optional. Opaque pagination token to go to next page based on previous
	// query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPoliciesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPoliciesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPoliciesResponse struct {
	// Optional opaque token for continuing pagination. `page_token` should be
	// set to this value for the next request to retrieve the next page of
	// results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The list of retrieved policies.
	Policies []PolicyInfo `json:"policies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPoliciesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPoliciesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPrivilegeAssignmentsRequest struct {
	// Full name of securable.
	FullName string `json:"-" url:"-"`
	// Specifies the maximum number of privilege assignments to return (page
	// length). Every PrivilegeAssignment present in a single page response is
	// guaranteed to contain all the privileges granted on the requested
	// Securable for the respective principal.
	//
	// If not set, page length is the server configured value. If set to -
	// lesser than 0: invalid parameter error - 0: page length is set to a
	// server configured value - lesser than 150 but greater than 0: invalid
	// parameter error (this is to ensure that server is able to return at least
	// one complete PrivilegeAssignment in a single page response) - greater
	// than (or equal to) 150: page length is the minimum of this value and a
	// server configured value
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	Principal string `json:"-" url:"principal,omitempty"`
	// Type of securable.
	SecurableType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPrivilegeAssignmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPrivilegeAssignmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPrivilegeAssignmentsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPrivilegeAssignmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPrivilegeAssignmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListQuotasRequest struct {
	// The number of quotas to return.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token for the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListQuotasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQuotasRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListQuotasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request.
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of returned QuotaInfos.
	Quotas []QuotaInfo `json:"quotas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListQuotasResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQuotasResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRefreshesRequest struct {
	// UC table name in format `catalog.schema.table_name`. table_name is case
	// insensitive and spaces are disallowed.
	TableName string `json:"-" url:"-"`
}

func (s *ListRefreshesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListRegisteredModelsRequest struct {
	// The identifier of the catalog under which to list registered models. If
	// specified, schema_name must be specified.
	CatalogName string `json:"-" url:"catalog_name,omitempty"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
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
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token to send for the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The identifier of the schema under which to list registered models. If
	// specified, catalog_name must be specified.
	SchemaName string `json:"-" url:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	NextPageToken string `json:"next_page_token,omitempty"`

	RegisteredModels []RegisteredModelInfo `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRegisteredModelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRegisteredModelsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSchemasRequest struct {
	// Parent catalog for schemas of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include schemas for which the principal can only access
	// selective metadata through the BROWSE privilege. Such schemas are
	// returned with browse_only set to true.
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of schemas to return. If not set, all the schemas are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of schema information objects.
	Schemas []SchemaInfo `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSchemasResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchemasResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSecretsRequest struct {
	// The name of the catalog under which to list secrets. Both
	// **catalog_name** and **schema_name** must be specified together.
	CatalogName string `json:"-" url:"catalog_name,omitempty"`
	// Maximum number of secrets to return.
	//
	// - If not specified, at most 1000 secrets are returned. - If set to a
	// value greater than 0, the page length is the minimum of this value and
	// 1000. - If set to 0, the page length is set to 1000. - If set to a value
	// less than 0, an invalid parameter error is returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token for the next page of results. Set it to the
	// **next_page_token** value from the previous response.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The name of the schema under which to list secrets. Both **catalog_name**
	// and **schema_name** must be specified together.
	SchemaName string `json:"-" url:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSecretsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSecretsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response message for ListSecrets.
type ListSecretsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. **page_token** should be set to this value for the next
	// request.
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of secret objects.
	Secrets []Secret `json:"secrets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSecretsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSecretsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSkillsRequest struct {
	// Maximum number of skills to return. Defaults to 200 when unset or 0; the
	// maximum is 200. Use `page_token` to retrieve additional pages.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token from a previous request.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Name of the parent schema. Format: `schemas/{catalog}.{schema}`. Each
	// `{...}` component is capped at 255 characters individually.
	//
	// Required: skill listing is schema-scoped, so `parent` must be set; an
	// unset or empty `parent` is rejected with INVALID_PARAMETER_VALUE.
	Parent string `json:"-" url:"parent"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSkillsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSkillsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for listing skills.
type ListSkillsResponse struct {
	// Pagination token for retrieving the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The list of skills.
	Skills []Skill `json:"skills,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSkillsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSkillsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListStorageCredentialsRequest struct {
	// Whether to include credentials not bound to the workspace. Effective only
	// if the user has permission to update the credential–workspace binding.
	IncludeUnbound bool `json:"-" url:"include_unbound,omitempty"`
	// Maximum number of storage credentials to return. If not set, all the
	// storage credentials are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	NextPageToken string `json:"next_page_token,omitempty"`

	StorageCredentials []StorageCredentialInfo `json:"storage_credentials,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListStorageCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStorageCredentialsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSummariesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include a manifest containing table capabilities in the
	// response.
	IncludeManifestCapabilities bool `json:"-" url:"include_manifest_capabilities,omitempty"`
	// Maximum number of summaries for tables to return. If not set, the page
	// length is set to a server configured value (10000, as of 1/5/2024). -
	// when set to a value greater than 0, the page length is the minimum of
	// this value and a server configured value (10000, as of 1/5/2024); - when
	// set to 0, the page length is set to a server configured value (10000, as
	// of 1/5/2024) (recommended); - when set to a value less than 0, an invalid
	// parameter error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// A sql LIKE pattern (% and _) for schema names. All schemas will be
	// returned if not set or empty.
	SchemaNamePattern string `json:"-" url:"schema_name_pattern,omitempty"`
	// A sql LIKE pattern (% and _) for table names. All tables will be returned
	// if not set or empty.
	TableNamePattern string `json:"-" url:"table_name_pattern,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSummariesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSummariesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSystemSchemasRequest struct {
	// Maximum number of schemas to return. - When set to 0, the page length is
	// set to a server configured value (recommended); - When set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - When set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all the schemas are returned (not
	// recommended).
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// The ID for the metastore in which the system schema resides.
	MetastoreId string `json:"-" url:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSystemSchemasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSystemSchemasRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSystemSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of system schema information objects.
	Schemas []SystemSchemaInfo `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSystemSchemasResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSystemSchemasResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTableSummariesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of table summaries.
	Tables []TableSummary `json:"tables,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTableSummariesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTableSummariesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTablesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for.
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Whether to include a manifest containing table capabilities in the
	// response.
	IncludeManifestCapabilities bool `json:"-" url:"include_manifest_capabilities,omitempty"`
	// Maximum number of tables to return. If not set, all the tables are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Whether to omit the columns of the table from the response or not.
	OmitColumns bool `json:"-" url:"omit_columns,omitempty"`
	// Whether to omit the properties of the table from the response or not.
	OmitProperties bool `json:"-" url:"omit_properties,omitempty"`
	// Whether to omit the username of the table (e.g. owner, updated_by,
	// created_by) from the response or not.
	OmitUsername bool `json:"-" url:"omit_username,omitempty"`
	// Opaque token to send for the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent schema of tables.
	SchemaName string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
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
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of table information objects.
	Tables []TableInfo `json:"tables,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTablesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTablesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListVolumesRequest struct {
	// The name of the catalog that contains the schema.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include volumes for which the principal can only access
	// selective metadata through the BROWSE privilege.
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
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
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token returned by a previous request. It must be included in the
	// request to retrieve the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The name of the schema whose volumes are listed.
	SchemaName string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
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
	NextPageToken string `json:"next_page_token,omitempty"`

	Volumes []VolumeInfo `json:"volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListVolumesResponseContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVolumesResponseContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MatchColumn struct {
	// Optional alias of the matched column.
	Alias string `json:"alias,omitempty"`
	// The condition expression used to match a table column.
	Condition string `json:"condition,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MatchColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MatchColumn) MarshalJSON() ([]byte, error) {
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

// Values returns all possible values for MatchType.
//
// There is no guarantee on the order of the values in the slice.
func (f *MatchType) Values() []MatchType {
	return []MatchType{
		MatchTypePrefixMatch,
	}
}

// Type always returns MatchType to satisfy [pflag.Value] interface
func (f *MatchType) Type() string {
	return "MatchType"
}

// A Unity Catalog securable that registers an MCP server through a Unity
// Catalog connection and exposes its tools for discovery, authorization, and
// invocation.
type McpService struct {
	// User-provided description.
	Comment string `json:"comment,omitempty"`
	// Connection, tool selectors, and rate limits. Required on Create. On
	// Update, provide this field when `update_mask` contains `config` or one of
	// its subpaths.
	Config *McpServiceConfig `json:"config,omitempty"`
	// Time the MCP service was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Username or service principal application ID of the principal that
	// created the MCP service.
	CreatedBy string `json:"created_by,omitempty"`
	// Username, group name, or service principal application ID of the MCP
	// service owner.
	EffectiveOwner string `json:"effective_owner,omitempty"`
	// Optimistic concurrency token returned on every read. To make an Update or
	// Delete conditional, pass the last-read value in that request's `etag`
	// field. In REST responses, this value is a base64 string; URL-encode it
	// when setting the `etag` query parameter.
	Etag string `json:"etag,omitempty"`
	// Unique identifier of the metastore that contains the MCP service.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Resource name of the MCP service. Format:
	// `mcp-services/{catalog}.{schema}.{mcp_service}`. Each `{...}` component
	// is capped at 255 characters individually. The server derives the name on
	// Create from `parent` and `mcp_service_id`; it is required and immutable
	// on Update, Get, and Delete.
	Name string `json:"name,omitempty"`
	// Time the MCP service was last modified.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Username or service principal application ID of the principal that last
	// updated the MCP service.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *McpService) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s McpService) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Operational configuration for an MCP service. Groups the source reference,
// tool selectors, and rate limits -- the fields that configure how the MCP
// service behaves at invocation time.
type McpServiceConfig struct {
	// Tool names or prefix patterns to expose from the MCP server. Use exact
	// tool names or prefix patterns such as `read_*`. An empty list exposes all
	// tools. At most 1,024 selectors are allowed, and each selector can contain
	// at most 256 characters.
	IncludeToolSelectors []string `json:"include_tool_selectors,omitempty"`
	// Rate limits for tool invocations. Supported scopes are user, group,
	// service principal, the service as a whole, and each user by default.
	// Request and token limits are supported. Empty when no rate limit is
	// configured.
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
	// Unity Catalog connection referencing the MCP server. Required on Create.
	SourceConnection *McpServiceConfigSourceConnection `json:"source_connection,omitempty"`
}

func (s *McpServiceConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Unity Catalog connection that points to the MCP server. On Create, provide
// `name` in the schema-scoped form
// `connections/{catalog}.{schema}.{connection}`. On read, the service populates
// the resolved connection metadata. If the connection is deleted, its reference
// remains visible so you can identify the broken dependency.
type McpServiceConfigSourceConnection struct {
	// Whether the referenced connection has been deleted. The MCP service keeps
	// the reference so callers can identify the broken dependency; tool
	// invocation fails until the source connection is updated.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// Resource name of the Unity Catalog connection used to access the MCP
	// server, in the form `connections/{catalog}.{schema}.{connection}`.
	Name string `json:"name"`
	// Options needed to build the U2M authorize request, returned as a flat
	// map. When set, it includes: `authorization_endpoint` (OAuth authorize
	// URL), `token_endpoint` (token-exchange URL), `oauth_scope`
	// (space-separated scopes to request), `client_id` (OAuth client id), and
	// `oauth_provider` (the OAuth provider).
	Options map[string]string `json:"options,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *McpServiceConfigSourceConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s McpServiceConfigSourceConnection) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A caller's per-user OAuth credential for an MCP service.
type McpServiceUserMappedCredential struct {
	// Token-expiry info for the credential, returned as a flat map:
	// `access_token_expiration` (always set) and `refresh_token_expiration`
	// (set when the credential has a refresh token). Both values are
	// timestamps.
	Options map[string]string `json:"options,omitempty"`
	// Provisioning state of the credential. `ACTIVE` means the caller is logged
	// in and the credential is usable; any other state means the login has not
	// completed.
	ProvisioningInfo *ProvisioningInfo `json:"provisioning_info,omitempty"`
}

func (s *McpServiceUserMappedCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Login input for an MCP service user credential. Carries the OAuth exchange
// fields as a flat map.
type McpServiceUserMappedCredentialLogin struct {
	// OAuth exchange fields: `pkce_verifier`, `authorization_code`, and
	// `oauth_redirect_uri`.
	Options map[string]string `json:"options,omitempty"`
}

func (s *McpServiceUserMappedCredentialLogin) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type MetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// deprecated. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// The unique ID of the metastore.
	MetastoreId string `json:"metastore_id"`
	// The unique ID of the Databricks workspace.
	WorkspaceId int64 `json:"workspace_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MetastoreAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetastoreAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MetastoreInfo struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud string `json:"cloud,omitempty"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope DeltaSharingScopeEnum `json:"delta_sharing_scope,omitempty"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled bool `json:"external_access_enabled,omitempty"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// Unique identifier of metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the metastore.
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string `json:"storage_root_credential_name,omitempty"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the metastore.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MetastoreInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetastoreInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Unity Catalog securable that stores authentication and request
// configuration for an external model provider, such as OpenAI, Azure OpenAI,
// or Amazon Bedrock. Model service destinations reference it to send requests
// to that provider.
//
// A model provider service can be referenced by multiple model services. A
// model service can route across multiple model provider services for traffic
// splitting or failover.
type ModelProviderService struct {
	// User-provided description.
	Comment string `json:"comment,omitempty"`
	// Provider authentication, exposed models, request-forwarding controls,
	// rate limits, and payload logging. Required on Create. On Update, it is
	// required only when `config` or one of its subpaths appears in
	// `update_mask`.
	Config *ModelProviderServiceConfig `json:"config,omitempty"`
	// Time the provider service was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Username of the user who created the model provider service.
	CreatedBy string `json:"created_by,omitempty"`
	// Username or group name of the model provider service owner.
	EffectiveOwner string `json:"effective_owner,omitempty"`
	// Optimistic concurrency token returned on every read. To make an Update or
	// Delete conditional, pass the last-read value in that request's `etag`
	// field. In REST responses, this value is a base64 string; URL-encode it
	// when setting the `etag` query parameter.
	Etag string `json:"etag,omitempty"`
	// Unique identifier of the metastore that contains the model provider
	// service.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Resource name of the provider service. Format:
	// `model-provider-services/{catalog}.{schema}.{model_provider_service}`.
	// Each `{...}` component is capped at 255 characters individually. The
	// server derives the name on Create from `parent` and
	// `model_provider_service_id`; it is required and immutable on Update, Get,
	// and Delete.
	Name string `json:"name,omitempty"`
	// Time the provider service was last modified.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Username of the user who last updated the model provider service.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderService) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderService) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Behavioral configuration for a ModelProviderService: provider authentication
// and provider-specific fields, the catalog of models this provider service can
// route to, and the passthrough policy that governs how request headers, query
// parameters, and unmanaged subpaths cross the trust boundary to the upstream
// provider.
type ModelProviderServiceConfig struct {
	// When true, allows models from the provider's built-in model catalog in
	// addition to explicit `targets`. When false, only models listed in
	// `targets` are routable. Defaults to false. CUSTOM providers have no
	// built-in model catalog and must leave this false and list their models in
	// `targets`.
	AllowAllTargets bool `json:"allow_all_targets,omitempty"`
	// Configuration for Amazon Bedrock. Set this field when `provider_type` is
	// `EXTERNAL_MODEL_PROVIDER_TYPE_AMAZON_BEDROCK`.
	AmazonBedrock *ModelProviderServiceConfigAmazonBedrockProviderConfig `json:"amazon_bedrock,omitempty"`
	// Configuration for Anthropic. Set this field when `provider_type` is
	// `EXTERNAL_MODEL_PROVIDER_TYPE_ANTHROPIC`.
	Anthropic *ModelProviderServiceConfigAnthropicProviderConfig `json:"anthropic,omitempty"`
	// Configuration for Azure OpenAI. Set this field when `provider_type` is
	// `EXTERNAL_MODEL_PROVIDER_TYPE_AZURE_OPENAI`.
	AzureOpenai *ModelProviderServiceConfigAzureOpenAiProviderConfig `json:"azure_openai,omitempty"`
	// Configuration for Amazon Bedrock Mantle, with the same fields as
	// `amazon_bedrock`. Set this field when `provider_type` is
	// `EXTERNAL_MODEL_PROVIDER_TYPE_BEDROCK_MANTLE`.
	BedrockMantle *ModelProviderServiceConfigAmazonBedrockProviderConfig `json:"bedrock_mantle,omitempty"`
	// Configuration for a custom OpenAI-compatible provider. Set this field
	// when `provider_type` is `EXTERNAL_MODEL_PROVIDER_TYPE_CUSTOM`.
	Custom *ModelProviderServiceConfigCustomProviderConfig `json:"custom,omitempty"`
	// Whether to forward incoming HTTP headers to the upstream provider.
	// Defaults to false and is configured for the entire provider service, not
	// per request. Upstream authentication is configured separately in the
	// provider-specific configuration.
	ForwardHeaders bool `json:"forward_headers,omitempty"`
	// Whether to forward incoming query parameters to the upstream provider.
	// Defaults to false and is configured for the entire provider service, not
	// per request.
	ForwardQueryParameters bool `json:"forward_query_parameters,omitempty"`
	// Whether to proxy paths that AI Gateway does not recognize as configured
	// provider-native API types. Defaults to false. When true, these paths are
	// forwarded unchanged to the upstream provider. When false, only recognized
	// API paths are served. Enabling this broadens the upstream API surface
	// exposed through the provider service. AI Gateway cannot evaluate Service
	// Policies on opaque unmanaged requests, so unmanaged requests are rejected
	// whenever Service Policies are attached to this securable.
	ForwardUnmanagedPaths bool `json:"forward_unmanaged_paths,omitempty"`
	// Configuration for Google Gemini Enterprise. Set this field when
	// `provider_type` is `EXTERNAL_MODEL_PROVIDER_TYPE_GEMINI_ENTERPRISE`.
	GeminiEnterprise *ModelProviderServiceConfigGeminiEnterpriseProviderConfig `json:"gemini_enterprise,omitempty"`
	// Payload logging configuration for requests sent directly to this provider
	// service. Requests routed through a model service are captured by that
	// model service's inference table instead.
	InferenceTable *InferenceTableConfig `json:"inference_table,omitempty"`
	// Configuration for Microsoft Foundry. Set this field when `provider_type`
	// is `EXTERNAL_MODEL_PROVIDER_TYPE_MICROSOFT_FOUNDRY`.
	MicrosoftFoundry *ModelProviderServiceConfigMicrosoftFoundryProviderConfig `json:"microsoft_foundry,omitempty"`
	// Configuration for OpenAI. Set this field when `provider_type` is
	// `EXTERNAL_MODEL_PROVIDER_TYPE_OPENAI`.
	Openai *ModelProviderServiceConfigOpenAiProviderConfig `json:"openai,omitempty"`
	// Pricing configuration for this provider service. An explicitly empty
	// configuration is retained on Create and Update and applies no adjustment.
	Pricing *ModelProviderServiceConfigProviderPricingConfig `json:"pricing,omitempty"`
	// External model provider. Required on Create and immutable thereafter. Set
	// the matching provider-specific configuration, such as `openai`,
	// `azure_openai`, or `amazon_bedrock`.
	ProviderType ModelProviderServiceConfigExternalModelProviderType `json:"provider_type,omitempty"`
	// Rate limits for requests sent directly to this provider service. Requests
	// routed through a model service use that model service's rate limits
	// instead.
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
	// Models and provider-native API types exposed by this provider service.
	// Each entry must include at least one `native_api_types` value. When
	// `allow_all_targets` is false, at least one entry is required and model
	// service destinations can reference only listed models. When
	// `allow_all_targets` is true, models from the provider's built-in model
	// catalog are also routable. CUSTOM providers must explicitly list their
	// models here and cannot enable `allow_all_targets`.
	Targets []ModelProviderServiceConfigModelTargetConfig `json:"targets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Amazon Bedrock provider configuration.
type ModelProviderServiceConfigAmazonBedrockProviderConfig struct {
	// Amazon Bedrock region and authentication configuration.
	Direct *ModelProviderServiceConfigAmazonBedrockProviderDirectConfig `json:"direct,omitempty"`
}

func (s *ModelProviderServiceConfigAmazonBedrockProviderConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Direct form of Amazon Bedrock provider config.
//
// Authentication is one of two mutually exclusive modes, exactly one of which
// must be supplied on Create: - Access keys: set `aws_access_key`, leave
// `service_credential` unset. - Unity Catalog service credential: set
// `service_credential.name` to the resource name `credentials/{name}`, leave
// `aws_access_key` unset. The credential value lives in Unity Catalog and is
// referenced by name, not held on this message. Setting more than one mode is
// rejected.
type ModelProviderServiceConfigAmazonBedrockProviderDirectConfig struct {
	// AWS access-key-pair authentication. Set `access_key_id` and
	// `secret_access_key.plaintext`. Mutually exclusive with
	// `service_credential`.
	AwsAccessKey *ModelProviderServiceConfigAwsAccessKey `json:"aws_access_key,omitempty"`
	// AWS region where the Bedrock endpoint is hosted (e.g., `us-east-1`).
	// Required on Create.
	Region string `json:"region,omitempty"`
	// Reference to a Unity Catalog service credential authorizing Bedrock
	// requests. On Create, supply `service_credential.name` in the form
	// `credentials/{name}`. Required on Create when using service-credential
	// authentication; mutually exclusive with `aws_access_key`. The credential
	// is referenced by name; its value is not carried here. Only supported on
	// AWS-hosted workspaces.
	ServiceCredential *ModelProviderServiceConfigServiceCredential `json:"service_credential,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfigAmazonBedrockProviderDirectConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfigAmazonBedrockProviderDirectConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Anthropic provider configuration. Exactly one of `direct` or `relayed` must
// be set on Create; the two are mutually exclusive.
type ModelProviderServiceConfigAnthropicProviderConfig struct {
	// Direct authentication with an API key supplied in
	// `direct.api_key.plaintext`. Required unless `relayed` is set.
	Direct *ModelProviderServiceConfigAnthropicProviderDirectConfig `json:"direct,omitempty"`
	// Relayed authentication. Each inference request supplies the caller's
	// OAuth token, which is forwarded to Anthropic. No Anthropic credential is
	// stored. Mutually exclusive with `direct`.
	Relayed *ModelProviderServiceConfigAnthropicProviderRelayedConfig `json:"relayed,omitempty"`
}

func (s *ModelProviderServiceConfigAnthropicProviderConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Direct form of Anthropic provider config.
type ModelProviderServiceConfigAnthropicProviderDirectConfig struct {
	// Anthropic API key. Required when creating the service. Supply the value
	// in `api_key.plaintext`.
	ApiKey *ModelProviderServiceConfigProviderSecret `json:"api_key,omitempty"`
}

func (s *ModelProviderServiceConfigAnthropicProviderDirectConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Relayed Anthropic provider configuration. Each inference request supplies the
// caller's OAuth token, which is forwarded to Anthropic. No Anthropic
// credential is stored.
type ModelProviderServiceConfigAnthropicProviderRelayedConfig struct {
}

func (s *ModelProviderServiceConfigAnthropicProviderRelayedConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// AWS access-key-pair auth for Amazon Bedrock: a SigV4-signing key pair.
type ModelProviderServiceConfigAwsAccessKey struct {
	// AWS access key ID. Required on Create when using access-key auth. Treated
	// as username-equivalent (not a secret value): round-trips on reads and is
	// scrubbed from audit logs.
	AccessKeyId string `json:"access_key_id,omitempty"`
	// AWS secret access key paired with `access_key_id`. Required when creating
	// a service with access-key authentication. Supply the value in
	// `secret_access_key.plaintext`.
	SecretAccessKey *ModelProviderServiceConfigProviderSecret `json:"secret_access_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfigAwsAccessKey) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfigAwsAccessKey) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Azure OpenAI provider configuration.
type ModelProviderServiceConfigAzureOpenAiProviderConfig struct {
	// Azure OpenAI endpoint and authentication configuration.
	Direct *ModelProviderServiceConfigAzureOpenAiProviderDirectConfig `json:"direct,omitempty"`
}

func (s *ModelProviderServiceConfigAzureOpenAiProviderConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Direct form of Azure OpenAI provider config. Exactly one of three
// mutually-exclusive auth modes must be supplied on Create: - API key: set
// `api_key`, leave `entra_service_principal` and `service_credential` unset. -
// Entra ID (service principal): set `entra_service_principal`, leave `api_key`
// and `service_credential` unset. - Unity Catalog service credential: set
// `service_credential.name` to the resource name `credentials/{name}`, leave
// `api_key` and `entra_service_principal` unset. The credential value lives in
// Unity Catalog and is referenced by name, not held on this message. Only
// supported on Azure-hosted workspaces. Setting more than one mode is rejected.
type ModelProviderServiceConfigAzureOpenAiProviderDirectConfig struct {
	// Azure OpenAI API key. Supply the value in `api_key.plaintext`. Mutually
	// exclusive with Entra ID and Unity Catalog service credential
	// authentication.
	ApiKey *ModelProviderServiceConfigProviderSecret `json:"api_key,omitempty"`
	// Full Azure OpenAI endpoint base URL, e.g.
	// `https://myresource.openai.azure.com`. Required on Create.
	BaseUrl string `json:"base_url,omitempty"`
	// Entra ID service-principal authentication. Set `tenant_id`, `client_id`,
	// and `client_secret.plaintext`. Mutually exclusive with `api_key` and
	// `service_credential`.
	EntraServicePrincipal *ModelProviderServiceConfigEntraServicePrincipal `json:"entra_service_principal,omitempty"`
	// Reference to a Unity Catalog service credential authorizing Azure OpenAI
	// requests. On Create, supply `service_credential.name` in the form
	// `credentials/{name}`. Required on Create when using service-credential
	// authentication; mutually exclusive with `api_key` and
	// `entra_service_principal`. The credential is referenced by name; its
	// value is not carried here. Only supported on Azure-hosted workspaces.
	ServiceCredential *ModelProviderServiceConfigServiceCredential `json:"service_credential,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfigAzureOpenAiProviderDirectConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfigAzureOpenAiProviderDirectConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Header-based API-key authentication for a custom provider: the secret is
// forwarded on outbound requests under a caller-chosen HTTP header, as
// `<api_key_name>: <api_key_value>`.
type ModelProviderServiceConfigCustomProviderApiKeyHeaderAuth struct {
	// HTTP header name that carries the API key on outbound requests (e.g.,
	// `Ocp-Apim-Subscription-Key`). The value forwarded under this header is
	// supplied via `api_key_value`.
	ApiKeyName string `json:"api_key_name,omitempty"`
	// Secret value forwarded under the `api_key_name` header on outbound
	// requests. Supply the value in `api_key_value.plaintext`.
	ApiKeyValue *ModelProviderServiceConfigProviderSecret `json:"api_key_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfigCustomProviderApiKeyHeaderAuth) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfigCustomProviderApiKeyHeaderAuth) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Custom OpenAI-compatible provider configuration.
type ModelProviderServiceConfigCustomProviderConfig struct {
	// Endpoint and authentication configuration for the custom provider.
	Direct *ModelProviderServiceConfigCustomProviderDirectConfig `json:"direct,omitempty"`
}

func (s *ModelProviderServiceConfigCustomProviderConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Direct form of a custom provider configuration. Set `api_key` to send the
// secret as an `Authorization` bearer token, or `header_auth` to forward it
// under a caller-chosen HTTP header.
type ModelProviderServiceConfigCustomProviderDirectConfig struct {
	// Bearer token forwarded in the `Authorization` header. Supply the value in
	// `api_key.plaintext`.
	ApiKey *ModelProviderServiceConfigProviderSecret `json:"api_key,omitempty"`
	// Endpoint URL of the OpenAI-compatible service (e.g.,
	// `https://api.example.com/v1`). Required on Create.
	BaseUrl string `json:"base_url,omitempty"`
	// Header-based API-key auth: the secret is forwarded on outbound requests
	// under a caller-chosen HTTP header rather than as an `Authorization`
	// bearer token. Set this instead of `api_key` for header auth.
	HeaderAuth *ModelProviderServiceConfigCustomProviderApiKeyHeaderAuth `json:"header_auth,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfigCustomProviderDirectConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfigCustomProviderDirectConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Entra ID (Azure AD) service-principal authentication. The `tenant_id` and
// `client_id` identify the service principal, and `client_secret` authenticates
// it. AI Gateway exchanges these credentials for an Entra bearer token for
// requests to Azure OpenAI or Microsoft Foundry.
type ModelProviderServiceConfigEntraServicePrincipal struct {
	// Entra ID client (application) ID. Required on Create.
	ClientId string `json:"client_id,omitempty"`
	// Entra ID client secret. Supply the value in `client_secret.plaintext`.
	ClientSecret *ModelProviderServiceConfigProviderSecret `json:"client_secret,omitempty"`
	// Entra ID (Azure AD) tenant ID. Required on Create.
	TenantId string `json:"tenant_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfigEntraServicePrincipal) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfigEntraServicePrincipal) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// External LLM provider for an EXTERNAL_FOUNDATION_MODEL destination.
type ModelProviderServiceConfigExternalModelProviderType string

const ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeAmazonBedrock ModelProviderServiceConfigExternalModelProviderType = `EXTERNAL_MODEL_PROVIDER_TYPE_AMAZON_BEDROCK`

const ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeAnthropic ModelProviderServiceConfigExternalModelProviderType = `EXTERNAL_MODEL_PROVIDER_TYPE_ANTHROPIC`

const ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeAzureOpenai ModelProviderServiceConfigExternalModelProviderType = `EXTERNAL_MODEL_PROVIDER_TYPE_AZURE_OPENAI`

const ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeBedrockMantle ModelProviderServiceConfigExternalModelProviderType = `EXTERNAL_MODEL_PROVIDER_TYPE_BEDROCK_MANTLE`

const ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeCustom ModelProviderServiceConfigExternalModelProviderType = `EXTERNAL_MODEL_PROVIDER_TYPE_CUSTOM`

const ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeGeminiEnterprise ModelProviderServiceConfigExternalModelProviderType = `EXTERNAL_MODEL_PROVIDER_TYPE_GEMINI_ENTERPRISE`

const ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeMicrosoftFoundry ModelProviderServiceConfigExternalModelProviderType = `EXTERNAL_MODEL_PROVIDER_TYPE_MICROSOFT_FOUNDRY`

const ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeOpenai ModelProviderServiceConfigExternalModelProviderType = `EXTERNAL_MODEL_PROVIDER_TYPE_OPENAI`

// String representation for [fmt.Print]
func (f *ModelProviderServiceConfigExternalModelProviderType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ModelProviderServiceConfigExternalModelProviderType) Set(v string) error {
	switch v {
	case `EXTERNAL_MODEL_PROVIDER_TYPE_AMAZON_BEDROCK`, `EXTERNAL_MODEL_PROVIDER_TYPE_ANTHROPIC`, `EXTERNAL_MODEL_PROVIDER_TYPE_AZURE_OPENAI`, `EXTERNAL_MODEL_PROVIDER_TYPE_BEDROCK_MANTLE`, `EXTERNAL_MODEL_PROVIDER_TYPE_CUSTOM`, `EXTERNAL_MODEL_PROVIDER_TYPE_GEMINI_ENTERPRISE`, `EXTERNAL_MODEL_PROVIDER_TYPE_MICROSOFT_FOUNDRY`, `EXTERNAL_MODEL_PROVIDER_TYPE_OPENAI`:
		*f = ModelProviderServiceConfigExternalModelProviderType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL_MODEL_PROVIDER_TYPE_AMAZON_BEDROCK", "EXTERNAL_MODEL_PROVIDER_TYPE_ANTHROPIC", "EXTERNAL_MODEL_PROVIDER_TYPE_AZURE_OPENAI", "EXTERNAL_MODEL_PROVIDER_TYPE_BEDROCK_MANTLE", "EXTERNAL_MODEL_PROVIDER_TYPE_CUSTOM", "EXTERNAL_MODEL_PROVIDER_TYPE_GEMINI_ENTERPRISE", "EXTERNAL_MODEL_PROVIDER_TYPE_MICROSOFT_FOUNDRY", "EXTERNAL_MODEL_PROVIDER_TYPE_OPENAI"`, v)
	}
}

// Values returns all possible values for ModelProviderServiceConfigExternalModelProviderType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ModelProviderServiceConfigExternalModelProviderType) Values() []ModelProviderServiceConfigExternalModelProviderType {
	return []ModelProviderServiceConfigExternalModelProviderType{
		ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeAmazonBedrock,
		ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeAnthropic,
		ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeAzureOpenai,
		ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeBedrockMantle,
		ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeCustom,
		ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeGeminiEnterprise,
		ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeMicrosoftFoundry,
		ModelProviderServiceConfigExternalModelProviderTypeExternalModelProviderTypeOpenai,
	}
}

// Type always returns ModelProviderServiceConfigExternalModelProviderType to satisfy [pflag.Value] interface
func (f *ModelProviderServiceConfigExternalModelProviderType) Type() string {
	return "ModelProviderServiceConfigExternalModelProviderType"
}

// Gemini Enterprise provider configuration.
type ModelProviderServiceConfigGeminiEnterpriseProviderConfig struct {
	// Gemini Enterprise project, region, and authentication configuration.
	Direct *ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig `json:"direct,omitempty"`
}

func (s *ModelProviderServiceConfigGeminiEnterpriseProviderConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Direct Gemini Enterprise provider configuration. An API key is required when
// creating the service.
type ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig struct {
	// Google Gemini Enterprise API key. Required when creating the service.
	// Supply the value in `api_key.plaintext`.
	ApiKey *ModelProviderServiceConfigProviderSecret `json:"api_key,omitempty"`
	// GCP project ID hosting the Gemini Enterprise endpoint. Required on
	// Create.
	ProjectId string `json:"project_id,omitempty"`
	// GCP region of the Gemini Enterprise endpoint (e.g., `us-central1`).
	// Required on Create.
	Region string `json:"region,omitempty"`
	// Reference to a Unity Catalog service credential authorizing Gemini
	// Enterprise requests. On Create, supply `service_credential.name` as
	// `credentials/{name}`; required when using service-credential
	// authentication and mutually exclusive with `api_key`. The credential is
	// referenced by name; its value is not carried here. On read, the resolved
	// `id` and `is_deleted` are also populated. Supported only on GCP-hosted
	// workspaces.
	ServiceCredential *ModelProviderServiceConfigServiceCredential `json:"service_credential,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Microsoft Foundry provider configuration.
type ModelProviderServiceConfigMicrosoftFoundryProviderConfig struct {
	// Microsoft Foundry endpoint and authentication configuration.
	Direct *ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig `json:"direct,omitempty"`
}

func (s *ModelProviderServiceConfigMicrosoftFoundryProviderConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Direct form of Microsoft Foundry provider config.
//
// Authentication is one of three mutually exclusive modes, exactly one of which
// must be supplied on Create: - API key: set `api_key`, leave
// `entra_service_principal` and `service_credential` unset. - Entra ID (service
// principal): set `entra_service_principal`, leave `api_key` and
// `service_credential` unset. AI Gateway exchanges these for an Entra bearer
// token on outbound requests via the OAuth2 client-credentials grant. - Unity
// Catalog service credential: set `service_credential.name` to the resource
// name `credentials/{name}`, leave `api_key` and `entra_service_principal`
// unset. The credential value lives in Unity Catalog and is referenced by name,
// not held on this message. Only supported on Azure-hosted workspaces. Setting
// more than one mode is rejected.
type ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig struct {
	// Microsoft Foundry API key. Supply the value in `api_key.plaintext`.
	// Mutually exclusive with Entra ID and Unity Catalog service credential
	// authentication.
	ApiKey *ModelProviderServiceConfigProviderSecret `json:"api_key,omitempty"`
	// Microsoft Foundry endpoint URL. Required on Create.
	BaseUrl string `json:"base_url,omitempty"`
	// Entra ID service-principal authentication. Set `tenant_id`, `client_id`,
	// and `client_secret.plaintext`. Mutually exclusive with `api_key` and
	// `service_credential`.
	EntraServicePrincipal *ModelProviderServiceConfigEntraServicePrincipal `json:"entra_service_principal,omitempty"`
	// Reference to a Unity Catalog service credential authorizing Microsoft
	// Foundry requests. On Create, supply `service_credential.name` in the form
	// `credentials/{name}`. Required on Create when using service-credential
	// authentication; mutually exclusive with `api_key` and
	// `entra_service_principal`. The credential is referenced by name; its
	// value is not carried here. Only supported on Azure-hosted workspaces.
	ServiceCredential *ModelProviderServiceConfigServiceCredential `json:"service_credential,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Model target configuration shared by model provider service targets and
// external model destinations on model services.
type ModelProviderServiceConfigModelTargetConfig struct {
	// Relative path appended to a custom provider's `base_url` for this model,
	// such as `serve/openai/chat`. Nonempty values are only supported for
	// custom model provider service targets; other provider types reject them.
	// Model service destinations do not support this field. Do not include URI
	// templates, queries, or fragments. When empty, `base_url` is used
	// unchanged.
	EndpointRoute string `json:"endpoint_route,omitempty"`
	// Provider-side model identifier, such as `gpt-5` or `claude-opus-4-7`.
	// This identifies a model at the upstream provider; it is not a Unity
	// Catalog model resource.
	Model string `json:"model"`
	// Provider-native API types supported by this model, such as
	// `openai/v1/chat/completions`. A model provider service target requires at
	// least one value. A model service destination may omit the list to use the
	// matching provider service target's types when available. AI Gateway uses
	// the selected types for translation. At most 64 entries of 256 characters
	// each are allowed.
	NativeApiTypes []string `json:"native_api_types,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfigModelTargetConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfigModelTargetConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// OpenAI provider configuration.
type ModelProviderServiceConfigOpenAiProviderConfig struct {
	// OpenAI configuration with an API key supplied in the request.
	Direct *ModelProviderServiceConfigOpenAiProviderDirectConfig `json:"direct,omitempty"`
}

func (s *ModelProviderServiceConfigOpenAiProviderConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Direct (inline-credentials) form of the OpenAI provider config.
type ModelProviderServiceConfigOpenAiProviderDirectConfig struct {
	// OpenAI API key. Required when creating the service. Supply the value in
	// `api_key.plaintext`.
	ApiKey *ModelProviderServiceConfigProviderSecret `json:"api_key,omitempty"`
	// Optional custom base URL. Defaults to `https://api.openai.com/v1`. Use
	// for OpenAI-API-compatible third-party endpoints or in-network proxies.
	BaseUrl string `json:"base_url,omitempty"`
	// Optional OpenAI organization ID. When set, the platform forwards it as
	// the `OpenAI-Organization` header.
	Organization string `json:"organization,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfigOpenAiProviderDirectConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfigOpenAiProviderDirectConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Pricing adjustments applied to this provider service's external-model spend
// estimates.
type ModelProviderServiceConfigProviderPricingConfig struct {
	// Provider-wide adjustment to external-model spend estimates in basis
	// points: 2000 = 20% off; -2000 = a 20% markup. Negative values have no
	// lower bound beyond the int32 range. Values cannot exceed 10000 (100%
	// off). An omitted value or zero leaves the estimate unchanged.
	DefaultDiscountBasisPoints int `json:"default_discount_basis_points,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfigProviderPricingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfigProviderPricingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A secret value supplied as part of an inline provider config. The caller
// supplies the value as inline `plaintext` on writes; the platform stores it
// encrypted. The `plaintext` field is `INPUT_ONLY` and never round-trips on
// reads.
type ModelProviderServiceConfigProviderSecret struct {
	// Inline plaintext credential. The value is write-only and never
	// round-trips on reads. Get and List responses omit `plaintext`; the
	// enclosing secret object remains present to indicate that a secret is
	// configured.
	Plaintext string `json:"plaintext,omitempty"`
	// Reference to a customer-owned UC Secret that carries this secret value.
	// The value is read at invoke time under the model provider service owner's
	// access and is never copied onto the model provider service, so rotating
	// the UC Secret takes effect with no change to the model provider service.
	// On Create, supply `secret_reference.name` as
	// `secrets/{catalog}.{schema}.{secret}`.
	SecretReference *ModelProviderServiceConfigSecretReference `json:"secret_reference,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelProviderServiceConfigProviderSecret) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelProviderServiceConfigProviderSecret) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Reference to a customer-owned UC Secret backing a secret-bearing provider
// field, in the `ProviderSecret.secret_reference` arm.
type ModelProviderServiceConfigSecretReference struct {
	// Resource name of the bound UC Secret, in the form
	// `secrets/{catalog}.{schema}.{secret}`. On Create the caller supplies the
	// name here. On read it reflects the secret's current name at read time.
	Name string `json:"name"`
}

func (s *ModelProviderServiceConfigSecretReference) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The customer-owned Unity Catalog service credential a model provider service
// uses to authenticate to its provider, referenced by name.
type ModelProviderServiceConfigServiceCredential struct {
	// Resource name of the bound Unity Catalog service credential, in the form
	// `credentials/{name}`. Supply this field when creating the service or
	// rebinding its credential. On read, it reflects the credential's current
	// name.
	Name string `json:"name"`
}

func (s *ModelProviderServiceConfigServiceCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A governed AI Gateway endpoint in Unity Catalog that routes inference
// requests to one or more destinations, such as a Databricks foundation model
// or an external model reached through a model provider service. Applies
// centralized access control, rate limits, and auditing to its traffic.
type ModelService struct {
	// User-provided description.
	Comment string `json:"comment,omitempty"`
	// Destinations, routing, rate limits, and payload logging configuration.
	// Required on Create. On Update, provide this field when `update_mask`
	// contains `config` or one of its subpaths.
	Config *ModelServiceConfig `json:"config,omitempty"`
	// Time the model service was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Username of the user who created the model service.
	CreatedBy string `json:"created_by,omitempty"`
	// Username or group name of the model service owner.
	EffectiveOwner string `json:"effective_owner,omitempty"`
	// Optimistic concurrency token returned on every read. To make an Update or
	// Delete conditional, pass the last-read value in that request's `etag`
	// field. In REST responses, this value is a base64 string; URL-encode it
	// when setting the `etag` query parameter.
	Etag string `json:"etag,omitempty"`
	// Unique identifier of the metastore that contains the model service.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Resource name of the model service. Format:
	// `model-services/{catalog}.{schema}.{model_service}`. Each `{...}`
	// component is capped at 255 characters individually. The server derives
	// the name on Create from `parent` and `model_service_id`; it is required
	// and immutable on Update, Get, and Delete.
	Name string `json:"name,omitempty"`
	// API types supported across this service's destinations, such as
	// `openai/v1/chat/completions`, `openai/v1/embeddings`, and
	// `mlflow/v1/chat/completions`. Derived from the backing models and
	// providers at read time.
	SupportedApiTypes []string `json:"supported_api_types,omitempty"`
	// Time the model service was last modified.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Username of the user who last updated the model service.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelService) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelService) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Operational configuration wrapped around the ModelService resource.
type ModelServiceConfig struct {
	// Configuration for logging request and response payloads to a Unity
	// Catalog inference table. Omit this field to leave payload logging off.
	InferenceTable *InferenceTableConfig `json:"inference_table,omitempty"`
	// Rate limits applied to requests routed through this model service.
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
	// Routing configuration: the primary destinations that receive traffic and
	// the fallback destinations tried when a primary destination fails.
	Routing *ModelServiceConfigRoutingConfig `json:"routing,omitempty"`
}

func (s *ModelServiceConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A destination the model service can route traffic to. Exactly one of the
// per-type configs inside `type_config` must be set, and it must match
// `destination_type`.
type ModelServiceConfigDestinationConfig struct {
	// Backing-model category. Provide the matching type-specific configuration
	// and leave the other type-specific configurations unset.
	DestinationType ModelServiceConfigDestinationConfigDestinationType `json:"destination_type"`
	// Configuration for an external model reached through a model provider
	// service.
	ExternalModelConfig *ModelServiceConfigExternalModelConfig `json:"external_model_config,omitempty"`
	// Whether the destination's backing model or model provider service has
	// been deleted. The destination remains visible so you can identify the
	// broken dependency. Requests cannot use this destination until the backing
	// resource is restored or the destination is replaced.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// Label that identifies this destination. Names are compared
	// case-insensitively and must be unique within the primary destinations and
	// within the fallback destinations.
	Name string `json:"name"`
	// Configuration for a pay-per-token Databricks foundation model.
	PayPerTokenConfig *ModelServiceConfigPayPerTokenConfig `json:"pay_per_token_config,omitempty"`
	// Configuration for a provisioned-throughput Databricks foundation model.
	ProvisionedThroughputConfig *ModelServiceConfigProvisionedThroughputConfig `json:"provisioned_throughput_config,omitempty"`
	// Percentage of primary traffic sent to this destination, from 0 to 100.
	// Required when there is more than one primary destination, in which case
	// the primary percentages must sum to 100; a single primary destination
	// receives all traffic. Fallback destinations are ordered and do not use
	// this field.
	TrafficPercentage int `json:"traffic_percentage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelServiceConfigDestinationConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelServiceConfigDestinationConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Backing-model category for a model service destination.
type ModelServiceConfigDestinationConfigDestinationType string

const ModelServiceConfigDestinationConfigDestinationTypeDestinationTypeExternalFoundationModel ModelServiceConfigDestinationConfigDestinationType = `DESTINATION_TYPE_EXTERNAL_FOUNDATION_MODEL`

const ModelServiceConfigDestinationConfigDestinationTypeDestinationTypePayPerTokenFoundationModel ModelServiceConfigDestinationConfigDestinationType = `DESTINATION_TYPE_PAY_PER_TOKEN_FOUNDATION_MODEL`

const ModelServiceConfigDestinationConfigDestinationTypeDestinationTypeProvisionedThroughputFoundationModel ModelServiceConfigDestinationConfigDestinationType = `DESTINATION_TYPE_PROVISIONED_THROUGHPUT_FOUNDATION_MODEL`

// String representation for [fmt.Print]
func (f *ModelServiceConfigDestinationConfigDestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ModelServiceConfigDestinationConfigDestinationType) Set(v string) error {
	switch v {
	case `DESTINATION_TYPE_EXTERNAL_FOUNDATION_MODEL`, `DESTINATION_TYPE_PAY_PER_TOKEN_FOUNDATION_MODEL`, `DESTINATION_TYPE_PROVISIONED_THROUGHPUT_FOUNDATION_MODEL`:
		*f = ModelServiceConfigDestinationConfigDestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DESTINATION_TYPE_EXTERNAL_FOUNDATION_MODEL", "DESTINATION_TYPE_PAY_PER_TOKEN_FOUNDATION_MODEL", "DESTINATION_TYPE_PROVISIONED_THROUGHPUT_FOUNDATION_MODEL"`, v)
	}
}

// Values returns all possible values for ModelServiceConfigDestinationConfigDestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ModelServiceConfigDestinationConfigDestinationType) Values() []ModelServiceConfigDestinationConfigDestinationType {
	return []ModelServiceConfigDestinationConfigDestinationType{
		ModelServiceConfigDestinationConfigDestinationTypeDestinationTypeExternalFoundationModel,
		ModelServiceConfigDestinationConfigDestinationTypeDestinationTypePayPerTokenFoundationModel,
		ModelServiceConfigDestinationConfigDestinationTypeDestinationTypeProvisionedThroughputFoundationModel,
	}
}

// Type always returns ModelServiceConfigDestinationConfigDestinationType to satisfy [pflag.Value] interface
func (f *ModelServiceConfigDestinationConfigDestinationType) Type() string {
	return "ModelServiceConfigDestinationConfigDestinationType"
}

// Configuration for an external-foundation-model destination. Provider auth and
// provider-specific cloud configuration are owned by a separate, governed
// ModelProviderService entity referenced via `model_provider_service`; the
// platform resolves the provider at invocation time.
type ModelServiceConfigExternalModelConfig struct {
	// Resource name of the governed ModelProviderService that owns provider
	// auth and provider-specific configuration. Format:
	// `model-provider-services/{catalog}.{schema}.{model_provider_service}`.
	// Each `{...}` component is capped at 255 characters individually.
	ModelProviderService string `json:"model_provider_service"`
	// Routing target for the destination: the provider-side model selected from
	// the referenced ModelProviderService's `targets` catalog, plus the unified
	// API types the platform should translate to/from at request time.
	Target ModelProviderServiceConfigModelTargetConfig `json:"target"`
}

func (s *ModelServiceConfigExternalModelConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Fallback routing applied after a primary destination fails. Fallback
// destinations are tried in the listed order.
type ModelServiceConfigFallbackConfig struct {
	// Fallback destinations, tried in the listed order. At most 5 are allowed.
	Destinations []ModelServiceConfigDestinationConfig `json:"destinations,omitempty"`
}

func (s *ModelServiceConfigFallbackConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Configuration for a pay-per-token foundation-model destination. Identifies
// the foundation model by its UC resource name; the platform resolves it to a
// Model Serving endpoint at request time.
type ModelServiceConfigPayPerTokenConfig struct {
	// Resource name of the Unity Catalog foundation model. Pay-per-token
	// foundation models live in the `system.ai` schema, so the name has the
	// form `models/system.ai.{model}`.
	Model string `json:"model"`
}

func (s *ModelServiceConfigPayPerTokenConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Configuration for a provisioned-throughput foundation-model destination.
// References a pre-existing Model Serving endpoint that serves the model;
// sizing (provisioned throughput, burst scaling, model version) is owned by the
// Model Serving endpoint itself, not by this message.
type ModelServiceConfigProvisionedThroughputConfig struct {
	// Resource name of the Unity Catalog model served by the backing endpoint,
	// in the form `models/{catalog}.{schema}.{model}`. Resolved from Model
	// Serving when the model service is created or updated.
	Model string `json:"model,omitempty"`
	// Name of the backing Model Serving endpoint serving the provisioned-
	// throughput foundation model, in the form `serving-endpoints/{name}`. The
	// same Unity Catalog model can be served on multiple Model Serving
	// endpoints with different throughput, regions, or configurations. The
	// caller selects the endpoint to which this destination routes. The
	// endpoint must exist at create time.
	ModelServingEndpoint string `json:"model_serving_endpoint"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelServiceConfigProvisionedThroughputConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelServiceConfigProvisionedThroughputConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Routing configuration for a model service, nesting destinations and fallback
// under a single sub-message.
type ModelServiceConfigRoutingConfig struct {
	// Primary routing destinations. At most 10 are allowed. At least one is
	// required on Create. On Update, provide this list when replacing the full
	// `config` or updating `config.routing.destinations`; other granular
	// routing updates do not require resending destinations. The intermediate
	// `config.routing` mask path is not supported.
	Destinations []ModelServiceConfigDestinationConfig `json:"destinations,omitempty"`
	// Fallback routing applied after a primary destination fails. Fallback
	// destinations are tried in the listed order.
	Fallback *ModelServiceConfigFallbackConfig `json:"fallback,omitempty"`
}

func (s *ModelServiceConfigRoutingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A numbered version of a registered model in Unity Catalog. A model version is
// identified by the full name of its parent registered model and its version
// number, or by an alias that points to it. See
// https://docs.databricks.com/aws/en/machine-learning/manage-model-lifecycle/.
type ModelVersionInfo struct {
	// List of aliases associated with the model version
	Aliases []RegisteredModelAlias `json:"aliases,omitempty"`
	// The name of the catalog containing the model version
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the model version
	Comment string `json:"comment,omitempty"`
	// Time at which the model version was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the model version
	CreatedBy string `json:"created_by,omitempty"`
	// The unique identifier of the model version
	Id string `json:"id,omitempty"`
	// The unique identifier of the metastore containing the model version
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the parent registered model of the model version, relative to
	// parent schema
	ModelName string `json:"model_name,omitempty"`
	// Model version dependencies, for feature-store packaged models
	ModelVersionDependencies *DependencyList `json:"model_version_dependencies,omitempty"`
	// MLflow run ID used when creating the model version, if ``source`` was
	// generated by an experiment run stored in an MLflow tracking server
	RunId string `json:"run_id,omitempty"`
	// ID of the Databricks workspace containing the MLflow run that generated
	// this model version, if applicable
	RunWorkspaceId int `json:"run_workspace_id,omitempty"`
	// The name of the schema containing the model version, relative to parent
	// catalog
	SchemaName string `json:"schema_name,omitempty"`
	// URI indicating the location of the source artifacts (files) for the model
	// version
	Source string `json:"source,omitempty"`
	// Current status of the model version. Newly created model versions start
	// in PENDING_REGISTRATION status, then move to READY status once the model
	// version files are uploaded and the model version is finalized. Only model
	// versions in READY status can be loaded for inference or served.
	Status ModelVersionInfoStatus `json:"status,omitempty"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `json:"storage_location,omitempty"`
	// Time at which the model version was last updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the model version last time
	UpdatedBy string `json:"updated_by,omitempty"`
	// Integer model version number, used to reference the model version in API
	// requests.
	Version int `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelVersionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelVersionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ModelVersionInfoStatus string

const ModelVersionInfoStatusFailedRegistration ModelVersionInfoStatus = `FAILED_REGISTRATION`

const ModelVersionInfoStatusModelVersionStatusUnknown ModelVersionInfoStatus = `MODEL_VERSION_STATUS_UNKNOWN`

const ModelVersionInfoStatusPendingRegistration ModelVersionInfoStatus = `PENDING_REGISTRATION`

const ModelVersionInfoStatusReady ModelVersionInfoStatus = `READY`

// String representation for [fmt.Print]
func (f *ModelVersionInfoStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ModelVersionInfoStatus) Set(v string) error {
	switch v {
	case `FAILED_REGISTRATION`, `MODEL_VERSION_STATUS_UNKNOWN`, `PENDING_REGISTRATION`, `READY`:
		*f = ModelVersionInfoStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_REGISTRATION", "MODEL_VERSION_STATUS_UNKNOWN", "PENDING_REGISTRATION", "READY"`, v)
	}
}

// Values returns all possible values for ModelVersionInfoStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *ModelVersionInfoStatus) Values() []ModelVersionInfoStatus {
	return []ModelVersionInfoStatus{
		ModelVersionInfoStatusFailedRegistration,
		ModelVersionInfoStatusModelVersionStatusUnknown,
		ModelVersionInfoStatusPendingRegistration,
		ModelVersionInfoStatusReady,
	}
}

// Type always returns ModelVersionInfoStatus to satisfy [pflag.Value] interface
func (f *ModelVersionInfoStatus) Type() string {
	return "ModelVersionInfoStatus"
}

type MonitorCronSchedule struct {
	// Read only field that indicates whether a schedule is paused or not.
	PauseStatus MonitorCronSchedulePauseStatus `json:"pause_status,omitempty"`
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// The timezone id (e.g., ``PST``) in which to evaluate the quartz
	// expression.
	TimezoneId string `json:"timezone_id"`
}

func (s *MonitorCronSchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Source link:
// https://src.dev.databricks.com/databricks/universe/-/blob/elastic-spark-common/api/messages/schedule.proto
// Monitoring workflow schedule pause status.
type MonitorCronSchedulePauseStatus string

const MonitorCronSchedulePauseStatusPaused MonitorCronSchedulePauseStatus = `PAUSED`

const MonitorCronSchedulePauseStatusUnpaused MonitorCronSchedulePauseStatus = `UNPAUSED`

const MonitorCronSchedulePauseStatusUnspecified MonitorCronSchedulePauseStatus = `UNSPECIFIED`

// String representation for [fmt.Print]
func (f *MonitorCronSchedulePauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorCronSchedulePauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`, `UNSPECIFIED`:
		*f = MonitorCronSchedulePauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED", "UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for MonitorCronSchedulePauseStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorCronSchedulePauseStatus) Values() []MonitorCronSchedulePauseStatus {
	return []MonitorCronSchedulePauseStatus{
		MonitorCronSchedulePauseStatusPaused,
		MonitorCronSchedulePauseStatusUnpaused,
		MonitorCronSchedulePauseStatusUnspecified,
	}
}

// Type always returns MonitorCronSchedulePauseStatus to satisfy [pflag.Value] interface
func (f *MonitorCronSchedulePauseStatus) Type() string {
	return "MonitorCronSchedulePauseStatus"
}

// Data classification related configuration.
type MonitorDataClassificationConfig struct {
	// Whether to enable data classification.
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	EmailAddresses []string `json:"email_addresses,omitempty"`
}

func (s *MonitorDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type MonitorInferenceLog struct {
	// List of granularities to use when aggregating data into time windows
	// based on their timestamp.
	Granularities []string `json:"granularities"`
	// Column for the label.
	LabelCol string `json:"label_col,omitempty"`
	// Column for the model identifier.
	ModelIdCol string `json:"model_id_col"`
	// Column for the prediction.
	PredictionCol string `json:"prediction_col"`
	// Column for prediction probabilities
	PredictionProbaCol string `json:"prediction_proba_col,omitempty"`
	// Problem type the model aims to solve.
	ProblemType MonitorInferenceLogProblemType `json:"problem_type"`
	// Column for the timestamp.
	TimestampCol string `json:"timestamp_col"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MonitorInferenceLog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorInferenceLog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Values returns all possible values for MonitorInferenceLogProblemType.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorInferenceLogProblemType) Values() []MonitorInferenceLogProblemType {
	return []MonitorInferenceLogProblemType{
		MonitorInferenceLogProblemTypeProblemTypeClassification,
		MonitorInferenceLogProblemTypeProblemTypeRegression,
	}
}

// Type always returns MonitorInferenceLogProblemType to satisfy [pflag.Value] interface
func (f *MonitorInferenceLogProblemType) Type() string {
	return "MonitorInferenceLogProblemType"
}

type MonitorInfo struct {
	// [Create:REQ Update:IGN] Field for specifying the absolute path to a
	// custom directory to store data-monitoring assets. Normally prepopulated
	// to a default user location via UI and Python APIs.
	AssetsDir string `json:"assets_dir,omitempty"`
	// [Create:OPT Update:OPT] Baseline table name. Baseline data is used to
	// compute drift from the data in the monitored `table_name`. The baseline
	// table and the monitored table shall have the same schema.
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// [Create:OPT Update:OPT] Custom metrics.
	CustomMetrics []MonitorMetric `json:"custom_metrics,omitempty"`
	// [Create:ERR Update:OPT] Id of dashboard that visualizes the computed
	// metrics. This can be empty if the monitor is in PENDING state.
	DashboardId string `json:"dashboard_id,omitempty"`
	// [Create:OPT Update:OPT] Data classification related config.
	DataClassificationConfig *MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`
	// [Create:ERR Update:IGN] Table that stores drift metrics data. Format:
	// `catalog.schema.table_name`.
	DriftMetricsTableName string `json:"drift_metrics_table_name"`

	InferenceLog *MonitorInferenceLog `json:"inference_log,omitempty"`
	// [Create:ERR Update:IGN] The latest error message for a monitor failure.
	LatestMonitorFailureMsg string `json:"latest_monitor_failure_msg,omitempty"`
	// [Create:ERR Update:IGN] Represents the current monitor configuration
	// version in use. The version will be represented in a numeric fashion
	// (1,2,3...). The field has flexibility to take on negative values, which
	// can indicate corrupted monitor_version numbers.
	MonitorVersion int64 `json:"monitor_version"`
	// [Create:OPT Update:OPT] Field for specifying notification settings.
	Notifications *MonitorNotifications `json:"notifications,omitempty"`
	// [Create:REQ Update:REQ] Schema where output tables are created. Needs to
	// be in 2-level format {catalog}.{schema}
	OutputSchemaName string `json:"output_schema_name"`
	// [Create:ERR Update:IGN] Table that stores profile metrics data. Format:
	// `catalog.schema.table_name`.
	ProfileMetricsTableName string `json:"profile_metrics_table_name"`
	// [Create:OPT Update:OPT] The monitor schedule.
	Schedule *MonitorCronSchedule `json:"schedule,omitempty"`
	// [Create:OPT Update:OPT] List of column expressions to slice data with for
	// targeted analysis. The data is grouped by each expression independently,
	// resulting in a separate slice for each predicate and its complements. For
	// example `slicing_exprs=[“col_1”, “col_2 > 10”]` will generate the
	// following slices: two slices for `col_2 > 10` (True and False), and one
	// slice per unique value in `col1`. For high-cardinality columns, only the
	// top 100 unique values by frequency will generate slices.
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	Snapshot *MonitorSnapshot `json:"snapshot,omitempty"`
	// [Create:ERR Update:IGN] The monitor status.
	Status MonitorInfoStatus `json:"status"`
	// [Create:ERR Update:IGN] UC table to monitor. Format:
	// `catalog.schema.table_name`
	TableName string `json:"table_name"`
	// Configuration for monitoring time series tables.
	TimeSeries *MonitorTimeSeries `json:"time_series,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MonitorInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Values returns all possible values for MonitorInfoStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorInfoStatus) Values() []MonitorInfoStatus {
	return []MonitorInfoStatus{
		MonitorInfoStatusMonitorStatusActive,
		MonitorInfoStatusMonitorStatusDeletePending,
		MonitorInfoStatusMonitorStatusError,
		MonitorInfoStatusMonitorStatusFailed,
		MonitorInfoStatusMonitorStatusPending,
	}
}

// Type always returns MonitorInfoStatus to satisfy [pflag.Value] interface
func (f *MonitorInfoStatus) Type() string {
	return "MonitorInfoStatus"
}

// Custom metric definition.
type MonitorMetric struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	Definition string `json:"definition"`
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	InputColumns []string `json:"input_columns"`
	// Name of the metric in the output tables.
	Name string `json:"name"`
	// The output type of the custom metric.
	OutputDataType string `json:"output_data_type"`
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
	Type MonitorMetricType `json:"type"`
}

func (s *MonitorMetric) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Can only be one of “\"CUSTOM_METRIC_TYPE_AGGREGATE\"“,
// “\"CUSTOM_METRIC_TYPE_DERIVED\"“, or “\"CUSTOM_METRIC_TYPE_DRIFT\"“. The
// “\"CUSTOM_METRIC_TYPE_AGGREGATE\"“ and “\"CUSTOM_METRIC_TYPE_DERIVED\"“
// metrics are computed on a single table, whereas the
// “\"CUSTOM_METRIC_TYPE_DRIFT\"“ compare metrics across baseline and input
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

// Values returns all possible values for MonitorMetricType.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorMetricType) Values() []MonitorMetricType {
	return []MonitorMetricType{
		MonitorMetricTypeCustomMetricTypeAggregate,
		MonitorMetricTypeCustomMetricTypeDerived,
		MonitorMetricTypeCustomMetricTypeDrift,
	}
}

// Type always returns MonitorMetricType to satisfy [pflag.Value] interface
func (f *MonitorMetricType) Type() string {
	return "MonitorMetricType"
}

type MonitorNotifications struct {
	// Destinations to send notifications on failure/timeout.
	OnFailure *MonitorDestination `json:"on_failure,omitempty"`
	// Destinations to send notifications on new classification tag detected.
	OnNewClassificationTagDetected *MonitorDestination `json:"on_new_classification_tag_detected,omitempty"`
}

func (s *MonitorNotifications) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type MonitorRefreshInfo struct {
	// Time at which refresh operation completed (milliseconds since 1/1/1970
	// UTC).
	EndTimeMs int64 `json:"end_time_ms,omitempty"`
	// An optional message to give insight into the current state of the job
	// (e.g. FAILURE messages).
	Message string `json:"message,omitempty"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"refresh_id"`
	// Time at which refresh operation was initiated (milliseconds since
	// 1/1/1970 UTC).
	StartTimeMs int64 `json:"start_time_ms"`
	// The current state of the refresh.
	State MonitorRefreshInfoState `json:"state"`
	// The method by which the refresh was triggered.
	Trigger MonitorRefreshInfoTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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

const MonitorRefreshInfoStateUnknown MonitorRefreshInfoState = `UNKNOWN`

// String representation for [fmt.Print]
func (f *MonitorRefreshInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorRefreshInfoState) Set(v string) error {
	switch v {
	case `CANCELED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCESS`, `UNKNOWN`:
		*f = MonitorRefreshInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "FAILED", "PENDING", "RUNNING", "SUCCESS", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for MonitorRefreshInfoState.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorRefreshInfoState) Values() []MonitorRefreshInfoState {
	return []MonitorRefreshInfoState{
		MonitorRefreshInfoStateCanceled,
		MonitorRefreshInfoStateFailed,
		MonitorRefreshInfoStatePending,
		MonitorRefreshInfoStateRunning,
		MonitorRefreshInfoStateSuccess,
		MonitorRefreshInfoStateUnknown,
	}
}

// Type always returns MonitorRefreshInfoState to satisfy [pflag.Value] interface
func (f *MonitorRefreshInfoState) Type() string {
	return "MonitorRefreshInfoState"
}

type MonitorRefreshInfoTrigger string

const MonitorRefreshInfoTriggerManual MonitorRefreshInfoTrigger = `MANUAL`

const MonitorRefreshInfoTriggerSchedule MonitorRefreshInfoTrigger = `SCHEDULE`

const MonitorRefreshInfoTriggerUnknownTrigger MonitorRefreshInfoTrigger = `UNKNOWN_TRIGGER`

// String representation for [fmt.Print]
func (f *MonitorRefreshInfoTrigger) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorRefreshInfoTrigger) Set(v string) error {
	switch v {
	case `MANUAL`, `SCHEDULE`, `UNKNOWN_TRIGGER`:
		*f = MonitorRefreshInfoTrigger(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANUAL", "SCHEDULE", "UNKNOWN_TRIGGER"`, v)
	}
}

// Values returns all possible values for MonitorRefreshInfoTrigger.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorRefreshInfoTrigger) Values() []MonitorRefreshInfoTrigger {
	return []MonitorRefreshInfoTrigger{
		MonitorRefreshInfoTriggerManual,
		MonitorRefreshInfoTriggerSchedule,
		MonitorRefreshInfoTriggerUnknownTrigger,
	}
}

// Type always returns MonitorRefreshInfoTrigger to satisfy [pflag.Value] interface
func (f *MonitorRefreshInfoTrigger) Type() string {
	return "MonitorRefreshInfoTrigger"
}

type MonitorRefreshListResponse struct {
	// List of refreshes.
	Refreshes []MonitorRefreshInfo `json:"refreshes,omitempty"`
}

func (s *MonitorRefreshListResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Snapshot analysis configuration
type MonitorSnapshot struct {
}

func (s *MonitorSnapshot) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Time series analysis configuration.
type MonitorTimeSeries struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``\"5 minutes\"``, ``\"30 minutes\"``, ``\"1 hour\"``, ``\"1 day\"``,
	// ``\"\u003cn\u003e week(s)\"``, ``\"1 month\"``, ``\"1 year\"``}.
	Granularities []string `json:"granularities"`
	// Column for the timestamp.
	TimestampCol string `json:"timestamp_col"`
}

func (s *MonitorTimeSeries) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type NamedTableConstraint struct {
	// The name of the constraint.
	Name string `json:"name"`
}

func (s *NamedTableConstraint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type NotificationDestination struct {
	// The identifier for the destination. This is the email address for EMAIL
	// destinations, the URL for URL destinations, or the unique Databricks
	// notification destination ID for all other external destinations.
	DestinationId string `json:"destination_id,omitempty"`
	// The type of the destination.
	DestinationType DestinationType `json:"destination_type,omitempty"`
	// This field is used to denote whether the destination is the email of the
	// owner of the securable object. The special destination cannot be assigned
	// to a securable and only represents the default destination of the
	// securable. The securable types that support default special destinations
	// are: "catalog", "external_location", "connection", "credential", and
	// "metastore". The **destination_type** of a **special_destination** is
	// always EMAIL.
	SpecialDestination SpecialDestination `json:"special_destination,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NotificationDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NotificationDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Online Table information.
type OnlineTable struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"name,omitempty"`
	// Specification of the online table.
	Spec *OnlineTableSpec `json:"spec,omitempty"`
	// Online Table data synchronization status
	Status *OnlineTableStatus `json:"status,omitempty"`
	// Data serving REST API URL for this table
	TableServingUrl string `json:"table_serving_url,omitempty"`
	// The provisioning state of the online table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState ProvisioningInfoState `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	PerformFullCopy bool `json:"perform_full_copy,omitempty"`
	// ID of the associated pipeline. Generated by the server - cannot be set by
	// the caller.
	PipelineId string `json:"pipeline_id,omitempty"`
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns []string `json:"primary_key_columns,omitempty"`
	// Pipeline runs continuously after generating the initial data.
	RunContinuously *OnlineTableSpecContinuousSchedulingPolicy `json:"run_continuously,omitempty"`
	// Pipeline stops after generating the initial data and can be triggered
	// later (manually, through a cron job or through data triggers)
	RunTriggered *OnlineTableSpecTriggeredSchedulingPolicy `json:"run_triggered,omitempty"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName string `json:"source_table_full_name,omitempty"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey string `json:"timeseries_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OnlineTableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineTableSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type OnlineTableSpecContinuousSchedulingPolicy struct {
}

func (s *OnlineTableSpecContinuousSchedulingPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type OnlineTableSpecTriggeredSchedulingPolicy struct {
}

func (s *OnlineTableSpecTriggeredSchedulingPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The state of an online table.
type OnlineTableState string

const OnlineTableStateOffline OnlineTableState = `OFFLINE`

const OnlineTableStateOfflineFailed OnlineTableState = `OFFLINE_FAILED`

const OnlineTableStateOnline OnlineTableState = `ONLINE`

const OnlineTableStateOnlineContinuousUpdate OnlineTableState = `ONLINE_CONTINUOUS_UPDATE`

const OnlineTableStateOnlineNoPendingUpdate OnlineTableState = `ONLINE_NO_PENDING_UPDATE`

const OnlineTableStateOnlinePipelineFailed OnlineTableState = `ONLINE_PIPELINE_FAILED`

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
	case `OFFLINE`, `OFFLINE_FAILED`, `ONLINE`, `ONLINE_CONTINUOUS_UPDATE`, `ONLINE_NO_PENDING_UPDATE`, `ONLINE_PIPELINE_FAILED`, `ONLINE_TRIGGERED_UPDATE`, `ONLINE_UPDATING_PIPELINE_RESOURCES`, `PROVISIONING`, `PROVISIONING_INITIAL_SNAPSHOT`, `PROVISIONING_PIPELINE_RESOURCES`:
		*f = OnlineTableState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OFFLINE", "OFFLINE_FAILED", "ONLINE", "ONLINE_CONTINUOUS_UPDATE", "ONLINE_NO_PENDING_UPDATE", "ONLINE_PIPELINE_FAILED", "ONLINE_TRIGGERED_UPDATE", "ONLINE_UPDATING_PIPELINE_RESOURCES", "PROVISIONING", "PROVISIONING_INITIAL_SNAPSHOT", "PROVISIONING_PIPELINE_RESOURCES"`, v)
	}
}

// Values returns all possible values for OnlineTableState.
//
// There is no guarantee on the order of the values in the slice.
func (f *OnlineTableState) Values() []OnlineTableState {
	return []OnlineTableState{
		OnlineTableStateOffline,
		OnlineTableStateOfflineFailed,
		OnlineTableStateOnline,
		OnlineTableStateOnlineContinuousUpdate,
		OnlineTableStateOnlineNoPendingUpdate,
		OnlineTableStateOnlinePipelineFailed,
		OnlineTableStateOnlineTriggeredUpdate,
		OnlineTableStateOnlineUpdatingPipelineResources,
		OnlineTableStateProvisioning,
		OnlineTableStateProvisioningInitialSnapshot,
		OnlineTableStateProvisioningPipelineResources,
	}
}

// Type always returns OnlineTableState to satisfy [pflag.Value] interface
func (f *OnlineTableState) Type() string {
	return "OnlineTableState"
}

// Status of an online table.
type OnlineTableStatus struct {
	ContinuousUpdateStatus *ContinuousUpdateStatus `json:"continuous_update_status,omitempty"`
	// The state of the online table.
	DetailedState OnlineTableState `json:"detailed_state,omitempty"`

	FailedStatus *FailedStatus `json:"failed_status,omitempty"`
	// A text description of the current state of the online table.
	Message string `json:"message,omitempty"`

	ProvisioningStatus *ProvisioningStatus `json:"provisioning_status,omitempty"`

	TriggeredUpdateStatus *TriggeredUpdateStatus `json:"triggered_update_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OnlineTableStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineTableStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Spec of an allowed option on a securable kind and its attributes. This is
// mostly used by UI to provide user friendly hints and descriptions in order to
// facilitate the securable creation process.
type OptionSpec struct {
	// For drop down / radio button selections, UI will want to know the
	// possible input values, it can also be used by other option types to limit
	// input selections.
	AllowedValues []string `json:"allowed_values,omitempty"`
	// The default value of the option, for example, value '443' for 'port'
	// option.
	DefaultValue string `json:"default_value,omitempty"`
	// A concise user facing description of what the input value of this option
	// should look like.
	Description string `json:"description,omitempty"`
	// The hint is used on the UI to suggest what the input value can possibly
	// be like, for example: example.com for 'host' option. Unlike default
	// value, it will not be applied automatically without user input.
	Hint string `json:"hint,omitempty"`
	// Indicates whether an option should be displayed with copy button on the
	// UI.
	IsCopiable bool `json:"is_copiable,omitempty"`
	// Indicates whether an option can be provided by users in the create/update
	// path of an entity.
	IsCreatable bool `json:"is_creatable,omitempty"`
	// Is the option value not user settable and is thus not shown on the UI.
	IsHidden bool `json:"is_hidden,omitempty"`
	// Specifies whether this option is safe to log, i.e. no sensitive
	// information.
	IsLoggable bool `json:"is_loggable,omitempty"`
	// Is the option required.
	IsRequired bool `json:"is_required,omitempty"`
	// Is the option value considered secret and thus redacted on the UI.
	IsSecret bool `json:"is_secret,omitempty"`
	// Is the option updatable by users.
	IsUpdatable bool `json:"is_updatable,omitempty"`
	// The unique name of the option.
	Name string `json:"name,omitempty"`
	// Specifies when the option value is displayed on the UI within the OAuth
	// flow.
	OauthStage OptionSpecOauthStage `json:"oauth_stage,omitempty"`
	// The type of the option.
	Type OptionSpecOptionType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OptionSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OptionSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// During the OAuth flow, specifies which stage the option should be displayed
// in the UI. OAUTH_STAGE_UNSPECIFIED is the default value for options unrelated
// to the OAuth flow. BEFORE_AUTHORIZATION_CODE corresponds to options necessary
// to initiate the OAuth process. BEFORE_ACCESS_TOKEN corresponds to options
// that are necessary to create a foreign connection, but that should be
// displayed after the authorization code has already been received.
type OptionSpecOauthStage string

const OptionSpecOauthStageBeforeAccessToken OptionSpecOauthStage = `BEFORE_ACCESS_TOKEN`

const OptionSpecOauthStageBeforeAuthorizationCode OptionSpecOauthStage = `BEFORE_AUTHORIZATION_CODE`

// String representation for [fmt.Print]
func (f *OptionSpecOauthStage) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OptionSpecOauthStage) Set(v string) error {
	switch v {
	case `BEFORE_ACCESS_TOKEN`, `BEFORE_AUTHORIZATION_CODE`:
		*f = OptionSpecOauthStage(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BEFORE_ACCESS_TOKEN", "BEFORE_AUTHORIZATION_CODE"`, v)
	}
}

// Values returns all possible values for OptionSpecOauthStage.
//
// There is no guarantee on the order of the values in the slice.
func (f *OptionSpecOauthStage) Values() []OptionSpecOauthStage {
	return []OptionSpecOauthStage{
		OptionSpecOauthStageBeforeAccessToken,
		OptionSpecOauthStageBeforeAuthorizationCode,
	}
}

// Type always returns OptionSpecOauthStage to satisfy [pflag.Value] interface
func (f *OptionSpecOauthStage) Type() string {
	return "OptionSpecOauthStage"
}

// Type of the option, we purposely follow JavaScript types so that the UI can
// map the options to JS types. https://www.w3schools.com/js/js_datatypes.asp
// Enum is a special case that it's just string with selections.
type OptionSpecOptionType string

const OptionSpecOptionTypeOptionBigint OptionSpecOptionType = `OPTION_BIGINT`

const OptionSpecOptionTypeOptionBoolean OptionSpecOptionType = `OPTION_BOOLEAN`

const OptionSpecOptionTypeOptionEnum OptionSpecOptionType = `OPTION_ENUM`

const OptionSpecOptionTypeOptionMultilineString OptionSpecOptionType = `OPTION_MULTILINE_STRING`

const OptionSpecOptionTypeOptionNumber OptionSpecOptionType = `OPTION_NUMBER`

const OptionSpecOptionTypeOptionServiceCredential OptionSpecOptionType = `OPTION_SERVICE_CREDENTIAL`

const OptionSpecOptionTypeOptionString OptionSpecOptionType = `OPTION_STRING`

// String representation for [fmt.Print]
func (f *OptionSpecOptionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OptionSpecOptionType) Set(v string) error {
	switch v {
	case `OPTION_BIGINT`, `OPTION_BOOLEAN`, `OPTION_ENUM`, `OPTION_MULTILINE_STRING`, `OPTION_NUMBER`, `OPTION_SERVICE_CREDENTIAL`, `OPTION_STRING`:
		*f = OptionSpecOptionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OPTION_BIGINT", "OPTION_BOOLEAN", "OPTION_ENUM", "OPTION_MULTILINE_STRING", "OPTION_NUMBER", "OPTION_SERVICE_CREDENTIAL", "OPTION_STRING"`, v)
	}
}

// Values returns all possible values for OptionSpecOptionType.
//
// There is no guarantee on the order of the values in the slice.
func (f *OptionSpecOptionType) Values() []OptionSpecOptionType {
	return []OptionSpecOptionType{
		OptionSpecOptionTypeOptionBigint,
		OptionSpecOptionTypeOptionBoolean,
		OptionSpecOptionTypeOptionEnum,
		OptionSpecOptionTypeOptionMultilineString,
		OptionSpecOptionTypeOptionNumber,
		OptionSpecOptionTypeOptionServiceCredential,
		OptionSpecOptionTypeOptionString,
	}
}

// Type always returns OptionSpecOptionType to satisfy [pflag.Value] interface
func (f *OptionSpecOptionType) Type() string {
	return "OptionSpecOptionType"
}

type PathOperation string

const PathOperationPathCreateTable PathOperation = `PATH_CREATE_TABLE`

const PathOperationPathRead PathOperation = `PATH_READ`

const PathOperationPathReadWrite PathOperation = `PATH_READ_WRITE`

// String representation for [fmt.Print]
func (f *PathOperation) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PathOperation) Set(v string) error {
	switch v {
	case `PATH_CREATE_TABLE`, `PATH_READ`, `PATH_READ_WRITE`:
		*f = PathOperation(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PATH_CREATE_TABLE", "PATH_READ", "PATH_READ_WRITE"`, v)
	}
}

// Values returns all possible values for PathOperation.
//
// There is no guarantee on the order of the values in the slice.
func (f *PathOperation) Values() []PathOperation {
	return []PathOperation{
		PathOperationPathCreateTable,
		PathOperationPathRead,
		PathOperationPathReadWrite,
	}
}

// Type always returns PathOperation to satisfy [pflag.Value] interface
func (f *PathOperation) Type() string {
	return "PathOperation"
}

type PermissionsChange struct {
	// The set of privileges to add.
	Add []Privilege `json:"add,omitempty"`
	// The principal whose privileges we are changing.
	Principal string `json:"principal,omitempty"`
	// The set of privileges to remove.
	Remove []Privilege `json:"remove,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PermissionsChange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionsChange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Progress information of the Online Table data synchronization pipeline.
type PipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	EstimatedCompletionTimeSeconds float64 `json:"estimated_completion_time_seconds,omitempty"`
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	LatestVersionCurrentlyProcessing int64 `json:"latest_version_currently_processing,omitempty"`
	// The completion ratio of this update. This is a number between 0 and 1.
	SyncProgressCompletion float64 `json:"sync_progress_completion,omitempty"`
	// The number of rows that have been synced in this update.
	SyncedRowCount int64 `json:"synced_row_count,omitempty"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	TotalRowCount int64 `json:"total_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineProgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineProgress) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A positional argument passed to a row filter or column mask function.
// Distinguishes between column references and literals.
type PolicyFunctionArgument struct {
	// A column reference.
	Column string `json:"column,omitempty"`
	// A constant literal.
	Constant string `json:"constant,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PolicyFunctionArgument) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PolicyFunctionArgument) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PolicyInfo struct {
	// Options for column mask policies. Valid only if `policy_type` is
	// `POLICY_TYPE_COLUMN_MASK`. Required on create and optional on update.
	// When specified on update, the new options will replace the existing
	// options as a whole.
	ColumnMask *ColumnMaskOptions `json:"column_mask,omitempty"`
	// Optional description of the policy.
	Comment string `json:"comment,omitempty"`
	// Time at which the policy was created, in epoch milliseconds. Output only.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of the user who created the policy. Output only.
	CreatedBy string `json:"created_by,omitempty"`
	// Options for deny policies. Valid only if `policy_type` is
	// `POLICY_TYPE_DENY`. Required on create and optional on update. When
	// specified on update, the new options will replace the existing options as
	// a whole.
	Deny *DenyOptions `json:"deny,omitempty"`
	// Optional list of user or group names that should be excluded from the
	// policy.
	ExceptPrincipals []string `json:"except_principals,omitempty"`
	// Type of securables that the policy should take effect on. Required on
	// create and optional on update.
	ForSecurableType SecurableType `json:"for_securable_type"`
	// Governed tag keys and values referenced by this policy's conditions and
	// function arguments. Output only.
	GovernedTags []GovernedTagReference `json:"governed_tags,omitempty"`
	// Options for grant policies. Valid only if `policy_type` is
	// `POLICY_TYPE_GRANT`. Required on create and optional on update. When
	// specified on update, the new options will replace the existing options as
	// a whole.
	Grant *GrantOptions `json:"grant,omitempty"`
	// Unique identifier of the policy. This field is output only and is
	// generated by the system.
	Id string `json:"id,omitempty"`
	// Optional list of condition expressions used to match table columns. Only
	// valid when `for_securable_type` is `TABLE`. When specified, the policy
	// only applies to tables whose columns satisfy all match conditions.
	MatchColumns []MatchColumn `json:"match_columns,omitempty"`
	// Name of the policy. Required on create and optional on update. To rename
	// the policy, set `name` to a different value on update.
	Name string `json:"name,omitempty"`
	// Full name of the securable on which the policy is defined. Required on
	// create.
	OnSecurableFullname string `json:"on_securable_fullname,omitempty"`
	// Type of the securable on which the policy is defined. Only `CATALOG`,
	// `SCHEMA` and `TABLE` are supported at this moment. Required on create.
	OnSecurableType SecurableType `json:"on_securable_type,omitempty"`
	// Type of the policy. Required on create.
	PolicyType PolicyType `json:"policy_type"`
	// Options for row filter policies. Valid only if `policy_type` is
	// `POLICY_TYPE_ROW_FILTER`. Required on create and optional on update. When
	// specified on update, the new options will replace the existing options as
	// a whole.
	RowFilter *RowFilterOptions `json:"row_filter,omitempty"`
	// List of user or group names that the policy applies to. Required on
	// create and optional on update.
	ToPrincipals []string `json:"to_principals"`
	// Time at which the policy was last modified, in epoch milliseconds. Output
	// only.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of the user who last modified the policy. Output only.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Optional condition when the policy should take effect.
	WhenCondition string `json:"when_condition,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PolicyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PolicyInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PolicyType string

const PolicyTypePolicyTypeColumnMask PolicyType = `POLICY_TYPE_COLUMN_MASK`

const PolicyTypePolicyTypeDeny PolicyType = `POLICY_TYPE_DENY`

const PolicyTypePolicyTypeGrant PolicyType = `POLICY_TYPE_GRANT`

const PolicyTypePolicyTypeRowFilter PolicyType = `POLICY_TYPE_ROW_FILTER`

// String representation for [fmt.Print]
func (f *PolicyType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PolicyType) Set(v string) error {
	switch v {
	case `POLICY_TYPE_COLUMN_MASK`, `POLICY_TYPE_DENY`, `POLICY_TYPE_GRANT`, `POLICY_TYPE_ROW_FILTER`:
		*f = PolicyType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "POLICY_TYPE_COLUMN_MASK", "POLICY_TYPE_DENY", "POLICY_TYPE_GRANT", "POLICY_TYPE_ROW_FILTER"`, v)
	}
}

// Values returns all possible values for PolicyType.
//
// There is no guarantee on the order of the values in the slice.
func (f *PolicyType) Values() []PolicyType {
	return []PolicyType{
		PolicyTypePolicyTypeColumnMask,
		PolicyTypePolicyTypeDeny,
		PolicyTypePolicyTypeGrant,
		PolicyTypePolicyTypeRowFilter,
	}
}

// Type always returns PolicyType to satisfy [pflag.Value] interface
func (f *PolicyType) Type() string {
	return "PolicyType"
}

type PrimaryKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns []string `json:"child_columns"`
	// The name of the constraint.
	Name string `json:"name"`
	// True if the constraint is RELY, false or unset if NORELY.
	Rely bool `json:"rely,omitempty"`
	// Column names that represent a timeseries.
	TimeseriesColumns []string `json:"timeseries_columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PrimaryKeyConstraint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrimaryKeyConstraint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Principal struct {
	// Databricks user, group or service principal ID.
	Id string `json:"id,omitempty"`
	// The type of principal that **id** identifies: a user, a group, or a
	// service principal.
	PrincipalType PrincipalType `json:"principal_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Principal) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Principal) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PrincipalType string

const PrincipalTypeGroupPrincipal PrincipalType = `GROUP_PRINCIPAL`

const PrincipalTypeServicePrincipal PrincipalType = `SERVICE_PRINCIPAL`

const PrincipalTypeUserPrincipal PrincipalType = `USER_PRINCIPAL`

// String representation for [fmt.Print]
func (f *PrincipalType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrincipalType) Set(v string) error {
	switch v {
	case `GROUP_PRINCIPAL`, `SERVICE_PRINCIPAL`, `USER_PRINCIPAL`:
		*f = PrincipalType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GROUP_PRINCIPAL", "SERVICE_PRINCIPAL", "USER_PRINCIPAL"`, v)
	}
}

// Values returns all possible values for PrincipalType.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrincipalType) Values() []PrincipalType {
	return []PrincipalType{
		PrincipalTypeGroupPrincipal,
		PrincipalTypeServicePrincipal,
		PrincipalTypeUserPrincipal,
	}
}

// Type always returns PrincipalType to satisfy [pflag.Value] interface
func (f *PrincipalType) Type() string {
	return "PrincipalType"
}

type Privilege string

const PrivilegeAccess Privilege = `ACCESS`

const PrivilegeAllPrivileges Privilege = `ALL_PRIVILEGES`

const PrivilegeApplyTag Privilege = `APPLY_TAG`

const PrivilegeBrowse Privilege = `BROWSE`

const PrivilegeCreate Privilege = `CREATE`

const PrivilegeCreateCatalog Privilege = `CREATE_CATALOG`

const PrivilegeCreateCleanRoom Privilege = `CREATE_CLEAN_ROOM`

const PrivilegeCreateConnection Privilege = `CREATE_CONNECTION`

const PrivilegeCreateExternalLocation Privilege = `CREATE_EXTERNAL_LOCATION`

const PrivilegeCreateExternalTable Privilege = `CREATE_EXTERNAL_TABLE`

const PrivilegeCreateExternalVolume Privilege = `CREATE_EXTERNAL_VOLUME`

const PrivilegeCreateFeature Privilege = `CREATE_FEATURE`

const PrivilegeCreateForeignCatalog Privilege = `CREATE_FOREIGN_CATALOG`

const PrivilegeCreateForeignSecurable Privilege = `CREATE_FOREIGN_SECURABLE`

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

const PrivilegeExecuteCleanRoomTask Privilege = `EXECUTE_CLEAN_ROOM_TASK`

const PrivilegeExternalUseLocation Privilege = `EXTERNAL_USE_LOCATION`

const PrivilegeExternalUseSchema Privilege = `EXTERNAL_USE_SCHEMA`

const PrivilegeManage Privilege = `MANAGE`

const PrivilegeManageAllowlist Privilege = `MANAGE_ALLOWLIST`

const PrivilegeModify Privilege = `MODIFY`

const PrivilegeModifyCleanRoom Privilege = `MODIFY_CLEAN_ROOM`

const PrivilegeReadFeature Privilege = `READ_FEATURE`

const PrivilegeReadFiles Privilege = `READ_FILES`

const PrivilegeReadMetadata Privilege = `READ_METADATA`

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
	case `ACCESS`, `ALL_PRIVILEGES`, `APPLY_TAG`, `BROWSE`, `CREATE`, `CREATE_CATALOG`, `CREATE_CLEAN_ROOM`, `CREATE_CONNECTION`, `CREATE_EXTERNAL_LOCATION`, `CREATE_EXTERNAL_TABLE`, `CREATE_EXTERNAL_VOLUME`, `CREATE_FEATURE`, `CREATE_FOREIGN_CATALOG`, `CREATE_FOREIGN_SECURABLE`, `CREATE_FUNCTION`, `CREATE_MANAGED_STORAGE`, `CREATE_MATERIALIZED_VIEW`, `CREATE_MODEL`, `CREATE_PROVIDER`, `CREATE_RECIPIENT`, `CREATE_SCHEMA`, `CREATE_SERVICE_CREDENTIAL`, `CREATE_SHARE`, `CREATE_STORAGE_CREDENTIAL`, `CREATE_TABLE`, `CREATE_VIEW`, `CREATE_VOLUME`, `EXECUTE`, `EXECUTE_CLEAN_ROOM_TASK`, `EXTERNAL_USE_LOCATION`, `EXTERNAL_USE_SCHEMA`, `MANAGE`, `MANAGE_ALLOWLIST`, `MODIFY`, `MODIFY_CLEAN_ROOM`, `READ_FEATURE`, `READ_FILES`, `READ_METADATA`, `READ_PRIVATE_FILES`, `READ_VOLUME`, `REFRESH`, `SELECT`, `SET_SHARE_PERMISSION`, `USAGE`, `USE_CATALOG`, `USE_CONNECTION`, `USE_MARKETPLACE_ASSETS`, `USE_PROVIDER`, `USE_RECIPIENT`, `USE_SCHEMA`, `USE_SHARE`, `WRITE_FILES`, `WRITE_PRIVATE_FILES`, `WRITE_VOLUME`:
		*f = Privilege(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACCESS", "ALL_PRIVILEGES", "APPLY_TAG", "BROWSE", "CREATE", "CREATE_CATALOG", "CREATE_CLEAN_ROOM", "CREATE_CONNECTION", "CREATE_EXTERNAL_LOCATION", "CREATE_EXTERNAL_TABLE", "CREATE_EXTERNAL_VOLUME", "CREATE_FEATURE", "CREATE_FOREIGN_CATALOG", "CREATE_FOREIGN_SECURABLE", "CREATE_FUNCTION", "CREATE_MANAGED_STORAGE", "CREATE_MATERIALIZED_VIEW", "CREATE_MODEL", "CREATE_PROVIDER", "CREATE_RECIPIENT", "CREATE_SCHEMA", "CREATE_SERVICE_CREDENTIAL", "CREATE_SHARE", "CREATE_STORAGE_CREDENTIAL", "CREATE_TABLE", "CREATE_VIEW", "CREATE_VOLUME", "EXECUTE", "EXECUTE_CLEAN_ROOM_TASK", "EXTERNAL_USE_LOCATION", "EXTERNAL_USE_SCHEMA", "MANAGE", "MANAGE_ALLOWLIST", "MODIFY", "MODIFY_CLEAN_ROOM", "READ_FEATURE", "READ_FILES", "READ_METADATA", "READ_PRIVATE_FILES", "READ_VOLUME", "REFRESH", "SELECT", "SET_SHARE_PERMISSION", "USAGE", "USE_CATALOG", "USE_CONNECTION", "USE_MARKETPLACE_ASSETS", "USE_PROVIDER", "USE_RECIPIENT", "USE_SCHEMA", "USE_SHARE", "WRITE_FILES", "WRITE_PRIVATE_FILES", "WRITE_VOLUME"`, v)
	}
}

// Values returns all possible values for Privilege.
//
// There is no guarantee on the order of the values in the slice.
func (f *Privilege) Values() []Privilege {
	return []Privilege{
		PrivilegeAccess,
		PrivilegeAllPrivileges,
		PrivilegeApplyTag,
		PrivilegeBrowse,
		PrivilegeCreate,
		PrivilegeCreateCatalog,
		PrivilegeCreateCleanRoom,
		PrivilegeCreateConnection,
		PrivilegeCreateExternalLocation,
		PrivilegeCreateExternalTable,
		PrivilegeCreateExternalVolume,
		PrivilegeCreateFeature,
		PrivilegeCreateForeignCatalog,
		PrivilegeCreateForeignSecurable,
		PrivilegeCreateFunction,
		PrivilegeCreateManagedStorage,
		PrivilegeCreateMaterializedView,
		PrivilegeCreateModel,
		PrivilegeCreateProvider,
		PrivilegeCreateRecipient,
		PrivilegeCreateSchema,
		PrivilegeCreateServiceCredential,
		PrivilegeCreateShare,
		PrivilegeCreateStorageCredential,
		PrivilegeCreateTable,
		PrivilegeCreateView,
		PrivilegeCreateVolume,
		PrivilegeExecute,
		PrivilegeExecuteCleanRoomTask,
		PrivilegeExternalUseLocation,
		PrivilegeExternalUseSchema,
		PrivilegeManage,
		PrivilegeManageAllowlist,
		PrivilegeModify,
		PrivilegeModifyCleanRoom,
		PrivilegeReadFeature,
		PrivilegeReadFiles,
		PrivilegeReadMetadata,
		PrivilegeReadPrivateFiles,
		PrivilegeReadVolume,
		PrivilegeRefresh,
		PrivilegeSelect,
		PrivilegeSetSharePermission,
		PrivilegeUsage,
		PrivilegeUseCatalog,
		PrivilegeUseConnection,
		PrivilegeUseMarketplaceAssets,
		PrivilegeUseProvider,
		PrivilegeUseRecipient,
		PrivilegeUseSchema,
		PrivilegeUseShare,
		PrivilegeWriteFiles,
		PrivilegeWritePrivateFiles,
		PrivilegeWriteVolume,
	}
}

// Type always returns Privilege to satisfy [pflag.Value] interface
func (f *Privilege) Type() string {
	return "Privilege"
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal string `json:"principal,omitempty"`
	// The privileges assigned to the principal.
	Privileges []Privilege `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Status of an asynchronously provisioned resource.
type ProvisioningInfo struct {
	// The provisioning state of the resource.
	State ProvisioningInfoState `json:"state,omitempty"`
}

func (s *ProvisioningInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ProvisioningInfoState string

const ProvisioningInfoStateActive ProvisioningInfoState = `ACTIVE`

const ProvisioningInfoStateDegraded ProvisioningInfoState = `DEGRADED`

const ProvisioningInfoStateDeleting ProvisioningInfoState = `DELETING`

const ProvisioningInfoStateFailed ProvisioningInfoState = `FAILED`

const ProvisioningInfoStateProvisioning ProvisioningInfoState = `PROVISIONING`

const ProvisioningInfoStateUpdating ProvisioningInfoState = `UPDATING`

// String representation for [fmt.Print]
func (f *ProvisioningInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ProvisioningInfoState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DEGRADED`, `DELETING`, `FAILED`, `PROVISIONING`, `UPDATING`:
		*f = ProvisioningInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DEGRADED", "DELETING", "FAILED", "PROVISIONING", "UPDATING"`, v)
	}
}

// Values returns all possible values for ProvisioningInfoState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ProvisioningInfoState) Values() []ProvisioningInfoState {
	return []ProvisioningInfoState{
		ProvisioningInfoStateActive,
		ProvisioningInfoStateDegraded,
		ProvisioningInfoStateDeleting,
		ProvisioningInfoStateFailed,
		ProvisioningInfoStateProvisioning,
		ProvisioningInfoStateUpdating,
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
	InitialPipelineSyncProgress *PipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
}

func (s *ProvisioningStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type QuotaInfo struct {
	// The timestamp that indicates when the quota count was last updated.
	LastRefreshedAt int64 `json:"last_refreshed_at,omitempty"`
	// Name of the parent resource. Returns metastore ID if the parent is a
	// metastore.
	ParentFullName string `json:"parent_full_name,omitempty"`
	// The quota parent securable type.
	ParentSecurableType SecurableType `json:"parent_securable_type,omitempty"`
	// The current usage of the resource quota.
	QuotaCount int `json:"quota_count,omitempty"`
	// The current limit of the resource quota.
	QuotaLimit int `json:"quota_limit,omitempty"`
	// The name of the quota.
	QuotaName string `json:"quota_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QuotaInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QuotaInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// R2 temporary credentials for API authentication. Read more at
// https://developers.cloudflare.com/r2/api/s3/tokens/.
type R2Credentials struct {
	// The access key ID that identifies the temporary credentials.
	AccessKeyId string `json:"access_key_id,omitempty"`
	// The secret access key associated with the access key.
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	// The generated JWT that users must pass to use the temporary credentials.
	SessionToken string `json:"session_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *R2Credentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s R2Credentials) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A rate limit applied to service requests. Leave `requests` or `tokens` unset
// to impose no limit on that dimension; set a value to cap that dimension
// within the renewal period.
type RateLimit struct {
	// Scope of the rate limit. Depending on this value, the limit applies to a
	// principal, the service as a whole, or each user by default.
	Key RateLimitRateLimitKey `json:"key"`
	// Principal this limit applies to: user email, group name, or service
	// principal application ID. Required when `key` applies to a user, group,
	// or service principal; otherwise it must be unset.
	Principal string `json:"principal,omitempty"`
	// Interval after which the request and token counters reset.
	RenewalPeriod RateLimitRateLimitRenewalPeriod `json:"renewal_period"`
	// Maximum requests allowed in one renewal period. Leave unset for no
	// request limit. Set to `0` to deny all requests.
	Requests int64 `json:"requests,omitempty"`
	// Maximum tokens allowed in one renewal period. Leave unset for no token
	// limit. Set to `0` to deny all requests.
	Tokens int64 `json:"tokens,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RateLimit) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RateLimit) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Scope key for a rate limit.
type RateLimitRateLimitKey string

const RateLimitRateLimitKeyRateLimitKeyService RateLimitRateLimitKey = `RATE_LIMIT_KEY_SERVICE`

const RateLimitRateLimitKeyRateLimitKeyServicePrincipal RateLimitRateLimitKey = `RATE_LIMIT_KEY_SERVICE_PRINCIPAL`

const RateLimitRateLimitKeyRateLimitKeyUser RateLimitRateLimitKey = `RATE_LIMIT_KEY_USER`

const RateLimitRateLimitKeyRateLimitKeyUserDefault RateLimitRateLimitKey = `RATE_LIMIT_KEY_USER_DEFAULT`

const RateLimitRateLimitKeyRateLimitKeyUserGroup RateLimitRateLimitKey = `RATE_LIMIT_KEY_USER_GROUP`

// String representation for [fmt.Print]
func (f *RateLimitRateLimitKey) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RateLimitRateLimitKey) Set(v string) error {
	switch v {
	case `RATE_LIMIT_KEY_SERVICE`, `RATE_LIMIT_KEY_SERVICE_PRINCIPAL`, `RATE_LIMIT_KEY_USER`, `RATE_LIMIT_KEY_USER_DEFAULT`, `RATE_LIMIT_KEY_USER_GROUP`:
		*f = RateLimitRateLimitKey(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "RATE_LIMIT_KEY_SERVICE", "RATE_LIMIT_KEY_SERVICE_PRINCIPAL", "RATE_LIMIT_KEY_USER", "RATE_LIMIT_KEY_USER_DEFAULT", "RATE_LIMIT_KEY_USER_GROUP"`, v)
	}
}

// Values returns all possible values for RateLimitRateLimitKey.
//
// There is no guarantee on the order of the values in the slice.
func (f *RateLimitRateLimitKey) Values() []RateLimitRateLimitKey {
	return []RateLimitRateLimitKey{
		RateLimitRateLimitKeyRateLimitKeyService,
		RateLimitRateLimitKeyRateLimitKeyServicePrincipal,
		RateLimitRateLimitKeyRateLimitKeyUser,
		RateLimitRateLimitKeyRateLimitKeyUserDefault,
		RateLimitRateLimitKeyRateLimitKeyUserGroup,
	}
}

// Type always returns RateLimitRateLimitKey to satisfy [pflag.Value] interface
func (f *RateLimitRateLimitKey) Type() string {
	return "RateLimitRateLimitKey"
}

// Renewal period for a rate limit.
type RateLimitRateLimitRenewalPeriod string

const RateLimitRateLimitRenewalPeriodRateLimitRenewalPeriodHour RateLimitRateLimitRenewalPeriod = `RATE_LIMIT_RENEWAL_PERIOD_HOUR`

const RateLimitRateLimitRenewalPeriodRateLimitRenewalPeriodMinute RateLimitRateLimitRenewalPeriod = `RATE_LIMIT_RENEWAL_PERIOD_MINUTE`

// String representation for [fmt.Print]
func (f *RateLimitRateLimitRenewalPeriod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RateLimitRateLimitRenewalPeriod) Set(v string) error {
	switch v {
	case `RATE_LIMIT_RENEWAL_PERIOD_HOUR`, `RATE_LIMIT_RENEWAL_PERIOD_MINUTE`:
		*f = RateLimitRateLimitRenewalPeriod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "RATE_LIMIT_RENEWAL_PERIOD_HOUR", "RATE_LIMIT_RENEWAL_PERIOD_MINUTE"`, v)
	}
}

// Values returns all possible values for RateLimitRateLimitRenewalPeriod.
//
// There is no guarantee on the order of the values in the slice.
func (f *RateLimitRateLimitRenewalPeriod) Values() []RateLimitRateLimitRenewalPeriod {
	return []RateLimitRateLimitRenewalPeriod{
		RateLimitRateLimitRenewalPeriodRateLimitRenewalPeriodHour,
		RateLimitRateLimitRenewalPeriodRateLimitRenewalPeriodMinute,
	}
}

// Type always returns RateLimitRateLimitRenewalPeriod to satisfy [pflag.Value] interface
func (f *RateLimitRateLimitRenewalPeriod) Type() string {
	return "RateLimitRateLimitRenewalPeriod"
}

type ReadVolumeRequest struct {
	// Whether to include volumes for which the principal can only access
	// selective metadata through the BROWSE privilege.
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// The fully qualified name of the volume.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ReadVolumeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReadVolumeRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegenerateDashboardRequest struct {
	// UC table name in format `catalog.schema.table_name`. This field
	// corresponds to the {full_table_name_arg} arg in the endpoint path.
	TableName string `json:"-" url:"-"`
	// Optional argument to specify the warehouse for dashboard regeneration. If
	// not specified, the first running warehouse will be used.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegenerateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegenerateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegenerateDashboardResponse struct {
	DashboardId string `json:"dashboard_id,omitempty"`
	// Parent folder is equivalent to {assets_dir}/{tableName}
	ParentFolder string `json:"parent_folder,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegenerateDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegenerateDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelAlias struct {
	// Name of the alias, e.g. 'champion' or 'latest_stable'
	AliasName string `json:"alias_name,omitempty"`
	// The name of the catalog containing the model version
	CatalogName string `json:"catalog_name,omitempty"`
	// The unique identifier of the alias
	Id string `json:"id,omitempty"`
	// The name of the parent registered model of the model version, relative to
	// parent schema
	ModelName string `json:"model_name,omitempty"`
	// The name of the schema containing the model version, relative to parent
	// catalog
	SchemaName string `json:"schema_name,omitempty"`
	// Integer version number of the model version to which this alias points.
	VersionNum int `json:"version_num,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelAlias) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelAlias) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelInfo struct {
	// List of aliases associated with the registered model
	Aliases []RegisteredModelAlias `json:"aliases,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog where the schema and the registered model reside
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the registered model
	Comment string `json:"comment,omitempty"`
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the registered model
	CreatedBy string `json:"created_by,omitempty"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the metastore
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the registered model
	Name string `json:"name,omitempty"`
	// The identifier of the user who owns the registered model
	Owner string `json:"owner,omitempty"`
	// The name of the schema where the registered model resides
	SchemaName string `json:"schema_name,omitempty"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `json:"storage_location,omitempty"`
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the registered model last time
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RowFilterOptions struct {
	// The fully qualified name of the row filter function. The function is
	// called on each row of the target table. It should return a boolean value
	// indicating whether the row should be visible to the user. Required on
	// create and update.
	FunctionName string `json:"function_name"`
	// Optional list of column aliases or constant literals to be passed as
	// arguments to the row filter function. The type of each column should
	// match the positional argument of the row filter function.
	Using []FunctionArgument `json:"using,omitempty"`
}

func (s *RowFilterOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RunRefreshRequest struct {
	// UC table name in format `catalog.schema.table_name`. table_name is case
	// insensitive and spaces are disallowed.
	TableName string `json:"-" url:"-"`
}

func (s *RunRefreshRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A schema (also called a database) is the second layer of Unity Catalog's
// three-level namespace and groups the tables, views, volumes, models, and
// functions under its parent catalog. Managed tables and volumes in the schema
// are stored under the schema's storage root when it has one, and otherwise
// under the storage of the parent catalog or the metastore.
type SchemaInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Name of parent catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// The type of the parent catalog, which the server copies from the catalog.
	CatalogType CatalogType `json:"catalog_type,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this schema was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of schema creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Dropped objects in the schema remain recoverable for this number of hours
	// before they are permanently purged. The value must be 0 or within the
	// range allowed for the metastore. Schemas that do not set their own value
	// inherit the parent catalog's value.
	CustomMaxRetentionHours int64 `json:"custom_max_retention_hours,omitempty"`
	// Indicates the predictive optimization setting that is in effect for the
	// schema once inherited values are resolved.
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `json:"effective_predictive_optimization_flag,omitempty"`
	// Whether predictive optimization is enabled for this schema and the
	// objects under it. Valid values are ENABLE, DISABLE, and INHERIT, where
	// INHERIT applies the parent catalog's setting. New schemas start as
	// INHERIT. It cannot be set when creating a schema; use the update method
	// instead.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Full name of schema, in form of __catalog_name__.__schema_name__.
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of schema, relative to parent catalog.
	Name string `json:"name,omitempty"`
	// The user or group that owns the schema. The creator becomes the owner,
	// and the owner can be changed only through the update method.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// The unique identifier (UUID) that the server assigns to the schema when
	// it is created.
	SchemaId string `json:"schema_id,omitempty"`
	// The server derives this full storage location URL for managed tables and
	// volumes within the schema from storage_root and the schema ID. It is
	// absent when the schema has no storage root of its own.
	StorageLocation string `json:"storage_location,omitempty"`
	// Managed tables and volumes in the schema are stored under this storage
	// root URL. The path must be inside an external location on which the
	// caller has the CREATE MANAGED STORAGE privilege. If omitted, managed data
	// is stored under the parent catalog's storage root, or under the
	// metastore's root storage location when the catalog has none. Schemas in
	// Delta Sharing catalogs cannot have a storage root, and schemas in a
	// federated Hive metastore catalog must set one.
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this schema was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified schema.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SchemaInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SchemaInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A secret stored in Unity Catalog. Secrets are three-level namespace objects
// (catalog.schema.secret) that securely store sensitive credential data such as
// passwords, tokens, and keys.
type Secret struct {
	// The name of the catalog where the schema and the secret reside.
	CatalogName string `json:"catalog_name"`
	// User-provided free-form text description of the secret.
	Comment string `json:"comment,omitempty"`
	// The time at which this secret was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The principal that created the secret.
	CreatedBy string `json:"created_by,omitempty"`
	// The effective owner of the secret, which may differ from the directly-set
	// **owner** due to inheritance.
	EffectiveOwner string `json:"effective_owner,omitempty"`
	// The secret value. Only populated in responses when you have the
	// **READ_SECRET** privilege and **include_value** is set to true in the
	// request. The maximum size is 60 KiB.
	EffectiveValue string `json:"effective_value,omitempty"`
	// User-provided expiration time of the secret. This field indicates when
	// the secret should no longer be used and may be displayed as a warning in
	// the UI. It is purely informational and does not trigger any automatic
	// actions or affect the secret's lifecycle.
	ExpireTime *time.Time `json:"expire_time,omitempty"`
	// The three-level (fully qualified) name of the secret, in the form of
	// **catalog_name.schema_name.secret_name**.
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of the metastore hosting the secret.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the secret, relative to its parent schema.
	Name string `json:"name"`
	// The owner of the secret. It cannot be set when creating a secret and
	// defaults to the creating principal. Set it in UpdateSecret to transfer
	// ownership of the secret to another principal.
	Owner string `json:"owner,omitempty"`
	// The name of the schema where the secret resides.
	SchemaName string `json:"schema_name"`
	// The time at which this secret was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// The principal that last updated the secret.
	UpdatedBy string `json:"updated_by,omitempty"`
	// The secret value to store, such as a password, token, or key. It is never
	// returned in responses. To read it, call GetSecret with **include_value**
	// set to true and use the **effective_value** field. The maximum size is 60
	// KiB before encryption.
	Value string `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Secret) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Secret) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Generic definition of a securable, which is uniquely defined in a metastore
// by its type and full name.
type Securable struct {
	// Required. The full name of the catalog/schema/table. Optional if
	// resource_name is present.
	FullName string `json:"full_name,omitempty"`
	// Optional. The name of the Share object that contains the securable when
	// the securable is getting shared in D2D Delta Sharing.
	ProviderShare string `json:"provider_share,omitempty"`
	// Required. The type of securable (catalog/schema/table). Optional if
	// resource_name is present.
	Type SecurableType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Securable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Securable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SecurableKind string

const SecurableKindTableDbStorage SecurableKind = `TABLE_DB_STORAGE`

const SecurableKindTableDelta SecurableKind = `TABLE_DELTA`

const SecurableKindTableDeltasharing SecurableKind = `TABLE_DELTASHARING`

const SecurableKindTableDeltasharingMutable SecurableKind = `TABLE_DELTASHARING_MUTABLE`

const SecurableKindTableDeltasharingOpenDirBased SecurableKind = `TABLE_DELTASHARING_OPEN_DIR_BASED`

const SecurableKindTableDeltaExternal SecurableKind = `TABLE_DELTA_EXTERNAL`

const SecurableKindTableDeltaIcebergDeltasharing SecurableKind = `TABLE_DELTA_ICEBERG_DELTASHARING`

const SecurableKindTableDeltaIcebergManaged SecurableKind = `TABLE_DELTA_ICEBERG_MANAGED`

const SecurableKindTableDeltaUniformHudiExternal SecurableKind = `TABLE_DELTA_UNIFORM_HUDI_EXTERNAL`

const SecurableKindTableDeltaUniformIcebergExternal SecurableKind = `TABLE_DELTA_UNIFORM_ICEBERG_EXTERNAL`

const SecurableKindTableDeltaUniformIcebergExternalDeltasharing SecurableKind = `TABLE_DELTA_UNIFORM_ICEBERG_EXTERNAL_DELTASHARING`

const SecurableKindTableDeltaUniformIcebergForeignDeltasharing SecurableKind = `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_DELTASHARING`

const SecurableKindTableDeltaUniformIcebergForeignHiveMetastoreExternal SecurableKind = `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_EXTERNAL`

const SecurableKindTableDeltaUniformIcebergForeignHiveMetastoreManaged SecurableKind = `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_MANAGED`

const SecurableKindTableDeltaUniformIcebergForeignSnowflake SecurableKind = `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_SNOWFLAKE`

const SecurableKindTableExternal SecurableKind = `TABLE_EXTERNAL`

const SecurableKindTableFeatureStore SecurableKind = `TABLE_FEATURE_STORE`

const SecurableKindTableFeatureStoreExternal SecurableKind = `TABLE_FEATURE_STORE_EXTERNAL`

const SecurableKindTableForeignBigquery SecurableKind = `TABLE_FOREIGN_BIGQUERY`

const SecurableKindTableForeignDatabricks SecurableKind = `TABLE_FOREIGN_DATABRICKS`

const SecurableKindTableForeignDeltasharing SecurableKind = `TABLE_FOREIGN_DELTASHARING`

const SecurableKindTableForeignDeltaDeltasharing SecurableKind = `TABLE_FOREIGN_DELTA_DELTASHARING`

const SecurableKindTableForeignHiveMetastore SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE`

const SecurableKindTableForeignHiveMetastoreDbfsExternal SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_EXTERNAL`

const SecurableKindTableForeignHiveMetastoreDbfsManaged SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_MANAGED`

const SecurableKindTableForeignHiveMetastoreDbfsShallowCloneExternal SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_EXTERNAL`

const SecurableKindTableForeignHiveMetastoreDbfsShallowCloneManaged SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_MANAGED`

const SecurableKindTableForeignHiveMetastoreDbfsView SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_VIEW`

const SecurableKindTableForeignHiveMetastoreExternal SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_EXTERNAL`

const SecurableKindTableForeignHiveMetastoreManaged SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_MANAGED`

const SecurableKindTableForeignHiveMetastoreShallowCloneExternal SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_EXTERNAL`

const SecurableKindTableForeignHiveMetastoreShallowCloneManaged SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_MANAGED`

const SecurableKindTableForeignHiveMetastoreView SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_VIEW`

const SecurableKindTableForeignMongodb SecurableKind = `TABLE_FOREIGN_MONGODB`

const SecurableKindTableForeignMysql SecurableKind = `TABLE_FOREIGN_MYSQL`

const SecurableKindTableForeignNetsuite SecurableKind = `TABLE_FOREIGN_NETSUITE`

const SecurableKindTableForeignOracle SecurableKind = `TABLE_FOREIGN_ORACLE`

const SecurableKindTableForeignPostgresql SecurableKind = `TABLE_FOREIGN_POSTGRESQL`

const SecurableKindTableForeignRedshift SecurableKind = `TABLE_FOREIGN_REDSHIFT`

const SecurableKindTableForeignSalesforce SecurableKind = `TABLE_FOREIGN_SALESFORCE`

const SecurableKindTableForeignSalesforceDataCloud SecurableKind = `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD`

const SecurableKindTableForeignSalesforceDataCloudFileSharing SecurableKind = `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING`

const SecurableKindTableForeignSalesforceDataCloudFileSharingView SecurableKind = `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING_VIEW`

const SecurableKindTableForeignSnowflake SecurableKind = `TABLE_FOREIGN_SNOWFLAKE`

const SecurableKindTableForeignSqldw SecurableKind = `TABLE_FOREIGN_SQLDW`

const SecurableKindTableForeignSqlserver SecurableKind = `TABLE_FOREIGN_SQLSERVER`

const SecurableKindTableForeignTeradata SecurableKind = `TABLE_FOREIGN_TERADATA`

const SecurableKindTableForeignWorkdayRaas SecurableKind = `TABLE_FOREIGN_WORKDAY_RAAS`

const SecurableKindTableIcebergUniformManaged SecurableKind = `TABLE_ICEBERG_UNIFORM_MANAGED`

const SecurableKindTableInternal SecurableKind = `TABLE_INTERNAL`

const SecurableKindTableManagedPostgresql SecurableKind = `TABLE_MANAGED_POSTGRESQL`

const SecurableKindTableMaterializedView SecurableKind = `TABLE_MATERIALIZED_VIEW`

const SecurableKindTableMaterializedViewDeltasharing SecurableKind = `TABLE_MATERIALIZED_VIEW_DELTASHARING`

const SecurableKindTableMetricView SecurableKind = `TABLE_METRIC_VIEW`

const SecurableKindTableMetricViewDeltasharing SecurableKind = `TABLE_METRIC_VIEW_DELTASHARING`

const SecurableKindTableOnlineVectorIndexDirect SecurableKind = `TABLE_ONLINE_VECTOR_INDEX_DIRECT`

const SecurableKindTableOnlineVectorIndexReplica SecurableKind = `TABLE_ONLINE_VECTOR_INDEX_REPLICA`

const SecurableKindTableOnlineView SecurableKind = `TABLE_ONLINE_VIEW`

const SecurableKindTableStandard SecurableKind = `TABLE_STANDARD`

const SecurableKindTableStreamingLiveTable SecurableKind = `TABLE_STREAMING_LIVE_TABLE`

const SecurableKindTableStreamingLiveTableDeltasharing SecurableKind = `TABLE_STREAMING_LIVE_TABLE_DELTASHARING`

const SecurableKindTableSystem SecurableKind = `TABLE_SYSTEM`

const SecurableKindTableSystemDeltasharing SecurableKind = `TABLE_SYSTEM_DELTASHARING`

const SecurableKindTableView SecurableKind = `TABLE_VIEW`

const SecurableKindTableViewDeltasharing SecurableKind = `TABLE_VIEW_DELTASHARING`

// String representation for [fmt.Print]
func (f *SecurableKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SecurableKind) Set(v string) error {
	switch v {
	case `TABLE_DB_STORAGE`, `TABLE_DELTA`, `TABLE_DELTASHARING`, `TABLE_DELTASHARING_MUTABLE`, `TABLE_DELTASHARING_OPEN_DIR_BASED`, `TABLE_DELTA_EXTERNAL`, `TABLE_DELTA_ICEBERG_DELTASHARING`, `TABLE_DELTA_ICEBERG_MANAGED`, `TABLE_DELTA_UNIFORM_HUDI_EXTERNAL`, `TABLE_DELTA_UNIFORM_ICEBERG_EXTERNAL`, `TABLE_DELTA_UNIFORM_ICEBERG_EXTERNAL_DELTASHARING`, `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_DELTASHARING`, `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_EXTERNAL`, `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_MANAGED`, `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_SNOWFLAKE`, `TABLE_EXTERNAL`, `TABLE_FEATURE_STORE`, `TABLE_FEATURE_STORE_EXTERNAL`, `TABLE_FOREIGN_BIGQUERY`, `TABLE_FOREIGN_DATABRICKS`, `TABLE_FOREIGN_DELTASHARING`, `TABLE_FOREIGN_DELTA_DELTASHARING`, `TABLE_FOREIGN_HIVE_METASTORE`, `TABLE_FOREIGN_HIVE_METASTORE_DBFS_EXTERNAL`, `TABLE_FOREIGN_HIVE_METASTORE_DBFS_MANAGED`, `TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_EXTERNAL`, `TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_MANAGED`, `TABLE_FOREIGN_HIVE_METASTORE_DBFS_VIEW`, `TABLE_FOREIGN_HIVE_METASTORE_EXTERNAL`, `TABLE_FOREIGN_HIVE_METASTORE_MANAGED`, `TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_EXTERNAL`, `TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_MANAGED`, `TABLE_FOREIGN_HIVE_METASTORE_VIEW`, `TABLE_FOREIGN_MONGODB`, `TABLE_FOREIGN_MYSQL`, `TABLE_FOREIGN_NETSUITE`, `TABLE_FOREIGN_ORACLE`, `TABLE_FOREIGN_POSTGRESQL`, `TABLE_FOREIGN_REDSHIFT`, `TABLE_FOREIGN_SALESFORCE`, `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD`, `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING`, `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING_VIEW`, `TABLE_FOREIGN_SNOWFLAKE`, `TABLE_FOREIGN_SQLDW`, `TABLE_FOREIGN_SQLSERVER`, `TABLE_FOREIGN_TERADATA`, `TABLE_FOREIGN_WORKDAY_RAAS`, `TABLE_ICEBERG_UNIFORM_MANAGED`, `TABLE_INTERNAL`, `TABLE_MANAGED_POSTGRESQL`, `TABLE_MATERIALIZED_VIEW`, `TABLE_MATERIALIZED_VIEW_DELTASHARING`, `TABLE_METRIC_VIEW`, `TABLE_METRIC_VIEW_DELTASHARING`, `TABLE_ONLINE_VECTOR_INDEX_DIRECT`, `TABLE_ONLINE_VECTOR_INDEX_REPLICA`, `TABLE_ONLINE_VIEW`, `TABLE_STANDARD`, `TABLE_STREAMING_LIVE_TABLE`, `TABLE_STREAMING_LIVE_TABLE_DELTASHARING`, `TABLE_SYSTEM`, `TABLE_SYSTEM_DELTASHARING`, `TABLE_VIEW`, `TABLE_VIEW_DELTASHARING`:
		*f = SecurableKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TABLE_DB_STORAGE", "TABLE_DELTA", "TABLE_DELTASHARING", "TABLE_DELTASHARING_MUTABLE", "TABLE_DELTASHARING_OPEN_DIR_BASED", "TABLE_DELTA_EXTERNAL", "TABLE_DELTA_ICEBERG_DELTASHARING", "TABLE_DELTA_ICEBERG_MANAGED", "TABLE_DELTA_UNIFORM_HUDI_EXTERNAL", "TABLE_DELTA_UNIFORM_ICEBERG_EXTERNAL", "TABLE_DELTA_UNIFORM_ICEBERG_EXTERNAL_DELTASHARING", "TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_DELTASHARING", "TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_EXTERNAL", "TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_MANAGED", "TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_SNOWFLAKE", "TABLE_EXTERNAL", "TABLE_FEATURE_STORE", "TABLE_FEATURE_STORE_EXTERNAL", "TABLE_FOREIGN_BIGQUERY", "TABLE_FOREIGN_DATABRICKS", "TABLE_FOREIGN_DELTASHARING", "TABLE_FOREIGN_DELTA_DELTASHARING", "TABLE_FOREIGN_HIVE_METASTORE", "TABLE_FOREIGN_HIVE_METASTORE_DBFS_EXTERNAL", "TABLE_FOREIGN_HIVE_METASTORE_DBFS_MANAGED", "TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_EXTERNAL", "TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_MANAGED", "TABLE_FOREIGN_HIVE_METASTORE_DBFS_VIEW", "TABLE_FOREIGN_HIVE_METASTORE_EXTERNAL", "TABLE_FOREIGN_HIVE_METASTORE_MANAGED", "TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_EXTERNAL", "TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_MANAGED", "TABLE_FOREIGN_HIVE_METASTORE_VIEW", "TABLE_FOREIGN_MONGODB", "TABLE_FOREIGN_MYSQL", "TABLE_FOREIGN_NETSUITE", "TABLE_FOREIGN_ORACLE", "TABLE_FOREIGN_POSTGRESQL", "TABLE_FOREIGN_REDSHIFT", "TABLE_FOREIGN_SALESFORCE", "TABLE_FOREIGN_SALESFORCE_DATA_CLOUD", "TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING", "TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING_VIEW", "TABLE_FOREIGN_SNOWFLAKE", "TABLE_FOREIGN_SQLDW", "TABLE_FOREIGN_SQLSERVER", "TABLE_FOREIGN_TERADATA", "TABLE_FOREIGN_WORKDAY_RAAS", "TABLE_ICEBERG_UNIFORM_MANAGED", "TABLE_INTERNAL", "TABLE_MANAGED_POSTGRESQL", "TABLE_MATERIALIZED_VIEW", "TABLE_MATERIALIZED_VIEW_DELTASHARING", "TABLE_METRIC_VIEW", "TABLE_METRIC_VIEW_DELTASHARING", "TABLE_ONLINE_VECTOR_INDEX_DIRECT", "TABLE_ONLINE_VECTOR_INDEX_REPLICA", "TABLE_ONLINE_VIEW", "TABLE_STANDARD", "TABLE_STREAMING_LIVE_TABLE", "TABLE_STREAMING_LIVE_TABLE_DELTASHARING", "TABLE_SYSTEM", "TABLE_SYSTEM_DELTASHARING", "TABLE_VIEW", "TABLE_VIEW_DELTASHARING"`, v)
	}
}

// Values returns all possible values for SecurableKind.
//
// There is no guarantee on the order of the values in the slice.
func (f *SecurableKind) Values() []SecurableKind {
	return []SecurableKind{
		SecurableKindTableDbStorage,
		SecurableKindTableDelta,
		SecurableKindTableDeltasharing,
		SecurableKindTableDeltasharingMutable,
		SecurableKindTableDeltasharingOpenDirBased,
		SecurableKindTableDeltaExternal,
		SecurableKindTableDeltaIcebergDeltasharing,
		SecurableKindTableDeltaIcebergManaged,
		SecurableKindTableDeltaUniformHudiExternal,
		SecurableKindTableDeltaUniformIcebergExternal,
		SecurableKindTableDeltaUniformIcebergExternalDeltasharing,
		SecurableKindTableDeltaUniformIcebergForeignDeltasharing,
		SecurableKindTableDeltaUniformIcebergForeignHiveMetastoreExternal,
		SecurableKindTableDeltaUniformIcebergForeignHiveMetastoreManaged,
		SecurableKindTableDeltaUniformIcebergForeignSnowflake,
		SecurableKindTableExternal,
		SecurableKindTableFeatureStore,
		SecurableKindTableFeatureStoreExternal,
		SecurableKindTableForeignBigquery,
		SecurableKindTableForeignDatabricks,
		SecurableKindTableForeignDeltasharing,
		SecurableKindTableForeignDeltaDeltasharing,
		SecurableKindTableForeignHiveMetastore,
		SecurableKindTableForeignHiveMetastoreDbfsExternal,
		SecurableKindTableForeignHiveMetastoreDbfsManaged,
		SecurableKindTableForeignHiveMetastoreDbfsShallowCloneExternal,
		SecurableKindTableForeignHiveMetastoreDbfsShallowCloneManaged,
		SecurableKindTableForeignHiveMetastoreDbfsView,
		SecurableKindTableForeignHiveMetastoreExternal,
		SecurableKindTableForeignHiveMetastoreManaged,
		SecurableKindTableForeignHiveMetastoreShallowCloneExternal,
		SecurableKindTableForeignHiveMetastoreShallowCloneManaged,
		SecurableKindTableForeignHiveMetastoreView,
		SecurableKindTableForeignMongodb,
		SecurableKindTableForeignMysql,
		SecurableKindTableForeignNetsuite,
		SecurableKindTableForeignOracle,
		SecurableKindTableForeignPostgresql,
		SecurableKindTableForeignRedshift,
		SecurableKindTableForeignSalesforce,
		SecurableKindTableForeignSalesforceDataCloud,
		SecurableKindTableForeignSalesforceDataCloudFileSharing,
		SecurableKindTableForeignSalesforceDataCloudFileSharingView,
		SecurableKindTableForeignSnowflake,
		SecurableKindTableForeignSqldw,
		SecurableKindTableForeignSqlserver,
		SecurableKindTableForeignTeradata,
		SecurableKindTableForeignWorkdayRaas,
		SecurableKindTableIcebergUniformManaged,
		SecurableKindTableInternal,
		SecurableKindTableManagedPostgresql,
		SecurableKindTableMaterializedView,
		SecurableKindTableMaterializedViewDeltasharing,
		SecurableKindTableMetricView,
		SecurableKindTableMetricViewDeltasharing,
		SecurableKindTableOnlineVectorIndexDirect,
		SecurableKindTableOnlineVectorIndexReplica,
		SecurableKindTableOnlineView,
		SecurableKindTableStandard,
		SecurableKindTableStreamingLiveTable,
		SecurableKindTableStreamingLiveTableDeltasharing,
		SecurableKindTableSystem,
		SecurableKindTableSystemDeltasharing,
		SecurableKindTableView,
		SecurableKindTableViewDeltasharing,
	}
}

// Type always returns SecurableKind to satisfy [pflag.Value] interface
func (f *SecurableKind) Type() string {
	return "SecurableKind"
}

// Manifest of a specific securable kind, or of an Atlas-native securable type,
// which has no kinds.
type SecurableKindManifest struct {
	// Privileges that can be assigned to the securable.
	AssignablePrivileges []string `json:"assignable_privileges,omitempty"`
	// A list of capabilities in the securable kind.
	Capabilities []string `json:"capabilities,omitempty"`
	// Detailed specs of allowed options.
	Options []OptionSpec `json:"options,omitempty"`
	// Securable kind to get manifest of. Unset for Atlas-native securable
	// types.
	SecurableKind SecurableKind `json:"securable_kind,omitempty"`
	// Securable Type of the kind.
	SecurableType SecurableType `json:"securable_type,omitempty"`
}

func (s *SecurableKindManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SecurablePermissions struct {
	// List of requested Unity Catalog permissions.
	Permissions []string `json:"permissions,omitempty"`
	// The securable on which the permissions are requested.
	Securable *Securable `json:"securable,omitempty"`
}

func (s *SecurablePermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The type of Unity Catalog securable.
type SecurableType string

const SecurableTypeAgentService SecurableType = `AGENT_SERVICE`

const SecurableTypeCatalog SecurableType = `CATALOG`

const SecurableTypeCleanRoom SecurableType = `CLEAN_ROOM`

const SecurableTypeConnection SecurableType = `CONNECTION`

const SecurableTypeCredential SecurableType = `CREDENTIAL`

const SecurableTypeExternalLocation SecurableType = `EXTERNAL_LOCATION`

const SecurableTypeExternalMetadata SecurableType = `EXTERNAL_METADATA`

const SecurableTypeFeature SecurableType = `FEATURE`

const SecurableTypeFunction SecurableType = `FUNCTION`

const SecurableTypeMcpService SecurableType = `MCP_SERVICE`

const SecurableTypeMetastore SecurableType = `METASTORE`

const SecurableTypeModel SecurableType = `MODEL`

const SecurableTypeModelProviderService SecurableType = `MODEL_PROVIDER_SERVICE`

const SecurableTypeModelService SecurableType = `MODEL_SERVICE`

const SecurableTypePipeline SecurableType = `PIPELINE`

const SecurableTypeProvider SecurableType = `PROVIDER`

const SecurableTypeRecipient SecurableType = `RECIPIENT`

const SecurableTypeSchema SecurableType = `SCHEMA`

const SecurableTypeShare SecurableType = `SHARE`

const SecurableTypeSkill SecurableType = `SKILL`

const SecurableTypeStagingTable SecurableType = `STAGING_TABLE`

const SecurableTypeStorageCredential SecurableType = `STORAGE_CREDENTIAL`

const SecurableTypeTable SecurableType = `TABLE`

const SecurableTypeVolume SecurableType = `VOLUME`

// String representation for [fmt.Print]
func (f *SecurableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SecurableType) Set(v string) error {
	switch v {
	case `AGENT_SERVICE`, `CATALOG`, `CLEAN_ROOM`, `CONNECTION`, `CREDENTIAL`, `EXTERNAL_LOCATION`, `EXTERNAL_METADATA`, `FEATURE`, `FUNCTION`, `MCP_SERVICE`, `METASTORE`, `MODEL`, `MODEL_PROVIDER_SERVICE`, `MODEL_SERVICE`, `PIPELINE`, `PROVIDER`, `RECIPIENT`, `SCHEMA`, `SHARE`, `SKILL`, `STAGING_TABLE`, `STORAGE_CREDENTIAL`, `TABLE`, `VOLUME`:
		*f = SecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AGENT_SERVICE", "CATALOG", "CLEAN_ROOM", "CONNECTION", "CREDENTIAL", "EXTERNAL_LOCATION", "EXTERNAL_METADATA", "FEATURE", "FUNCTION", "MCP_SERVICE", "METASTORE", "MODEL", "MODEL_PROVIDER_SERVICE", "MODEL_SERVICE", "PIPELINE", "PROVIDER", "RECIPIENT", "SCHEMA", "SHARE", "SKILL", "STAGING_TABLE", "STORAGE_CREDENTIAL", "TABLE", "VOLUME"`, v)
	}
}

// Values returns all possible values for SecurableType.
//
// There is no guarantee on the order of the values in the slice.
func (f *SecurableType) Values() []SecurableType {
	return []SecurableType{
		SecurableTypeAgentService,
		SecurableTypeCatalog,
		SecurableTypeCleanRoom,
		SecurableTypeConnection,
		SecurableTypeCredential,
		SecurableTypeExternalLocation,
		SecurableTypeExternalMetadata,
		SecurableTypeFeature,
		SecurableTypeFunction,
		SecurableTypeMcpService,
		SecurableTypeMetastore,
		SecurableTypeModel,
		SecurableTypeModelProviderService,
		SecurableTypeModelService,
		SecurableTypePipeline,
		SecurableTypeProvider,
		SecurableTypeRecipient,
		SecurableTypeSchema,
		SecurableTypeShare,
		SecurableTypeSkill,
		SecurableTypeStagingTable,
		SecurableTypeStorageCredential,
		SecurableTypeTable,
		SecurableTypeVolume,
	}
}

// Type always returns SecurableType to satisfy [pflag.Value] interface
func (f *SecurableType) Type() string {
	return "SecurableType"
}

type SetArtifactAllowlist struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers []ArtifactMatcher `json:"artifact_matchers"`
	// The artifact type of the allowlist.
	ArtifactType ArtifactType `json:"-" url:"-"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of the user who set the artifact allowlist.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SetArtifactAllowlist) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SetArtifactAllowlist) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SetRegisteredModelAliasRequest struct {
	// The name of the alias
	Alias string `json:"-" url:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
	// The version number of the model version to which the alias points
	VersionNum int `json:"version_num"`
}

func (s *SetRegisteredModelAliasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A Skill is an agentskills.io bundle registered in Unity Catalog. Clients
// transfer bundle bytes through the Files API. FinalizeSkill reads the uploaded
// SKILL.md and projects its frontmatter onto the Skill metadata.
type Skill struct {
	// Name from the most recently successfully finalized SKILL.md. It may
	// differ from the final component of the Skill resource name. Unset until
	// FinalizeSkill succeeds.
	BundleName string `json:"bundle_name,omitempty"`
	// User-provided comment for the skill, which you can change with
	// UpdateSkill. It is separate from `description`, which the server reads
	// from the SKILL.md frontmatter, and it is preserved across bundle
	// re-uploads and FinalizeSkill. When `comment` is in the update mask,
	// omitting it clears the field, while an explicitly empty string is
	// retained.
	Comment string `json:"comment,omitempty"`
	// Time the skill was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Username or service principal application ID of the principal that
	// created the skill.
	CreatedBy string `json:"created_by,omitempty"`
	// Description from the most recently successfully finalized SKILL.md. Unset
	// until FinalizeSkill succeeds.
	Description string `json:"description,omitempty"`
	// Username, group name, or service principal application ID of the skill
	// owner.
	EffectiveOwner string `json:"effective_owner,omitempty"`
	// Optimistic concurrency token returned on every read. To make an Update or
	// Delete conditional, pass the last-read value in that request's `etag`
	// field. In REST responses, this value is a base64 string; URL-encode it
	// when setting the `etag` query parameter.
	Etag string `json:"etag,omitempty"`
	// Time of the most recent successful FinalizeSkill. Unset until one
	// succeeds.
	FinalizeTime *time.Time `json:"finalize_time,omitempty"`
	// Unique identifier of the metastore that contains the skill.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Resource name of the skill. Format: `skills/{catalog}.{schema}.{skill}`.
	// Each `{...}` component is capped at 255 characters individually. The
	// server derives the name on Create from `parent` and `skill_id`; it is
	// required and immutable on Update, Get, and Delete.
	Name string `json:"name,omitempty"`
	// Time of the most recent Skill metadata mutation. Uploading bundle files
	// alone does not change this value.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Username or service principal application ID of the principal that last
	// updated the skill.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Skill) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Skill) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SpecialDestination string

const SpecialDestinationSpecialDestinationCatalogOwner SpecialDestination = `SPECIAL_DESTINATION_CATALOG_OWNER`

const SpecialDestinationSpecialDestinationConnectionOwner SpecialDestination = `SPECIAL_DESTINATION_CONNECTION_OWNER`

const SpecialDestinationSpecialDestinationCredentialOwner SpecialDestination = `SPECIAL_DESTINATION_CREDENTIAL_OWNER`

const SpecialDestinationSpecialDestinationExternalLocationOwner SpecialDestination = `SPECIAL_DESTINATION_EXTERNAL_LOCATION_OWNER`

const SpecialDestinationSpecialDestinationFunctionOwner SpecialDestination = `SPECIAL_DESTINATION_FUNCTION_OWNER`

const SpecialDestinationSpecialDestinationMetastoreOwner SpecialDestination = `SPECIAL_DESTINATION_METASTORE_OWNER`

const SpecialDestinationSpecialDestinationRegisteredModelOwner SpecialDestination = `SPECIAL_DESTINATION_REGISTERED_MODEL_OWNER`

const SpecialDestinationSpecialDestinationSchemaOwner SpecialDestination = `SPECIAL_DESTINATION_SCHEMA_OWNER`

const SpecialDestinationSpecialDestinationTableOwner SpecialDestination = `SPECIAL_DESTINATION_TABLE_OWNER`

const SpecialDestinationSpecialDestinationVolumeOwner SpecialDestination = `SPECIAL_DESTINATION_VOLUME_OWNER`

// String representation for [fmt.Print]
func (f *SpecialDestination) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SpecialDestination) Set(v string) error {
	switch v {
	case `SPECIAL_DESTINATION_CATALOG_OWNER`, `SPECIAL_DESTINATION_CONNECTION_OWNER`, `SPECIAL_DESTINATION_CREDENTIAL_OWNER`, `SPECIAL_DESTINATION_EXTERNAL_LOCATION_OWNER`, `SPECIAL_DESTINATION_FUNCTION_OWNER`, `SPECIAL_DESTINATION_METASTORE_OWNER`, `SPECIAL_DESTINATION_REGISTERED_MODEL_OWNER`, `SPECIAL_DESTINATION_SCHEMA_OWNER`, `SPECIAL_DESTINATION_TABLE_OWNER`, `SPECIAL_DESTINATION_VOLUME_OWNER`:
		*f = SpecialDestination(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SPECIAL_DESTINATION_CATALOG_OWNER", "SPECIAL_DESTINATION_CONNECTION_OWNER", "SPECIAL_DESTINATION_CREDENTIAL_OWNER", "SPECIAL_DESTINATION_EXTERNAL_LOCATION_OWNER", "SPECIAL_DESTINATION_FUNCTION_OWNER", "SPECIAL_DESTINATION_METASTORE_OWNER", "SPECIAL_DESTINATION_REGISTERED_MODEL_OWNER", "SPECIAL_DESTINATION_SCHEMA_OWNER", "SPECIAL_DESTINATION_TABLE_OWNER", "SPECIAL_DESTINATION_VOLUME_OWNER"`, v)
	}
}

// Values returns all possible values for SpecialDestination.
//
// There is no guarantee on the order of the values in the slice.
func (f *SpecialDestination) Values() []SpecialDestination {
	return []SpecialDestination{
		SpecialDestinationSpecialDestinationCatalogOwner,
		SpecialDestinationSpecialDestinationConnectionOwner,
		SpecialDestinationSpecialDestinationCredentialOwner,
		SpecialDestinationSpecialDestinationExternalLocationOwner,
		SpecialDestinationSpecialDestinationFunctionOwner,
		SpecialDestinationSpecialDestinationMetastoreOwner,
		SpecialDestinationSpecialDestinationRegisteredModelOwner,
		SpecialDestinationSpecialDestinationSchemaOwner,
		SpecialDestinationSpecialDestinationTableOwner,
		SpecialDestinationSpecialDestinationVolumeOwner,
	}
}

// Type always returns SpecialDestination to satisfy [pflag.Value] interface
func (f *SpecialDestination) Type() string {
	return "SpecialDestination"
}

// Server-Side Encryption properties for clients communicating with AWS s3.
type SseEncryptionDetails struct {
	// Sets the value of the 'x-amz-server-side-encryption' header in S3
	// request.
	Algorithm SseEncryptionDetailsAlgorithm `json:"algorithm,omitempty"`
	// Optional. The ARN of the SSE-KMS key used with the S3 location, when
	// algorithm = "SSE-KMS". Sets the value of the
	// 'x-amz-server-side-encryption-aws-kms-key-id' header.
	AwsKmsKeyArn string `json:"aws_kms_key_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SseEncryptionDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SseEncryptionDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Values returns all possible values for SseEncryptionDetailsAlgorithm.
//
// There is no guarantee on the order of the values in the slice.
func (f *SseEncryptionDetailsAlgorithm) Values() []SseEncryptionDetailsAlgorithm {
	return []SseEncryptionDetailsAlgorithm{
		SseEncryptionDetailsAlgorithmAwsSseKms,
		SseEncryptionDetailsAlgorithmAwsSseS3,
	}
}

// Type always returns SseEncryptionDetailsAlgorithm to satisfy [pflag.Value] interface
func (f *SseEncryptionDetailsAlgorithm) Type() string {
	return "SseEncryptionDetailsAlgorithm"
}

// A storage credential represents an authentication and authorization mechanism
// for accessing data stored on your cloud tenant. It holds exactly one cloud
// credential, for example an AWS IAM role, an Azure managed identity or service
// principal, or a GCP service account. Only the field for the selected
// credential type is set in requests and responses.
type StorageCredentialInfo struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleResponse `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityResponse `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// Time at which this credential was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountResponse `json:"databricks_gcp_service_account,omitempty"`
	// The full name of the credential.
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the credential.
	Id string `json:"id,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of the parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name string `json:"name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Whether this credential is the current metastore's root storage
	// credential. Only applicable when purpose is **STORAGE**.
	UsedForManagedStorage bool `json:"used_for_managed_storage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StorageCredentialInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StorageCredentialInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A system schema of a metastore, with its current enablement state. System
// schemas hold Databricks-provided system tables, such as audit logs and
// billing usage. They live in the `system` catalog unless they are enabled in
// another supported catalog. See
// https://docs.databricks.com/aws/en/admin/system-tables/.
type SystemSchemaInfo struct {
	// Name of the system schema, without the catalog prefix.
	Schema string `json:"schema"`
	// The current enablement state of the system schema. Possible values are
	// `AVAILABLE` (the schema can be enabled), `ENABLE_INITIALIZED` (enablement
	// is in progress), `ENABLE_COMPLETED` (the schema is enabled),
	// `DISABLE_INITIALIZED` (disablement is in progress), `UNAVAILABLE` (the
	// schema cannot be enabled at the moment) and `MANAGED` (Databricks enables
	// the schema automatically).
	State string `json:"state"`
}

func (s *SystemSchemaInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SystemType string

const SystemTypeAmazonRedshift SystemType = `AMAZON_REDSHIFT`

const SystemTypeAzureSynapse SystemType = `AZURE_SYNAPSE`

const SystemTypeConfluent SystemType = `CONFLUENT`

const SystemTypeDatabricks SystemType = `DATABRICKS`

const SystemTypeGoogleBigquery SystemType = `GOOGLE_BIGQUERY`

const SystemTypeKafka SystemType = `KAFKA`

const SystemTypeLooker SystemType = `LOOKER`

const SystemTypeMicrosoftFabric SystemType = `MICROSOFT_FABRIC`

const SystemTypeMicrosoftSqlServer SystemType = `MICROSOFT_SQL_SERVER`

const SystemTypeMongodb SystemType = `MONGODB`

const SystemTypeMysql SystemType = `MYSQL`

const SystemTypeOracle SystemType = `ORACLE`

const SystemTypeOther SystemType = `OTHER`

const SystemTypePostgresql SystemType = `POSTGRESQL`

const SystemTypePowerBi SystemType = `POWER_BI`

const SystemTypeSalesforce SystemType = `SALESFORCE`

const SystemTypeSap SystemType = `SAP`

const SystemTypeServicenow SystemType = `SERVICENOW`

const SystemTypeSnowflake SystemType = `SNOWFLAKE`

const SystemTypeStreamNative SystemType = `STREAM_NATIVE`

const SystemTypeTableau SystemType = `TABLEAU`

const SystemTypeTeradata SystemType = `TERADATA`

const SystemTypeWorkday SystemType = `WORKDAY`

// String representation for [fmt.Print]
func (f *SystemType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SystemType) Set(v string) error {
	switch v {
	case `AMAZON_REDSHIFT`, `AZURE_SYNAPSE`, `CONFLUENT`, `DATABRICKS`, `GOOGLE_BIGQUERY`, `KAFKA`, `LOOKER`, `MICROSOFT_FABRIC`, `MICROSOFT_SQL_SERVER`, `MONGODB`, `MYSQL`, `ORACLE`, `OTHER`, `POSTGRESQL`, `POWER_BI`, `SALESFORCE`, `SAP`, `SERVICENOW`, `SNOWFLAKE`, `STREAM_NATIVE`, `TABLEAU`, `TERADATA`, `WORKDAY`:
		*f = SystemType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AMAZON_REDSHIFT", "AZURE_SYNAPSE", "CONFLUENT", "DATABRICKS", "GOOGLE_BIGQUERY", "KAFKA", "LOOKER", "MICROSOFT_FABRIC", "MICROSOFT_SQL_SERVER", "MONGODB", "MYSQL", "ORACLE", "OTHER", "POSTGRESQL", "POWER_BI", "SALESFORCE", "SAP", "SERVICENOW", "SNOWFLAKE", "STREAM_NATIVE", "TABLEAU", "TERADATA", "WORKDAY"`, v)
	}
}

// Values returns all possible values for SystemType.
//
// There is no guarantee on the order of the values in the slice.
func (f *SystemType) Values() []SystemType {
	return []SystemType{
		SystemTypeAmazonRedshift,
		SystemTypeAzureSynapse,
		SystemTypeConfluent,
		SystemTypeDatabricks,
		SystemTypeGoogleBigquery,
		SystemTypeKafka,
		SystemTypeLooker,
		SystemTypeMicrosoftFabric,
		SystemTypeMicrosoftSqlServer,
		SystemTypeMongodb,
		SystemTypeMysql,
		SystemTypeOracle,
		SystemTypeOther,
		SystemTypePostgresql,
		SystemTypePowerBi,
		SystemTypeSalesforce,
		SystemTypeSap,
		SystemTypeServicenow,
		SystemTypeSnowflake,
		SystemTypeStreamNative,
		SystemTypeTableau,
		SystemTypeTeradata,
		SystemTypeWorkday,
	}
}

// Type always returns SystemType to satisfy [pflag.Value] interface
func (f *SystemType) Type() string {
	return "SystemType"
}

// A table constraint, as defined by *one* of the following fields being set:
// __primary_key_constraint__, __foreign_key_constraint__,
// __named_table_constraint__.
type TableConstraint struct {
	ForeignKeyConstraint *ForeignKeyConstraint `json:"foreign_key_constraint,omitempty"`

	NamedTableConstraint *NamedTableConstraint `json:"named_table_constraint,omitempty"`

	PrimaryKeyConstraint *PrimaryKeyConstraint `json:"primary_key_constraint,omitempty"`
}

func (s *TableConstraint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A table that is dependent on a SQL object.
type TableDependency struct {
	// Full name of the dependent table, in the form of
	// __catalog_name__.__schema_name__.__table_name__.
	TableFullName string `json:"table_full_name"`
}

func (s *TableDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type TableExistsResponse struct {
	// Whether the table exists or not.
	TableExists bool `json:"table_exists,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableExistsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableExistsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A table resides in the third layer of Unity Catalog's three-level namespace
// and holds rows of data organized in columns. Besides managed and external
// tables, it covers views, materialized views, streaming tables, metric views,
// foreign tables and shallow clones, which table_type distinguishes.
type TableInfo struct {
	// The AWS access point to use when accessing S3 for the storage location of
	// this table.
	AccessPoint string `json:"access_point,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Name of parent catalog. Required for creation using catalog and schema
	// names.
	CatalogName string `json:"catalog_name,omitempty"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns []ColumnInfo `json:"columns,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this table was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of table creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique ID of the Data Access Configuration to use with the table data.
	DataAccessConfigurationId string `json:"data_access_configuration_id,omitempty"`
	// The format of the data that backs the table, such as DELTA, PARQUET, CSV
	// or JSON. Foreign tables report the external system that holds them, for
	// example SNOWFLAKE_FORMAT or MYSQL_FORMAT.
	DataSourceFormat DataSourceFormat `json:"data_source_format,omitempty"`
	// Time at which this table was deleted, in epoch milliseconds. Field is
	// omitted if table is not deleted.
	DeletedAt int64 `json:"deleted_at,omitempty"`
	// Information pertaining to current state of the delta table.
	DeltaRuntimePropertiesKvpairs *DeltaRuntimePropertiesKvPairs `json:"delta_runtime_properties_kvpairs,omitempty"`
	// The predictive optimization setting that is in effect for the table once
	// inherited values are resolved.
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `json:"effective_predictive_optimization_flag,omitempty"`
	// Whether predictive optimization is enabled for this table. Valid values
	// are ENABLE, DISABLE and INHERIT, where INHERIT applies the setting of the
	// parent schema. It is ignored when creating a table. See
	// https://docs.databricks.com/aws/en/optimizations/predictive-optimization
	// for details.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`

	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of table, relative to parent schema.
	Name string `json:"name,omitempty"`
	// The user or group that owns the table. The principal that creates the
	// table becomes its owner.
	Owner string `json:"owner,omitempty"`
	// The pipeline ID of the table. Applicable for tables created by pipelines
	// (Materialized View, Streaming Table, etc.).
	PipelineId string `json:"pipeline_id,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// The row filter function applied to the table, if any. A row filter is a
	// SQL user-defined function that decides which rows a query returns. See
	// https://docs.databricks.com/aws/en/tables/row-and-column-filters for
	// details.
	RowFilter *TableRowFilter `json:"row_filter,omitempty"`
	// Name of parent schema relative to its parent catalog. Required for
	// creation using catalog and schema names.
	SchemaName string `json:"schema_name,omitempty"`
	// SecurableKindManifest of table, including capabilities the table has.
	SecurableKindManifest *SecurableKindManifest `json:"securable_kind_manifest,omitempty"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// Name of the storage credential, when a storage credential is configured
	// for use with this table.
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// The cloud storage URL under which the data files of the table are stored.
	// For an EXTERNAL table, it must be set on creation to a path within an
	// existing external location. For a MANAGED table, the server assigns a
	// path under the managed storage location of the parent schema, catalog or
	// metastore and returns it in this field. Views have no storage location.
	StorageLocation string `json:"storage_location,omitempty"`
	// List of table constraints. Note: this field is not set in the output of
	// the __listTables__ API.
	TableConstraints []TableConstraint `json:"table_constraints,omitempty"`
	// The unique identifier of the table.
	TableId string `json:"table_id,omitempty"`
	// The type of the table. MANAGED tables store their data in storage that
	// Unity Catalog manages, EXTERNAL tables store it in a registered external
	// location, FOREIGN tables mirror objects of an external database through a
	// connection, and the shallow clone types reference the data files of a
	// base table. VIEW, MATERIALIZED_VIEW, STREAMING_TABLE and METRIC_VIEW are
	// defined by a view definition. See
	// https://docs.databricks.com/aws/en/tables/ for an overview of the table
	// types.
	TableType TableType `json:"table_type,omitempty"`
	// Time at which this table was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// View definition SQL (when __table_type__ is **VIEW**,
	// **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
	ViewDefinition string `json:"view_definition,omitempty"`
	// View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**,
	// **STREAMING_TABLE**) - when DependencyList is None, the dependency is not
	// provided; - when DependencyList is an empty list, the dependency is
	// provided but is empty; - when DependencyList is not an empty list,
	// dependencies are provided and recorded. Note: this field is not set in
	// the output of the __listTables__ API.
	ViewDependencies *DependencyList `json:"view_dependencies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableOperation string

const TableOperationRead TableOperation = `READ`

const TableOperationReadWrite TableOperation = `READ_WRITE`

// String representation for [fmt.Print]
func (f *TableOperation) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableOperation) Set(v string) error {
	switch v {
	case `READ`, `READ_WRITE`:
		*f = TableOperation(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "READ", "READ_WRITE"`, v)
	}
}

// Values returns all possible values for TableOperation.
//
// There is no guarantee on the order of the values in the slice.
func (f *TableOperation) Values() []TableOperation {
	return []TableOperation{
		TableOperationRead,
		TableOperationReadWrite,
	}
}

// Type always returns TableOperation to satisfy [pflag.Value] interface
func (f *TableOperation) Type() string {
	return "TableOperation"
}

type TableRowFilter struct {
	// The full name of the row filter SQL UDF.
	FunctionName string `json:"function_name"`
	// The list of table columns or literals to be passed as additional
	// arguments to a row filter function, carrying the type (column reference
	// vs constant literal) of each argument. Deprecated: use input_column_names
	// instead.
	InputArguments []PolicyFunctionArgument `json:"input_arguments,omitempty"`
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	InputColumnNames []string `json:"input_column_names"`
}

func (s *TableRowFilter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type TableSummary struct {
	// The full name of the table.
	FullName string `json:"full_name,omitempty"`
	// SecurableKindManifest of table, including capabilities the table has.
	SecurableKindManifest *SecurableKindManifest `json:"securable_kind_manifest,omitempty"`
	// The type of the table, with the same meaning as table_type on TableInfo.
	TableType TableType `json:"table_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableSummary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableSummary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableType string

const TableTypeExternal TableType = `EXTERNAL`

const TableTypeExternalShallowClone TableType = `EXTERNAL_SHALLOW_CLONE`

const TableTypeForeign TableType = `FOREIGN`

const TableTypeManaged TableType = `MANAGED`

const TableTypeManagedShallowClone TableType = `MANAGED_SHALLOW_CLONE`

const TableTypeMaterializedView TableType = `MATERIALIZED_VIEW`

const TableTypeMetricView TableType = `METRIC_VIEW`

const TableTypeStreamingTable TableType = `STREAMING_TABLE`

const TableTypeView TableType = `VIEW`

// String representation for [fmt.Print]
func (f *TableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableType) Set(v string) error {
	switch v {
	case `EXTERNAL`, `EXTERNAL_SHALLOW_CLONE`, `FOREIGN`, `MANAGED`, `MANAGED_SHALLOW_CLONE`, `MATERIALIZED_VIEW`, `METRIC_VIEW`, `STREAMING_TABLE`, `VIEW`:
		*f = TableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "EXTERNAL_SHALLOW_CLONE", "FOREIGN", "MANAGED", "MANAGED_SHALLOW_CLONE", "MATERIALIZED_VIEW", "METRIC_VIEW", "STREAMING_TABLE", "VIEW"`, v)
	}
}

// Values returns all possible values for TableType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TableType) Values() []TableType {
	return []TableType{
		TableTypeExternal,
		TableTypeExternalShallowClone,
		TableTypeForeign,
		TableTypeManaged,
		TableTypeManagedShallowClone,
		TableTypeMaterializedView,
		TableTypeMetricView,
		TableTypeStreamingTable,
		TableTypeView,
	}
}

// Type always returns TableType to satisfy [pflag.Value] interface
func (f *TableType) Type() string {
	return "TableType"
}

// Enum representing the source type of a tag assignment
type TagAssignmentSourceType string

const TagAssignmentSourceTypeTagAssignmentSourceTypeSystemDataClassification TagAssignmentSourceType = `TAG_ASSIGNMENT_SOURCE_TYPE_SYSTEM_DATA_CLASSIFICATION`

// String representation for [fmt.Print]
func (f *TagAssignmentSourceType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TagAssignmentSourceType) Set(v string) error {
	switch v {
	case `TAG_ASSIGNMENT_SOURCE_TYPE_SYSTEM_DATA_CLASSIFICATION`:
		*f = TagAssignmentSourceType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TAG_ASSIGNMENT_SOURCE_TYPE_SYSTEM_DATA_CLASSIFICATION"`, v)
	}
}

// Values returns all possible values for TagAssignmentSourceType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TagAssignmentSourceType) Values() []TagAssignmentSourceType {
	return []TagAssignmentSourceType{
		TagAssignmentSourceTypeTagAssignmentSourceTypeSystemDataClassification,
	}
}

// Type always returns TagAssignmentSourceType to satisfy [pflag.Value] interface
func (f *TagAssignmentSourceType) Type() string {
	return "TagAssignmentSourceType"
}

// An expression that introspects tags at query time.
type TagIntrospectionExpression struct {
	// Extracts the value of a column-level tag.
	ColumnTagValue *ColumnTagValueExtraction `json:"column_tag_value,omitempty"`
	// Extracts the value of a securable-level tag.
	TagValue *TagValueExtraction `json:"tag_value,omitempty"`
}

func (s *TagIntrospectionExpression) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type TagKeyValue struct {
	// name of the tag
	Key string `json:"key,omitempty"`
	// value of the tag associated with the key, could be optional
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TagKeyValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TagKeyValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Extracts the value of a securable-level tag: get_tag_value("tagKey").
type TagValueExtraction struct {
	// 1024 matches the max_length on FunctionArgument.constant above.
	TagKey string `json:"tag_key"`
}

func (s *TagValueExtraction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type TemporaryCredentials struct {
	// Temporary AWS credentials that can be used to sign AWS API requests.
	AwsTempCredentials *AwsCredentials `json:"aws_temp_credentials,omitempty"`
	// A Microsoft Entra ID (formerly Azure Active Directory) token that is
	// issued for the credential's identity.
	AzureAad *AzureActiveDirectoryToken `json:"azure_aad,omitempty"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// A GCP OAuth 2.0 access token that is issued for the credential's service
	// account.
	GcpOauthToken *GcpOauthToken `json:"gcp_oauth_token,omitempty"`
	// Temporary Cloudflare R2 credentials that can be used with the R2
	// S3-compatible API.
	R2TempCredentials *R2Credentials `json:"r2_temp_credentials,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TemporaryCredentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TemporaryCredentials) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
type TriggeredUpdateStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp string `json:"timestamp,omitempty"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress *PipelineProgress `json:"triggered_update_progress,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TriggeredUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TriggeredUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UnassignRequest struct {
	// Query for the ID of the metastore to delete.
	MetastoreId string `json:"-" url:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *UnassignRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to restore a soft-deleted metastore by ID.
type UndropMetastoreByIdRequest struct {
	// Required. Id of the metastore to undrop.
	Id string `json:"id"`
}

func (s *UndropMetastoreByIdRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateAccessRequestDestinationsRequest struct {
	// The access request destinations to assign to the securable. For each
	// destination, a **destination_id** and **destination_type** must be
	// defined.
	AccessRequestDestinations AccessRequestDestinations `json:"access_request_destinations"`
	// The paths of the field mask are not evaluated, because an update always
	// replaces the full list of destinations.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateAccessRequestDestinationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateAccountsMetastore struct {
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope DeltaSharingScopeEnum `json:"delta_sharing_scope,omitempty"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled bool `json:"external_access_enabled,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateAccountsMetastore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateAccountsMetastore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A storage credential represents an authentication and authorization mechanism
// for accessing data stored on your cloud tenant. It holds exactly one cloud
// credential, for example an AWS IAM role, an Azure managed identity or service
// principal, or a GCP service account. Only the field for the selected
// credential type is set in requests and responses.
type UpdateAccountsStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleRequest `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityResponse `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateAccountsStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateAccountsStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateCatalog struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Dropped objects in the catalog remain recoverable for this number of
	// hours before they are permanently purged. The value applies to standard
	// catalogs only, must be 0 or within the range allowed for the metastore,
	// and is inherited by schemas and tables that do not set their own value.
	CustomMaxRetentionHours int64 `json:"custom_max_retention_hours,omitempty"`
	// Whether predictive optimization is enabled for this catalog and the
	// objects under it. Valid values are ENABLE, DISABLE, and INHERIT, where
	// INHERIT applies the account-level setting. Standard catalogs start as
	// INHERIT and other catalog types as DISABLE. It cannot be set when
	// creating a catalog; use the update method instead.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Whether the catalog is accessible from all workspaces attached to the
	// metastore (OPEN) or only from the workspaces bound to it (ISOLATED). It
	// cannot be set when creating a catalog; change it with the update method
	// and manage the bound workspaces with the workspace bindings API. See
	// [Workspace-catalog binding].
	//
	// [Workspace-catalog binding]: https://docs.databricks.com/aws/en/catalogs/binding
	IsolationMode CatalogIsolationMode `json:"isolation_mode,omitempty"`
	// Customer-managed key (CMK) settings control the encryption at rest of the
	// catalog's managed data.
	ManagedEncryptionSettings *EncryptionSettings `json:"managed_encryption_settings,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`
	// New name for the catalog.
	NewName string `json:"new_name,omitempty"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options,omitempty"`
	// Username of current owner of catalog.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateCatalogWorkspaceBindingsResponse struct {
	// The IDs of the workspaces that are bound to the catalog.
	Workspaces []int64 `json:"workspaces,omitempty"`
}

func (s *UpdateCatalogWorkspaceBindingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateConnection struct {
	// Environment settings for the connection, such as the Java libraries that
	// it needs and the environment version.
	EnvironmentSettings *EnvironmentSettings `json:"environment_settings,omitempty"`
	// Name of the connection.
	Name string `json:"-" url:"-"`
	// New name for the connection.
	NewName string `json:"new_name,omitempty"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options"`
	// Username of the current owner of the connection. The creator becomes the
	// owner at creation, and ownership can be transferred on update.
	Owner string `json:"owner,omitempty"`
	// A map of option names to UC Secret references. Keys are connection option
	// names (same as in OptionsKVPairs) and values are UC Secret fully
	// qualified names.
	Secrets map[string]string `json:"secrets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateConnection) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateCredentialRequest struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentity `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force bool `json:"force,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Name of the credential.
	NameArg string `json:"-" url:"-"`
	// New name of credential.
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`
	// Supply true to this argument to skip validation of the updated
	// credential.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateEntityTagAssignmentRequest struct {
	// The fully qualified name of the entity to which the tag is assigned
	EntityName string `json:"-" url:"-"`
	// The type of the entity to which the tag is assigned.
	EntityType string `json:"-" url:"-"`

	TagAssignment EntityTagAssignment `json:"tag_assignment"`
	// The key of the tag
	TagKey string `json:"-" url:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateEntityTagAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateExternalLineageRelationshipRequest struct {
	ExternalLineageRelationship UpdateRequestExternalLineage `json:"external_lineage_relationship"`
	// The fields of the relationship to update. The field mask must be a single
	// string, with multiple fields separated by commas (no spaces), for example
	// `columns,properties`. The `source` and `target` identify the relationship
	// and cannot be changed.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateExternalLineageRelationshipRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateExternalLocation struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the storage credential used with this location.
	CredentialName string `json:"credential_name,omitempty"`
	// The effective value of `enable_file_events` after applying server-side
	// defaults.
	EffectiveEnableFileEvents bool `json:"effective_enable_file_events,omitempty"`
	// The effective file event queue configuration after applying server-side
	// defaults. Always populated when a queue is provisioned, regardless of
	// whether the user explicitly set `enable_file_events`. Use this field
	// instead of `file_event_queue` for reading the actual queue state.
	EffectiveFileEventQueue *FileEventQueue `json:"effective_file_event_queue,omitempty"`
	// Whether to enable file events on this external location. Defaults to
	// `true`. Set to `false` to disable file events. The actual applied value
	// may differ due to server-side defaults; check
	// `effective_enable_file_events` for the effective state.
	EnableFileEvents bool `json:"enable_file_events,omitempty"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback bool `json:"fallback,omitempty"`
	// File event queue settings. When `enable_file_events` is `true`, exactly
	// one of the queue options must be set. Must be omitted when
	// `enable_file_events` is `false`.
	FileEventQueue *FileEventQueue `json:"file_event_queue,omitempty"`
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	Force bool `json:"force,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces. Cannot be set when creating an external
	// location.
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Name of the external location.
	Name string `json:"-" url:"-"`
	// New name for the external location.
	NewName string `json:"new_name,omitempty"`
	// The owner of the external location.
	Owner string `json:"owner,omitempty"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// URL of the cloud storage path that the external location points to,
	// including the scheme (for example `s3://`, `abfss://`, or `gs://`).
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateExternalLocation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateExternalLocation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateExternalMetadataRequest struct {
	ExternalMetadata ExternalMetadata `json:"external_metadata"`
	// Name of the external metadata object. The name is unique within the
	// metastore, must follow the [Unity Catalog naming requirements], and is
	// stored in lowercase. It cannot be changed after creation, and external
	// lineage relationships refer to the object by this name.
	//
	// [Unity Catalog naming requirements]: https://docs.databricks.com/aws/en/data-governance/unity-catalog/requirements#naming-requirements
	Name string `json:"-" url:"-"`
	// Comma-separated list of the fields to update, for example
	// `description,properties`. Only the listed fields change. The other fields
	// keep their current values. The mask must not be empty.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateExternalMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateFunction struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function_name__).
	Name string `json:"-" url:"-"`
	// Username of current owner of the function.
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateMcpServiceRequest struct {
	// Optimistic concurrency token from the most recent read. When set, the
	// update succeeds only if the resource has not changed. Leave unset for an
	// unconditional update. For REST requests, URL-encode the base64 string
	// returned by the API when setting the `etag` query parameter.
	Etag string `json:"-" url:"etag,omitempty"`
	// The MCP service with the updated field values. `name` identifies the
	// resource (`mcp-services/{catalog}.{schema}.{mcp_service}`); only fields
	// listed in `update_mask` are applied.
	McpService McpService `json:"mcp_service"`
	// Resource name of the MCP service. Format:
	// `mcp-services/{catalog}.{schema}.{mcp_service}`. Each `{...}` component
	// is capped at 255 characters individually. The server derives the name on
	// Create from `parent` and `mcp_service_id`; it is required and immutable
	// on Update, Get, and Delete.
	Name string `json:"-" url:"-"`
	// Fields to update. Use `config` to replace the entire configuration. The
	// replacement must include every required field; any optional field you
	// omit is cleared. To preserve sibling fields, use one or more granular
	// paths: `comment`, `config.source_connection.name`,
	// `config.include_tool_selectors`, or `config.rate_limits`.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateMcpServiceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMcpServiceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateMetastore struct {
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope DeltaSharingScopeEnum `json:"delta_sharing_scope,omitempty"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled bool `json:"external_access_enabled,omitempty"`
	// Unique ID of the metastore.
	Id string `json:"-" url:"-"`
	// New name for the metastore.
	NewName string `json:"new_name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateMetastore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMetastore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// deprecated. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// The unique ID of the metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateMetastoreAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMetastoreAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateModelProviderServiceRequest struct {
	// Optimistic concurrency token from the most recent read. When set, the
	// update succeeds only if the resource has not changed. Leave unset for an
	// unconditional update. For REST requests, URL-encode the base64 string
	// returned by the API when setting the `etag` query parameter.
	Etag string `json:"-" url:"etag,omitempty"`
	// The model provider service with the updated field values. `name`
	// identifies the resource
	// (`model-provider-services/{catalog}.{schema}.{model_provider_service}`);
	// only fields listed in `update_mask` are applied.
	ModelProviderService ModelProviderService `json:"model_provider_service"`
	// Resource name of the provider service. Format:
	// `model-provider-services/{catalog}.{schema}.{model_provider_service}`.
	// Each `{...}` component is capped at 255 characters individually. The
	// server derives the name on Create from `parent` and
	// `model_provider_service_id`; it is required and immutable on Update, Get,
	// and Delete.
	Name string `json:"-" url:"-"`
	// Fields to update. Use `config` to replace the entire configuration. The
	// replacement must include every required field; any optional field you
	// omit is cleared. To preserve sibling fields, use one or more granular
	// paths: `comment`; `config.provider` to replace the active
	// provider-specific value (for example, `config.openai`; the mask path
	// remains `config.provider`); `config.allow_all_targets`, `config.targets`,
	// `config.forward_headers`, `config.forward_query_parameters`,
	// `config.forward_unmanaged_paths`, `config.rate_limits`,
	// `config.pricing.default_discount_basis_points`, or
	// `config.inference_table`. For a pricing update, an explicitly empty
	// `pricing` object clears the discount but retains the object; omitting
	// `pricing` clears it. `config.pricing` is not a supported mask path. The
	// provider type is immutable. A `config` or `config.provider` replacement
	// that carries no authentication material preserves the existing
	// authentication binding; input-only plaintext does not need to be read
	// back and re-sent.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateModelProviderServiceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateModelProviderServiceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateModelServiceRequest struct {
	// Optimistic concurrency token from the most recent read. When set, the
	// update succeeds only if the resource has not changed. Leave unset for an
	// unconditional update. For REST requests, URL-encode the base64 string
	// returned by the API when setting the `etag` query parameter.
	Etag string `json:"-" url:"etag,omitempty"`
	// The model service with the updated field values. `name` identifies the
	// resource (`model-services/{catalog}.{schema}.{model_service}`); only
	// fields listed in `update_mask` are applied.
	ModelService ModelService `json:"model_service"`
	// Resource name of the model service. Format:
	// `model-services/{catalog}.{schema}.{model_service}`. Each `{...}`
	// component is capped at 255 characters individually. The server derives
	// the name on Create from `parent` and `model_service_id`; it is required
	// and immutable on Update, Get, and Delete.
	Name string `json:"-" url:"-"`
	// Fields to update. Use `config` for a broad configuration replacement. The
	// replacement must include every required field; optional fields in its
	// replacement scope are cleared when omitted. To preserve sibling fields,
	// use one or more granular paths: `comment`, `config.routing.destinations`,
	// `config.routing.fallback.destinations`, `config.rate_limits`, or
	// `config.inference_table`. Intermediate paths such as `config.routing` and
	// `config.routing.fallback` are not supported.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateModelServiceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateModelServiceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateModelVersionRequest struct {
	// List of aliases associated with the model version
	Aliases []RegisteredModelAlias `json:"aliases,omitempty"`
	// The name of the catalog containing the model version
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the model version
	Comment string `json:"comment,omitempty"`
	// Time at which the model version was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the model version
	CreatedBy string `json:"created_by,omitempty"`
	// The three-level (fully qualified) name of the parent registered model, in
	// the form `catalog_name.schema_name.model_name`.
	FullName string `json:"-" url:"-"`
	// The unique identifier of the model version
	Id string `json:"id,omitempty"`
	// The unique identifier of the metastore containing the model version
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the parent registered model of the model version, relative to
	// parent schema
	ModelName string `json:"model_name,omitempty"`
	// Model version dependencies, for feature-store packaged models
	ModelVersionDependencies *DependencyList `json:"model_version_dependencies,omitempty"`
	// MLflow run ID used when creating the model version, if ``source`` was
	// generated by an experiment run stored in an MLflow tracking server
	RunId string `json:"run_id,omitempty"`
	// ID of the Databricks workspace containing the MLflow run that generated
	// this model version, if applicable
	RunWorkspaceId int `json:"run_workspace_id,omitempty"`
	// The name of the schema containing the model version, relative to parent
	// catalog
	SchemaName string `json:"schema_name,omitempty"`
	// URI indicating the location of the source artifacts (files) for the model
	// version
	Source string `json:"source,omitempty"`
	// Current status of the model version. Newly created model versions start
	// in PENDING_REGISTRATION status, then move to READY status once the model
	// version files are uploaded and the model version is finalized. Only model
	// versions in READY status can be loaded for inference or served.
	Status ModelVersionInfoStatus `json:"status,omitempty"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `json:"storage_location,omitempty"`
	// Time at which the model version was last updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the model version last time
	UpdatedBy string `json:"updated_by,omitempty"`
	// The integer version number of the model version
	Version int `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateModelVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateMonitor struct {
	// [Create:OPT Update:OPT] Baseline table name. Baseline data is used to
	// compute drift from the data in the monitored `table_name`. The baseline
	// table and the monitored table shall have the same schema.
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// [Create:OPT Update:OPT] Custom metrics.
	CustomMetrics []MonitorMetric `json:"custom_metrics,omitempty"`
	// [Create:ERR Update:OPT] Id of dashboard that visualizes the computed
	// metrics. This can be empty if the monitor is in PENDING state.
	DashboardId string `json:"dashboard_id,omitempty"`
	// [Create:OPT Update:OPT] Data classification related config.
	DataClassificationConfig *MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`

	InferenceLog *MonitorInferenceLog `json:"inference_log,omitempty"`
	// [Create:ERR Update:IGN] The latest error message for a monitor failure.
	LatestMonitorFailureMsg string `json:"latest_monitor_failure_msg,omitempty"`
	// [Create:OPT Update:OPT] Field for specifying notification settings.
	Notifications *MonitorNotifications `json:"notifications,omitempty"`
	// [Create:REQ Update:REQ] Schema where output tables are created. Needs to
	// be in 2-level format {catalog}.{schema}
	OutputSchemaName string `json:"output_schema_name"`
	// [Create:OPT Update:OPT] The monitor schedule.
	Schedule *MonitorCronSchedule `json:"schedule,omitempty"`
	// [Create:OPT Update:OPT] List of column expressions to slice data with for
	// targeted analysis. The data is grouped by each expression independently,
	// resulting in a separate slice for each predicate and its complements. For
	// example `slicing_exprs=[“col_1”, “col_2 > 10”]` will generate the
	// following slices: two slices for `col_2 > 10` (True and False), and one
	// slice per unique value in `col1`. For high-cardinality columns, only the
	// top 100 unique values by frequency will generate slices.
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	Snapshot *MonitorSnapshot `json:"snapshot,omitempty"`
	// UC table name in format `catalog.schema.table_name`. This field
	// corresponds to the {full_table_name_arg} arg in the endpoint path.
	TableName string `json:"-" url:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries *MonitorTimeSeries `json:"time_series,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateMonitor) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMonitor) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdatePermissions struct {
	// Array of permissions change objects.
	Changes []PermissionsChange `json:"changes,omitempty"`
	// Full name of the securable. Tables, views, volumes and functions use the
	// three-level name `catalog.schema.object`; a metastore uses its metastore
	// ID.
	FullName string `json:"-" url:"-"`
	// Whether to omit the privilege assignments from the response. When false,
	// the response lists all direct privilege assignments on the securable
	// after the update.
	OmitPermissionsInResponse bool `json:"omit_permissions_in_response,omitempty"`
	// Type of the securable, such as `catalog`, `schema`, `table` or `volume`.
	// The value is case-insensitive.
	SecurableType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdatePermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdatePermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdatePermissionsResponse struct {
	// The privileges assigned to each principal
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

func (s *UpdatePermissionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdatePolicyRequest struct {
	// Required. The name of the policy to update.
	Name string `json:"-" url:"-"`
	// Required. The fully qualified name of the securable to update the policy
	// for.
	OnSecurableFullname string `json:"-" url:"-"`
	// Required. The type of the securable to update the policy for.
	OnSecurableType string `json:"-" url:"-"`
	// Optional fields to update. This is the request body for updating a
	// policy. Use `update_mask` field to specify which fields in the request is
	// to be updated. - If `update_mask` is empty or "*", all specified fields
	// will be updated. - If `update_mask` is specified, only the fields
	// specified in the `update_mask` will be updated. If a field is specified
	// in `update_mask` and not set in the request, the field will be cleared.
	// Users can use the update mask to explicitly unset optional fields such as
	// `exception_principals` and `when_condition`.
	PolicyInfo PolicyInfo `json:"policy_info"`
	// Optional. The update mask field for specifying user intentions on which
	// fields to update in the request.
	UpdateMask string `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdatePolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdatePolicyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateRegisteredModelRequest struct {
	// List of aliases associated with the registered model
	Aliases []RegisteredModelAlias `json:"aliases,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog where the schema and the registered model reside
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the registered model
	Comment string `json:"comment,omitempty"`
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the registered model
	CreatedBy string `json:"created_by,omitempty"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
	// The unique identifier of the metastore
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the registered model
	Name string `json:"name,omitempty"`
	// New name for the registered model.
	NewName string `json:"new_name,omitempty"`
	// The identifier of the user who owns the registered model
	Owner string `json:"owner,omitempty"`
	// The name of the schema where the registered model resides
	SchemaName string `json:"schema_name,omitempty"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `json:"storage_location,omitempty"`
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the registered model last time
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A lineage relationship between a source object and a target object, where at
// least one of the two is an external metadata object. The other object can be
// a Unity Catalog table, a cloud storage path, a model version, or another
// external metadata object. A relationship is identified by its source and
// target; optional column mappings and key-value properties describe the data
// flow between them.
type UpdateRequestExternalLineage struct {
	// List of column relationships between source and target objects. The
	// combined size of all source and target column names is limited.
	Columns []ColumnRelationship `json:"columns,omitempty"`
	// Unique identifier of the external lineage relationship.
	Id string `json:"id,omitempty"`
	// Key-value properties associated with the external lineage relationship.
	// The combined length of all keys and values must not exceed 5000
	// characters.
	Properties map[string]string `json:"properties,omitempty"`
	// Source object of the external lineage relationship.
	Source ExternalLineageObject `json:"source"`
	// Target object of the external lineage relationship.
	Target ExternalLineageObject `json:"target"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateRequestExternalLineage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRequestExternalLineage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateSchema struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Dropped objects in the schema remain recoverable for this number of hours
	// before they are permanently purged. The value must be 0 or within the
	// range allowed for the metastore. Schemas that do not set their own value
	// inherit the parent catalog's value.
	CustomMaxRetentionHours int64 `json:"custom_max_retention_hours,omitempty"`
	// Whether predictive optimization is enabled for this schema and the
	// objects under it. Valid values are ENABLE, DISABLE, and INHERIT, where
	// INHERIT applies the parent catalog's setting. New schemas start as
	// INHERIT. It cannot be set when creating a schema; use the update method
	// instead.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Full name of the schema, in the form catalog_name.schema_name.
	FullName string `json:"-" url:"-"`
	// New name for the schema.
	NewName string `json:"new_name,omitempty"`
	// The user or group that owns the schema. The creator becomes the owner,
	// and the owner can be changed only through the update method.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateSchema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateSchema) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateSecretRequest struct {
	// The three-level (fully qualified) name of the secret (for example,
	// **catalog_name.schema_name.secret_name**).
	FullName string `json:"-" url:"-"`
	// The secret object containing the fields to update. Only fields specified
	// in **update_mask** will be updated.
	Secret Secret `json:"secret"`
	// The field mask specifying which fields of the secret to update. - If
	// **update_mask** is **"*"**, all fields specified in **secret** are
	// updated. - If **update_mask** specifies one or more fields, only those
	// fields are updated. Each specified field must be set in **secret**.
	// Supported fields: **value**, **comment**, **owner**, **expire_time**. To
	// change the secret name, delete and recreate the secret.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateSecretRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateSkillRequest struct {
	// Optimistic concurrency token from the most recent read. When set, the
	// update succeeds only if the resource has not changed. Leave unset for an
	// unconditional update. For REST requests, URL-encode the base64 string
	// returned by the API when setting the `etag` query parameter.
	Etag string `json:"-" url:"etag,omitempty"`
	// Resource name of the skill. Format: `skills/{catalog}.{schema}.{skill}`.
	// Each `{...}` component is capped at 255 characters individually. The
	// server derives the name on Create from `parent` and `skill_id`; it is
	// required and immutable on Update, Get, and Delete.
	Name string `json:"-" url:"-"`
	// The skill with the updated field values. `name` identifies the resource
	// (`skills/{catalog}.{schema}.{skill}`); only fields listed in
	// `update_mask` are applied.
	Skill Skill `json:"skill"`
	// Comma-separated list of `skill` fields to update. `comment` is the only
	// field that can be updated.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateSkillRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateSkillRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleRequest `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityResponse `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	// Force update even if there are dependent external locations or external
	// tables.
	Force bool `json:"force,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Name of the storage credential.
	Name string `json:"-" url:"-"`
	// New name for the storage credential.
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`
	// Supplying true to this argument skips validation of the updated
	// credential.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateTableRequest struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`
	// The user or group that owns the table. The principal that creates the
	// table becomes its owner.
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateTableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateVolumeRequestContent struct {
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`
	// The three-level (fully qualified) name of the volume
	Name string `json:"-" url:"-"`
	// The new name for the volume. Set it to rename the volume.
	NewName string `json:"new_name,omitempty"`
	// The user or group that owns the volume. The user who creates a volume
	// becomes its owner; the owner can be changed afterwards with an update.
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateVolumeRequestContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateVolumeRequestContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateWorkspaceBindings struct {
	// The IDs of the workspaces to bind to the catalog.
	AssignWorkspaces []int64 `json:"assign_workspaces,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`
	// The IDs of the workspaces to unbind from the catalog.
	UnassignWorkspaces []int64 `json:"unassign_workspaces,omitempty"`
}

func (s *UpdateWorkspaceBindings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceBindingsParameters struct {
	// List of workspace bindings to add. If a binding for the workspace already
	// exists with a different binding_type, adding it again with a new
	// binding_type will update the existing binding (e.g., from READ_WRITE to
	// READ_ONLY).
	Add []WorkspaceBinding `json:"add,omitempty"`
	// List of workspace bindings to remove.
	Remove []WorkspaceBinding `json:"remove,omitempty"`
	// The name of the securable.
	SecurableName string `json:"-" url:"-"`
	// The type of the securable to bind to a workspace (catalog,
	// storage_credential, credential, or external_location).
	SecurableType string `json:"-" url:"-"`
}

func (s *UpdateWorkspaceBindingsParameters) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A list of workspace IDs that are bound to the securable
type UpdateWorkspaceBindingsResponse struct {
	// List of workspace bindings.
	Bindings []WorkspaceBinding `json:"bindings,omitempty"`
}

func (s *UpdateWorkspaceBindingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Access request URLs shared by every workspace in a metastore.
type UrlConfiguration struct {
	// The URL template for access requests on catalogs and the securables
	// inside them, such as schemas, tables, volumes, functions, and registered
	// models. Placeholders are replaced with percent-encoded values from the
	// requested securable: `{catalog}` with the catalog name (`main`),
	// `{catalog.schema}` with the catalog and schema names (`main.sales`),
	// `{catalog.schema.object}` with the full name (`main.sales.orders`),
	// `{securable_type}` with the securable type (`TABLE`), and
	// `{securable_id}` with the securable ID. Name placeholders leave out
	// levels the securable does not have, so `{catalog.schema.object}` is
	// `main.sales` for a schema. When the deepest name placeholder stops above
	// the requested securable, such as `{catalog.schema}` for a table,
	// `{securable_type}` and `{securable_id}` describe that parent. If the
	// filled-in URL is longer than 4,096 characters, the **redirect_url** is
	// used instead.
	//
	// Must start with `http://` or `https://`, contain only printable ASCII
	// characters without spaces or backslashes, and be at most 2,048
	// characters. Braces are only allowed around a placeholder.
	DynamicUrl string `json:"dynamic_url,omitempty"`
	// The URL for access requests on the metastore and on securables outside
	// catalogs, such as external locations, connections, and credentials. Also
	// used for catalogs and the securables inside them when **dynamic_url** is
	// not set or cannot produce a valid URL. Must start with `http://` or
	// `https://`, contain only printable ASCII characters without spaces or
	// backslashes, and be at most 2,048 characters.
	RedirectUrl string `json:"redirect_url,omitempty"`
	// The version of the placeholder syntax used in **dynamic_url**. The only
	// supported version is `1`.
	SyntaxVersion int `json:"syntax_version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UrlConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UrlConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ValidateCredentialRequest struct {
	// Deprecated: create a credential and pass its credential_name instead.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// Deprecated: create a credential and pass its credential_name instead.
	AzureManagedIdentity *AzureManagedIdentity `json:"azure_managed_identity,omitempty"`
	// Required. The name of an existing credential or long-lived cloud
	// credential to validate.
	CredentialName string `json:"credential_name,omitempty"`
	// Deprecated: create a credential and pass its credential_name instead.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	// The name of an existing external location to validate. Only applicable
	// for storage credentials (purpose is **STORAGE**.)
	ExternalLocationName string `json:"external_location_name,omitempty"`
	// The purpose of the credential. This should only be used when the
	// credential is specified.
	Purpose CredentialPurpose `json:"purpose,omitempty"`
	// Whether the credential is only usable for read operations. Only
	// applicable for storage credentials (purpose is **STORAGE**.)
	ReadOnly bool `json:"read_only,omitempty"`
	// The external location url to validate. Only applicable when purpose is
	// **STORAGE**.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ValidateCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ValidateCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage. Only
	// applicable for when purpose is **STORAGE**.
	IsDir bool `json:"isDir,omitempty"`
	// The results of the validation check.
	Results []CredentialValidationResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ValidateCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateCredentialResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A enum represents the result of the file operation
type ValidateCredentialResult string

const ValidateCredentialResultFail ValidateCredentialResult = `FAIL`

const ValidateCredentialResultPass ValidateCredentialResult = `PASS`

const ValidateCredentialResultSkip ValidateCredentialResult = `SKIP`

// String representation for [fmt.Print]
func (f *ValidateCredentialResult) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ValidateCredentialResult) Set(v string) error {
	switch v {
	case `FAIL`, `PASS`, `SKIP`:
		*f = ValidateCredentialResult(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAIL", "PASS", "SKIP"`, v)
	}
}

// Values returns all possible values for ValidateCredentialResult.
//
// There is no guarantee on the order of the values in the slice.
func (f *ValidateCredentialResult) Values() []ValidateCredentialResult {
	return []ValidateCredentialResult{
		ValidateCredentialResultFail,
		ValidateCredentialResultPass,
		ValidateCredentialResultSkip,
	}
}

// Type always returns ValidateCredentialResult to satisfy [pflag.Value] interface
func (f *ValidateCredentialResult) Type() string {
	return "ValidateCredentialResult"
}

type ValidateStorageCredential struct {
	// Deprecated: create a storage credential and pass its
	// storage_credential_name instead. The AWS IAM role configuration.
	AwsIamRole *AwsIamRoleRequest `json:"aws_iam_role,omitempty"`
	// Deprecated: create a storage credential and pass its
	// storage_credential_name instead. The Azure managed identity
	// configuration.
	AzureManagedIdentity *AzureManagedIdentityRequest `json:"azure_managed_identity,omitempty"`
	// Deprecated: create a storage credential and pass its
	// storage_credential_name instead. The Azure service principal
	// configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Deprecated: create a storage credential and pass its
	// storage_credential_name instead. The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// Deprecated: create a storage credential and pass its
	// storage_credential_name instead. The Databricks created GCP service
	// account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	// The name of an existing external location to validate.
	ExternalLocationName string `json:"external_location_name,omitempty"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// Required. The name of an existing credential or long-lived cloud
	// credential to validate.
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// The external location url to validate.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ValidateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ValidateStorageCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage.
	IsDir bool `json:"isDir,omitempty"`
	// The results of the validation check.
	Results []ValidationResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ValidateStorageCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateStorageCredentialResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message string `json:"message,omitempty"`
	// The operation tested.
	Operation ValidationResultOperation `json:"operation,omitempty"`
	// The results of the tested operation.
	Result ValidationResultResult `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ValidationResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidationResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A enum represents the file operation performed on the external location with
// the storage credential
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

// Values returns all possible values for ValidationResultOperation.
//
// There is no guarantee on the order of the values in the slice.
func (f *ValidationResultOperation) Values() []ValidationResultOperation {
	return []ValidationResultOperation{
		ValidationResultOperationDelete,
		ValidationResultOperationList,
		ValidationResultOperationPathExists,
		ValidationResultOperationRead,
		ValidationResultOperationWrite,
	}
}

// Type always returns ValidationResultOperation to satisfy [pflag.Value] interface
func (f *ValidationResultOperation) Type() string {
	return "ValidationResultOperation"
}

// A enum represents the result of the file operation
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

// Values returns all possible values for ValidationResultResult.
//
// There is no guarantee on the order of the values in the slice.
func (f *ValidationResultResult) Values() []ValidationResultResult {
	return []ValidationResultResult{
		ValidationResultResultFail,
		ValidationResultResultPass,
		ValidationResultResultSkip,
	}
}

// Type always returns ValidationResultResult to satisfy [pflag.Value] interface
func (f *ValidationResultResult) Type() string {
	return "ValidationResultResult"
}

type VolumeInfo struct {
	// The AWS access point to use when accessing S3 for the storage location of
	// this volume.
	AccessPoint string `json:"access_point,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog where the schema and the volume are
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`
	// Time at which the volume was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the volume
	CreatedBy string `json:"created_by,omitempty"`

	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// The three-level (fully qualified) name of the volume, in the form
	// catalog_name.schema_name.name.
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the metastore
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the volume
	Name string `json:"name,omitempty"`
	// The user or group that owns the volume. The user who creates a volume
	// becomes its owner; the owner can be changed afterwards with an update.
	Owner string `json:"owner,omitempty"`
	// The name of the schema where the volume is
	SchemaName string `json:"schema_name,omitempty"`
	// The cloud storage URL where the files of the volume are stored. For an
	// external volume, it must be set on creation to a path within an existing
	// external location. For a managed volume, it must be omitted on creation;
	// the volume is placed under the managed storage location of the parent
	// schema, catalog, or metastore, and the assigned path is returned in this
	// field.
	StorageLocation string `json:"storage_location,omitempty"`
	// Time at which the volume was last updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the volume last time
	UpdatedBy string `json:"updated_by,omitempty"`
	// The unique identifier of the volume
	VolumeId string `json:"volume_id,omitempty"`
	// The type of the volume. An external volume is located in the specified
	// external location. A managed volume is located in the default location
	// which is specified by the parent schema, or the parent catalog, or the
	// Metastore. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/aws/en/volumes/managed-vs-external
	VolumeType VolumeType `json:"volume_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *VolumeInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VolumeInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VolumeOperation string

const VolumeOperationReadVolume VolumeOperation = `READ_VOLUME`

const VolumeOperationWriteVolume VolumeOperation = `WRITE_VOLUME`

// String representation for [fmt.Print]
func (f *VolumeOperation) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VolumeOperation) Set(v string) error {
	switch v {
	case `READ_VOLUME`, `WRITE_VOLUME`:
		*f = VolumeOperation(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "READ_VOLUME", "WRITE_VOLUME"`, v)
	}
}

// Values returns all possible values for VolumeOperation.
//
// There is no guarantee on the order of the values in the slice.
func (f *VolumeOperation) Values() []VolumeOperation {
	return []VolumeOperation{
		VolumeOperationReadVolume,
		VolumeOperationWriteVolume,
	}
}

// Type always returns VolumeOperation to satisfy [pflag.Value] interface
func (f *VolumeOperation) Type() string {
	return "VolumeOperation"
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

// Values returns all possible values for VolumeType.
//
// There is no guarantee on the order of the values in the slice.
func (f *VolumeType) Values() []VolumeType {
	return []VolumeType{
		VolumeTypeExternal,
		VolumeTypeManaged,
	}
}

// Type always returns VolumeType to satisfy [pflag.Value] interface
func (f *VolumeType) Type() string {
	return "VolumeType"
}

// A binding between a workspace and a securable. When the isolation mode of the
// securable is `ISOLATED`, only bound workspaces can access it.
type WorkspaceBinding struct {
	// One of READ_WRITE/READ_ONLY. Default is READ_WRITE. Only catalogs support
	// `BINDING_TYPE_READ_ONLY`; other securable types accept only
	// `BINDING_TYPE_READ_WRITE`.
	BindingType WorkspaceBindingBindingType `json:"binding_type,omitempty"`
	// The ID of the workspace that the binding applies to.
	WorkspaceId int64 `json:"workspace_id"`
}

func (s *WorkspaceBinding) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Using `BINDING_TYPE_` prefix here to avoid conflict with `TableOperation`
// enum in `credentials_common.proto`.
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

// Values returns all possible values for WorkspaceBindingBindingType.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceBindingBindingType) Values() []WorkspaceBindingBindingType {
	return []WorkspaceBindingBindingType{
		WorkspaceBindingBindingTypeBindingTypeReadOnly,
		WorkspaceBindingBindingTypeBindingTypeReadWrite,
	}
}

// Type always returns WorkspaceBindingBindingType to satisfy [pflag.Value] interface
func (f *WorkspaceBindingBindingType) Type() string {
	return "WorkspaceBindingBindingType"
}
