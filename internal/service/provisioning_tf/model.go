// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package provisioning_tf

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AwsCredentials struct {
	StsRole []StsRole `tfsdk:"sts_role" tf:"optional,object"`
}

func (newState *AwsCredentials) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsCredentials) {
}

func (newState *AwsCredentials) SyncEffectiveFieldsDuringRead(existingState AwsCredentials) {
}

type AwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias types.String `tfsdk:"key_alias" tf:"optional"`
	// The AWS KMS key's Amazon Resource Name (ARN).
	KeyArn types.String `tfsdk:"key_arn" tf:""`
	// The AWS KMS key region.
	KeyRegion types.String `tfsdk:"key_region" tf:""`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to `false`.
	ReuseKeyForClusterVolumes types.Bool `tfsdk:"reuse_key_for_cluster_volumes" tf:"optional"`
}

func (newState *AwsKeyInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsKeyInfo) {
}

func (newState *AwsKeyInfo) SyncEffectiveFieldsDuringRead(existingState AwsKeyInfo) {
}

type AzureWorkspaceInfo struct {
	// Azure Resource Group name
	ResourceGroup types.String `tfsdk:"resource_group" tf:"optional"`
	// Azure Subscription ID
	SubscriptionId types.String `tfsdk:"subscription_id" tf:"optional"`
}

func (newState *AzureWorkspaceInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureWorkspaceInfo) {
}

func (newState *AzureWorkspaceInfo) SyncEffectiveFieldsDuringRead(existingState AzureWorkspaceInfo) {
}

// The general workspace configurations that are specific to cloud providers.
type CloudResourceContainer struct {
	// The general workspace configurations that are specific to Google Cloud.
	Gcp []CustomerFacingGcpCloudResourceContainer `tfsdk:"gcp" tf:"optional,object"`
}

func (newState *CloudResourceContainer) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloudResourceContainer) {
}

func (newState *CloudResourceContainer) SyncEffectiveFieldsDuringRead(existingState CloudResourceContainer) {
}

type CreateAwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias types.String `tfsdk:"key_alias" tf:"optional"`
	// The AWS KMS key's Amazon Resource Name (ARN). Note that the key's AWS
	// region is inferred from the ARN.
	KeyArn types.String `tfsdk:"key_arn" tf:""`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. To not use this key also for encrypting EBS volumes,
	// set this to `false`.
	ReuseKeyForClusterVolumes types.Bool `tfsdk:"reuse_key_for_cluster_volumes" tf:"optional"`
}

func (newState *CreateAwsKeyInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAwsKeyInfo) {
}

func (newState *CreateAwsKeyInfo) SyncEffectiveFieldsDuringRead(existingState CreateAwsKeyInfo) {
}

type CreateCredentialAwsCredentials struct {
	StsRole []CreateCredentialStsRole `tfsdk:"sts_role" tf:"optional,object"`
}

func (newState *CreateCredentialAwsCredentials) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialAwsCredentials) {
}

func (newState *CreateCredentialAwsCredentials) SyncEffectiveFieldsDuringRead(existingState CreateCredentialAwsCredentials) {
}

type CreateCredentialRequest struct {
	AwsCredentials []CreateCredentialAwsCredentials `tfsdk:"aws_credentials" tf:"object"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name" tf:""`
}

func (newState *CreateCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialRequest) {
}

func (newState *CreateCredentialRequest) SyncEffectiveFieldsDuringRead(existingState CreateCredentialRequest) {
}

type CreateCredentialStsRole struct {
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn types.String `tfsdk:"role_arn" tf:"optional"`
}

func (newState *CreateCredentialStsRole) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialStsRole) {
}

func (newState *CreateCredentialStsRole) SyncEffectiveFieldsDuringRead(existingState CreateCredentialStsRole) {
}

type CreateCustomerManagedKeyRequest struct {
	AwsKeyInfo []CreateAwsKeyInfo `tfsdk:"aws_key_info" tf:"optional,object"`

	GcpKeyInfo []CreateGcpKeyInfo `tfsdk:"gcp_key_info" tf:"optional,object"`
	// The cases that the key can be used for.
	UseCases []types.String `tfsdk:"use_cases" tf:""`
}

func (newState *CreateCustomerManagedKeyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCustomerManagedKeyRequest) {
}

func (newState *CreateCustomerManagedKeyRequest) SyncEffectiveFieldsDuringRead(existingState CreateCustomerManagedKeyRequest) {
}

type CreateGcpKeyInfo struct {
	// The GCP KMS key's resource name
	KmsKeyId types.String `tfsdk:"kms_key_id" tf:""`
}

func (newState *CreateGcpKeyInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateGcpKeyInfo) {
}

func (newState *CreateGcpKeyInfo) SyncEffectiveFieldsDuringRead(existingState CreateGcpKeyInfo) {
}

type CreateNetworkRequest struct {
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	GcpNetworkInfo []GcpNetworkInfo `tfsdk:"gcp_network_info" tf:"optional,object"`
	// The human-readable name of the network configuration.
	NetworkName types.String `tfsdk:"network_name" tf:""`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds []types.String `tfsdk:"security_group_ids" tf:"optional"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds []types.String `tfsdk:"subnet_ids" tf:"optional"`
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	VpcEndpoints []NetworkVpcEndpoints `tfsdk:"vpc_endpoints" tf:"optional,object"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId types.String `tfsdk:"vpc_id" tf:"optional"`
}

func (newState *CreateNetworkRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateNetworkRequest) {
}

func (newState *CreateNetworkRequest) SyncEffectiveFieldsDuringRead(existingState CreateNetworkRequest) {
}

type CreateStorageConfigurationRequest struct {
	// Root S3 bucket information.
	RootBucketInfo []RootBucketInfo `tfsdk:"root_bucket_info" tf:"object"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName types.String `tfsdk:"storage_configuration_name" tf:""`
}

func (newState *CreateStorageConfigurationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateStorageConfigurationRequest) {
}

func (newState *CreateStorageConfigurationRequest) SyncEffectiveFieldsDuringRead(existingState CreateStorageConfigurationRequest) {
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId types.String `tfsdk:"aws_vpc_endpoint_id" tf:"optional"`
	// The Google Cloud specific information for this Private Service Connect
	// endpoint.
	GcpVpcEndpointInfo []GcpVpcEndpointInfo `tfsdk:"gcp_vpc_endpoint_info" tf:"optional,object"`
	// The AWS region in which this VPC endpoint object exists.
	Region types.String `tfsdk:"region" tf:"optional"`
	// The human-readable name of the storage configuration.
	VpcEndpointName types.String `tfsdk:"vpc_endpoint_name" tf:""`
}

func (newState *CreateVpcEndpointRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVpcEndpointRequest) {
}

func (newState *CreateVpcEndpointRequest) SyncEffectiveFieldsDuringRead(existingState CreateVpcEndpointRequest) {
}

type CreateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane.
	AwsRegion types.String `tfsdk:"aws_region" tf:"optional"`
	// The cloud provider which the workspace uses. For Google Cloud workspaces,
	// always set this field to `gcp`.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// The general workspace configurations that are specific to cloud
	// providers.
	CloudResourceContainer []CloudResourceContainer `tfsdk:"cloud_resource_container" tf:"optional,object"`
	// ID of the workspace's credential configuration object.
	CredentialsId types.String `tfsdk:"credentials_id" tf:"optional"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for the web application and REST APIs is
	// `<workspace-deployment-name>.cloud.databricks.com`. For example, if the
	// deployment name is `abcsales`, your workspace URL will be
	// `https://abcsales.cloud.databricks.com`. Hyphens are allowed. This
	// property supports only the set of characters that are allowed in a
	// subdomain.
	//
	// To set this value, you must have a deployment name prefix. Contact your
	// Databricks account team to add an account deployment name prefix to your
	// account.
	//
	// Workspace deployment names follow the account prefix and a hyphen. For
	// example, if your account's deployment prefix is `acme` and the workspace
	// deployment name is `workspace-1`, the JSON response for the
	// `deployment_name` field becomes `acme-workspace-1`. The workspace URL
	// would be `acme-workspace-1.cloud.databricks.com`.
	//
	// You can also set the `deployment_name` to the reserved keyword `EMPTY` if
	// you want the deployment name to only include the deployment prefix. For
	// example, if your account's deployment prefix is `acme` and the workspace
	// deployment name is `EMPTY`, the `deployment_name` becomes `acme` only and
	// the workspace URL is `acme.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	//
	// If a new workspace omits this property, the server generates a unique
	// deployment name for you with the pattern `dbc-xxxxxxxx-xxxx`.
	DeploymentName types.String `tfsdk:"deployment_name" tf:"optional"`
	// The network settings for the workspace. The configurations are only for
	// Databricks-managed VPCs. It is ignored if you specify a customer-managed
	// VPC in the `network_id` field.", All the IP range configurations must be
	// mutually exclusive. An attempt to create a workspace fails if Databricks
	// detects an IP range overlap.
	//
	// Specify custom IP ranges in CIDR format. The IP ranges for these fields
	// must not overlap, and all IP addresses must be entirely within the
	// following ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`,
	// `192.168.0.0/16`, and `240.0.0.0/4`.
	//
	// The sizes of these IP ranges affect the maximum number of nodes for the
	// workspace.
	//
	// **Important**: Confirm the IP ranges used by your Databricks workspace
	// before creating the workspace. You cannot change them after your
	// workspace is deployed. If the IP address ranges for your Databricks are
	// too small, IP exhaustion can occur, causing your Databricks jobs to fail.
	// To determine the address range sizes that you need, Databricks provides a
	// calculator as a Microsoft Excel spreadsheet. See [calculate subnet sizes
	// for a new workspace].
	//
	// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
	GcpManagedNetworkConfig []GcpManagedNetworkConfig `tfsdk:"gcp_managed_network_config" tf:"optional,object"`
	// The configurations for the GKE cluster of a Databricks workspace.
	GkeConfig []GkeConfig `tfsdk:"gke_config" tf:"optional,object"`
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	Location types.String `tfsdk:"location" tf:"optional"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to help protect and control access to the
	// workspace's notebooks, secrets, Databricks SQL queries, and query
	// history. The provided key configuration object property `use_cases` must
	// contain `MANAGED_SERVICES`.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id" tf:"optional"`

	NetworkId types.String `tfsdk:"network_id" tf:"optional"`
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing].
	//
	// [AWS Pricing]: https://databricks.com/product/aws-pricing
	PricingTier types.String `tfsdk:"pricing_tier" tf:"optional"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. This ID must be specified for customers using [AWS
	// PrivateLink] for either front-end (user-to-workspace connection),
	// back-end (data plane to control plane connection), or both connection
	// types.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id" tf:"optional"`
	// The ID of the workspace's storage configuration object.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id" tf:"optional"`
	// The ID of the workspace's storage encryption key configuration object.
	// This is used to encrypt the workspace's root S3 bucket (root DBFS and
	// system data) and, optionally, cluster EBS volumes. The provided key
	// configuration object property `use_cases` must contain `STORAGE`.
	StorageCustomerManagedKeyId types.String `tfsdk:"storage_customer_managed_key_id" tf:"optional"`
	// The workspace's human-readable name.
	WorkspaceName types.String `tfsdk:"workspace_name" tf:""`
}

func (newState *CreateWorkspaceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWorkspaceRequest) {
}

func (newState *CreateWorkspaceRequest) SyncEffectiveFieldsDuringRead(existingState CreateWorkspaceRequest) {
}

type Credential struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`

	AwsCredentials []AwsCredentials `tfsdk:"aws_credentials" tf:"optional,object"`
	// Time in epoch milliseconds when the credential was created.
	CreationTime          types.Int64 `tfsdk:"creation_time" tf:"optional"`
	EffectiveCreationTime types.Int64 `tfsdk:"effective_creation_time" tf:"computed,optional"`
	// Databricks credential configuration ID.
	CredentialsId types.String `tfsdk:"credentials_id" tf:"optional"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name" tf:"optional"`
}

func (newState *Credential) SyncEffectiveFieldsDuringCreateOrUpdate(plan Credential) {
	newState.EffectiveCreationTime = newState.CreationTime
	newState.CreationTime = plan.CreationTime
}

func (newState *Credential) SyncEffectiveFieldsDuringRead(existingState Credential) {
	newState.EffectiveCreationTime = existingState.EffectiveCreationTime
	if existingState.EffectiveCreationTime.ValueInt64() == newState.CreationTime.ValueInt64() {
		newState.CreationTime = existingState.CreationTime
	}
}

// The general workspace configurations that are specific to Google Cloud.
type CustomerFacingGcpCloudResourceContainer struct {
	// The Google Cloud project ID, which the workspace uses to instantiate
	// cloud resources for your workspace.
	ProjectId types.String `tfsdk:"project_id" tf:"optional"`
}

func (newState *CustomerFacingGcpCloudResourceContainer) SyncEffectiveFieldsDuringCreateOrUpdate(plan CustomerFacingGcpCloudResourceContainer) {
}

func (newState *CustomerFacingGcpCloudResourceContainer) SyncEffectiveFieldsDuringRead(existingState CustomerFacingGcpCloudResourceContainer) {
}

type CustomerManagedKey struct {
	// The Databricks account ID that holds the customer-managed key.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`

	AwsKeyInfo []AwsKeyInfo `tfsdk:"aws_key_info" tf:"optional,object"`
	// Time in epoch milliseconds when the customer key was created.
	CreationTime          types.Int64 `tfsdk:"creation_time" tf:"optional"`
	EffectiveCreationTime types.Int64 `tfsdk:"effective_creation_time" tf:"computed,optional"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId types.String `tfsdk:"customer_managed_key_id" tf:"optional"`

	GcpKeyInfo []GcpKeyInfo `tfsdk:"gcp_key_info" tf:"optional,object"`
	// The cases that the key can be used for.
	UseCases []types.String `tfsdk:"use_cases" tf:"optional"`
}

func (newState *CustomerManagedKey) SyncEffectiveFieldsDuringCreateOrUpdate(plan CustomerManagedKey) {
	newState.EffectiveCreationTime = newState.CreationTime
	newState.CreationTime = plan.CreationTime
}

func (newState *CustomerManagedKey) SyncEffectiveFieldsDuringRead(existingState CustomerManagedKey) {
	newState.EffectiveCreationTime = existingState.EffectiveCreationTime
	if existingState.EffectiveCreationTime.ValueInt64() == newState.CreationTime.ValueInt64() {
		newState.CreationTime = existingState.CreationTime
	}
}

// Delete credential configuration
type DeleteCredentialRequest struct {
	// Databricks Account API credential configuration ID
	CredentialsId types.String `tfsdk:"-"`
}

func (newState *DeleteCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCredentialRequest) {
}

func (newState *DeleteCredentialRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCredentialRequest) {
}

// Delete encryption key configuration
type DeleteEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId types.String `tfsdk:"-"`
}

func (newState *DeleteEncryptionKeyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteEncryptionKeyRequest) {
}

func (newState *DeleteEncryptionKeyRequest) SyncEffectiveFieldsDuringRead(existingState DeleteEncryptionKeyRequest) {
}

// Delete a network configuration
type DeleteNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId types.String `tfsdk:"-"`
}

func (newState *DeleteNetworkRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteNetworkRequest) {
}

func (newState *DeleteNetworkRequest) SyncEffectiveFieldsDuringRead(existingState DeleteNetworkRequest) {
}

// Delete a private access settings object
type DeletePrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (newState *DeletePrivateAccesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePrivateAccesRequest) {
}

func (newState *DeletePrivateAccesRequest) SyncEffectiveFieldsDuringRead(existingState DeletePrivateAccesRequest) {
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

// Delete storage configuration
type DeleteStorageRequest struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"-"`
}

func (newState *DeleteStorageRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteStorageRequest) {
}

func (newState *DeleteStorageRequest) SyncEffectiveFieldsDuringRead(existingState DeleteStorageRequest) {
}

// Delete VPC endpoint configuration
type DeleteVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId types.String `tfsdk:"-"`
}

func (newState *DeleteVpcEndpointRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteVpcEndpointRequest) {
}

func (newState *DeleteVpcEndpointRequest) SyncEffectiveFieldsDuringRead(existingState DeleteVpcEndpointRequest) {
}

// Delete a workspace
type DeleteWorkspaceRequest struct {
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *DeleteWorkspaceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWorkspaceRequest) {
}

func (newState *DeleteWorkspaceRequest) SyncEffectiveFieldsDuringRead(existingState DeleteWorkspaceRequest) {
}

type GcpKeyInfo struct {
	// The GCP KMS key's resource name
	KmsKeyId types.String `tfsdk:"kms_key_id" tf:""`
}

func (newState *GcpKeyInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpKeyInfo) {
}

func (newState *GcpKeyInfo) SyncEffectiveFieldsDuringRead(existingState GcpKeyInfo) {
}

// The network settings for the workspace. The configurations are only for
// Databricks-managed VPCs. It is ignored if you specify a customer-managed VPC
// in the `network_id` field.", All the IP range configurations must be mutually
// exclusive. An attempt to create a workspace fails if Databricks detects an IP
// range overlap.
//
// Specify custom IP ranges in CIDR format. The IP ranges for these fields must
// not overlap, and all IP addresses must be entirely within the following
// ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`, `192.168.0.0/16`, and
// `240.0.0.0/4`.
//
// The sizes of these IP ranges affect the maximum number of nodes for the
// workspace.
//
// **Important**: Confirm the IP ranges used by your Databricks workspace before
// creating the workspace. You cannot change them after your workspace is
// deployed. If the IP address ranges for your Databricks are too small, IP
// exhaustion can occur, causing your Databricks jobs to fail. To determine the
// address range sizes that you need, Databricks provides a calculator as a
// Microsoft Excel spreadsheet. See [calculate subnet sizes for a new
// workspace].
//
// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
type GcpManagedNetworkConfig struct {
	// The IP range from which to allocate GKE cluster pods. No bigger than `/9`
	// and no smaller than `/21`.
	GkeClusterPodIpRange types.String `tfsdk:"gke_cluster_pod_ip_range" tf:"optional"`
	// The IP range from which to allocate GKE cluster services. No bigger than
	// `/16` and no smaller than `/27`.
	GkeClusterServiceIpRange types.String `tfsdk:"gke_cluster_service_ip_range" tf:"optional"`
	// The IP range from which to allocate GKE cluster nodes. No bigger than
	// `/9` and no smaller than `/29`.
	SubnetCidr types.String `tfsdk:"subnet_cidr" tf:"optional"`
}

func (newState *GcpManagedNetworkConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpManagedNetworkConfig) {
}

func (newState *GcpManagedNetworkConfig) SyncEffectiveFieldsDuringRead(existingState GcpManagedNetworkConfig) {
}

// The Google Cloud specific information for this network (for example, the VPC
// ID, subnet ID, and secondary IP ranges).
type GcpNetworkInfo struct {
	// The Google Cloud project ID of the VPC network.
	NetworkProjectId types.String `tfsdk:"network_project_id" tf:""`
	// The name of the secondary IP range for pods. A Databricks-managed GKE
	// cluster uses this IP range for its pods. This secondary IP range can be
	// used by only one workspace.
	PodIpRangeName types.String `tfsdk:"pod_ip_range_name" tf:""`
	// The name of the secondary IP range for services. A Databricks-managed GKE
	// cluster uses this IP range for its services. This secondary IP range can
	// be used by only one workspace.
	ServiceIpRangeName types.String `tfsdk:"service_ip_range_name" tf:""`
	// The ID of the subnet associated with this network.
	SubnetId types.String `tfsdk:"subnet_id" tf:""`
	// The Google Cloud region of the workspace data plane (for example,
	// `us-east4`).
	SubnetRegion types.String `tfsdk:"subnet_region" tf:""`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId types.String `tfsdk:"vpc_id" tf:""`
}

func (newState *GcpNetworkInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpNetworkInfo) {
}

func (newState *GcpNetworkInfo) SyncEffectiveFieldsDuringRead(existingState GcpNetworkInfo) {
}

// The Google Cloud specific information for this Private Service Connect
// endpoint.
type GcpVpcEndpointInfo struct {
	// Region of the PSC endpoint.
	EndpointRegion types.String `tfsdk:"endpoint_region" tf:""`
	// The Google Cloud project ID of the VPC network where the PSC connection
	// resides.
	ProjectId types.String `tfsdk:"project_id" tf:""`
	// The unique ID of this PSC connection.
	PscConnectionId types.String `tfsdk:"psc_connection_id" tf:"optional"`
	// The name of the PSC endpoint in the Google Cloud project.
	PscEndpointName types.String `tfsdk:"psc_endpoint_name" tf:""`
	// The service attachment this PSC connection connects to.
	ServiceAttachmentId types.String `tfsdk:"service_attachment_id" tf:"optional"`
}

func (newState *GcpVpcEndpointInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpVpcEndpointInfo) {
}

func (newState *GcpVpcEndpointInfo) SyncEffectiveFieldsDuringRead(existingState GcpVpcEndpointInfo) {
}

// Get credential configuration
type GetCredentialRequest struct {
	// Databricks Account API credential configuration ID
	CredentialsId types.String `tfsdk:"-"`
}

func (newState *GetCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCredentialRequest) {
}

func (newState *GetCredentialRequest) SyncEffectiveFieldsDuringRead(existingState GetCredentialRequest) {
}

// Get encryption key configuration
type GetEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId types.String `tfsdk:"-"`
}

func (newState *GetEncryptionKeyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetEncryptionKeyRequest) {
}

func (newState *GetEncryptionKeyRequest) SyncEffectiveFieldsDuringRead(existingState GetEncryptionKeyRequest) {
}

// Get a network configuration
type GetNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId types.String `tfsdk:"-"`
}

func (newState *GetNetworkRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetNetworkRequest) {
}

func (newState *GetNetworkRequest) SyncEffectiveFieldsDuringRead(existingState GetNetworkRequest) {
}

// Get a private access settings object
type GetPrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (newState *GetPrivateAccesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPrivateAccesRequest) {
}

func (newState *GetPrivateAccesRequest) SyncEffectiveFieldsDuringRead(existingState GetPrivateAccesRequest) {
}

// Get storage configuration
type GetStorageRequest struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"-"`
}

func (newState *GetStorageRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetStorageRequest) {
}

func (newState *GetStorageRequest) SyncEffectiveFieldsDuringRead(existingState GetStorageRequest) {
}

// Get a VPC endpoint configuration
type GetVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId types.String `tfsdk:"-"`
}

func (newState *GetVpcEndpointRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetVpcEndpointRequest) {
}

func (newState *GetVpcEndpointRequest) SyncEffectiveFieldsDuringRead(existingState GetVpcEndpointRequest) {
}

// Get a workspace
type GetWorkspaceRequest struct {
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *GetWorkspaceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceRequest) {
}

func (newState *GetWorkspaceRequest) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceRequest) {
}

// The configurations for the GKE cluster of a Databricks workspace.
type GkeConfig struct {
	// Specifies the network connectivity types for the GKE nodes and the GKE
	// master network.
	//
	// Set to `PRIVATE_NODE_PUBLIC_MASTER` for a private GKE cluster for the
	// workspace. The GKE nodes will not have public IPs.
	//
	// Set to `PUBLIC_NODE_PUBLIC_MASTER` for a public GKE cluster. The nodes of
	// a public GKE cluster have public IP addresses.
	ConnectivityType types.String `tfsdk:"connectivity_type" tf:"optional"`
	// The IP range from which to allocate GKE cluster master resources. This
	// field will be ignored if GKE private cluster is not enabled.
	//
	// It must be exactly as big as `/28`.
	MasterIpRange types.String `tfsdk:"master_ip_range" tf:"optional"`
}

func (newState *GkeConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan GkeConfig) {
}

func (newState *GkeConfig) SyncEffectiveFieldsDuringRead(existingState GkeConfig) {
}

type Network struct {
	// The Databricks account ID associated with this network configuration.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// Time in epoch milliseconds when the network was created.
	CreationTime          types.Int64 `tfsdk:"creation_time" tf:"optional"`
	EffectiveCreationTime types.Int64 `tfsdk:"effective_creation_time" tf:"computed,optional"`
	// Array of error messages about the network configuration.
	ErrorMessages          []NetworkHealth `tfsdk:"error_messages" tf:"optional"`
	EffectiveErrorMessages []NetworkHealth `tfsdk:"effective_error_messages" tf:"computed,optional"`
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	GcpNetworkInfo []GcpNetworkInfo `tfsdk:"gcp_network_info" tf:"optional,object"`
	// The Databricks network configuration ID.
	NetworkId types.String `tfsdk:"network_id" tf:"optional"`
	// The human-readable name of the network configuration.
	NetworkName types.String `tfsdk:"network_name" tf:"optional"`

	SecurityGroupIds []types.String `tfsdk:"security_group_ids" tf:"optional"`

	SubnetIds []types.String `tfsdk:"subnet_ids" tf:"optional"`
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	VpcEndpoints []NetworkVpcEndpoints `tfsdk:"vpc_endpoints" tf:"optional,object"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId types.String `tfsdk:"vpc_id" tf:"optional"`
	// The status of this network configuration object in terms of its use in a
	// workspace: * `UNATTACHED`: Unattached. * `VALID`: Valid. * `BROKEN`:
	// Broken. * `WARNED`: Warned.
	VpcStatus          types.String `tfsdk:"vpc_status" tf:"optional"`
	EffectiveVpcStatus types.String `tfsdk:"effective_vpc_status" tf:"computed,optional"`
	// Array of warning messages about the network configuration.
	WarningMessages          []NetworkWarning `tfsdk:"warning_messages" tf:"optional"`
	EffectiveWarningMessages []NetworkWarning `tfsdk:"effective_warning_messages" tf:"computed,optional"`
	// Workspace ID associated with this network configuration.
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:"optional"`
}

func (newState *Network) SyncEffectiveFieldsDuringCreateOrUpdate(plan Network) {
	newState.EffectiveCreationTime = newState.CreationTime
	newState.CreationTime = plan.CreationTime
	newState.EffectiveVpcStatus = newState.VpcStatus
	newState.VpcStatus = plan.VpcStatus
}

func (newState *Network) SyncEffectiveFieldsDuringRead(existingState Network) {
	newState.EffectiveCreationTime = existingState.EffectiveCreationTime
	if existingState.EffectiveCreationTime.ValueInt64() == newState.CreationTime.ValueInt64() {
		newState.CreationTime = existingState.CreationTime
	}
	newState.EffectiveVpcStatus = existingState.EffectiveVpcStatus
	if existingState.EffectiveVpcStatus.ValueString() == newState.VpcStatus.ValueString() {
		newState.VpcStatus = existingState.VpcStatus
	}
}

type NetworkHealth struct {
	// Details of the error.
	ErrorMessage types.String `tfsdk:"error_message" tf:"optional"`
	// The AWS resource associated with this error: credentials, VPC, subnet,
	// security group, or network ACL.
	ErrorType types.String `tfsdk:"error_type" tf:"optional"`
}

func (newState *NetworkHealth) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkHealth) {
}

func (newState *NetworkHealth) SyncEffectiveFieldsDuringRead(existingState NetworkHealth) {
}

// If specified, contains the VPC endpoints used to allow cluster communication
// from this VPC over [AWS PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	DataplaneRelay []types.String `tfsdk:"dataplane_relay" tf:""`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	RestApi []types.String `tfsdk:"rest_api" tf:""`
}

func (newState *NetworkVpcEndpoints) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkVpcEndpoints) {
}

func (newState *NetworkVpcEndpoints) SyncEffectiveFieldsDuringRead(existingState NetworkVpcEndpoints) {
}

type NetworkWarning struct {
	// Details of the warning.
	WarningMessage types.String `tfsdk:"warning_message" tf:"optional"`
	// The AWS resource associated with this warning: a subnet or a security
	// group.
	WarningType types.String `tfsdk:"warning_type" tf:"optional"`
}

func (newState *NetworkWarning) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkWarning) {
}

func (newState *NetworkWarning) SyncEffectiveFieldsDuringRead(existingState NetworkWarning) {
}

type PrivateAccessSettings struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// An array of Databricks VPC endpoint IDs.
	AllowedVpcEndpointIds []types.String `tfsdk:"allowed_vpc_endpoint_ids" tf:"optional"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// * `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see `allowed_vpc_endpoint_ids`.
	PrivateAccessLevel types.String `tfsdk:"private_access_level" tf:"optional"`
	// Databricks private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id" tf:"optional"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName types.String `tfsdk:"private_access_settings_name" tf:"optional"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled types.Bool `tfsdk:"public_access_enabled" tf:"optional"`
	// The cloud region for workspaces attached to this private access settings
	// object.
	Region types.String `tfsdk:"region" tf:"optional"`
}

func (newState *PrivateAccessSettings) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrivateAccessSettings) {
}

func (newState *PrivateAccessSettings) SyncEffectiveFieldsDuringRead(existingState PrivateAccessSettings) {
}

type ReplaceResponse struct {
}

func (newState *ReplaceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReplaceResponse) {
}

func (newState *ReplaceResponse) SyncEffectiveFieldsDuringRead(existingState ReplaceResponse) {
}

// Root S3 bucket information.
type RootBucketInfo struct {
	// The name of the S3 bucket.
	BucketName types.String `tfsdk:"bucket_name" tf:"optional"`
}

func (newState *RootBucketInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RootBucketInfo) {
}

func (newState *RootBucketInfo) SyncEffectiveFieldsDuringRead(existingState RootBucketInfo) {
}

type StorageConfiguration struct {
	// The Databricks account ID that hosts the credential.
	AccountId          types.String `tfsdk:"account_id" tf:"optional"`
	EffectiveAccountId types.String `tfsdk:"effective_account_id" tf:"computed,optional"`
	// Time in epoch milliseconds when the storage configuration was created.
	CreationTime          types.Int64 `tfsdk:"creation_time" tf:"optional"`
	EffectiveCreationTime types.Int64 `tfsdk:"effective_creation_time" tf:"computed,optional"`
	// Root S3 bucket information.
	RootBucketInfo []RootBucketInfo `tfsdk:"root_bucket_info" tf:"optional,object"`
	// Databricks storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id" tf:"optional"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName types.String `tfsdk:"storage_configuration_name" tf:"optional"`
}

func (newState *StorageConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan StorageConfiguration) {
	newState.EffectiveAccountId = newState.AccountId
	newState.AccountId = plan.AccountId
	newState.EffectiveCreationTime = newState.CreationTime
	newState.CreationTime = plan.CreationTime
}

func (newState *StorageConfiguration) SyncEffectiveFieldsDuringRead(existingState StorageConfiguration) {
	newState.EffectiveAccountId = existingState.EffectiveAccountId
	if existingState.EffectiveAccountId.ValueString() == newState.AccountId.ValueString() {
		newState.AccountId = existingState.AccountId
	}
	newState.EffectiveCreationTime = existingState.EffectiveCreationTime
	if existingState.EffectiveCreationTime.ValueInt64() == newState.CreationTime.ValueInt64() {
		newState.CreationTime = existingState.CreationTime
	}
}

type StsRole struct {
	// The external ID that needs to be trusted by the cross-account role. This
	// is always your Databricks account ID.
	ExternalId types.String `tfsdk:"external_id" tf:"optional"`
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn types.String `tfsdk:"role_arn" tf:"optional"`
}

func (newState *StsRole) SyncEffectiveFieldsDuringCreateOrUpdate(plan StsRole) {
}

func (newState *StsRole) SyncEffectiveFieldsDuringRead(existingState StsRole) {
}

type UpdateResponse struct {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse) {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringRead(existingState UpdateResponse) {
}

type UpdateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane (for example, `us-west-2`).
	// This parameter is available only for updating failed workspaces.
	AwsRegion types.String `tfsdk:"aws_region" tf:"optional"`
	// ID of the workspace's credential configuration object. This parameter is
	// available for updating both failed and running workspaces.
	CredentialsId types.String `tfsdk:"credentials_id" tf:"optional"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This parameter is available only for updating failed workspaces.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id" tf:"optional"`

	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id" tf:"optional"`
	// The ID of the workspace's network configuration object. Used only if you
	// already use a customer-managed VPC. For failed workspaces only, you can
	// switch from a Databricks-managed VPC to a customer-managed VPC by
	// updating the workspace to add a network configuration ID.
	NetworkId types.String `tfsdk:"network_id" tf:"optional"`
	// The ID of the workspace's storage configuration object. This parameter is
	// available only for updating failed workspaces.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id" tf:"optional"`
	// The ID of the key configuration object for workspace storage. This
	// parameter is available for updating both failed and running workspaces.
	StorageCustomerManagedKeyId types.String `tfsdk:"storage_customer_managed_key_id" tf:"optional"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *UpdateWorkspaceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceRequest) {
}

func (newState *UpdateWorkspaceRequest) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceRequest) {
}

type UpsertPrivateAccessSettingsRequest struct {
	// An array of Databricks VPC endpoint IDs. This is the Databricks ID that
	// is returned when registering the VPC endpoint configuration in your
	// Databricks account. This is not the ID of the VPC endpoint in AWS.
	//
	// Only used when `private_access_level` is set to `ENDPOINT`. This is an
	// allow list of VPC endpoints that in your account that can connect to your
	// workspace over AWS PrivateLink.
	//
	// If hybrid access to your workspace is enabled by setting
	// `public_access_enabled` to `true`, this control only works for
	// PrivateLink connections. To control how your workspace is accessed via
	// public internet, see [IP access lists].
	//
	// [IP access lists]: https://docs.databricks.com/security/network/ip-access-list.html
	AllowedVpcEndpointIds []types.String `tfsdk:"allowed_vpc_endpoint_ids" tf:"optional"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// * `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see `allowed_vpc_endpoint_ids`.
	PrivateAccessLevel types.String `tfsdk:"private_access_level" tf:"optional"`
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"-"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName types.String `tfsdk:"private_access_settings_name" tf:""`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled types.Bool `tfsdk:"public_access_enabled" tf:"optional"`
	// The cloud region for workspaces associated with this private access
	// settings object.
	Region types.String `tfsdk:"region" tf:""`
}

func (newState *UpsertPrivateAccessSettingsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpsertPrivateAccessSettingsRequest) {
}

func (newState *UpsertPrivateAccessSettingsRequest) SyncEffectiveFieldsDuringRead(existingState UpsertPrivateAccessSettingsRequest) {
}

type VpcEndpoint struct {
	// The Databricks account ID that hosts the VPC endpoint configuration.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// The AWS Account in which the VPC endpoint object exists.
	AwsAccountId types.String `tfsdk:"aws_account_id" tf:"optional"`
	// The ID of the Databricks [endpoint service] that this VPC endpoint is
	// connected to. For a list of endpoint service IDs for each supported AWS
	// region, see the [Databricks PrivateLink documentation].
	//
	// [Databricks PrivateLink documentation]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	AwsEndpointServiceId types.String `tfsdk:"aws_endpoint_service_id" tf:"optional"`
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId types.String `tfsdk:"aws_vpc_endpoint_id" tf:"optional"`
	// The Google Cloud specific information for this Private Service Connect
	// endpoint.
	GcpVpcEndpointInfo []GcpVpcEndpointInfo `tfsdk:"gcp_vpc_endpoint_info" tf:"optional,object"`
	// The AWS region in which this VPC endpoint object exists.
	Region types.String `tfsdk:"region" tf:"optional"`
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	State types.String `tfsdk:"state" tf:"optional"`
	// This enumeration represents the type of Databricks VPC [endpoint service]
	// that was used when creating this VPC endpoint.
	//
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	UseCase types.String `tfsdk:"use_case" tf:"optional"`
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	VpcEndpointId types.String `tfsdk:"vpc_endpoint_id" tf:"optional"`
	// The human-readable name of the storage configuration.
	VpcEndpointName types.String `tfsdk:"vpc_endpoint_name" tf:"optional"`
}

func (newState *VpcEndpoint) SyncEffectiveFieldsDuringCreateOrUpdate(plan VpcEndpoint) {
}

func (newState *VpcEndpoint) SyncEffectiveFieldsDuringRead(existingState VpcEndpoint) {
}

type Workspace struct {
	// Databricks account ID.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// The AWS region of the workspace data plane (for example, `us-west-2`).
	AwsRegion types.String `tfsdk:"aws_region" tf:"optional"`

	AzureWorkspaceInfo []AzureWorkspaceInfo `tfsdk:"azure_workspace_info" tf:"optional,object"`
	// The cloud name. This field always has the value `gcp`.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// The general workspace configurations that are specific to cloud
	// providers.
	CloudResourceContainer []CloudResourceContainer `tfsdk:"cloud_resource_container" tf:"optional,object"`
	// Time in epoch milliseconds when the workspace was created.
	CreationTime          types.Int64 `tfsdk:"creation_time" tf:"optional"`
	EffectiveCreationTime types.Int64 `tfsdk:"effective_creation_time" tf:"computed,optional"`
	// ID of the workspace's credential configuration object.
	CredentialsId types.String `tfsdk:"credentials_id" tf:"optional"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `<deployment-name>.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	DeploymentName types.String `tfsdk:"deployment_name" tf:"optional"`
	// The network settings for the workspace. The configurations are only for
	// Databricks-managed VPCs. It is ignored if you specify a customer-managed
	// VPC in the `network_id` field.", All the IP range configurations must be
	// mutually exclusive. An attempt to create a workspace fails if Databricks
	// detects an IP range overlap.
	//
	// Specify custom IP ranges in CIDR format. The IP ranges for these fields
	// must not overlap, and all IP addresses must be entirely within the
	// following ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`,
	// `192.168.0.0/16`, and `240.0.0.0/4`.
	//
	// The sizes of these IP ranges affect the maximum number of nodes for the
	// workspace.
	//
	// **Important**: Confirm the IP ranges used by your Databricks workspace
	// before creating the workspace. You cannot change them after your
	// workspace is deployed. If the IP address ranges for your Databricks are
	// too small, IP exhaustion can occur, causing your Databricks jobs to fail.
	// To determine the address range sizes that you need, Databricks provides a
	// calculator as a Microsoft Excel spreadsheet. See [calculate subnet sizes
	// for a new workspace].
	//
	// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
	GcpManagedNetworkConfig []GcpManagedNetworkConfig `tfsdk:"gcp_managed_network_config" tf:"optional,object"`
	// The configurations for the GKE cluster of a Databricks workspace.
	GkeConfig []GkeConfig `tfsdk:"gke_config" tf:"optional,object"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location types.String `tfsdk:"location" tf:"optional"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id" tf:"optional"`
	// The network configuration ID that is attached to the workspace. This
	// field is available only if the network is a customer-managed network.
	NetworkId types.String `tfsdk:"network_id" tf:"optional"`
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing].
	//
	// [AWS Pricing]: https://databricks.com/product/aws-pricing
	PricingTier types.String `tfsdk:"pricing_tier" tf:"optional"`
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
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id" tf:"optional"`
	// ID of the workspace's storage configuration object.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id" tf:"optional"`
	// ID of the key configuration for encrypting workspace storage.
	StorageCustomerManagedKeyId types.String `tfsdk:"storage_customer_managed_key_id" tf:"optional"`
	// A unique integer ID for the workspace
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:"optional"`
	// The human-readable name of the workspace.
	WorkspaceName types.String `tfsdk:"workspace_name" tf:"optional"`
	// The status of the workspace. For workspace creation, usually it is set to
	// `PROVISIONING` initially. Continue to check the status until the status
	// is `RUNNING`.
	WorkspaceStatus          types.String `tfsdk:"workspace_status" tf:"optional"`
	EffectiveWorkspaceStatus types.String `tfsdk:"effective_workspace_status" tf:"computed,optional"`
	// Message describing the current workspace status.
	WorkspaceStatusMessage          types.String `tfsdk:"workspace_status_message" tf:"optional"`
	EffectiveWorkspaceStatusMessage types.String `tfsdk:"effective_workspace_status_message" tf:"computed,optional"`
}

func (newState *Workspace) SyncEffectiveFieldsDuringCreateOrUpdate(plan Workspace) {
	newState.EffectiveCreationTime = newState.CreationTime
	newState.CreationTime = plan.CreationTime
	newState.EffectiveWorkspaceStatus = newState.WorkspaceStatus
	newState.WorkspaceStatus = plan.WorkspaceStatus
	newState.EffectiveWorkspaceStatusMessage = newState.WorkspaceStatusMessage
	newState.WorkspaceStatusMessage = plan.WorkspaceStatusMessage
}

func (newState *Workspace) SyncEffectiveFieldsDuringRead(existingState Workspace) {
	newState.EffectiveCreationTime = existingState.EffectiveCreationTime
	if existingState.EffectiveCreationTime.ValueInt64() == newState.CreationTime.ValueInt64() {
		newState.CreationTime = existingState.CreationTime
	}
	newState.EffectiveWorkspaceStatus = existingState.EffectiveWorkspaceStatus
	if existingState.EffectiveWorkspaceStatus.ValueString() == newState.WorkspaceStatus.ValueString() {
		newState.WorkspaceStatus = existingState.WorkspaceStatus
	}
	newState.EffectiveWorkspaceStatusMessage = existingState.EffectiveWorkspaceStatusMessage
	if existingState.EffectiveWorkspaceStatusMessage.ValueString() == newState.WorkspaceStatusMessage.ValueString() {
		newState.WorkspaceStatusMessage = existingState.WorkspaceStatusMessage
	}
}
