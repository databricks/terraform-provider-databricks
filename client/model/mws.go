package model

// StsRole is the object that contains cross account role arn and external app id
type StsRole struct {
	RoleArn    string `json:"role_arn,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
}

// AwsCredentials is the object that points to the cross account role
type AwsCredentials struct {
	StsRole *StsRole `json:"sts_role,omitempty"`
}

// MWSCredentials is the object that contains all the information for the credentials to create a workspace
type MWSCredentials struct {
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

// MWSStorageConfigurations is the object that contains all the information for the root storage bucket
type MWSStorageConfigurations struct {
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

// MWSNetwork is the object that contains all the information for BYOVPC
type MWSNetwork struct {
	NetworkID        string          `json:"network_id,omitempty"`
	NetworkName      string          `json:"network_name,omitempty"`
	VPCID            string          `json:"vpc_id,omitempty"`
	SubnetIds        []string        `json:"subnet_ids,omitempty"`
	SecurityGroupIds []string        `json:"security_group_ids,omitempty"`
	VPCStatus        string          `json:"vpc_status,omitempty"`
	ErrorMessages    []NetworkHealth `json:"error_messages,omitempty"`
	WorkspaceID      int64           `json:"workspace_id,omitempty"`
	AccountID        string          `json:"account_id,omitempty"`
	CreationTime     int64           `json:"creation_time,omitempty"`
}

// AwsKeyInfo has information about the KMS key for BYOK
type AwsKeyInfo struct {
	KeyArn    string `json:"key_arn,omitempty"`
	KeyAlias  string `json:"key_alias,omitempty"`
	KeyRegion string `json:"key_region,omitempty"`
}

// MWSCustomerManagedKey contains key information and metadata for BYOK for E2
type MWSCustomerManagedKey struct {
	CustomerManagedKeyID string      `json:"customer_managed_key_id,omitempty"`
	AwsKeyInfo           *AwsKeyInfo `json:"aws_key_info,omitempty"`
	AccountID            string      `json:"account_id,omitempty"`
	CreationTime         int64       `json:"creation_time,omitempty"`
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

// ContainsWorkspaceState given a list of workspaceStates and the search state
// it will return true if it found the search state
func ContainsWorkspaceState(workspaceStates []string, searchState string) bool {
	for _, state := range workspaceStates {
		if state == searchState {
			return true
		}
	}
	return false
}

// MWSWorkspace is the object that contains all the information for deploying a E2 workspace
type MWSWorkspace struct {
	WorkspaceID            int64  `json:"workspace_id,omitempty"`
	WorkspaceName          string `json:"workspace_name,omitempty"`
	DeploymentName         string `json:"deployment_name,omitempty"`
	AwsRegion              string `json:"aws_region,omitempty"`
	CredentialsID          string `json:"credentials_id,omitempty"`
	StorageConfigurationID string `json:"storage_configuration_id,omitempty"`
	NetworkID              string `json:"network_id,omitempty"`
	CustomerManagedKeyID   string `json:"customer_managed_key_id,omitempty"`
	IsNoPublicIpEnabled    bool   `json:"is_no_public_ip_enabled,omitempty"`
	AccountID              string `json:"account_id,omitempty"`
	WorkspaceStatus        string `json:"workspace_status,omitempty"`
	WorkspaceStatusMessage string `json:"workspace_status_message,omitempty"`
	CreationTime           int64  `json:"creation_time,omitempty"`
}
