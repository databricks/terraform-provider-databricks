package pools

import (
	"context"
	"strings"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// InstancePoolAwsAttributes contains aws attributes for AWS Databricks deployments for instance pools
type InstancePoolAwsAttributes struct {
	Availability        clusters.Availability `json:"availability,omitempty" tf:"force_new"`
	ZoneID              string                `json:"zone_id,omitempty" tf:"computed,force_new"`
	SpotBidPricePercent int32                 `json:"spot_bid_price_percent,omitempty" tf:"force_new"`
}

// InstancePoolAzureAttributes contains aws attributes for Azure Databricks deployments for instance pools
// https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/instance-pools#clusterinstancepoolazureattributes
type InstancePoolAzureAttributes struct {
	Availability    clusters.Availability `json:"availability,omitempty" tf:"force_new"`
	SpotBidMaxPrice float64               `json:"spot_bid_max_price,omitempty" tf:"force_new"`
}

// InstancePoolGcpAttributes contains aws attributes for GCP Databricks deployments for instance pools
// https://docs.gcp.databricks.com/dev-tools/api/latest/instance-pools.html#instancepoolgcpattributes
type InstancePoolGcpAttributes struct {
	Availability  clusters.Availability `json:"gcp_availability,omitempty" tf:"force_new"`
	LocalSsdCount int32                 `json:"local_ssd_count,omitempty" tf:"computed,force_new"`
	ZoneID        string                `json:"zone_id,omitempty" tf:"computed,force_new"`
}

// InstancePoolDiskType contains disk type information for each of the different cloud service providers
type InstancePoolDiskType struct {
	AzureDiskVolumeType string `json:"azure_disk_volume_type,omitempty" tf:"force_new"`
	EbsVolumeType       string `json:"ebs_volume_type,omitempty" tf:"force_new"`
}

// InstancePoolDiskSpec contains disk size, type and count information for the pool
type InstancePoolDiskSpec struct {
	DiskType  *InstancePoolDiskType `json:"disk_type,omitempty"`
	DiskCount int32                 `json:"disk_count,omitempty"`
	DiskSize  int32                 `json:"disk_size,omitempty"`
}

type AwsAllocationStrategy string

const (
	// AwsAllocationStrategyLowestPrice is allocation type for AWS Instance fleets
	AwsAllocationStrategyLowestPrice = "LOWEST_PRICE"
	// AwsAllocationStrategyCapacityOptimized is allocation type for AWS Instance fleets
	AwsAllocationStrategyCapacityOptimized = "CAPACITY_OPTIMIZED"
)

type AwsFleetOption struct {
	AllocationStrategy      AwsAllocationStrategy `json:"allocation_strategy" tf:"force_new,suppress_diff"`
	InstancePoolsToUseCount int32                 `json:"instance_pools_to_use_count,omitempty" tf:"suppress_diff"`
}

type AwsFleetLaunchTemplateOverride struct {
	AvailabilityZone string `json:"availability_zone" tf:"force_new,suppress_diff"`
	InstanceType     string `json:"instance_type" tf:"force_new,suppress_diff"`
}

type AwsInstancePoolFleetAttributes struct {
	FleetSpotOption             *AwsFleetOption                  `json:"fleet_spot_option,omitempty" tf:"force_new,suppress_diff,conflicts:fleet_on_demand_option"`
	FleetOnDemandOption         *AwsFleetOption                  `json:"fleet_on_demand_option,omitempty" tf:"force_new,suppress_diff,conflicts:fleet_spot_option"`
	FleetLaunchTemplateOverride []AwsFleetLaunchTemplateOverride `json:"launch_template_overrides" tf:"suppress_diff,force_new,slice_set,alias:launch_template_override"`
}

// InstancePool describes the instance pool object on Databricks
type InstancePool struct {
	InstancePoolID                     string                          `json:"instance_pool_id,omitempty" tf:"computed"`
	InstancePoolName                   string                          `json:"instance_pool_name"`
	MinIdleInstances                   int32                           `json:"min_idle_instances,omitempty"`
	MaxCapacity                        int32                           `json:"max_capacity,omitempty" tf:"suppress_diff"`
	IdleInstanceAutoTerminationMinutes int32                           `json:"idle_instance_autotermination_minutes"`
	AwsAttributes                      *InstancePoolAwsAttributes      `json:"aws_attributes,omitempty" tf:"force_new,suppress_diff"`
	AwsInstancePoolFleetAttributes     *AwsInstancePoolFleetAttributes `json:"instance_pool_fleet_attributes,omitempty" tf:"force_new,suppress_diff,conflicts:node_type_id"`
	AzureAttributes                    *InstancePoolAzureAttributes    `json:"azure_attributes,omitempty" tf:"force_new,suppress_diff"`
	GcpAttributes                      *InstancePoolGcpAttributes      `json:"gcp_attributes,omitempty" tf:"force_new,suppress_diff"`
	NodeTypeID                         string                          `json:"node_type_id,omitempty" tf:"suppress_diff,force_new,conflicts:instance_pool_fleet_attributes"`
	CustomTags                         map[string]string               `json:"custom_tags" tf:"optional"`
	EnableElasticDisk                  bool                            `json:"enable_elastic_disk,omitempty" tf:"force_new,suppress_diff"`
	DiskSpec                           *InstancePoolDiskSpec           `json:"disk_spec,omitempty" tf:"force_new"`
	PreloadedSparkVersions             []string                        `json:"preloaded_spark_versions,omitempty" tf:"force_new"`
	PreloadedDockerImages              []clusters.DockerImage          `json:"preloaded_docker_images,omitempty" tf:"force_new,slice_set,alias:preloaded_docker_image"`
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
	InstancePoolID                     string                           `json:"instance_pool_id,omitempty" tf:"computed"`
	InstancePoolName                   string                           `json:"instance_pool_name"`
	MinIdleInstances                   int32                            `json:"min_idle_instances,omitempty"`
	MaxCapacity                        int32                            `json:"max_capacity,omitempty"`
	AwsAttributes                      *InstancePoolAwsAttributes       `json:"aws_attributes,omitempty"`
	AwsInstancePoolFleetAttributes     []AwsInstancePoolFleetAttributes `json:"instance_pool_fleet_attributes,omitempty"`
	AzureAttributes                    *InstancePoolAzureAttributes     `json:"azure_attributes,omitempty"`
	GcpAttributes                      *InstancePoolGcpAttributes       `json:"gcp_attributes,omitempty"`
	NodeTypeID                         string                           `json:"node_type_id,omitempty"`
	DefaultTags                        map[string]string                `json:"default_tags,omitempty" tf:"computed"`
	CustomTags                         map[string]string                `json:"custom_tags,omitempty"`
	IdleInstanceAutoTerminationMinutes int32                            `json:"idle_instance_autotermination_minutes"`
	EnableElasticDisk                  bool                             `json:"enable_elastic_disk,omitempty"`
	DiskSpec                           *InstancePoolDiskSpec            `json:"disk_spec,omitempty"`
	PreloadedSparkVersions             []string                         `json:"preloaded_spark_versions,omitempty"`
	State                              string                           `json:"state,omitempty"`
	Stats                              *InstancePoolStats               `json:"stats,omitempty"`
	PreloadedDockerImages              []clusters.DockerImage           `json:"preloaded_docker_images,omitempty" tf:"slice_set,alias:preloaded_docker_image"`
}

// InstancePoolList shows list of instance pools
type InstancePoolList struct {
	InstancePools []InstancePoolAndStats `json:"instance_pools"`
}

// NewInstancePoolsAPI creates InstancePoolsAPI instance from provider meta
func NewInstancePoolsAPI(ctx context.Context, m any) InstancePoolsAPI {
	return InstancePoolsAPI{m.(*common.DatabricksClient), ctx}
}

// InstancePoolsAPI exposes the instance pools api
type InstancePoolsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates the instance pool to given the instance pool configuration
func (a InstancePoolsAPI) Create(instancePool InstancePool) (InstancePoolAndStats, error) {
	var instancePoolInfo InstancePoolAndStats
	err := a.client.Post(a.context, "/instance-pools/create", instancePool, &instancePoolInfo)
	return instancePoolInfo, err
}

// Update edits the configuration of a instance pool to match the provided attributes and size
func (a InstancePoolsAPI) Update(ip InstancePool) error {
	return a.client.Post(a.context, "/instance-pools/edit", ip, nil)
}

// Read retrieves the information for a instance pool given its identifier
func (a InstancePoolsAPI) Read(instancePoolID string) (ip InstancePool, err error) {
	err = a.client.Get(a.context, "/instance-pools/get", map[string]string{
		"instance_pool_id": instancePoolID,
	}, &ip)
	return
}

// List retrieves the list of existing instance pools
func (a InstancePoolsAPI) List() (ipl InstancePoolList, err error) {
	err = a.client.Get(a.context, "/instance-pools/list", nil, &ipl)
	return
}

// Delete terminates a instance pool given its ID
func (a InstancePoolsAPI) Delete(instancePoolID string) error {
	return a.client.Post(a.context, "/instance-pools/delete", map[string]string{
		"instance_pool_id": instancePoolID,
	}, nil)
}

// ResourceInstancePool ...
func ResourceInstancePool() common.Resource {
	s := common.StructToSchema(InstancePool{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["enable_elastic_disk"].Default = true
		s["aws_attributes"].ConflictsWith = []string{"azure_attributes", "gcp_attributes"}
		s["azure_attributes"].ConflictsWith = []string{"aws_attributes", "gcp_attributes"}
		s["gcp_attributes"].ConflictsWith = []string{"azure_attributes", "aws_attributes"}
		if v, err := common.SchemaPath(s, "aws_attributes", "availability"); err == nil {
			v.Default = clusters.AwsAvailabilitySpot
			v.ValidateFunc = validation.StringInSlice([]string{
				clusters.AwsAvailabilityOnDemand,
				clusters.AwsAvailabilitySpot,
			}, false)
		}
		if v, err := common.SchemaPath(s, "aws_attributes", "spot_bid_price_percent"); err == nil {
			v.Default = 100
		}
		common.MustSchemaPath(s, "aws_attributes", "zone_id").DiffSuppressFunc = func(k, oldValue, newValue string, d *schema.ResourceData) bool {
			return oldValue != "" && strings.ToLower(newValue) == "auto"
		}

		if v, err := common.SchemaPath(s, "azure_attributes", "availability"); err == nil {
			v.Default = clusters.AzureAvailabilityOnDemand
			v.ValidateFunc = validation.StringInSlice([]string{
				clusters.AzureAvailabilitySpot,
				clusters.AzureAvailabilityOnDemand,
			}, false)
		}
		if v, err := common.SchemaPath(s, "gcp_attributes", "gcp_availability"); err == nil {
			v.Default = clusters.GcpAvailabilityOnDemand
			v.ValidateFunc = validation.StringInSlice([]string{
				clusters.GcpAvailabilityOnDemand,
				clusters.GcpAvailabilityPreemptible,
				clusters.GcpAvailabilityPreemptibleWithFallback,
			}, false)
		}
		if v, err := common.SchemaPath(s, "disk_spec", "disk_type", "azure_disk_volume_type"); err == nil {
			// nolint
			v.ValidateFunc = validation.StringInSlice([]string{
				clusters.AzureDiskVolumeTypePremium,
				clusters.AzureDiskVolumeTypeStandard,
			}, false)
		}
		if v, err := common.SchemaPath(s, "disk_spec", "disk_type", "ebs_volume_type"); err == nil {
			// nolint
			v.ValidateFunc = validation.StringInSlice([]string{
				clusters.EbsVolumeTypeGeneralPurposeSsd,
				clusters.EbsVolumeTypeThroughputOptimizedHdd,
			}, false)
		}
		if v, err := common.SchemaPath(s, "instance_pool_fleet_attributes", "fleet_on_demand_option", "allocation_strategy"); err == nil {
			v.ValidateFunc = validation.StringInSlice([]string{AwsAllocationStrategyLowestPrice}, false)
		}
		if v, err := common.SchemaPath(s, "instance_pool_fleet_attributes", "fleet_spot_option", "allocation_strategy"); err == nil {
			v.ValidateFunc = validation.StringInSlice([]string{
				AwsAllocationStrategyLowestPrice,
				AwsAllocationStrategyCapacityOptimized,
			}, false)
		}
		if v, err := common.SchemaPath(s, "preloaded_docker_image", "url"); err == nil {
			v.ForceNew = true
		}
		if v, err := common.SchemaPath(s, "preloaded_docker_image", "basic_auth"); err == nil {
			v.ForceNew = true
		}
		if v, err := common.SchemaPath(s, "preloaded_docker_image", "basic_auth", "username"); err == nil {
			v.ForceNew = true
		}
		if v, err := common.SchemaPath(s, "preloaded_docker_image", "basic_auth", "password"); err == nil {
			v.ForceNew = true
		}

		return s
	})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ip InstancePool
			common.DataToStructPointer(d, s, &ip)
			instancePoolInfo, err := NewInstancePoolsAPI(ctx, c).Create(ip)
			if err != nil {
				return err
			}
			d.SetId(instancePoolInfo.InstancePoolID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ip, err := NewInstancePoolsAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(ip, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ip InstancePool
			common.DataToStructPointer(d, s, &ip)
			ip.InstancePoolID = d.Id()
			return NewInstancePoolsAPI(ctx, c).Update(ip)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewInstancePoolsAPI(ctx, c).Delete(d.Id())
		},
	}
}
