package model

// AutoScale is a struct the describes auto scaling for clusters
type AutoScale struct {
	MinWorkers int32 `json:"min_workers,omitempty"`
	MaxWorkers int32 `json:"max_workers,omitempty"`
}

// AwsAvailability is a type for describing AWS availability on cluster nodes
type AwsAvailability string

const (
	// AwsAvailabilitySpot is spot instance type for clusters
	AwsAvailabilitySpot = "SPOT"
	// AwsAvailabilityOnDemand is OnDemand instance type for clusters
	AwsAvailabilityOnDemand = "ON_DEMAND"
	// AwsAvailabilitySpotWithFallback is Spot instance type for clusters with option
	// to fallback into on-demand if instance cannot be acquired
	AwsAvailabilitySpotWithFallback = "SPOT_WITH_FALLBACK"
)

// AzureDiskVolumeType is disk type on azure vms
type AzureDiskVolumeType string

const (
	// AzureDiskVolumeTypeStandard is for standard local redundant storage
	AzureDiskVolumeTypeStandard = "STANDARD_LRS"
	// AzureDiskVolumeTypePremium is for premium local redundant storage
	AzureDiskVolumeTypePremium = "PREMIUM_LRS"
)

// EbsVolumeType is disk type on aws vms
type EbsVolumeType string

const (
	// EbsVolumeTypeGeneralPurposeSsd is general purpose ssd (starts at 32 gb)
	EbsVolumeTypeGeneralPurposeSsd = "GENERAL_PURPOSE_SSD"
	// EbsVolumeTypeThroughputOptimizedHdd is throughput optimized hdd (starts at 500 gb)
	EbsVolumeTypeThroughputOptimizedHdd = "THROUGHPUT_OPTIMIZED_HDD"
)

// ClusterState is for describing possible cluster states
type ClusterState string

const (
	// ClusterStatePending is for PENDING state
	ClusterStatePending = "PENDING"

	// ClusterStateRunning is for RUNNING state
	ClusterStateRunning = "RUNNING"

	// ClusterStateRestarting is for RESTARTING state
	ClusterStateRestarting = "RESTARTING"

	// ClusterStateResizing is for RESIZING state
	ClusterStateResizing = "RESIZING"

	// ClusterStateTerminating is for TERMINATING state
	ClusterStateTerminating = "TERMINATING"

	// ClusterStateTerminated is for TERMINATED state
	ClusterStateTerminated = "TERMINATED"

	// ClusterStateError is for ERROR state
	ClusterStateError = "ERROR"

	// ClusterStateUnknown is for UNKNOWN state
	ClusterStateUnknown = "UNKNOWN"
)

// ClusterStateNonRunnable is a list of states in which the cluster cannot go back into running by itself
// without user intervention
var ClusterStateNonRunnable = []ClusterState{ClusterStateTerminating, ClusterStateTerminated, ClusterStateError, ClusterStateUnknown}

// ClusterStateNonTerminating is a list of states in which the cluster cannot go back into terminated by itself
//// without user intervention
var ClusterStateNonTerminating = []ClusterState{ClusterStatePending, ClusterStateRunning, ClusterStateRestarting, ClusterStateResizing, ClusterStateUnknown}

// ContainsClusterState given a set of cluster states and a search state it will return true if the state is in the
// given set
func ContainsClusterState(clusterStates []ClusterState, searchState ClusterState) bool {
	for _, state := range clusterStates {
		if state == searchState {
			return true
		}
	}
	return false
}

// ZonesInfo encapsulates the zone information from the zones api call
type ZonesInfo struct {
	Zones       []string `json:"zones,omitempty"`
	DefaultZone string   `json:"default_zone,omitempty"`
}

// AwsAttributes encapsulates the aws attributes for aws based clusters
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

// DbfsStorageInfo contains the destination string for DBFS
type DbfsStorageInfo struct {
	Destination string `json:"destination,omitempty"`
}

// S3StorageInfo contains the struct for when storing files in S3
type S3StorageInfo struct {
	Destination      string `json:"destination,omitempty"`
	Region           string `json:"region,omitempty"`
	Endpoint         string `json:"endpoint,omitempty"`
	EnableEncryption bool   `json:"enable_encryption,omitempty"`
	EncryptionType   string `json:"encryption_type,omitempty"`
	KmsKey           string `json:"kms_key,omitempty"`
	CannedACL        string `json:"canned_acl,omitempty"`
}

// StorageInfo contains the struct for either DBFS or S3 storage depending on which one is relevant.
type StorageInfo struct {
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty"`
	S3   *S3StorageInfo   `json:"s3,omitempty"`
}

// SparkNodeAwsAttributes is the struct that determines if the node is a spot instance or not
type SparkNodeAwsAttributes struct {
	IsSpot bool `json:"is_spot,omitempty"`
}

// SparkNode encapsulates all the attributes of a node that is part of a databricks cluster
type SparkNode struct {
	PrivateIP         string                  `json:"private_ip,omitempty"`
	PublicDNS         string                  `json:"public_dns,omitempty"`
	NodeID            string                  `json:"node_id,omitempty"`
	InstanceID        string                  `json:"instance_id,omitempty"`
	StartTimestamp    int64                   `json:"start_timestamp,omitempty"`
	NodeAwsAttributes *SparkNodeAwsAttributes `json:"node_aws_attributes,omitempty"`
	HostPrivateIP     string                  `json:"host_private_ip,omitempty"`
}

// TerminationReason encapsulates the termination code and potential parameters
type TerminationReason struct {
	Code       string            `json:"code,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

// LogSyncStatus encapsulates when the cluster logs were last delivered.
type LogSyncStatus struct {
	LastAttempted int64  `json:"last_attempted,omitempty"`
	LastException string `json:"last_exception,omitempty"`
}

// ClusterCloudProviderNodeInfo encapsulates the existing quota available from the cloud service provider.
type ClusterCloudProviderNodeInfo struct {
	Status             []string `json:"status,omitempty"`
	AvailableCoreQuota float32  `json:"available_core_quota,omitempty"`
	TotalCoreQuota     float32  `json:"total_core_quota,omitempty"`
}

// NodeType encapsulates information about a give node when using the list-node-types api
type NodeType struct {
	NodeTypeID     string                        `json:"node_type_id,omitempty"`
	MemoryMb       int32                         `json:"memory_mb,omitempty"`
	NumCores       float32                       `json:"num_cores,omitempty"`
	Description    string                        `json:"description,omitempty"`
	InstanceTypeID string                        `json:"instance_type_id,omitempty"`
	IsDeprecated   bool                          `json:"is_deprecated,omitempty"`
	NodeInfo       *ClusterCloudProviderNodeInfo `json:"node_info,omitempty"`
}

// DockerBasicAuth contains the auth information when fetching containers
type DockerBasicAuth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// DockerImage contains the image url and the auth for DCS
type DockerImage struct {
	URL       string           `json:"url,omitempty"`
	BasicAuth *DockerBasicAuth `json:"basic_auth,omitempty"`
}

// Cluster contains the information when trying to submit api calls or editing a cluster
type Cluster struct {
	ClusterID              string            `json:"cluster_id,omitempty"`
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
	InstancePoolID         string            `json:"instance_pool_id,omitempty"`
	SingleUserName         string            `json:"single_user_name,omitempty"`
	IdempotencyToken       string            `json:"idempotency_token,omitempty"`
}

// ClusterInfo contains the information when getting cluster info from the get request.
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
	InstancePoolID         string             `json:"instance_pool_id,omitempty"`
	SingleUserName         string             `json:"single_user_name,omitempty"`
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
