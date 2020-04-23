package databricks

import (
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	"strings"
)

func resourceInstancePool() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstancePoolCreate,
		Read:   resourceInstancePoolRead,
		Update: resourceInstancePoolUpdate,
		Delete: resourceInstancePoolDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"instance_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"min_idle_instances": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"max_capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"idle_instance_autotermination_minutes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"aws_attributes": &schema.Schema{
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
			"node_type_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_tags": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
			"custom_tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
			"enable_elastic_disk": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},
			"disk_spec": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ebs_volume_type": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Default:  nil,
						},
						"azure_disk_volume_type": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Default:  nil,
						},
						"disk_count": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Default:  nil,
						},
						"disk_size": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Default:  nil,
						},
					},
				},
			},
			"preloaded_spark_versions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			//TODO: Determine what this does from a state management perspective
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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

func resourceInstancePoolCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)

	var instancePool model.InstancePool
	var instancePoolAwsAttributes model.InstancePoolAwsAttributes
	var instancePoolDiskSpec model.InstancePoolDiskSpec
	var instancePoolDiskSpecDiskType model.InstancePoolDiskType
	instancePool.InstancePoolName = d.Get("instance_pool_name").(string)
	instancePool.MinIdleInstances = int32(d.Get("min_idle_instances").(int))
	instancePool.MaxCapacity = int32(d.Get("max_capacity").(int))
	instancePool.IdleInstanceAutoTerminationMinutes = int32(d.Get("idle_instance_autotermination_minutes").(int))

	if awsAttributesSchema, ok := d.GetOk("aws_attributes"); ok {
		awsAttributesMap := getMapFromOneItemList(awsAttributesSchema)
		if availability, ok := awsAttributesMap["availability"]; ok {
			instancePoolAwsAttributes.Availability = model.AwsAvailability(availability.(string))
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
		diskSpecMap := diskSpec.(map[string]interface{})
		if ebsVolumeType, ok := diskSpecMap["ebs_volume_type"]; ok {
			instancePoolDiskSpecDiskType.EbsVolumeType = ebsVolumeType.(string)
		}
		if azureDiskVolumeType, ok := diskSpecMap["azure_disk_volume_type"]; ok {
			instancePoolDiskSpecDiskType.AzureDiskVolumeType = azureDiskVolumeType.(string)
		}
		instancePoolDiskSpec.DiskType = &instancePoolDiskSpecDiskType

		if diskCount, ok := diskSpecMap["disk_count"]; ok {
			intVal, err := strconv.Atoi(diskCount.(string))
			if err != nil {
				return err
			}
			instancePoolDiskSpec.DiskCount = int32(intVal)
		}
		if diskSize, ok := diskSpecMap["disk_size"]; ok {
			intVal, err := strconv.Atoi(diskSize.(string))
			if err != nil {
				return err
			}
			instancePoolDiskSpec.DiskSize = int32(intVal)
		}
		instancePool.DiskSpec = &instancePoolDiskSpec
	}

	if sparkVersions, ok := d.GetOk("preloaded_spark_versions"); ok {
		instancePool.PreloadedSparkVersions = sparkVersions.([]string)
	}

	instancePoolInfo, err := client.InstancePools().Create(instancePool)
	if err != nil {
		return err
	}
	d.SetId(instancePoolInfo.InstancePoolID)
	return resourceInstancePoolRead(d, m)
}

func resourceInstancePoolRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(service.DBApiClient)
	instancePoolInfo, err := client.InstancePools().Read(id)
	if err != nil {
		if isInstancePoolMissing(err.Error(), id) {
			log.Printf("Missing instance pool with id: %s.", id)
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
		diskSpec := map[string]interface{}{}
		if instancePoolInfo.DiskSpec.DiskType != nil {

			if instancePoolInfo.DiskSpec.DiskCount >= 0 {
				diskSpec["disk_count"] = strconv.FormatInt(int64(instancePoolInfo.DiskSpec.DiskCount), 10)
			}
			if instancePoolInfo.DiskSpec.DiskSize >= 0 {
				diskSpec["disk_size"] = strconv.FormatInt(int64(instancePoolInfo.DiskSpec.DiskSize), 10)
			}

		}
		if instancePoolInfo.DiskSpec.DiskType.EbsVolumeType != "" {
			diskSpec["ebs_volume_type"] = instancePoolInfo.DiskSpec.DiskType.EbsVolumeType
		}
		if instancePoolInfo.DiskSpec.DiskType.AzureDiskVolumeType != "" {
			diskSpec["azure_disk_volume_type"] = instancePoolInfo.DiskSpec.DiskType.AzureDiskVolumeType
		}
		err = d.Set("disk_spec", diskSpec)
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
	client := m.(service.DBApiClient)

	var instancePoolInfo model.InstancePoolInfo
	instancePoolInfo.InstancePoolName = d.Get("instance_pool_name").(string)
	instancePoolInfo.MinIdleInstances = int32(d.Get("min_idle_instances").(int))
	instancePoolInfo.MaxCapacity = int32(d.Get("max_capacity").(int))
	instancePoolInfo.IdleInstanceAutoTerminationMinutes = int32(d.Get("idle_instance_autotermination_minutes").(int))
	instancePoolInfo.InstancePoolID = id
	instancePoolInfo.NodeTypeID = d.Get("node_type_id").(string)

	err := client.InstancePools().Update(instancePoolInfo)
	if err != nil {
		return err
	}
	return resourceInstancePoolUpdate(d, m)
}

func resourceInstancePoolDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	id := d.Id()
	err := client.InstancePools().Delete(id)
	return err
}

func isInstancePoolMissing(errorMsg, resourceID string) bool {
	return strings.Contains(errorMsg, "RESOURCE_DOES_NOT_EXIST") &&
		strings.Contains(errorMsg, fmt.Sprintf("Can't find an instance pool with id: %s", resourceID))
}
