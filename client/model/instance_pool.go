package model

type InstancePoolAwsAttributes struct {
	Availability        AwsAvailability `json:"availability,omitempty"`
	ZoneId              string          `json:"zone_id,omitempty"`
	SpotBidPricePercent int32           `json:"spot_bid_price_percent,omitempty"`
}

type InstancePoolDiskType struct {
	AzureDiskVolumeType string `json:"azure_disk_volume_type,omitempty"`
	EbsVolumeType       string `json:"ebs_volume_type,omitempty"`
}

type InstancePoolDiskSpec struct {
	DiskType  *InstancePoolDiskType `json:"disk_type,omitempty"`
	DiskCount int32                 `json:"disk_count,omitempty"`
	DiskSize  int32                 `json:"disk_size,omitempty"`
}

type InstancePool struct {
	InstancePoolName                   string                     `json:"instance_pool_name,omitempty"`
	MinIdleInstances                   int32                      `json:"min_idle_instances,omitempty"`
	MaxCapacity                        int32                      `json:"max_capacity,omitempty"`
	AwsAttributes                      *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	NodeTypeId                         string                     `json:"node_type_id,omitempty"`
	CustomTags                         map[string]string          `json:"custom_tags,omitempty"`
	IdleInstanceAutoTerminationMinutes int32                      `json:"idle_instance_autotermination_minutes,omitempty"`
	EnableElasticDisk                  bool                       `json:"enable_elastic_disk,omitempty"`
	DiskSpec                           *InstancePoolDiskSpec      `json:"disk_spec,omitempty"`
	PreloadedSparkVersions             []string                   `json:"preloaded_spark_versions,omitempty"`
}

type InstancePoolState string

const (
	InstancePoolStateActive  InstancePoolState = "ACTIVE"
	InstancePoolStateDeleted InstancePoolState = "DELETED"
)

type InstancePoolStats struct {
	UsedCount        int32 `json:"used_count,omitempty"`
	IdleCount        int32 `json:"idle_count,omitempty"`
	PendingUsedCount int32 `json:"pending_used_count,omitempty"`
	PendingIdleCount int32 `json:"pending_idle_count,omitempty"`
}

type InstancePoolInfo struct {
	InstancePoolId                     string                     `json:"instance_pool_id,omitempty"`
	InstancePoolName                   string                     `json:"instance_pool_name,omitempty"`
	MinIdleInstances                   int32                      `json:"min_idle_instances,omitempty"`
	MaxCapacity                        int32                      `json:"max_capacity,omitempty"`
	AwsAttributes                      *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	NodeTypeId                         string                     `json:"node_type_id,omitempty"`
	DefaultTags                        map[string]string          `json:"default_tags,omitempty"`
	CustomTags                         map[string]string          `json:"custom_tags,omitempty"`
	IdleInstanceAutoTerminationMinutes int32                      `json:"idle_instance_autotermination_minutes,omitempty"`
	EnableElasticDisk                  bool                       `json:"enable_elastic_disk,omitempty"`
	DiskSpec                           *InstancePoolDiskSpec      `json:"disk_spec,omitempty"`
	PreloadedSparkVersions             []string                   `json:"preloaded_spark_versions,omitempty"`
	State                              InstancePoolState          `json:"state,omitempty"`
	Stats                              *InstancePoolStats         `json:"stats,omitempty"`
}
