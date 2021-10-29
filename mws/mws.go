package mws

import "encoding/json"

// StsRole is the object that contains cross account role arn and external app id
type StsRole struct {
	RoleArn    string `json:"role_arn,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
}

// AwsCredentials is the object that points to the cross account role
type AwsCredentials struct {
	StsRole *StsRole `json:"sts_role,omitempty"`
}

// Credentials is the object that contains all the information for the credentials to create a workspace
type Credentials struct {
	CredentialsID   string          `json:"credentials_id,omitempty"`
	CredentialsName string          `json:"credentials_name,omitempty"`
	AwsCredentials  *AwsCredentials `json:"aws_credentials,omitempty"`
	AccountID       string          `json:"account_id,omitempty"`
	CreationTime    int64           `json:"creation_time,omitempty"`
}

// RootBucketInfo points to a bucket name
type RootBucketInfo struct {
	BucketName string `json:"bucket_name,omitempty"`
}

// StorageConfiguration is the object that contains all the information for the root storage bucket
type StorageConfiguration struct {
	StorageConfigurationID   string          `json:"storage_configuration_id,omitempty"`
	StorageConfigurationName string          `json:"storage_configuration_name,omitempty"`
	RootBucketInfo           *RootBucketInfo `json:"root_bucket_info,omitempty"`
	AccountID                string          `json:"account_id,omitempty"`
	CreationTime             int64           `json:"creation_time,omitempty"`
}

// NetworkHealth is the object that contains all the error message when attaching a network to workspace
type NetworkHealth struct {
	ErrorType    string `json:"error_type,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// NetworkVPCEndpoints is the object that contains VPC endpoints of a network
type NetworkVPCEndpoints struct {
	RestAPI           []string `json:"rest_api" tf:"slice_set"`
	DataplaneRelayAPI []string `json:"dataplane_relay" tf:"slice_set"`
}

// Network is the object that contains all the information for BYOVPC
type Network struct {
	AccountID        string               `json:"account_id"`
	NetworkID        string               `json:"network_id,omitempty" tf:"computed"`
	NetworkName      string               `json:"network_name"`
	VPCID            string               `json:"vpc_id"`
	SubnetIds        []string             `json:"subnet_ids" tf:"slice_set"`
	VPCEndpoints     *NetworkVPCEndpoints `json:"vpc_endpoints,omitempty" tf:"computed"`
	SecurityGroupIds []string             `json:"security_group_ids" tf:"slice_set"`
	VPCStatus        string               `json:"vpc_status,omitempty" tf:"computed"`
	ErrorMessages    []NetworkHealth      `json:"error_messages,omitempty" tf:"computed"`
	WorkspaceID      int64                `json:"workspace_id,omitempty" tf:"computed"`
	CreationTime     int64                `json:"creation_time,omitempty" tf:"computed"`
}

// List of workspace statuses for provisioning the workspace
const (
	WorkspaceStatusNotProvisioned = "NOT_PROVISIONED"
	WorkspaceStatusProvisioning   = "PROVISIONING"
	WorkspaceStatusRunning        = "RUNNING"
	WorkspaceStatusFailed         = "FAILED"
	WorkspaceStatusCanceled       = "CANCELLED"
)

// WorkspaceStatusesNonRunnable is a list of statuses in which the workspace is not runnable
var WorkspaceStatusesNonRunnable = []string{WorkspaceStatusCanceled, WorkspaceStatusFailed}

type GCP struct {
	ProjectID string `json:"project_id"`
}

type CloudResourceBucket struct {
	GCP *GCP `json:"gcp"`
}

type GCPManagedNetworkConfig struct {
	SubnetCIDR               string `json:"subnet_cidr"`
	GKEClusterPodIPRange     string `json:"gke_cluster_pod_ip_range"`
	GKEClusterServiceIPRange string `json:"gke_cluster_service_ip_range"`
}

type GCPCommonNetworkConfig struct {
	GKEConnectivityType     string `json:"gke_connectivity_type"`
	GKEClusterMasterIPRange string `json:"gke_cluster_master_ip_range"`
}

type GCPNetwork struct {
	GCPManagedNetworkConfig *GCPManagedNetworkConfig `json:"gcp_managed_network_config"`
	GCPCommonNetworkConfig  *GCPCommonNetworkConfig  `json:"gcp_common_network_config"`
}

// Workspace is the object that contains all the information for deploying a workspace
type Workspace struct {
	AccountID                           string `json:"account_id"`
	WorkspaceName                       string `json:"workspace_name"`
	DeploymentName                      string `json:"deployment_name,omitempty"`
	AwsRegion                           string `json:"aws_region,omitempty"`               // required for AWS, not allowed for GCP
	CredentialsID                       string `json:"credentials_id,omitempty"`           // required for AWS, not allowed for GCP
	CustomerManagedKeyID                string `json:"customer_managed_key_id,omitempty"`  // just for compatibility, will be removed
	StorageConfigurationID              string `json:"storage_configuration_id,omitempty"` // required for AWS, not allowed for GCP
	ManagedServicesCustomerManagedKeyID string `json:"managed_services_customer_managed_key_id,omitempty"`
	StorageCustomerManagedKeyID         string `json:"storage_customer_managed_key_id,omitempty"`
	PricingTier                         string `json:"pricing_tier,omitempty" tf:"computed"`
	PrivateAccessSettingsID             string `json:"private_access_settings_id,omitempty"`
	NetworkID                           string `json:"network_id,omitempty"`
	IsNoPublicIPEnabled                 bool   `json:"is_no_public_ip_enabled"`
	WorkspaceID                         int64  `json:"workspace_id,omitempty" tf:"computed"`
	WorkspaceURL                        string `json:"workspace_url,omitempty" tf:"computed"`
	WorkspaceStatus                     string `json:"workspace_status,omitempty" tf:"computed"`
	WorkspaceStatusMessage              string `json:"workspace_status_message,omitempty" tf:"computed"`
	CreationTime                        int64  `json:"creation_time,omitempty" tf:"computed"`

	ExternalCustomerInfo *externalCustomerInfo `json:"external_customer_info,omitempty"`

	CloudResourceBucket *CloudResourceBucket `json:"cloud_resource_bucket,omitempty"`
	Network             *GCPNetwork          `json:"network,omitempty"`
	Cloud               string               `json:"cloud,omitempty" tf:"computed"`
	Location            string               `json:"location,omitempty"`
}

// this type alias hack is required for Marshaller to work without an infinite loop
type aWorkspace Workspace

// MarshalJSON is required to overcome the limitations of `omitempty` usage with reflect_resource.go
// for workspace creation in Accounts API for AWS and GCP. It exits early on AWS and picks only
// the relevant fields for GCP.
func (w *Workspace) MarshalJSON() ([]byte, error) {
	if w.Cloud != "gcp" {
		return json.Marshal(aWorkspace(*w))
	}
	workspaceCreationRequest := map[string]interface{}{
		"account_id":            w.AccountID,
		"cloud":                 w.Cloud,
		"cloud_resource_bucket": w.CloudResourceBucket,
		"location":              w.Location,
		"workspace_name":        w.WorkspaceName,
	}
	if w.Network != nil {
		workspaceCreationRequest["network"] = w.Network
	}
	return json.Marshal(workspaceCreationRequest)
}

// VPCEndpoint is the object that contains all the information for registering an VPC endpoint
type VPCEndpoint struct {
	VPCEndpointID           string `json:"vpc_endpoint_id,omitempty" tf:"computed"`
	AwsVPCEndpointID        string `json:"aws_vpc_endpoint_id"`
	AccountID               string `json:"account_id,omitempty"`
	VPCEndpointName         string `json:"vpc_endpoint_name"`
	AwsVPCEndpointServiceID string `json:"aws_endpoint_service_id,omitempty" tf:"computed"`
	AWSAccountID            string `json:"aws_account_id,omitempty" tf:"computed"`
	UseCase                 string `json:"use_case,omitempty" tf:"computed"`
	Region                  string `json:"region"`
	State                   string `json:"state,omitempty" tf:"computed"`
}

// PrivateAccessSettings (PAS) is the object that contains all the information for creating an PrivateAccessSettings (PAS)
type PrivateAccessSettings struct {
	AccountID             string   `json:"account_id,omitempty"`
	PasID                 string   `json:"private_access_settings_id,omitempty" tf:"computed"`
	PasName               string   `json:"private_access_settings_name"`
	Region                string   `json:"region"`
	Status                string   `json:"status,omitempty" tf:"computed"`
	PublicAccessEnabled   bool     `json:"public_access_enabled,omitempty"`
	PrivateAccessLevel    string   `json:"private_access_level,omitempty" tf:"default:ACCOUNT"`
	AllowedVpcEndpointIDS []string `json:"allowed_vpc_endpoint_ids,omitempty"`
}

type externalCustomerInfo struct {
	CustomerName              string `json:"customer_name"`
	AuthoritativeUserEmail    string `json:"authoritative_user_email"`
	AuthoritativeUserFullName string `json:"authoritative_user_full_name"`
}
