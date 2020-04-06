package model

import "errors"

type AutoScale struct {
	MinWorkers int32 `json:"min_workers,omitempty"`
	MaxWorkers int32 `json:"max_workers,omitempty"`
}

type AwsAvailability string

const (
	AwsAvailabilitySpot             = "SPOT"
	AwsAvailabilityOnDemand         = "ON_DEMAND"
	AwsAvailabilitySpotWithFallback = "SPOT_WITH_FALLBACK"
)

func GetAwsAvailability(val string) (AwsAvailability, error) {
	switch val {
	case "SPOT":
		return AwsAvailabilitySpot, nil
	case "ON_DEMAND":
		return AwsAvailabilityOnDemand, nil
	case "SPOT_WITH_FALLBACK":
		return AwsAvailabilitySpotWithFallback, nil
	}
	return "", errors.New("No Match!")
}

type AzureDiskVolumeType string

const (
	AzureDiskVolumeTypeStandard = "STANDARD_LRS"
	AzureDiskVolumeTypePremium  = "PREMIUM_LRS"
)

func GetAzureDiskVolumeType(val string) (AzureDiskVolumeType, error) {
	switch val {
	case "STANDARD_LRS":
		return AzureDiskVolumeTypeStandard, nil
	case "PREMIUM_LRS":
		return AzureDiskVolumeTypePremium, nil
	}
	return "", errors.New("No Match!")
}

type EbsVolumeType string

const (
	EbsVolumeTypeGeneralPurposeSsd      = "GENERAL_PURPOSE_SSD"
	EbsVolumeTypeThroughputOptimizedHdd = "THROUGHPUT_OPTIMIZED_HDD"
)

func GetEbsVolumeType(val string) (EbsVolumeType, error) {
	switch val {
	case "GENERAL_PURPOSE_SSD":
		return EbsVolumeTypeGeneralPurposeSsd, nil
	case "THROUGHPUT_OPTIMIZED_HDD":
		return EbsVolumeTypeThroughputOptimizedHdd, nil
	}
	return "", errors.New("No Match!")
}

type ClusterState string

const (
	ClusterStatePending     = "PENDING"
	ClusterStateRunning     = "RUNNING"
	ClusterStateRestarting  = "RESTARTING"
	ClusterStateResizing    = "RESIZING"
	ClusterStateTerminating = "TERMINATING"
	ClusterStateTerminated  = "TERMINATED"
	ClusterStateError       = "ERROR"
	ClusterStateUnknown     = "UNKNOWN"
)

var ClusterStateNonRunnable []ClusterState = []ClusterState{ClusterStateTerminating, ClusterStateTerminated, ClusterStateError, ClusterStateUnknown}
var ClusterStateNonTerminating []ClusterState = []ClusterState{ClusterStatePending, ClusterStateRunning, ClusterStateRestarting, ClusterStateResizing}

func ContainsClusterState(clusterStates []ClusterState, searchState ClusterState) bool {
	for _, state := range clusterStates {
		if state == searchState {
			return true
		}
	}
	return false
}

type AwsAttributes struct {
	FirstOnDemand       int32           `json:"first_on_demand,omitempty"`
	Availability        AwsAvailability `json:"availability,omitempty"`
	ZoneID              string          `json:"zone_id,omitempty"`
	InstanceProfileArn  string          `json:"instance_profile_arn,omitempty"`
	SpotBidPricePercent int32           `json:"spot_bid_price_percent,omitempty"`
	EbsVolumeType       EbsVolumeType   `json:"ebs_volume_type,omitempty"`
	EbsVolumeCount      int32           `json:"ebs_volume_count,omitempty"`
	EbsVolumeSize       int32           `json:"ebs_volume_size,omitempty"`
}

type DbfsStorageInfo struct {
	Destination string `json:"destination,omitempty"`
}

type S3StorageInfo struct {
	Destination      string `json:"destination,omitempty"`
	Region           string `json:"region,omitempty"`
	Endpoint         string `json:"endpoint,omitempty"`
	EnableEncryption bool   `json:"enable_encryption,omitempty"`
	EncryptionType   string `json:"encryption_type,omitempty"`
	KmsKey           string `json:"kms_key,omitempty"`
	CannedACL        string `json:"canned_acl,omitempty"`
}

type StorageInfo struct {
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty"`
	S3   *S3StorageInfo   `json:"s3,omitempty"`
}

type SparkNodeAwsAttributes struct {
	IsSpot bool `json:"is_spot,omitempty"`
}

type SparkNode struct {
	PrivateIP         string                  `json:"private_ip,omitempty"`
	PublicDNS         string                  `json:"public_dns,omitempty"`
	NodeID            string                  `json:"node_id,omitempty"`
	InstanceID        string                  `json:"instance_id,omitempty"`
	StartTimestamp    int64                   `json:"start_timestamp,omitempty"`
	NodeAwsAttributes *SparkNodeAwsAttributes `json:"node_aws_attributes,omitempty"`
	HostPrivateIP     string                  `json:"host_private_ip,omitempty"`
}

type TerminationCode string

const (
	TerminationCodeUserRequest                = "USER_REQUEST"
	TerminationCodeJobFinished                = "JOB_FINISHED"
	TerminationCodeInactivity                 = "INACTIVITY"
	TerminationCodeCloudProviderShutdown      = "CLOUD_PROVIDER_SHUTDOWN"
	TerminationCodeCommunicationLost          = "COMMUNICATION_LOST"
	TerminationCodeCloudProviderLaunchFailure = "CLOUD_PROVIDER_LAUNCH_FAILURE"
	TerminationCodeSparkStartupFailure        = "SPARK_STARTUP_FAILURE"
	TerminationCodeInvalidArgument            = "INVALID_ARGUMENT"
	TerminationCodeUnexpectedLaunchFailure    = "UNEXPECTED_LAUNCH_FAILURE"
	TerminationCodeInternalError              = "INTERNAL_ERROR"
	TerminationCodeInstanceUnreachable        = "INSTANCE_UNREACHABLE"
	TerminationCodeRequestRejected            = "REQUEST_REJECTED"
	TerminationCodeInitScriptFailure          = "INIT_SCRIPT_FAILURE"
	TerminationCodeTrialExpired               = "TRIAL_EXPIRED"
)

type TerminationParameter string

const (
	TerminationParameterUsername                 = "username"
	TerminationParameterAwsAPIErrorCode          = "aws_api_error_code"
	TerminationParameterAwsInstanceStateReason   = "aws_instance_state_reason"
	TerminationParameterAwsSpotRequestStatus     = "aws_spot_request_status"
	TerminationParameterAwsSpotRequestFaultCode  = "aws_spot_request_fault_code"
	TerminationParameterAwsImpairedStatusDetails = "aws_impaired_status_details"
	TerminationParameterAwsInstanceStatusEvent   = "aws_instance_status_event"
	TerminationParameterAwsErrorMessage          = "aws_error_message"
	TerminationParameterDatabricksErrorMessage   = "databricks_error_message"
	TerminationParameterInactivityDurationMin    = "inactivity_duration_min"
	TerminationParameterInstanceID               = "instance_id"
)

type TerminationReason struct {
	Code       TerminationCode   `json:"code,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

type LogSyncStatus struct {
	LastAttempted int64  `json:"last_attempted,omitempty"`
	LastException string `json:"last_exception,omitempty"`
}

type ClusterCloudProviderNodeStatus string

const (
	ClusterCloudProviderNodeStatusNotEnabledOnSubscription = "NotEnabledOnSubscription"
	ClusterCloudProviderNodeStatusNotAvailableInRegion     = "NotAvailableInRegion"
)

type ClusterCloudProviderNodeInfo struct {
	Status             []ClusterCloudProviderNodeStatus `json:"status,omitempty"`
	AvailableCoreQuota float32                          `json:"available_core_quota,omitempty"`
	TotalCoreQuota     float32                          `json:"total_core_quota,omitempty"`
}

type NodeType struct {
	NodeTypeID     string                        `json:"node_type_id,omitempty"`
	MemoryMb       int32                         `json:"memory_mb,omitempty"`
	NumCores       float32                       `json:"num_cores,omitempty"`
	Description    string                        `json:"description,omitempty"`
	InstanceTypeID string                        `json:"instance_type_id,omitempty"`
	IsDeprecated   bool                          `json:"is_deprecated,omitempty"`
	NodeInfo       *ClusterCloudProviderNodeInfo `json:"node_info,omitempty"`
}

type DockerBasicAuth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type DockerImage struct {
	Url       string           `json:"url,omitempty"`
	BasicAuth *DockerBasicAuth `json:"basic_auth,omitempty"`
}

type Cluster struct {
	ClusterId              string            `json:"cluster_id,omitempty"`
	NumWorkers             int32             `json:"num_workers,omitempty"`
	Autoscale              *AutoScale        `json:"autoscale,omitempty"`
	ClusterName            string            `json:"cluster_name,omitempty"`
	SparkVersion           string            `json:"spark_version,omitempty"`
	SparkConf              map[string]string `json:"spark_conf,omitempty"`
	AwsAttributes          *AwsAttributes    `json:"aws_attributes,omitempty"`
	NodeTypeID             string            `json:"node_type_id,omitempty"`
	DriverNodeTypeID       string            `json:"driver_node_type_id,omitempty"`
	SSHPublicKeys          []string          `json:"ssh_public_keys,omitempty"`
	CustomTags             map[string]string `json:"custom_tags,omitempty"`
	ClusterLogConf         *StorageInfo      `json:"cluster_log_conf,omitempty"`
	InitScripts            []StorageInfo     `json:"init_scripts,omitempty"`
	DockerImage            *DockerImage      `json:"docker_image,omitempty"`
	SparkEnvVars           map[string]string `json:"spark_env_vars,omitempty"`
	AutoterminationMinutes int32             `json:"autotermination_minutes,omitempty"`
	EnableElasticDisk      bool              `json:"enable_elastic_disk,omitempty"`
	InstancePoolId         string            `json:"instance_pool_id,omitempty"`
	IdempotencyToken       string            `json:"idempotency_token,omitempty"`
}

type ClusterInfo struct {
	NumWorkers             int32              `json:"num_workers,omitempty"`
	AutoScale              *AutoScale         `json:"autoscale,omitempty"`
	ClusterID              string             `json:"cluster_id,omitempty"`
	CreatorUserName        string             `json:"creator_user_name,omitempty"`
	Driver                 *SparkNode         `json:"driver,omitempty"`
	Executors              []SparkNode        `json:"executors,omitempty"`
	SparkContextID         int64              `json:"spark_context_id,omitempty"`
	JdbcPort               int32              `json:"jdbc_port,omitempty"`
	ClusterName            string             `json:"cluster_name,omitempty"`
	SparkVersion           string             `json:"spark_version,omitempty"`
	SparkConf              map[string]string  `json:"spark_conf,omitempty"`
	AwsAttributes          *AwsAttributes     `json:"aws_attributes,omitempty"`
	NodeTypeID             string             `json:"node_type_id,omitempty"`
	DriverNodeTypeID       string             `json:"driver_node_type_id,omitempty"`
	SSHPublicKeys          []string           `json:"ssh_public_keys,omitempty"`
	CustomTags             map[string]string  `json:"custom_tags,omitempty"`
	ClusterLogConf         *StorageInfo       `json:"cluster_log_conf,omitempty"`
	InitScripts            []StorageInfo      `json:"init_scripts,omitempty"`
	SparkEnvVars           map[string]string  `json:"spark_env_vars,omitempty"`
	AutoterminationMinutes int32              `json:"autotermination_minutes,omitempty"`
	EnableElasticDisk      bool               `json:"enable_elastic_disk,omitempty"`
	InstancePoolId         string             `json:"instance_pool_id,omitempty"`
	ClusterSource          AwsAvailability    `json:"cluster_source,omitempty"`
	DockerImage            *DockerImage       `json:"docker_image,omitempty"`
	State                  ClusterState       `json:"state,omitempty"`
	StateMessage           string             `json:"state_message,omitempty"`
	StartTime              int64              `json:"start_time,omitempty"`
	TerminateTime          int64              `json:"terminate_time,omitempty"`
	LastStateLossTime      int64              `json:"last_state_loss_time,omitempty"`
	LastActivityTime       int64              `json:"last_activity_time,omitempty"`
	ClusterMemoryMb        int64              `json:"cluster_memory_mb,omitempty"`
	ClusterCores           float32            `json:"cluster_cores,omitempty"`
	DefaultTags            map[string]string  `json:"default_tags,omitempty"`
	ClusterLogStatus       *LogSyncStatus     `json:"cluster_log_status,omitempty"`
	TerminationReason      *TerminationReason `json:"termination_reason,omitempty"`
}
