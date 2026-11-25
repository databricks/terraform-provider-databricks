// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AwsCredentials struct {
	StsRole *StsRole `json:"sts_role,omitempty"`
}

func (s *AwsCredentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias string `json:"key_alias,omitempty"`
	// The AWS KMS key's Amazon Resource Name (ARN).
	KeyArn string `json:"key_arn"`
	// The AWS KMS key region.
	KeyRegion string `json:"key_region"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to true or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to false.
	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsKeyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsKeyInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AzureKeyInfo struct {
	// The Disk Encryption Set id that is used to represent the key info used
	// for Managed Disk BYOK use case
	DiskEncryptionSetId string `json:"disk_encryption_set_id,omitempty"`
	// The structure to store key access credential This is set if the Managed
	// Identity is being used to access the Azure Key Vault key.
	KeyAccessConfiguration *KeyAccessConfiguration `json:"key_access_configuration,omitempty"`
	// The name of the key in KeyVault.
	KeyName string `json:"key_name,omitempty"`
	// The base URI of the KeyVault.
	KeyVaultUri string `json:"key_vault_uri,omitempty"`
	// The tenant id where the KeyVault lives.
	TenantId string `json:"tenant_id,omitempty"`
	// The current key version.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureKeyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureKeyInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AzureWorkspaceInfo struct {
	// Azure Resource Group name
	ResourceGroup string `json:"resource_group,omitempty"`
	// Azure Subscription ID
	SubscriptionId string `json:"subscription_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureWorkspaceInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureWorkspaceInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CloudResourceContainer struct {
	Gcp *CustomerFacingGcpCloudResourceContainer `json:"gcp,omitempty"`
}

func (s *CloudResourceContainer) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateAwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias string `json:"key_alias,omitempty"`
	// The AWS KMS key's Amazon Resource Name (ARN).
	KeyArn string `json:"key_arn"`
	// The AWS KMS key region.
	KeyRegion string `json:"key_region,omitempty"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to true or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to false.
	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateAwsKeyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAwsKeyInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateAzureKeyInfo struct {
	// The Disk Encryption Set id that is used to represent the key info used
	// for Managed Disk BYOK use case
	DiskEncryptionSetId string `json:"disk_encryption_set_id,omitempty"`
	// The structure to store key access credential This is set if the Managed
	// Identity is being used to access the Azure Key Vault key.
	KeyAccessConfiguration *KeyAccessConfiguration `json:"key_access_configuration,omitempty"`
	// The name of the key in KeyVault.
	KeyName string `json:"key_name,omitempty"`
	// The base URI of the KeyVault.
	KeyVaultUri string `json:"key_vault_uri,omitempty"`
	// The tenant id where the KeyVault lives.
	TenantId string `json:"tenant_id,omitempty"`
	// The current key version.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateAzureKeyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAzureKeyInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCredentialAwsCredentials struct {
	StsRole *CreateCredentialStsRole `json:"sts_role,omitempty"`
}

func (s *CreateCredentialAwsCredentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateCredentialRequest struct {
	AwsCredentials CreateCredentialAwsCredentials `json:"aws_credentials"`
	// The human-readable name of the credential configuration object.
	CredentialsName string `json:"credentials_name"`
}

func (s *CreateCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateCredentialStsRole struct {
	// The Amazon Resource Name (ARN) of the cross account IAM role.
	RoleArn string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCredentialStsRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentialStsRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCustomerManagedKeyRequest struct {
	AwsKeyInfo *CreateAwsKeyInfo `json:"aws_key_info,omitempty"`

	AzureKeyInfo *CreateAzureKeyInfo `json:"azure_key_info,omitempty"`

	GcpKeyInfo *CreateGcpKeyInfo `json:"gcp_key_info,omitempty"`
	// The cases that the key can be used for.
	UseCases []KeyUseCase `json:"use_cases"`
}

func (s *CreateCustomerManagedKeyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateGcpKeyInfo struct {
	// Globally unique service account email that has access to the KMS key. The
	// service account exists within the Databricks CP project.
	GcpServiceAccount *GcpServiceAccount `json:"gcp_service_account,omitempty"`
	// Globally unique kms key resource id of the form
	// projects/testProjectId/locations/us-east4/keyRings/gcpCmkKeyRing/cryptoKeys/cmk-eastus4
	KmsKeyId string `json:"kms_key_id"`
	// When true, Databricks will not use OAuth to grant the service account
	// access to the KMS key. The customer is responsible for granting access
	// manually.
	Manual bool `json:"manual,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateGcpKeyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateGcpKeyInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateNetworkRequest struct {
	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`
	// The human-readable name of the network configuration.
	NetworkName string `json:"network_name,omitempty"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds []string `json:"subnet_ids,omitempty"`

	VpcEndpoints *NetworkVpcEndpoints `json:"vpc_endpoints,omitempty"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId string `json:"vpc_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateNetworkRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateNetworkRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePrivateAccessSettingsRequest struct {
	// An array of Databricks VPC endpoint IDs. This is the Databricks ID
	// returned when registering the VPC endpoint configuration in your
	// Databricks account. This is not the ID of the VPC endpoint in AWS. Only
	// used when private_access_level is set to ENDPOINT. This is an allow list
	// of VPC endpoints registered in your Databricks account that can connect
	// to your workspace over AWS PrivateLink. Note: If hybrid access to your
	// workspace is enabled by setting public_access_enabled to true, this
	// control only works for PrivateLink connections. To control how your
	// workspace is accessed via public internet, see IP access lists.
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see allowed_vpc_endpoint_ids.
	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName string `json:"private_access_settings_name,omitempty"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify false, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify true, which means that public access is
	// enabled.
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The AWS region for workspaces attached to this private access settings
	// object.
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePrivateAccessSettingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePrivateAccessSettingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateStorageConfigurationRequest struct {
	// Optional IAM role that is used to access the workspace catalog which is
	// created during workspace creation for UC by Default. If a storage
	// configuration with this field populated is used to create a workspace,
	// then a workspace catalog is created together with the workspace. The
	// workspace catalog shares the root bucket with internal workspace storage
	// (including DBFS root) but uses a dedicated bucket path prefix.
	RoleArn string `json:"role_arn,omitempty"`
	// Root S3 bucket information.
	RootBucketInfo RootBucketInfo `json:"root_bucket_info"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName string `json:"storage_configuration_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateStorageConfigurationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateStorageConfigurationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id,omitempty"`
	// The cloud info of this vpc endpoint.
	GcpVpcEndpointInfo *GcpVpcEndpointInfo `json:"gcp_vpc_endpoint_info,omitempty"`
	// The region in which this VPC endpoint object exists.
	Region string `json:"region,omitempty"`
	// The human-readable name of the storage configuration.
	VpcEndpointName string `json:"vpc_endpoint_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateVpcEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateVpcEndpointRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateWorkspaceRequest struct {
	AwsRegion string `json:"aws_region,omitempty"`
	// DEPRECATED: This field is being ignored by the server and will be removed
	// in the future. The cloud name. This field always has the value `gcp`.
	Cloud string `json:"cloud,omitempty"`

	CloudResourceContainer *CloudResourceContainer `json:"cloud_resource_container,omitempty"`
	// If the compute mode is `SERVERLESS`, a serverless workspace is created
	// that comes pre-configured with serverless compute and default storage,
	// providing a fully-managed, enterprise-ready SaaS experience. This means
	// you don't need to provide any resources managed by you, such as
	// credentials, storage, or network. If the compute mode is `HYBRID` (which
	// is the default option), a classic workspace is created that uses
	// customer-managed resources.
	ComputeMode CustomerFacingComputeMode `json:"compute_mode,omitempty"`
	// ID of the workspace's credential configuration object.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for the web application and REST APIs is
	// <workspace-deployment-name>.cloud.databricks.com. For example, if the
	// deployment name is abcsales, your workspace URL will be
	// https://abcsales.cloud.databricks.com. Hyphens are allowed. This property
	// supports only the set of characters that are allowed in a subdomain. To
	// set this value, you must have a deployment name prefix. Contact your
	// Databricks account team to add an account deployment name prefix to your
	// account. Workspace deployment names follow the account prefix and a
	// hyphen. For example, if your account's deployment prefix is acme and the
	// workspace deployment name is workspace-1, the JSON response for the
	// deployment_name field becomes acme-workspace-1. The workspace URL would
	// be acme-workspace-1.cloud.databricks.com. You can also set the
	// deployment_name to the reserved keyword EMPTY if you want the deployment
	// name to only include the deployment prefix. For example, if your
	// account's deployment prefix is acme and the workspace deployment name is
	// EMPTY, the deployment_name becomes acme only and the workspace URL is
	// acme.cloud.databricks.com. This value must be unique across all
	// non-deleted deployments across all AWS regions. If a new workspace omits
	// this property, the server generates a unique deployment name for you with
	// the pattern dbc-xxxxxxxx-xxxx.
	DeploymentName string `json:"deployment_name,omitempty"`

	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`

	GkeConfig *GkeConfig `json:"gke_config,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location string `json:"location,omitempty"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to help protect and control access to the
	// workspace's notebooks, secrets, Databricks SQL queries, and query
	// history. The provided key configuration object property use_cases must
	// contain MANAGED_SERVICES.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`
	// The object ID of network connectivity config. Once assigned, the
	// workspace serverless compute resources use the same set of stable IP CIDR
	// blocks and optional private link to access your resources.
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The ID of the workspace's network configuration object. To use AWS
	// PrivateLink, this field is required.
	NetworkId string `json:"network_id,omitempty"`

	PricingTier PricingTier `json:"pricing_tier,omitempty"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. You must specify this ID if you are using [AWS PrivateLink]
	// for either front-end (user-to-workspace connection), back-end (data plane
	// to control plane connection), or both connection types. Before
	// configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// ID of the workspace's storage configuration object.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The ID of the workspace's storage encryption key configuration object.
	// This is used to encrypt the workspace's root S3 bucket (root DBFS and
	// system data) and, optionally, cluster EBS volumes. The provided key
	// configuration object property use_cases must contain STORAGE.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// The human-readable name of the workspace.
	WorkspaceName string `json:"workspace_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Credential struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`

	AwsCredentials *AwsCredentials `json:"aws_credentials,omitempty"`
	// Time in epoch milliseconds when the credential was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Databricks credential configuration ID.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The human-readable name of the credential configuration object.
	CredentialsName string `json:"credentials_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Credential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Credential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Corresponds to compute mode defined here:
// https://src.dev.databricks.com/databricks/universe@9076536b18479afd639d1c1f9dd5a59f72215e69/-/blob/central/api/common.proto?L872
type CustomerFacingComputeMode string

const CustomerFacingComputeModeHybrid CustomerFacingComputeMode = `HYBRID`

const CustomerFacingComputeModeServerless CustomerFacingComputeMode = `SERVERLESS`

// String representation for [fmt.Print]
func (f *CustomerFacingComputeMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomerFacingComputeMode) Set(v string) error {
	switch v {
	case `HYBRID`, `SERVERLESS`:
		*f = CustomerFacingComputeMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "HYBRID", "SERVERLESS"`, v)
	}
}

// Values returns all possible values for CustomerFacingComputeMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomerFacingComputeMode) Values() []CustomerFacingComputeMode {
	return []CustomerFacingComputeMode{
		CustomerFacingComputeModeHybrid,
		CustomerFacingComputeModeServerless,
	}
}

// Type always returns CustomerFacingComputeMode to satisfy [pflag.Value] interface
func (f *CustomerFacingComputeMode) Type() string {
	return "CustomerFacingComputeMode"
}

type CustomerFacingGcpCloudResourceContainer struct {
	ProjectId string `json:"project_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingGcpCloudResourceContainer) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingGcpCloudResourceContainer) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CustomerFacingStorageMode string

const CustomerFacingStorageModeCustomerHosted CustomerFacingStorageMode = `CUSTOMER_HOSTED`

const CustomerFacingStorageModeDefaultStorage CustomerFacingStorageMode = `DEFAULT_STORAGE`

// String representation for [fmt.Print]
func (f *CustomerFacingStorageMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomerFacingStorageMode) Set(v string) error {
	switch v {
	case `CUSTOMER_HOSTED`, `DEFAULT_STORAGE`:
		*f = CustomerFacingStorageMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CUSTOMER_HOSTED", "DEFAULT_STORAGE"`, v)
	}
}

// Values returns all possible values for CustomerFacingStorageMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomerFacingStorageMode) Values() []CustomerFacingStorageMode {
	return []CustomerFacingStorageMode{
		CustomerFacingStorageModeCustomerHosted,
		CustomerFacingStorageModeDefaultStorage,
	}
}

// Type always returns CustomerFacingStorageMode to satisfy [pflag.Value] interface
func (f *CustomerFacingStorageMode) Type() string {
	return "CustomerFacingStorageMode"
}

type CustomerManagedKey struct {
	// The Databricks account ID that holds the customer-managed key.
	AccountId string `json:"account_id,omitempty"`

	AwsKeyInfo *AwsKeyInfo `json:"aws_key_info,omitempty"`

	AzureKeyInfo *AzureKeyInfo `json:"azure_key_info,omitempty"`
	// Time in epoch milliseconds when the customer key was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId string `json:"customer_managed_key_id,omitempty"`

	GcpKeyInfo *GcpKeyInfo `json:"gcp_key_info,omitempty"`
	// The cases that the key can be used for.
	UseCases []KeyUseCase `json:"use_cases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerManagedKey) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerManagedKey) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteCredentialRequest struct {
	// Databricks Account API credential configuration ID
	CredentialsId string `json:"-" url:"-"`
}

func (s *DeleteCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string `json:"-" url:"-"`
}

func (s *DeleteEncryptionKeyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId string `json:"-" url:"-"`
}

func (s *DeleteNetworkRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeletePrivateAccesRequest struct {
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

func (s *DeletePrivateAccesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteStorageRequest struct {
	StorageConfigurationId string `json:"-" url:"-"`
}

func (s *DeleteStorageRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteVpcEndpointRequest struct {
	VpcEndpointId string `json:"-" url:"-"`
}

func (s *DeleteVpcEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteWorkspaceRequest struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *DeleteWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EndpointUseCase string

const EndpointUseCaseDataplaneRelayAccess EndpointUseCase = `DATAPLANE_RELAY_ACCESS`

const EndpointUseCaseGeneralAccess EndpointUseCase = `GENERAL_ACCESS`

const EndpointUseCaseWorkspaceAccess EndpointUseCase = `WORKSPACE_ACCESS`

// String representation for [fmt.Print]
func (f *EndpointUseCase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointUseCase) Set(v string) error {
	switch v {
	case `DATAPLANE_RELAY_ACCESS`, `GENERAL_ACCESS`, `WORKSPACE_ACCESS`:
		*f = EndpointUseCase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATAPLANE_RELAY_ACCESS", "GENERAL_ACCESS", "WORKSPACE_ACCESS"`, v)
	}
}

// Values returns all possible values for EndpointUseCase.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointUseCase) Values() []EndpointUseCase {
	return []EndpointUseCase{
		EndpointUseCaseDataplaneRelayAccess,
		EndpointUseCaseGeneralAccess,
		EndpointUseCaseWorkspaceAccess,
	}
}

// Type always returns EndpointUseCase to satisfy [pflag.Value] interface
func (f *EndpointUseCase) Type() string {
	return "EndpointUseCase"
}

// ErrorType and WarningType are used to represent the type of error or warning
// by NetworkHealth and NetworkWarning defined in
// central/api/accounts/accounts.proto
type ErrorType string

const ErrorTypeCredentials ErrorType = `credentials`

const ErrorTypeNetworkAcl ErrorType = `networkAcl`

const ErrorTypeSecurityGroup ErrorType = `securityGroup`

const ErrorTypeSubnet ErrorType = `subnet`

const ErrorTypeVpc ErrorType = `vpc`

// String representation for [fmt.Print]
func (f *ErrorType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ErrorType) Set(v string) error {
	switch v {
	case `credentials`, `networkAcl`, `securityGroup`, `subnet`, `vpc`:
		*f = ErrorType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "credentials", "networkAcl", "securityGroup", "subnet", "vpc"`, v)
	}
}

// Values returns all possible values for ErrorType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ErrorType) Values() []ErrorType {
	return []ErrorType{
		ErrorTypeCredentials,
		ErrorTypeNetworkAcl,
		ErrorTypeSecurityGroup,
		ErrorTypeSubnet,
		ErrorTypeVpc,
	}
}

// Type always returns ErrorType to satisfy [pflag.Value] interface
func (f *ErrorType) Type() string {
	return "ErrorType"
}

// The shared network config for GCP workspace. This object has common network
// configurations that are network attributions of a workspace. DEPRECATED. Use
// GkeConfig instead.
type GcpCommonNetworkConfig struct {
	// The IP range that will be used to allocate GKE cluster master resources
	// from. This field must not be set if
	// gke_cluster_type=PUBLIC_NODE_PUBLIC_MASTER.
	GkeClusterMasterIpRange string `json:"gke_cluster_master_ip_range,omitempty"`
	// The type of network connectivity of the GKE cluster.
	GkeConnectivityType GkeConfigConnectivityType `json:"gke_connectivity_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpCommonNetworkConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpCommonNetworkConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GcpKeyInfo struct {
	// Globally unique service account email that has access to the KMS key. The
	// service account exists within the Databricks CP project.
	GcpServiceAccount *GcpServiceAccount `json:"gcp_service_account,omitempty"`
	// Globally unique kms key resource id of the form
	// projects/testProjectId/locations/us-east4/keyRings/gcpCmkKeyRing/cryptoKeys/cmk-eastus4
	KmsKeyId string `json:"kms_key_id"`
	// When true, Databricks will not use OAuth to grant the service account
	// access to the KMS key. The customer is responsible for granting access
	// manually.
	Manual bool `json:"manual,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpKeyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpKeyInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The network configuration for the workspace.
type GcpManagedNetworkConfig struct {
	// The IP range that will be used to allocate GKE cluster Pods from.
	GkeClusterPodIpRange string `json:"gke_cluster_pod_ip_range,omitempty"`
	// The IP range that will be used to allocate GKE cluster Services from.
	GkeClusterServiceIpRange string `json:"gke_cluster_service_ip_range,omitempty"`
	// The IP range which will be used to allocate GKE cluster nodes from. Note:
	// Pods, services and master IP range must be mutually exclusive.
	SubnetCidr string `json:"subnet_cidr,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpManagedNetworkConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpManagedNetworkConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GcpNetworkInfo struct {
	// The GCP project ID for network resources. This project is where the VPC
	// and subnet resides.
	NetworkProjectId string `json:"network_project_id"`
	// Name of the secondary range within the subnet that will be used by GKE as
	// Pod IP range. This is BYO VPC specific. DB VPC uses
	// network.getGcpManagedNetworkConfig.getGkeClusterPodIpRange
	PodIpRangeName string `json:"pod_ip_range_name"`
	// Name of the secondary range within the subnet that will be used by GKE as
	// Service IP range.
	ServiceIpRangeName string `json:"service_ip_range_name"`
	// The customer-provided Subnet ID that will be available to Clusters in
	// Workspaces using this Network.
	SubnetId string `json:"subnet_id"`

	SubnetRegion string `json:"subnet_region"`
	// The customer-provided VPC ID.
	VpcId string `json:"vpc_id"`
}

func (s *GcpNetworkInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GcpServiceAccount struct {
	ServiceAccountEmail string `json:"service_account_email,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpServiceAccount) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpServiceAccount) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GcpVpcEndpointInfo struct {
	EndpointRegion string `json:"endpoint_region"`

	ProjectId string `json:"project_id"`

	PscConnectionId string `json:"psc_connection_id,omitempty"`

	PscEndpointName string `json:"psc_endpoint_name"`

	ServiceAttachmentId string `json:"service_attachment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpVpcEndpointInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpVpcEndpointInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetCredentialRequest struct {
	// Credential configuration ID
	CredentialsId string `json:"-" url:"-"`
}

func (s *GetCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string `json:"-" url:"-"`
}

func (s *GetEncryptionKeyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId string `json:"-" url:"-"`
}

func (s *GetNetworkRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPrivateAccesRequest struct {
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

func (s *GetPrivateAccesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetStorageRequest struct {
	StorageConfigurationId string `json:"-" url:"-"`
}

func (s *GetStorageRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId string `json:"-" url:"-"`
}

func (s *GetVpcEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceRequest struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *GetWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The configurations of the GKE cluster used by the GCP workspace.
type GkeConfig struct {
	// The type of network connectivity of the GKE cluster.
	ConnectivityType GkeConfigConnectivityType `json:"connectivity_type,omitempty"`
	// The IP range that will be used to allocate GKE cluster master resources
	// from. This field must not be set if
	// gke_cluster_type=PUBLIC_NODE_PUBLIC_MASTER.
	MasterIpRange string `json:"master_ip_range,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GkeConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GkeConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Specifies the network connectivity types for the GKE nodes and the GKE master
// network.
//
// Set to `PRIVATE_NODE_PUBLIC_MASTER` for a private GKE cluster for the
// workspace. The GKE nodes will not have public IPs.
//
// Set to `PUBLIC_NODE_PUBLIC_MASTER` for a public GKE cluster. The nodes of a
// public GKE cluster have public IP addresses.
type GkeConfigConnectivityType string

const GkeConfigConnectivityTypePrivateNodePublicMaster GkeConfigConnectivityType = `PRIVATE_NODE_PUBLIC_MASTER`

const GkeConfigConnectivityTypePublicNodePublicMaster GkeConfigConnectivityType = `PUBLIC_NODE_PUBLIC_MASTER`

// String representation for [fmt.Print]
func (f *GkeConfigConnectivityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GkeConfigConnectivityType) Set(v string) error {
	switch v {
	case `PRIVATE_NODE_PUBLIC_MASTER`, `PUBLIC_NODE_PUBLIC_MASTER`:
		*f = GkeConfigConnectivityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PRIVATE_NODE_PUBLIC_MASTER", "PUBLIC_NODE_PUBLIC_MASTER"`, v)
	}
}

// Values returns all possible values for GkeConfigConnectivityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *GkeConfigConnectivityType) Values() []GkeConfigConnectivityType {
	return []GkeConfigConnectivityType{
		GkeConfigConnectivityTypePrivateNodePublicMaster,
		GkeConfigConnectivityTypePublicNodePublicMaster,
	}
}

// Type always returns GkeConfigConnectivityType to satisfy [pflag.Value] interface
func (f *GkeConfigConnectivityType) Type() string {
	return "GkeConfigConnectivityType"
}

// The credential ID that is used to access the key vault.
type KeyAccessConfiguration struct {
	CredentialId string `json:"credential_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KeyAccessConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KeyAccessConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type KeyUseCase string

const KeyUseCaseManagedServices KeyUseCase = `MANAGED_SERVICES`

const KeyUseCaseStorage KeyUseCase = `STORAGE`

// String representation for [fmt.Print]
func (f *KeyUseCase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *KeyUseCase) Set(v string) error {
	switch v {
	case `MANAGED_SERVICES`, `STORAGE`:
		*f = KeyUseCase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANAGED_SERVICES", "STORAGE"`, v)
	}
}

// Values returns all possible values for KeyUseCase.
//
// There is no guarantee on the order of the values in the slice.
func (f *KeyUseCase) Values() []KeyUseCase {
	return []KeyUseCase{
		KeyUseCaseManagedServices,
		KeyUseCaseStorage,
	}
}

// Type always returns KeyUseCase to satisfy [pflag.Value] interface
func (f *KeyUseCase) Type() string {
	return "KeyUseCase"
}

type Network struct {
	// The Databricks account ID associated with this network configuration.
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when the network was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Array of error messages about the network configuration.
	ErrorMessages []NetworkHealth `json:"error_messages,omitempty"`

	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`
	// The Databricks network configuration ID.
	NetworkId string `json:"network_id,omitempty"`
	// The human-readable name of the network configuration.
	NetworkName string `json:"network_name,omitempty"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds []string `json:"subnet_ids,omitempty"`

	VpcEndpoints *NetworkVpcEndpoints `json:"vpc_endpoints,omitempty"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId string `json:"vpc_id,omitempty"`

	VpcStatus VpcStatus `json:"vpc_status,omitempty"`
	// Array of warning messages about the network configuration.
	WarningMessages []NetworkWarning `json:"warning_messages,omitempty"`
	// Workspace ID associated with this network configuration.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Network) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Network) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NetworkHealth struct {
	// Details of the error.
	ErrorMessage string `json:"error_message,omitempty"`

	ErrorType ErrorType `json:"error_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NetworkHealth) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NetworkHealth) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	DataplaneRelay []string `json:"dataplane_relay,omitempty"`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	RestApi []string `json:"rest_api,omitempty"`
}

func (s *NetworkVpcEndpoints) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type NetworkWarning struct {
	// Details of the warning.
	WarningMessage string `json:"warning_message,omitempty"`

	WarningType WarningType `json:"warning_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NetworkWarning) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NetworkWarning) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PricingTier string

const PricingTierCommunityEdition PricingTier = `COMMUNITY_EDITION`

const PricingTierDedicated PricingTier = `DEDICATED`

const PricingTierEnterprise PricingTier = `ENTERPRISE`

const PricingTierPremium PricingTier = `PREMIUM`

const PricingTierStandard PricingTier = `STANDARD`

const PricingTierUnknown PricingTier = `UNKNOWN`

// String representation for [fmt.Print]
func (f *PricingTier) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PricingTier) Set(v string) error {
	switch v {
	case `COMMUNITY_EDITION`, `DEDICATED`, `ENTERPRISE`, `PREMIUM`, `STANDARD`, `UNKNOWN`:
		*f = PricingTier(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COMMUNITY_EDITION", "DEDICATED", "ENTERPRISE", "PREMIUM", "STANDARD", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for PricingTier.
//
// There is no guarantee on the order of the values in the slice.
func (f *PricingTier) Values() []PricingTier {
	return []PricingTier{
		PricingTierCommunityEdition,
		PricingTierDedicated,
		PricingTierEnterprise,
		PricingTierPremium,
		PricingTierStandard,
		PricingTierUnknown,
	}
}

// Type always returns PricingTier to satisfy [pflag.Value] interface
func (f *PricingTier) Type() string {
	return "PricingTier"
}

type PrivateAccessLevel string

const PrivateAccessLevelAccount PrivateAccessLevel = `ACCOUNT`

const PrivateAccessLevelEndpoint PrivateAccessLevel = `ENDPOINT`

// String representation for [fmt.Print]
func (f *PrivateAccessLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrivateAccessLevel) Set(v string) error {
	switch v {
	case `ACCOUNT`, `ENDPOINT`:
		*f = PrivateAccessLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACCOUNT", "ENDPOINT"`, v)
	}
}

// Values returns all possible values for PrivateAccessLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrivateAccessLevel) Values() []PrivateAccessLevel {
	return []PrivateAccessLevel{
		PrivateAccessLevelAccount,
		PrivateAccessLevelEndpoint,
	}
}

// Type always returns PrivateAccessLevel to satisfy [pflag.Value] interface
func (f *PrivateAccessLevel) Type() string {
	return "PrivateAccessLevel"
}

// *
type PrivateAccessSettings struct {
	// The Databricks account ID that hosts the private access settings.
	AccountId string `json:"account_id,omitempty"`
	// An array of Databricks VPC endpoint IDs. This is the Databricks ID that
	// is returned when registering the VPC endpoint configuration in your
	// Databricks account. This is not the ID of the VPC endpoint in AWS. Only
	// used when private_access_level is set to ENDPOINT. This is an allow list
	// of VPC endpoints that in your account that can connect to your workspace
	// over AWS PrivateLink. If hybrid access to your workspace is enabled by
	// setting public_access_enabled to true, this control only works for
	// PrivateLink connections. To control how your workspace is accessed via
	// public internet, see IP access lists.
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see allowed_vpc_endpoint_ids.
	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`
	// Databricks private access settings ID.
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName string `json:"private_access_settings_name,omitempty"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify false, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify true, which means that public access is
	// enabled.
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The AWS region for workspaces attached to this private access settings
	// object.
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PrivateAccessSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrivateAccessSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ReplacePrivateAccessSettingsRequest struct {
	// Properties of the new private access settings object.
	CustomerFacingPrivateAccessSettings PrivateAccessSettings `json:"customer_facing_private_access_settings"`
	// Databricks private access settings ID.
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

func (s *ReplacePrivateAccessSettingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RootBucketInfo struct {
	// Name of the S3 bucket
	BucketName string `json:"bucket_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RootBucketInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RootBucketInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StorageConfiguration struct {
	// The Databricks account ID associated with this storage configuration.
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when the storage configuration was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Optional IAM role that is used to access the workspace catalog which is
	// created during workspace creation for UC by Default. If a storage
	// configuration with this field populated is used to create a workspace,
	// then a workspace catalog is created together with the workspace. The
	// workspace catalog shares the root bucket with internal workspace storage
	// (including DBFS root) but uses a dedicated bucket path prefix.
	RoleArn string `json:"role_arn,omitempty"`
	// The root bucket information for the storage configuration.
	RootBucketInfo *RootBucketInfo `json:"root_bucket_info,omitempty"`
	// Databricks storage configuration ID.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName string `json:"storage_configuration_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StorageConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StorageConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StsRole struct {
	// The Amazon Resource Name (ARN) of the cross account IAM role.
	RoleArn string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StsRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StsRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateWorkspaceRequest struct {
	CustomerFacingWorkspace Workspace `json:"customer_facing_workspace"`
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
	UpdateMask string `json:"-" url:"update_mask,omitempty"`
	// A unique integer ID for the workspace
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// *
type VpcEndpoint struct {
	// The Databricks account ID that hosts the VPC endpoint configuration.
	AccountId string `json:"account_id,omitempty"`
	// The AWS Account in which the VPC endpoint object exists.
	AwsAccountId string `json:"aws_account_id,omitempty"`
	// The ID of the Databricks [endpoint service] that this VPC endpoint is
	// connected to. For a list of endpoint service IDs for each supported AWS
	// region, see the [Databricks PrivateLink documentation].
	//
	// [Databricks PrivateLink documentation]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	AwsEndpointServiceId string `json:"aws_endpoint_service_id,omitempty"`
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id,omitempty"`
	// The cloud info of this vpc endpoint. Info for a GCP vpc endpoint.
	GcpVpcEndpointInfo *GcpVpcEndpointInfo `json:"gcp_vpc_endpoint_info,omitempty"`
	// The AWS region in which this VPC endpoint object exists.
	Region string `json:"region,omitempty"`
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	State string `json:"state,omitempty"`
	// This enumeration represents the type of Databricks VPC endpoint service
	// that was used when creating this VPC endpoint. If the VPC endpoint
	// connects to the Databricks control plane for either the front-end
	// connection or the back-end REST API connection, the value is
	// GENERAL_ACCESS. If the VPC endpoint connects to the Databricks workspace
	// for the back-end secure cluster connectivity relay, the value is
	// DATAPLANE_RELAY_ACCESS.
	UseCase EndpointUseCase `json:"use_case,omitempty"`
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`
	// The human-readable name of the storage configuration.
	VpcEndpointName string `json:"vpc_endpoint_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *VpcEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VpcEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VpcStatus string

const VpcStatusBroken VpcStatus = `BROKEN`

const VpcStatusUnattached VpcStatus = `UNATTACHED`

const VpcStatusValid VpcStatus = `VALID`

const VpcStatusWarned VpcStatus = `WARNED`

// String representation for [fmt.Print]
func (f *VpcStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VpcStatus) Set(v string) error {
	switch v {
	case `BROKEN`, `UNATTACHED`, `VALID`, `WARNED`:
		*f = VpcStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BROKEN", "UNATTACHED", "VALID", "WARNED"`, v)
	}
}

// Values returns all possible values for VpcStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *VpcStatus) Values() []VpcStatus {
	return []VpcStatus{
		VpcStatusBroken,
		VpcStatusUnattached,
		VpcStatusValid,
		VpcStatusWarned,
	}
}

// Type always returns VpcStatus to satisfy [pflag.Value] interface
func (f *VpcStatus) Type() string {
	return "VpcStatus"
}

type WarningType string

const WarningTypeSecurityGroup WarningType = `securityGroup`

const WarningTypeSubnet WarningType = `subnet`

// String representation for [fmt.Print]
func (f *WarningType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WarningType) Set(v string) error {
	switch v {
	case `securityGroup`, `subnet`:
		*f = WarningType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "securityGroup", "subnet"`, v)
	}
}

// Values returns all possible values for WarningType.
//
// There is no guarantee on the order of the values in the slice.
func (f *WarningType) Values() []WarningType {
	return []WarningType{
		WarningTypeSecurityGroup,
		WarningTypeSubnet,
	}
}

// Type always returns WarningType to satisfy [pflag.Value] interface
func (f *WarningType) Type() string {
	return "WarningType"
}

type Workspace struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`

	AwsRegion string `json:"aws_region,omitempty"`

	AzureWorkspaceInfo *AzureWorkspaceInfo `json:"azure_workspace_info,omitempty"`
	// The cloud name. This field can have values like `azure`, `gcp`.
	Cloud string `json:"cloud,omitempty"`

	CloudResourceContainer *CloudResourceContainer `json:"cloud_resource_container,omitempty"`
	// The compute mode of the workspace.
	ComputeMode CustomerFacingComputeMode `json:"compute_mode,omitempty"`
	// Time in epoch milliseconds when the workspace was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the workspace's credential configuration object.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags map[string]string `json:"custom_tags,omitempty"`

	DeploymentName string `json:"deployment_name,omitempty"`
	// A client owned field used to indicate the workspace status that the
	// client expects to be in. For now this is only used to unblock Temporal
	// workflow for GCP least privileged workspace.
	ExpectedWorkspaceStatus WorkspaceStatus `json:"expected_workspace_status,omitempty"`

	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`

	GkeConfig *GkeConfig `json:"gke_config,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location string `json:"location,omitempty"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`
	// The network configuration for the workspace.
	//
	// DEPRECATED. Use `network_id` instead.
	Network *WorkspaceNetwork `json:"network,omitempty"`
	// The object ID of network connectivity config.
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// If this workspace is BYO VPC, then the network_id will be populated. If
	// this workspace is not BYO VPC, then the network_id will be empty.
	NetworkId string `json:"network_id,omitempty"`

	PricingTier PricingTier `json:"pricing_tier,omitempty"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. You must specify this ID if you are using [AWS PrivateLink]
	// for either front-end (user-to-workspace connection), back-end (data plane
	// to control plane connection), or both connection types.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// ID of the workspace's storage configuration object.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// ID of the key configuration for encrypting workspace storage.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// The storage mode of the workspace.
	StorageMode CustomerFacingStorageMode `json:"storage_mode,omitempty"`
	// A unique integer ID for the workspace
	WorkspaceId int64 `json:"workspace_id,omitempty"`
	// The human-readable name of the workspace.
	WorkspaceName string `json:"workspace_name,omitempty"`
	// The status of a workspace
	WorkspaceStatus WorkspaceStatus `json:"workspace_status,omitempty"`
	// Message describing the current workspace status.
	WorkspaceStatusMessage string `json:"workspace_status_message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Workspace) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Workspace) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The network configuration for workspaces.
type WorkspaceNetwork struct {
	// The shared network config for GCP workspace. This object has common
	// network configurations that are network attributions of a workspace. This
	// object is input-only.
	GcpCommonNetworkConfig *GcpCommonNetworkConfig `json:"gcp_common_network_config,omitempty"`
	// The mutually exclusive network deployment modes. The option decides which
	// network mode the workspace will use. The network config for GCP workspace
	// with Databricks managed network. This object is input-only and will not
	// be provided when listing workspaces.
	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`
	// The ID of the network object, if the workspace is a BYOVPC workspace.
	// This should apply to workspaces on all clouds in internal services. In
	// accounts-rest-api, user will use workspace.network_id for input and
	// output instead. Currently (2021-06-19) the network ID is only used by
	// GCP.
	NetworkId string `json:"network_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceNetwork) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceNetwork) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The different statuses of a workspace. The following represents the current
// set of valid transitions from status to status: NOT_PROVISIONED ->
// PROVISIONING -> CANCELLED PROVISIONING -> RUNNING -> FAILED -> CANCELLED
// (note that this transition is disallowed in the MultiWorkspace Project)
// RUNNING -> PROVISIONING -> BANNED -> CANCELLED FAILED -> PROVISIONING ->
// CANCELLED BANNED -> RUNNING -> CANCELLED Note that a transition from any
// state to itself is also valid.
type WorkspaceStatus string

const WorkspaceStatusBanned WorkspaceStatus = `BANNED`

const WorkspaceStatusCancelling WorkspaceStatus = `CANCELLING`

const WorkspaceStatusFailed WorkspaceStatus = `FAILED`

const WorkspaceStatusNotProvisioned WorkspaceStatus = `NOT_PROVISIONED`

const WorkspaceStatusProvisioning WorkspaceStatus = `PROVISIONING`

const WorkspaceStatusRunning WorkspaceStatus = `RUNNING`

// String representation for [fmt.Print]
func (f *WorkspaceStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceStatus) Set(v string) error {
	switch v {
	case `BANNED`, `CANCELLING`, `FAILED`, `NOT_PROVISIONED`, `PROVISIONING`, `RUNNING`:
		*f = WorkspaceStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BANNED", "CANCELLING", "FAILED", "NOT_PROVISIONED", "PROVISIONING", "RUNNING"`, v)
	}
}

// Values returns all possible values for WorkspaceStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceStatus) Values() []WorkspaceStatus {
	return []WorkspaceStatus{
		WorkspaceStatusBanned,
		WorkspaceStatusCancelling,
		WorkspaceStatusFailed,
		WorkspaceStatusNotProvisioned,
		WorkspaceStatusProvisioning,
		WorkspaceStatusRunning,
	}
}

// Type always returns WorkspaceStatus to satisfy [pflag.Value] interface
func (f *WorkspaceStatus) Type() string {
	return "WorkspaceStatus"
}
