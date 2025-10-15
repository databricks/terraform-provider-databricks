package mws

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/docs"
)

func getGkeDeprecationMessage(fieldName string, docOptions docs.DocOptions) string {
	return fmt.Sprintf(
		"%s is deprecated and will be removed in a future release. For more information, review the documentation at %s",
		fieldName,
		docs.DocumentationUrl(docOptions),
	)
}

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

// GcpNetworkInfo is the object that configures byovpc settings for gcp
type GcpNetworkInfo struct {
	NetworkProjectId   string `json:"network_project_id" tf:"force_new"`
	VpcId              string `json:"vpc_id" tf:"force_new"`
	SubnetId           string `json:"subnet_id" tf:"force_new"`
	SubnetRegion       string `json:"subnet_region" tf:"force_new"`
	PodIpRangeName     string `json:"pod_ip_range_name,omitempty"`
	ServiceIpRangeName string `json:"service_ip_range_name,omitempty"`
}

// Network is the object that contains all the information for BYOVPC
type Network struct {
	common.Namespace
	AccountID        string               `json:"account_id" tf:"force_new"`
	NetworkID        string               `json:"network_id,omitempty" tf:"computed"`
	NetworkName      string               `json:"network_name" tf:"force_new"`
	VPCID            string               `json:"vpc_id,omitempty" tf:"force_new"`
	SubnetIds        []string             `json:"subnet_ids,omitempty" tf:"slice_set,force_new"`
	VPCEndpoints     *NetworkVPCEndpoints `json:"vpc_endpoints,omitempty" tf:"computed,force_new"`
	SecurityGroupIds []string             `json:"security_group_ids,omitempty" tf:"slice_set,force_new"`
	VPCStatus        string               `json:"vpc_status,omitempty" tf:"computed"`
	ErrorMessages    []NetworkHealth      `json:"error_messages,omitempty" tf:"computed"`
	WorkspaceID      int64                `json:"workspace_id,omitempty" tf:"computed"`
	CreationTime     int64                `json:"creation_time,omitempty" tf:"computed"`
	GcpNetworkInfo   *GcpNetworkInfo      `json:"gcp_network_info,omitempty"`
}

// GcpVpcEndpointInfo is the objecy that configures GCP Private Service Connect endpoint.
type GcpVpcEndpointInfo struct {
	PscConnectionId     string `json:"psc_connection_id,omitempty" tf:"computed"`
	ProjectId           string `json:"project_id"`
	PscEndpointName     string `json:"psc_endpoint_name"`
	EndpointRegion      string `json:"endpoint_region"`
	ServiceAttachmentId string `json:"service_attachment_id,omitempty" tf:"computed"`
}

// VPCEndpoint is the object that contains all the information for registering an VPC endpoint
type VPCEndpoint struct {
	common.Namespace
	VPCEndpointID           string              `json:"vpc_endpoint_id,omitempty" tf:"computed"`
	AwsVPCEndpointID        string              `json:"aws_vpc_endpoint_id,omitempty"`
	AccountID               string              `json:"account_id,omitempty"`
	VPCEndpointName         string              `json:"vpc_endpoint_name"`
	AwsVPCEndpointServiceID string              `json:"aws_endpoint_service_id,omitempty" tf:"computed"`
	AWSAccountID            string              `json:"aws_account_id,omitempty" tf:"computed"`
	UseCase                 string              `json:"use_case,omitempty" tf:"computed"`
	Region                  string              `json:"region,omitempty"`
	State                   string              `json:"state,omitempty" tf:"computed"`
	GcpVpcEndpointInfo      *GcpVpcEndpointInfo `json:"gcp_vpc_endpoint_info,omitempty"`
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

	ForceSendFields []string `json:"-"`
}

func (s *PrivateAccessSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrivateAccessSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
