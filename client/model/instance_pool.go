package model

// InstancePoolAwsAttributes contains aws attributes for AWS Databricks deployments for instance pools
type InstancePoolAwsAttributes struct {
	Availability        AwsAvailability `json:"availability,omitempty"`
	ZoneID              string          `json:"zone_id,omitempty"`
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
	InstancePoolName                   string                     `json:"instance_pool_name,omitempty"`
	MinIdleInstances                   int32                      `json:"min_idle_instances,omitempty"`
	MaxCapacity                        int32                      `json:"max_capacity,omitempty"`
	AwsAttributes                      *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	NodeTypeID                         string                     `json:"node_type_id,omitempty"`
	CustomTags                         map[string]string          `json:"custom_tags,omitempty"`
	IdleInstanceAutoTerminationMinutes int32                      `json:"idle_instance_autotermination_minutes,omitempty"`
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

// InstancePoolInfo encapsulates a get response from the GET api for instance pools on Databricks
type InstancePoolInfo struct {
	InstancePoolID                     string                     `json:"instance_pool_id,omitempty"`
	InstancePoolName                   string                     `json:"instance_pool_name,omitempty"`
	MinIdleInstances                   int32                      `json:"min_idle_instances,omitempty"`
	MaxCapacity                        int32                      `json:"max_capacity,omitempty"`
	AwsAttributes                      *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	NodeTypeID                         string                     `json:"node_type_id,omitempty"`
	DefaultTags                        map[string]string          `json:"default_tags,omitempty"`
	CustomTags                         map[string]string          `json:"custom_tags,omitempty"`
	IdleInstanceAutoTerminationMinutes int32                      `json:"idle_instance_autotermination_minutes,omitempty"`
	EnableElasticDisk                  bool                       `json:"enable_elastic_disk,omitempty"`
	DiskSpec                           *InstancePoolDiskSpec      `json:"disk_spec,omitempty"`
	PreloadedSparkVersions             []string                   `json:"preloaded_spark_versions,omitempty"`
	State                              string                     `json:"state,omitempty"`
	Stats                              *InstancePoolStats         `json:"stats,omitempty"`
}
