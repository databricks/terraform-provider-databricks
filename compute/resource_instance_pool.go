package compute

import (
	"context"
	"log"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// NewInstancePoolsAPI creates InstancePoolsAPI instance from provider meta
func NewInstancePoolsAPI(m interface{}) InstancePoolsAPI {
	return InstancePoolsAPI{m.(*common.DatabricksClient), context.TODO()}
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
func (a InstancePoolsAPI) Update(instancePoolInfo InstancePoolAndStats) error {
	return a.client.Post(a.context, "/instance-pools/edit", instancePoolInfo, nil)
}

// Read retrieves the information for a instance pool given its identifier
func (a InstancePoolsAPI) Read(instancePoolID string) (InstancePoolAndStats, error) {
	var instancePoolInfo InstancePoolAndStats
	err := a.client.Get(a.context, "/instance-pools/get", map[string]string{
		"instance_pool_id": instancePoolID,
	}, &instancePoolInfo)
	return instancePoolInfo, err
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

func ResourceInstancePool() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstancePoolCreate,
		Read:   resourceInstancePoolRead,
		Update: resourceInstancePoolUpdate,
		Delete: resourceInstancePoolDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"instance_pool_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"min_idle_instances": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_capacity": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"idle_instance_autotermination_minutes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"aws_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"availability": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "SPOT",
							ForceNew: true,
						},
						"zone_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"spot_bid_price_percent": {
							Type:     schema.TypeInt,
							Default:  "100",
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"node_type_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_tags": {
				Deprecated: "`default_tags` are going to be removed in v0.3",
				Type:       schema.TypeMap,
				Computed:   true,
			},
			"custom_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
			"enable_elastic_disk": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},
			"disk_spec": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ebs_volume_type": {
							Deprecated:    "`ebs_volume_type` is going to be moved to `disk_type` sub-block in 0.3.",
							Type:          schema.TypeString,
							Optional:      true,
							ForceNew:      true,
							ConflictsWith: []string{"disk_spec.0.azure_disk_volume_type"},
							ValidateFunc: validation.StringInSlice(
								[]string{
									EbsVolumeTypeGeneralPurposeSsd,
									EbsVolumeTypeThroughputOptimizedHdd,
								}, false),
						},
						"azure_disk_volume_type": {
							Deprecated:    "`azure_disk_volume_type` is going to be moved to `disk_type` sub-block in 0.3.",
							Type:          schema.TypeString,
							Optional:      true,
							ForceNew:      true,
							ConflictsWith: []string{"disk_spec.0.ebs_volume_type"},
							ValidateFunc: validation.StringInSlice(
								[]string{
									AzureDiskVolumeTypePremium,
									AzureDiskVolumeTypeStandard,
								}, false),
						},
						"disk_count": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"disk_size": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"preloaded_spark_versions": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			//TODO: Determine what this does from a state management perspective
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
		},
	}
}

func convertMapStringInterfaceToStringString(m map[string]interface{}) map[string]string {
	response := make(map[string]string)
	for k, v := range m {
		if v != nil {
			response[k] = v.(string)
		}
	}
	return response
}

func getMapFromOneItemList(input interface{}) map[string]interface{} {
	inputList := input.([]interface{})
	if len(inputList) >= 1 {
		return inputList[0].(map[string]interface{})
	}
	return nil
}

func resourceInstancePoolCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)

	var instancePool InstancePool
	var instancePoolAwsAttributes InstancePoolAwsAttributes
	var instancePoolDiskSpec InstancePoolDiskSpec
	var instancePoolDiskSpecDiskType InstancePoolDiskType
	instancePool.InstancePoolName = d.Get("instance_pool_name").(string)
	instancePool.MinIdleInstances = int32(d.Get("min_idle_instances").(int))
	instancePool.IdleInstanceAutoTerminationMinutes = int32(d.Get("idle_instance_autotermination_minutes").(int))

	if maxCapacity, ok := d.GetOk("max_capacity"); ok {
		instancePool.MaxCapacity = int32(maxCapacity.(int))
	}

	if awsAttributesSchema, ok := d.GetOk("aws_attributes"); ok {
		awsAttributesMap := getMapFromOneItemList(awsAttributesSchema)
		if availability, ok := awsAttributesMap["availability"]; ok {
			instancePoolAwsAttributes.Availability = AwsAvailability(availability.(string))
		}
		if zoneID, ok := awsAttributesMap["zone_id"]; ok {
			instancePoolAwsAttributes.ZoneID = zoneID.(string)
		}
		if spotBidPricePercent, ok := awsAttributesMap["spot_bid_price_percent"]; ok {
			//val, _ := strconv.ParseInt(spotBidPricePercent.(string), 10, 32)
			instancePoolAwsAttributes.SpotBidPricePercent = int32(spotBidPricePercent.(int))
		}
		instancePool.AwsAttributes = &instancePoolAwsAttributes
	}

	if nodeTypeID, ok := d.GetOk("node_type_id"); ok {
		instancePool.NodeTypeID = nodeTypeID.(string)
	}

	if customTags, ok := d.GetOk("custom_tags"); ok {
		tags := customTags.(map[string]interface{})
		instancePool.CustomTags = convertMapStringInterfaceToStringString(tags)
	}

	if enableElasticDisk, ok := d.GetOk("enable_elastic_disk"); ok {
		instancePool.EnableElasticDisk = enableElasticDisk.(bool)
	}

	if diskSpec, ok := d.GetOk("disk_spec"); ok {
		diskSpecList := diskSpec.([]interface{})
		if len(diskSpecList) > 0 {
			diskSpecMap := diskSpecList[0].(map[string]interface{})
			if ebsVolumeType, ok := diskSpecMap["ebs_volume_type"]; ok {
				instancePoolDiskSpecDiskType.EbsVolumeType = ebsVolumeType.(string)
			}
			if azureDiskVolumeType, ok := diskSpecMap["azure_disk_volume_type"]; ok {
				instancePoolDiskSpecDiskType.AzureDiskVolumeType = azureDiskVolumeType.(string)
			}
			instancePoolDiskSpec.DiskType = &instancePoolDiskSpecDiskType

			if diskCount, ok := diskSpecMap["disk_count"]; ok {
				instancePoolDiskSpec.DiskCount = int32(diskCount.(int))
			}
			if diskSize, ok := diskSpecMap["disk_size"]; ok {
				instancePoolDiskSpec.DiskSize = int32(diskSize.(int))
			}
			instancePool.DiskSpec = &instancePoolDiskSpec
		}
	}

	if sparkVersions, ok := d.GetOk("preloaded_spark_versions"); ok {
		instancePool.PreloadedSparkVersions = internal.ConvertListInterfaceToString(sparkVersions.([]interface{}))
	}

	instancePoolInfo, err := NewInstancePoolsAPI(client).Create(instancePool)
	if err != nil {
		return err
	}
	d.SetId(instancePoolInfo.InstancePoolID)
	return resourceInstancePoolRead(d, m)
}

func resourceInstancePoolRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	instancePoolInfo, err := NewInstancePoolsAPI(client).Read(id)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}

	err = d.Set("instance_pool_name", instancePoolInfo.InstancePoolName)
	if err != nil {
		return err
	}
	err = d.Set("min_idle_instances", int(instancePoolInfo.MinIdleInstances))
	if err != nil {
		return err
	}

	err = d.Set("max_capacity", int(instancePoolInfo.MaxCapacity))
	if err != nil {
		return err
	}

	err = d.Set("idle_instance_autotermination_minutes", int(instancePoolInfo.IdleInstanceAutoTerminationMinutes))
	if err != nil {
		return err
	}

	if instancePoolInfo.AwsAttributes != nil {
		awsAtts := map[string]interface{}{}
		//if _, ok := d.GetOk("aws_attributes.availability"); ok {
		awsAtts["availability"] = string(instancePoolInfo.AwsAttributes.Availability)
		//}
		//if _, ok := d.GetOk("aws_attributes.zone_id"); ok {
		awsAtts["zone_id"] = instancePoolInfo.AwsAttributes.ZoneID
		//}
		//if _, ok := d.GetOk("aws_attributes.spot_bid_price_percent"); ok {
		awsAtts["spot_bid_price_percent"] = int(instancePoolInfo.AwsAttributes.SpotBidPricePercent)
		//}
		awsAttsList := []map[string]interface{}{awsAtts}
		err = d.Set("aws_attributes", awsAttsList)
		if err != nil {
			return err
		}
	}

	err = d.Set("node_type_id", instancePoolInfo.NodeTypeID)
	if err != nil {
		return err
	}

	err = d.Set("enable_elastic_disk", instancePoolInfo.EnableElasticDisk)
	if err != nil {
		return err
	}

	if instancePoolInfo.DiskSpec != nil {
		diskSpecList := []interface{}{}
		diskSpecListItem := map[string]interface{}{}
		if instancePoolInfo.DiskSpec.DiskType != nil {
			if instancePoolInfo.DiskSpec.DiskCount >= 0 {
				diskSpecListItem["disk_count"] = instancePoolInfo.DiskSpec.DiskCount
			}
			if instancePoolInfo.DiskSpec.DiskSize >= 0 {
				diskSpecListItem["disk_size"] = instancePoolInfo.DiskSpec.DiskSize
			}
		}
		if instancePoolInfo.DiskSpec.DiskType.EbsVolumeType != "" {
			diskSpecListItem["ebs_volume_type"] = instancePoolInfo.DiskSpec.DiskType.EbsVolumeType
		}
		if instancePoolInfo.DiskSpec.DiskType.AzureDiskVolumeType != "" {
			diskSpecListItem["azure_disk_volume_type"] = instancePoolInfo.DiskSpec.DiskType.AzureDiskVolumeType
		}
		diskSpecList = append(diskSpecList, diskSpecListItem)
		err = d.Set("disk_spec", diskSpecList)
		if err != nil {
			return err
		}
	}

	if len(instancePoolInfo.CustomTags) > 0 {
		err = d.Set("custom_tags", instancePoolInfo.CustomTags)
	}

	if len(instancePoolInfo.DefaultTags) > 0 {
		err = d.Set("default_tags", instancePoolInfo.DefaultTags)
	}

	if len(instancePoolInfo.PreloadedSparkVersions) > 0 {
		err = d.Set("preloaded_spark_versions", instancePoolInfo.PreloadedSparkVersions)
	}
	d.SetId(id)
	return err
}

func resourceInstancePoolUpdate(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)

	var instancePoolInfo InstancePoolAndStats
	instancePoolInfo.InstancePoolName = d.Get("instance_pool_name").(string)
	instancePoolInfo.MinIdleInstances = int32(d.Get("min_idle_instances").(int))
	instancePoolInfo.IdleInstanceAutoTerminationMinutes = int32(d.Get("idle_instance_autotermination_minutes").(int))
	instancePoolInfo.InstancePoolID = id
	instancePoolInfo.NodeTypeID = d.Get("node_type_id").(string)

	if maxCapacity, ok := d.GetOk("max_capacity"); ok {
		instancePoolInfo.MaxCapacity = int32(maxCapacity.(int))
	}

	err := NewInstancePoolsAPI(client).Update(instancePoolInfo)
	if err != nil {
		return err
	}
	return resourceInstancePoolRead(d, m)
}

func resourceInstancePoolDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	id := d.Id()
	err := NewInstancePoolsAPI(client).Delete(id)
	return err
}
