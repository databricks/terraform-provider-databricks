package compute

import (
	"fmt"
	"sort"
)

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
	// ClusterStatePending Indicates that a cluster is in the process of being created.
	ClusterStatePending = "PENDING"
	// ClusterStateRunning Indicates that a cluster has been started and is ready for use.
	ClusterStateRunning = "RUNNING"
	// ClusterStateRestarting Indicates that a cluster is in the process of restarting.
	ClusterStateRestarting = "RESTARTING"
	// ClusterStateResizing Indicates that a cluster is in the process of adding or removing nodes.
	ClusterStateResizing = "RESIZING"
	// ClusterStateTerminating Indicates that a cluster is in the process of being destroyed.
	ClusterStateTerminating = "TERMINATING"
	// ClusterStateTerminated Indicates that a cluster has been successfully destroyed.
	ClusterStateTerminated = "TERMINATED"
	// ClusterStateError This state is not used anymore. It was used to indicate a cluster
	// that failed to be created. Terminating and Terminated are used instead.
	ClusterStateError = "ERROR"
	// ClusterStateUnknown Indicates that a cluster is in an unknown state. A cluster should never be in this state.
	ClusterStateUnknown = "UNKNOWN"
)

var stateMachine = map[ClusterState][]ClusterState{
	ClusterStatePending:     {ClusterStateRunning, ClusterStateTerminating},
	ClusterStateRunning:     {ClusterStateResizing, ClusterStateRestarting, ClusterStateTerminating},
	ClusterStateRestarting:  {ClusterStateRunning, ClusterStateTerminating},
	ClusterStateResizing:    {ClusterStateRunning, ClusterStateTerminating},
	ClusterStateTerminating: {ClusterStateTerminated},
}

// CanReach returns true if cluster state can reach desired state
func (state ClusterState) CanReach(desired ClusterState) bool {
	if state == desired {
		return true
	}
	visited := map[ClusterState]bool{}
	queue := []ClusterState{state}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if _, ok := visited[current]; ok {
			continue
		}
		adjacent, ok := stateMachine[current]
		visited[current] = true
		if !ok {
			return false
		}
		for _, possible := range adjacent {
			if possible == desired {
				return true
			}
			queue = append(queue, possible)
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
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclusterattributes
type AwsAttributes struct {
	FirstOnDemand       int32           `json:"first_on_demand,omitempty" tf:"computed"`
	Availability        AwsAvailability `json:"availability,omitempty" tf:"computed"`
	ZoneID              string          `json:"zone_id,omitempty" tf:"computed"`
	InstanceProfileArn  string          `json:"instance_profile_arn,omitempty"`
	SpotBidPricePercent int32           `json:"spot_bid_price_percent,omitempty" tf:"computed"`
	EbsVolumeType       EbsVolumeType   `json:"ebs_volume_type,omitempty" tf:"computed"`
	EbsVolumeCount      int32           `json:"ebs_volume_count,omitempty" tf:"computed"`
	EbsVolumeSize       int32           `json:"ebs_volume_size,omitempty" tf:"computed"`
}

// DbfsStorageInfo contains the destination string for DBFS
type DbfsStorageInfo struct {
	Destination string `json:"destination"`
}

// S3StorageInfo contains the struct for when storing files in S3
type S3StorageInfo struct {
	// TODO: add instance profile validation check + prefix validation
	Destination      string `json:"destination"`
	Region           string `json:"region,omitempty" tf:"group:location"`
	Endpoint         string `json:"endpoint,omitempty" tf:"group:location"`
	EnableEncryption bool   `json:"enable_encryption,omitempty"`
	EncryptionType   string `json:"encryption_type,omitempty"`
	KmsKey           string `json:"kms_key,omitempty"`
	CannedACL        string `json:"canned_acl,omitempty"`
}

// StorageInfo contains the struct for either DBFS or S3 storage depending on which one is relevant.
type StorageInfo struct {
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty" tf:"group:storage"`
	S3   *S3StorageInfo   `json:"s3,omitempty" tf:"group:storage"`
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
	Type       string            `json:"type,omitempty"`
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

// NodeInstanceType encapsulates information about a specific node type
type NodeInstanceType struct {
	InstanceTypeID      string `json:"instance_type_id,omitempty"`
	LocalDisks          int32  `json:"local_disks,omitempty"`
	LocalDiskSizeGB     int32  `json:"local_disk_size_gb,omitempty"`
	LocalNVMeDisks      int32  `json:"local_nvme_disks,omitempty"`
	LocalNVMeDiskSizeGB int32  `json:"local_nvme_disk_size_gb,omitempty"`
}

// NodeType encapsulates information about a given node when using the list-node-types api
type NodeType struct {
	NodeTypeID            string                        `json:"node_type_id,omitempty"`
	MemoryMB              int32                         `json:"memory_mb,omitempty"`
	NumCores              float32                       `json:"num_cores,omitempty"`
	NumGPUs               int32                         `json:"num_gpus,omitempty"`
	SupportEBSVolumes     bool                          `json:"support_ebs_volumes,omitempty"`
	IsIOCacheEnabled      bool                          `json:"is_io_cache_enabled,omitempty"`
	SupportPortForwarding bool                          `json:"support_port_forwarding,omitempty"`
	Description           string                        `json:"description,omitempty"`
	Category              string                        `json:"category,omitempty"`
	InstanceTypeID        string                        `json:"instance_type_id,omitempty"`
	IsDeprecated          bool                          `json:"is_deprecated,omitempty"`
	IsHidden              bool                          `json:"is_hidden,omitempty"`
	SupportClusterTags    bool                          `json:"support_cluster_tags,omitempty"`
	DisplayOrder          int32                         `json:"display_order,omitempty"`
	NodeInfo              *ClusterCloudProviderNodeInfo `json:"node_info,omitempty"`
	NodeInstanceType      *NodeInstanceType             `json:"node_instance_type,omitempty"`
}

// DockerBasicAuth contains the auth information when fetching containers
type DockerBasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DockerImage contains the image url and the auth for DCS
type DockerImage struct {
	URL       string           `json:"url"`
	BasicAuth *DockerBasicAuth `json:"basic_auth,omitempty"`
}

// Cluster contains the information when trying to submit api calls or editing a cluster
type Cluster struct {
	ClusterID   string `json:"cluster_id,omitempty"`
	ClusterName string `json:"cluster_name,omitempty"`

	SparkVersion              string     `json:"spark_version"` // TODO: perhaps make a default
	NumWorkers                int32      `json:"num_workers,omitempty" tf:"group:size"`
	Autoscale                 *AutoScale `json:"autoscale,omitempty" tf:"group:size"`
	EnableElasticDisk         bool       `json:"enable_elastic_disk,omitempty" tf:"computed"`
	EnableLocalDiskEncryption bool       `json:"enable_local_disk_encryption,omitempty"`

	NodeTypeID             string         `json:"node_type_id,omitempty" tf:"group:node_type,computed"`
	DriverNodeTypeID       string         `json:"driver_node_type_id,omitempty" tf:"conflicts:instance_pool_id,computed"`
	InstancePoolID         string         `json:"instance_pool_id,omitempty" tf:"group:node_type"`
	PolicyID               string         `json:"policy_id,omitempty"`
	AwsAttributes          *AwsAttributes `json:"aws_attributes,omitempty" tf:"conflicts:instance_pool_id"`
	AutoterminationMinutes int32          `json:"autotermination_minutes,omitempty"`

	SparkConf    map[string]string `json:"spark_conf,omitempty"`
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	CustomTags   map[string]string `json:"custom_tags,omitempty"`

	SSHPublicKeys  []string      `json:"ssh_public_keys,omitempty" tf:"max_items:10"`
	InitScripts    []StorageInfo `json:"init_scripts,omitempty" tf:"max_items:10"` // TODO: tf:alias
	ClusterLogConf *StorageInfo  `json:"cluster_log_conf,omitempty"`
	DockerImage    *DockerImage  `json:"docker_image,omitempty"`

	SingleUserName   string `json:"single_user_name,omitempty"`
	IdempotencyToken string `json:"idempotency_token,omitempty"`
}

// ClusterList shows existing clusters
type ClusterList struct {
	Clusters []ClusterInfo `json:"clusters,omitempty"`
}

// ClusterInfo contains the information when getting cluster info from the get request.
type ClusterInfo struct {
	NumWorkers                int32              `json:"num_workers,omitempty"`
	AutoScale                 *AutoScale         `json:"autoscale,omitempty"`
	ClusterID                 string             `json:"cluster_id,omitempty"`
	CreatorUserName           string             `json:"creator_user_name,omitempty"`
	Driver                    *SparkNode         `json:"driver,omitempty"`
	Executors                 []SparkNode        `json:"executors,omitempty"`
	SparkContextID            int64              `json:"spark_context_id,omitempty"`
	JdbcPort                  int32              `json:"jdbc_port,omitempty"`
	ClusterName               string             `json:"cluster_name,omitempty"`
	SparkVersion              string             `json:"spark_version"`
	SparkConf                 map[string]string  `json:"spark_conf,omitempty"`
	AwsAttributes             *AwsAttributes     `json:"aws_attributes,omitempty"`
	NodeTypeID                string             `json:"node_type_id,omitempty"`
	DriverNodeTypeID          string             `json:"driver_node_type_id,omitempty"`
	SSHPublicKeys             []string           `json:"ssh_public_keys,omitempty"`
	CustomTags                map[string]string  `json:"custom_tags,omitempty"`
	ClusterLogConf            *StorageInfo       `json:"cluster_log_conf,omitempty"`
	InitScripts               []StorageInfo      `json:"init_scripts,omitempty"`
	SparkEnvVars              map[string]string  `json:"spark_env_vars,omitempty"`
	AutoterminationMinutes    int32              `json:"autotermination_minutes,omitempty"`
	EnableElasticDisk         bool               `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption bool               `json:"enable_local_disk_encryption,omitempty"`
	InstancePoolID            string             `json:"instance_pool_id,omitempty"`
	PolicyID                  string             `json:"policy_id,omitempty"`
	SingleUserName            string             `json:"single_user_name,omitempty"`
	ClusterSource             AwsAvailability    `json:"cluster_source,omitempty"`
	DockerImage               *DockerImage       `json:"docker_image,omitempty"`
	State                     ClusterState       `json:"state"`
	StateMessage              string             `json:"state_message,omitempty"`
	StartTime                 int64              `json:"start_time,omitempty"`
	TerminateTime             int64              `json:"terminate_time,omitempty"`
	LastStateLossTime         int64              `json:"last_state_loss_time,omitempty"`
	LastActivityTime          int64              `json:"last_activity_time,omitempty"`
	ClusterMemoryMb           int64              `json:"cluster_memory_mb,omitempty"`
	ClusterCores              float32            `json:"cluster_cores,omitempty"`
	DefaultTags               map[string]string  `json:"default_tags"`
	ClusterLogStatus          *LogSyncStatus     `json:"cluster_log_status,omitempty"`
	TerminationReason         *TerminationReason `json:"termination_reason,omitempty"`
}

// IsRunningOrResizing returns true if cluster is running or resizing
func (ci *ClusterInfo) IsRunningOrResizing() bool {
	return ci.State == ClusterStateRunning || ci.State == ClusterStateResizing
}

// ClusterID holds cluster ID
type ClusterID struct {
	ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
}

// ClusterPolicy defines cluster policy
type ClusterPolicy struct {
	PolicyID           string `json:"policy_id,omitempty"`
	Name               string `json:"name"`
	Definition         string `json:"definition"`
	CreatedAtTimeStamp int64  `json:"created_at_timestamp"`
}

// ClusterPolicyCreate is the endity used for request
type ClusterPolicyCreate struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

// CommandResults is the out put when the command finishes in API 1.2
type CommandResults struct {
	ResultType   string      `json:"resultType,omitempty"`
	Summary      string      `json:"summary,omitempty"`
	Cause        string      `json:"cause,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Schema       interface{} `json:"schema,omitempty"`
	Truncated    bool        `json:"truncated,omitempty"`
	IsJSONSchema bool        `json:"isJsonSchema,omitempty"`
}

// Command is the struct that contains what the 1.2 api returns for the commands api
type Command struct {
	ID      string          `json:"id,omitempty"`
	Status  string          `json:"status,omitempty"`
	Results *CommandResults `json:"results,omitempty"`
}

// InstancePoolAwsAttributes contains aws attributes for AWS Databricks deployments for instance pools
type InstancePoolAwsAttributes struct {
	Availability        AwsAvailability `json:"availability,omitempty"`
	ZoneID              string          `json:"zone_id"`
	SpotBidPricePercent int32           `json:"spot_bid_price_percent,omitempty"`
}

// InstancePoolDiskType contains disk type information for each of the different cloud service providers
type InstancePoolDiskType struct {
	AzureDiskVolumeType string `json:"azure_disk_volume_type,omitempty"`
	EbsVolumeType       string `json:"ebs_volume_type,omitempty"`
}

// InstancePoolDiskSpec contains disk size, type and count information for the pool
type InstancePoolDiskSpec struct {
	DiskType  *InstancePoolDiskType `json:"disk_type,omitempty"`
	DiskCount int32                 `json:"disk_count,omitempty"`
	DiskSize  int32                 `json:"disk_size,omitempty"`
}

// InstancePool describes the instance pool object on Databricks
type InstancePool struct {
	InstancePoolID                     string                     `json:"instance_pool_id,omitempty" tf:"computed"`
	InstancePoolName                   string                     `json:"instance_pool_name"`
	MinIdleInstances                   int32                      `json:"min_idle_instances,omitempty"`
	MaxCapacity                        int32                      `json:"max_capacity,omitempty"`
	IdleInstanceAutoTerminationMinutes int32                      `json:"idle_instance_autotermination_minutes"`
	AwsAttributes                      *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	NodeTypeID                         string                     `json:"node_type_id"`
	CustomTags                         map[string]string          `json:"custom_tags,omitempty"`
	EnableElasticDisk                  bool                       `json:"enable_elastic_disk,omitempty"`
	DiskSpec                           *InstancePoolDiskSpec      `json:"disk_spec,omitempty"`
	PreloadedSparkVersions             []string                   `json:"preloaded_spark_versions,omitempty"`
}

// InstancePoolStats contains the stats on a given pool
type InstancePoolStats struct {
	UsedCount        int32 `json:"used_count,omitempty"`
	IdleCount        int32 `json:"idle_count,omitempty"`
	PendingUsedCount int32 `json:"pending_used_count,omitempty"`
	PendingIdleCount int32 `json:"pending_idle_count,omitempty"`
}

// InstancePoolAndStats encapsulates a get response from the GET api for instance pools on Databricks
type InstancePoolAndStats struct {
	InstancePoolID                     string                     `json:"instance_pool_id,omitempty" tf:"computed"`
	InstancePoolName                   string                     `json:"instance_pool_name"`
	MinIdleInstances                   int32                      `json:"min_idle_instances,omitempty"`
	MaxCapacity                        int32                      `json:"max_capacity,omitempty"`
	AwsAttributes                      *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	NodeTypeID                         string                     `json:"node_type_id"`
	DefaultTags                        map[string]string          `json:"default_tags,omitempty" tf:"computed"`
	CustomTags                         map[string]string          `json:"custom_tags,omitempty"`
	IdleInstanceAutoTerminationMinutes int32                      `json:"idle_instance_autotermination_minutes"`
	EnableElasticDisk                  bool                       `json:"enable_elastic_disk,omitempty"`
	DiskSpec                           *InstancePoolDiskSpec      `json:"disk_spec,omitempty"`
	PreloadedSparkVersions             []string                   `json:"preloaded_spark_versions,omitempty"`
	State                              string                     `json:"state,omitempty"`
	Stats                              *InstancePoolStats         `json:"stats,omitempty"`
}

// InstancePoolList shows list of instance pools
type InstancePoolList struct {
	InstancePools []InstancePoolAndStats `json:"instance_pools"`
}

// NodeTypeList contains a list of node types
type NodeTypeList struct {
	NodeTypes []NodeType `json:"node_types,omitempty"`
}

// Sort NodeTypes within this struct
func (l *NodeTypeList) Sort() {
	sort.Slice(l.NodeTypes, func(i, j int) bool {
		if l.NodeTypes[i].IsDeprecated != l.NodeTypes[j].IsDeprecated {
			return !l.NodeTypes[i].IsDeprecated
		}
		if l.NodeTypes[i].NodeInstanceType != nil &&
			l.NodeTypes[j].NodeInstanceType != nil {
			if l.NodeTypes[i].NodeInstanceType.LocalDisks !=
				l.NodeTypes[j].NodeInstanceType.LocalDisks {
				return l.NodeTypes[i].NodeInstanceType.LocalDisks <
					l.NodeTypes[j].NodeInstanceType.LocalDisks
			}
			if l.NodeTypes[i].NodeInstanceType.LocalDiskSizeGB !=
				l.NodeTypes[j].NodeInstanceType.LocalDiskSizeGB {
				return l.NodeTypes[i].NodeInstanceType.LocalDiskSizeGB <
					l.NodeTypes[j].NodeInstanceType.LocalDiskSizeGB
			}
		}
		if l.NodeTypes[i].MemoryMB != l.NodeTypes[j].MemoryMB {
			return l.NodeTypes[i].MemoryMB < l.NodeTypes[j].MemoryMB
		}
		if l.NodeTypes[i].NumCores != l.NodeTypes[j].NumCores {
			return l.NodeTypes[i].NumCores < l.NodeTypes[j].NumCores
		}
		if l.NodeTypes[i].NumGPUs != l.NodeTypes[j].NumGPUs {
			return l.NodeTypes[i].NumGPUs < l.NodeTypes[j].NumGPUs
		}
		return l.NodeTypes[i].InstanceTypeID < l.NodeTypes[j].InstanceTypeID
	})
}

// NotebookTask contains the information for notebook jobs
type NotebookTask struct {
	NotebookPath   string            `json:"notebook_path"`
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
}

// SparkPythonTask contains the information for python jobs
type SparkPythonTask struct {
	PythonFile string   `json:"python_file"`
	Parameters []string `json:"parameters,omitempty"`
}

// SparkJarTask contains the information for jar jobs
type SparkJarTask struct {
	JarURI        string   `json:"jar_uri,omitempty"`
	MainClassName string   `json:"main_class_name,omitempty"`
	Parameters    []string `json:"parameters,omitempty"`
}

// SparkSubmitTask contains the information for spark submit jobs
type SparkSubmitTask struct {
	Parameters []string `json:"parameters,omitempty"`
}

// JobEmailNotifications contains the information for email notifications after job completion
type JobEmailNotifications struct {
	OnStart               []string `json:"on_start,omitempty"`
	OnSuccess             []string `json:"on_success,omitempty"`
	OnFailure             []string `json:"on_failure,omitempty"`
	NoAlertForSkippedRuns bool     `json:"no_alert_for_skipped_runs,omitempty"`
}

// CronSchedule contains the information for the quartz cron expression
type CronSchedule struct {
	QuartzCronExpression string `json:"quartz_cron_expression"`
	TimezoneID           string `json:"timezone_id"`
	PauseStatus          string `json:"pause_status,omitempty" tf:"computed"`
}

// JobSettings contains the information for configuring a job on databricks
type JobSettings struct {
	Name string `json:"name,omitempty" tf:"default:Untitled"`

	ExistingClusterID string   `json:"existing_cluster_id,omitempty" tf:"group:cluster_type"`
	NewCluster        *Cluster `json:"new_cluster,omitempty" tf:"group:cluster_type"`

	NotebookTask    *NotebookTask    `json:"notebook_task,omitempty" tf:"group:task_type"`
	SparkJarTask    *SparkJarTask    `json:"spark_jar_task,omitempty" tf:"group:task_type"`
	SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty" tf:"group:task_type"`
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty" tf:"group:task_type"`

	Libraries              []Library     `json:"libraries,omitempty" tf:"slice_set,alias:library"`
	TimeoutSeconds         int32         `json:"timeout_seconds,omitempty"`
	MaxRetries             int32         `json:"max_retries,omitempty"`
	MinRetryIntervalMillis int32         `json:"min_retry_interval_millis,omitempty"`
	RetryOnTimeout         bool          `json:"retry_on_timeout,omitempty"`
	Schedule               *CronSchedule `json:"schedule,omitempty"`
	MaxConcurrentRuns      int32         `json:"max_concurrent_runs,omitempty"`

	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
}

// Job contains the information when using a GET request from the Databricks Jobs api
type Job struct {
	JobID           int64        `json:"job_id,omitempty"`
	CreatorUserName string       `json:"creator_user_name,omitempty"`
	Settings        *JobSettings `json:"settings,omitempty"`
	CreatedTime     int64        `json:"created_time,omitempty"`
}

// ID returns job id as string
func (j Job) ID() string {
	return fmt.Sprintf("%d", j.JobID)
}

// UpdateJobRequest ...
type UpdateJobRequest struct {
	JobID       int64        `json:"job_id,omitempty" url:"job_id,omitempty"`
	NewSettings *JobSettings `json:"new_settings,omitempty" url:"new_settings,omitempty"`
}

// PyPi is a python library hosted on PYPI
type PyPi struct {
	Package string `json:"package"`
	Repo    string `json:"repo,omitempty"`
}

// Maven is a jar library hosted on Maven
type Maven struct {
	Coordinates string   `json:"coordinates"`
	Repo        string   `json:"repo,omitempty"`
	Exclusions  []string `json:"exclusions,omitempty"`
}

// Cran is a R library hosted on Maven
type Cran struct {
	Package string `json:"package"`
	Repo    string `json:"repo,omitempty"`
}

// SortOrder - constants for API
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterlistorder
type SortOrder string

// constants for SortOrder
const (
	SortDescending SortOrder = "DESC"
	SortAscending  SortOrder = "ASC"
)

// ClusterEventType - constants for API
type ClusterEventType string

// Constants for Event Types
const (
	EvTypeCreating            ClusterEventType = "CREATING"
	EvTypeDidNotExpandDisk    ClusterEventType = "DID_NOT_EXPAND_DISK"
	EvTypeExpandedDisk        ClusterEventType = "EXPANDED_DISK"
	EvTypeFailedToExpandDisk  ClusterEventType = "FAILED_TO_EXPAND_DISK"
	EvTypeInitScriptsStarting ClusterEventType = "INIT_SCRIPTS_STARTING"
	EvTypeInitScriptsFinished ClusterEventType = "INIT_SCRIPTS_FINISHED"
	EvTypeStarting            ClusterEventType = "STARTING"
	EvTypeRestarting          ClusterEventType = "RESTARTING"
	EvTypeTerminating         ClusterEventType = "TERMINATING"
	EvTypeEdited              ClusterEventType = "EDITED"
	EvTypeRunning             ClusterEventType = "RUNNING"
	EvTypeResizing            ClusterEventType = "RESIZING"
	EvTypeUpsizeCompleted     ClusterEventType = "UPSIZE_COMPLETED"
	EvTypeNodesLost           ClusterEventType = "NODES_LOST"
	EvTypeDriverHealthy       ClusterEventType = "DRIVER_HEALTHY"
	EvTypeDriverUnavailable   ClusterEventType = "DRIVER_UNAVAILABLE"
	EvTypeSparkException      ClusterEventType = "SPARK_EXCEPTION"
	EvTypeDriverNotResponding ClusterEventType = "DRIVER_NOT_RESPONDING"
	EvTypeDbfsDown            ClusterEventType = "DBFS_DOWN"
	EvTypeMetastoreDown       ClusterEventType = "METASTORE_DOWN"
	EvTypeNodeBlacklisted     ClusterEventType = "NODE_BLACKLISTED"
	EvTypePinned              ClusterEventType = "PINNED"
	EvTypeUnpinned            ClusterEventType = "UNPINNED"
)

// EventsRequest - request structure
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#request-structure
type EventsRequest struct {
	ClusterID  string             `json:"cluster_id"`
	StartTime  int64              `json:"start_time,omitempty"`
	EndTime    int64              `json:"end_time,omitempty"`
	Order      SortOrder          `json:"order,omitempty"`
	EventTypes []ClusterEventType `json:"event_types,omitempty"`
	Offset     int64              `json:"offset,omitempty"`
	Limit      int64              `json:"limit,omitempty"`
	MaxItems   uint               `json:"-"`
}

// ClusterSize is structure to keep
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclustersize
type ClusterSize struct {
	NumWorkers int32      `json:"num_workers"`
	AutoScale  *AutoScale `json:"autoscale"`
}

// ResizeCause holds reason for resizing
type ResizeCause string

// EventDetails - details about specific events
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#clustereventseventdetails
type EventDetails struct {
	CurrentNumWorkers   int32              `json:"current_num_workers,omitempty"`
	TargetNumWorkers    int32              `json:"target_num_workers,omitempty"`
	PreviousAttributes  *AwsAttributes     `json:"previous_attributes,omitempty"`
	Attributes          *AwsAttributes     `json:"attributes,omitempty"`
	PreviousClusterSize *ClusterSize       `json:"previous_cluster_size,omitempty"`
	ClusterSize         *ClusterSize       `json:"cluster_size,omitempty"`
	ResizeCause         *ResizeCause       `json:"cause,omitempty"`
	Reason              *TerminationReason `json:"reason,omitempty"`
	User                string             `json:"user"`
}

// ClusterEvent - event information
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#clustereventsclusterevent
type ClusterEvent struct {
	ClusterID string           `json:"cluster_id"`
	Timestamp int64            `json:"timestamp"`
	Type      ClusterEventType `json:"type"`
	Details   EventDetails     `json:"details"`
}

// EventsResponse - answer from API
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#response-structure
type EventsResponse struct {
	Events     []ClusterEvent `json:"events"`
	NextPage   *EventsRequest `json:"next_page"`
	TotalCount int64          `json:"total_count"`
}
