// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"context"
)

// These APIs manage credential configurations for this workspace. Databricks
// needs access to a cross-account service IAM role in your AWS account so that
// Databricks can deploy clusters in the appropriate VPC for the new workspace.
// A credential configuration encapsulates this role information, and its ID is
// used when creating a new workspace.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CredentialsService interface {

	// Creates a Databricks credential configuration that represents cloud
	// cross-account credentials for a specified account. Databricks uses this
	// to set up network infrastructure properly to host Databricks clusters.
	// For your AWS IAM role, you need to trust the External ID (the Databricks
	// Account API account ID) in the returned credential object, and configure
	// the required access policy.
	//
	// Save the response's `credentials_id` field, which is the ID for your new
	// credential configuration object.
	//
	// For information about how to create a new workspace with this API, see
	// [Create a new workspace using the Account API]
	//
	// [Create a new workspace using the Account API]: http://docs.databricks.com/administration-guide/account-api/new-workspace.html
	Create(ctx context.Context, request CreateCredentialRequest) (*Credential, error)

	// Deletes a Databricks credential configuration object for an account, both
	// specified by ID. You cannot delete a credential that is associated with
	// any workspace.
	Delete(ctx context.Context, request DeleteCredentialRequest) (*Credential, error)

	// Gets a Databricks credential configuration object for an account, both
	// specified by ID.
	Get(ctx context.Context, request GetCredentialRequest) (*Credential, error)

	// List Databricks credential configuration objects for an account,
	// specified by ID.
	List(ctx context.Context) ([]Credential, error)
}

// These APIs manage encryption key configurations for this workspace
// (optional). A key configuration encapsulates the AWS KMS key information and
// some information about how the key configuration can be used. There are two
// possible uses for key configurations:
//
// * Managed services: A key configuration can be used to encrypt a workspace's
// notebook and secret data in the control plane, as well as Databricks SQL
// queries and query history. * Storage: A key configuration can be used to
// encrypt a workspace's DBFS and EBS data in the data plane.
//
// In both of these cases, the key configuration's ID is used when creating a
// new workspace. This Preview feature is available if your account is on the E2
// version of the platform. Updating a running workspace with workspace storage
// encryption requires that the workspace is on the E2 version of the platform.
// If you have an older workspace, it might not be on the E2 version of the
// platform. If you are not sure, contact your Databricks representative.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type EncryptionKeysService interface {

	// Creates a customer-managed key configuration object for an account,
	// specified by ID. This operation uploads a reference to a customer-managed
	// key to Databricks. If the key is assigned as a workspace's
	// customer-managed key for managed services, Databricks uses the key to
	// encrypt the workspaces notebooks and secrets in the control plane, in
	// addition to Databricks SQL queries and query history. If it is specified
	// as a workspace's customer-managed key for workspace storage, the key
	// encrypts the workspace's root S3 bucket (which contains the workspace's
	// root DBFS and system data) and, optionally, cluster EBS volume data.
	//
	// **Important**: Customer-managed keys are supported only for some
	// deployment types, subscription types, and AWS regions that currently
	// support creation of Databricks workspaces.
	//
	// This operation is available only if your account is on the E2 version of
	// the platform or on a select custom plan that allows multiple workspaces
	// per account.
	//
	// **GCP only**: To create a customer-managed key on GCP, you must include
	// the `X-Databricks-GCP-SA-Access-Token` HTTP header in your request. This
	// header must contain a Google Cloud OAuth access token with the
	// `cloud-platform` scope. The Google identity associated with the token
	// must also have the `setIamPermissions` and `getIamPermissions` IAM
	// permissions on the key resource. For details on obtaining this token, see
	// [Authenticate with Google ID tokens].
	//
	// [Authenticate with Google ID tokens]: https://docs.databricks.com/gcp/en/dev-tools/auth/authentication-google-id.html
	Create(ctx context.Context, request CreateCustomerManagedKeyRequest) (*CustomerManagedKey, error)

	// Deletes a customer-managed key configuration object for an account. You
	// cannot delete a configuration that is associated with a running
	// workspace.
	Delete(ctx context.Context, request DeleteEncryptionKeyRequest) (*CustomerManagedKey, error)

	// Gets a customer-managed key configuration object for an account,
	// specified by ID. This operation uploads a reference to a customer-managed
	// key to Databricks. If assigned as a workspace's customer-managed key for
	// managed services, Databricks uses the key to encrypt the workspaces
	// notebooks and secrets in the control plane, in addition to Databricks SQL
	// queries and query history. If it is specified as a workspace's
	// customer-managed key for storage, the key encrypts the workspace's root
	// S3 bucket (which contains the workspace's root DBFS and system data) and,
	// optionally, cluster EBS volume data.
	//
	// **Important**: Customer-managed keys are supported only for some
	// deployment types, subscription types, and AWS regions.
	//
	// This operation is available only if your account is on the E2 version of
	// the platform.",
	Get(ctx context.Context, request GetEncryptionKeyRequest) (*CustomerManagedKey, error)

	// Lists Databricks customer-managed key configurations for an account.
	List(ctx context.Context) ([]CustomerManagedKey, error)
}

// These APIs manage network configurations for customer-managed VPCs
// (optional). Its ID is used when creating a new workspace if you use
// customer-managed VPCs.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type NetworksService interface {

	// Creates a Databricks network configuration that represents an VPC and its
	// resources. The VPC will be used for new Databricks clusters. This
	// requires a pre-existing VPC and subnets.
	Create(ctx context.Context, request CreateNetworkRequest) (*Network, error)

	// Deletes a Databricks network configuration, which represents a cloud VPC
	// and its resources. You cannot delete a network that is associated with a
	// workspace.
	//
	// This operation is available only if your account is on the E2 version of
	// the platform.
	Delete(ctx context.Context, request DeleteNetworkRequest) (*Network, error)

	// Gets a Databricks network configuration, which represents a cloud VPC and
	// its resources.
	Get(ctx context.Context, request GetNetworkRequest) (*Network, error)

	// Lists Databricks network configurations for an account.
	List(ctx context.Context) ([]Network, error)
}

// These APIs manage private access settings for this account.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type PrivateAccessService interface {

	// Creates a private access settings configuration, which represents network
	// access restrictions for workspace resources. Private access settings
	// configure whether workspaces can be accessed from the public internet or
	// only from private endpoints.
	Create(ctx context.Context, request CreatePrivateAccessSettingsRequest) (*PrivateAccessSettings, error)

	// Deletes a Databricks private access settings configuration, both
	// specified by ID.
	Delete(ctx context.Context, request DeletePrivateAccesRequest) (*PrivateAccessSettings, error)

	// Gets a Databricks private access settings configuration, both specified
	// by ID.
	Get(ctx context.Context, request GetPrivateAccesRequest) (*PrivateAccessSettings, error)

	// Lists Databricks private access settings for an account.
	List(ctx context.Context) ([]PrivateAccessSettings, error)

	// Updates an existing private access settings object, which specifies how
	// your workspace is accessed over AWS PrivateLink. To use AWS PrivateLink,
	// a workspace must have a private access settings object referenced by ID
	// in the workspace's private_access_settings_id property. This operation
	// completely overwrites your existing private access settings object
	// attached to your workspaces. All workspaces attached to the private
	// access settings are affected by any change. If public_access_enabled,
	// private_access_level, or allowed_vpc_endpoint_ids are updated, effects of
	// these changes might take several minutes to propagate to the workspace
	// API. You can share one private access settings object with multiple
	// workspaces in a single account. However, private access settings are
	// specific to AWS regions, so only workspaces in the same AWS region can
	// use a given private access settings object. Before configuring
	// PrivateLink, read the Databricks article about PrivateLink.
	Replace(ctx context.Context, request ReplacePrivateAccessSettingsRequest) (*PrivateAccessSettings, error)
}

// These APIs manage storage configurations for this workspace. A root storage
// S3 bucket in your account is required to store objects like cluster logs,
// notebook revisions, and job results. You can also use the root storage S3
// bucket for storage of non-production DBFS data. A storage configuration
// encapsulates this bucket information, and its ID is used when creating a new
// workspace.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type StorageService interface {

	// Creates a Databricks storage configuration for an account.
	Create(ctx context.Context, request CreateStorageConfigurationRequest) (*StorageConfiguration, error)

	// Deletes a Databricks storage configuration. You cannot delete a storage
	// configuration that is associated with any workspace.
	Delete(ctx context.Context, request DeleteStorageRequest) (*StorageConfiguration, error)

	// Gets a Databricks storage configuration for an account, both specified by
	// ID.
	Get(ctx context.Context, request GetStorageRequest) (*StorageConfiguration, error)

	// Lists Databricks storage configurations for an account, specified by ID.
	List(ctx context.Context) ([]StorageConfiguration, error)
}

// These APIs manage VPC endpoint configurations for this account.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type VpcEndpointsService interface {

	// Creates a VPC endpoint configuration, which represents a [VPC endpoint]
	// object in AWS used to communicate privately with Databricks over [AWS
	// PrivateLink].
	//
	// After you create the VPC endpoint configuration, the Databricks [endpoint
	// service] automatically accepts the VPC endpoint.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [VPC endpoint]: https://docs.aws.amazon.com/vpc/latest/privatelink/vpc-endpoints.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/privatelink-share-your-services.html
	Create(ctx context.Context, request CreateVpcEndpointRequest) (*VpcEndpoint, error)

	// Deletes a Databricks VPC endpoint configuration. You cannot delete a VPC
	// endpoint configuration that is associated with any workspace.
	Delete(ctx context.Context, request DeleteVpcEndpointRequest) (*VpcEndpoint, error)

	// Gets a VPC endpoint configuration, which represents a [VPC endpoint]
	// object in AWS used to communicate privately with Databricks over [AWS
	// PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink
	// [VPC endpoint]: https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html
	Get(ctx context.Context, request GetVpcEndpointRequest) (*VpcEndpoint, error)

	// Lists Databricks VPC endpoint configurations for an account.
	List(ctx context.Context) ([]VpcEndpoint, error)
}

// These APIs manage workspaces for this account. A Databricks workspace is an
// environment for accessing all of your Databricks assets. The workspace
// organizes objects (notebooks, libraries, and experiments) into folders, and
// provides access to data and computational resources such as clusters and
// jobs.
//
// These endpoints are available if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspacesService interface {

	// Creates a new workspace using a credential configuration and a storage
	// configuration, an optional network configuration (if using a
	// customer-managed VPC), an optional managed services key configuration (if
	// using customer-managed keys for managed services), and an optional
	// storage key configuration (if using customer-managed keys for storage).
	// The key configurations used for managed services and storage encryption
	// can be the same or different.
	//
	// Important: This operation is asynchronous. A response with HTTP status
	// code 200 means the request has been accepted and is in progress, but does
	// not mean that the workspace deployed successfully and is running. The
	// initial workspace status is typically PROVISIONING. Use the workspace ID
	// (workspace_id) field in the response to identify the new workspace and
	// make repeated GET requests with the workspace ID and check its status.
	// The workspace becomes available when the status changes to RUNNING.
	//
	// You can share one customer-managed VPC with multiple workspaces in a
	// single account. It is not required to create a new VPC for each
	// workspace. However, you cannot reuse subnets or Security Groups between
	// workspaces. If you plan to share one VPC with multiple workspaces, make
	// sure you size your VPC and subnets accordingly. Because a Databricks
	// Account API network configuration encapsulates this information, you
	// cannot reuse a Databricks Account API network configuration across
	// workspaces.
	//
	// For information about how to create a new workspace with this API
	// including error handling, see [Create a new workspace using the Account
	// API].
	//
	// Important: Customer-managed VPCs, PrivateLink, and customer-managed keys
	// are supported on a limited set of deployment and subscription types. If
	// you have questions about availability, contact your Databricks
	// representative.
	//
	// This operation is available only if your account is on the E2 version of
	// the platform or on a select custom plan that allows multiple workspaces
	// per account.
	//
	// [Create a new workspace using the Account API]: http://docs.databricks.com/administration-guide/account-api/new-workspace.html
	Create(ctx context.Context, request CreateWorkspaceRequest) (*Workspace, error)

	// Deletes a Databricks workspace, both specified by ID.
	Delete(ctx context.Context, request DeleteWorkspaceRequest) (*Workspace, error)

	// Gets information including status for a Databricks workspace, specified
	// by ID. In the response, the `workspace_status` field indicates the
	// current status. After initial workspace creation (which is asynchronous),
	// make repeated `GET` requests with the workspace ID and check its status.
	// The workspace becomes available when the status changes to `RUNNING`. For
	// information about how to create a new workspace with this API **including
	// error handling**, see [Create a new workspace using the Account API].
	//
	// [Create a new workspace using the Account API]: http://docs.databricks.com/administration-guide/account-api/new-workspace.html
	Get(ctx context.Context, request GetWorkspaceRequest) (*Workspace, error)

	// Lists Databricks workspaces for an account.
	List(ctx context.Context) ([]Workspace, error)

	// Updates a workspace.
	Update(ctx context.Context, request UpdateWorkspaceRequest) (*Workspace, error)
}
