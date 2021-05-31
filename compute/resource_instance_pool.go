package compute

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// NewInstancePoolsAPI creates InstancePoolsAPI instance from provider meta
func NewInstancePoolsAPI(ctx context.Context, m interface{}) InstancePoolsAPI {
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
func ResourceInstancePool() *schema.Resource {
	s := common.StructToSchema(InstancePool{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["aws_attributes"].ForceNew = true
		s["node_type_id"].ForceNew = true
		s["custom_tags"].ForceNew = true
		s["preloaded_spark_versions"].ForceNew = true
		s["preloaded_docker_image"].ForceNew = true
		s["azure_attributes"].ForceNew = true
		s["disk_spec"].ForceNew = true
		s["enable_elastic_disk"].ForceNew = true
		s["enable_elastic_disk"].Default = true
		s["aws_attributes"].ConflictsWith = []string{"azure_attributes"}
		s["azure_attributes"].ConflictsWith = []string{"aws_attributes"}
		s["aws_attributes"].DiffSuppressFunc = common.MakeEmptyBlockSuppressFunc("aws_attributes.#")
		s["azure_attributes"].DiffSuppressFunc = common.MakeEmptyBlockSuppressFunc("azure_attributes.#")
		if v, err := common.SchemaPath(s, "aws_attributes", "availability"); err == nil {
			v.ForceNew = true
			v.Default = AwsAvailabilitySpot
			v.ValidateFunc = validation.StringInSlice([]string{
				AwsAvailabilityOnDemand,
				AwsAvailabilitySpot,
			}, false)
		}
		if v, err := common.SchemaPath(s, "aws_attributes", "zone_id"); err == nil {
			v.ForceNew = true
		}
		if v, err := common.SchemaPath(s, "aws_attributes", "spot_bid_price_percent"); err == nil {
			v.ForceNew = true
			v.Default = 100
		}
		if v, err := common.SchemaPath(s, "azure_attributes", "availability"); err == nil {
			v.ForceNew = true
			v.Default = AzureAvailabilityOnDemand
			v.ValidateFunc = validation.StringInSlice([]string{
				AzureAvailabilitySpot,
				AzureAvailabilityOnDemand,
			}, false)
		}
		if v, err := common.SchemaPath(s, "azure_attributes", "spot_bid_max_price"); err == nil {
			v.ForceNew = true
		}
		if v, err := common.SchemaPath(s, "disk_spec", "disk_type", "azure_disk_volume_type"); err == nil {
			v.ForceNew = true
			// nolint
			v.ValidateFunc = validation.StringInSlice([]string{
				AzureDiskVolumeTypePremium,
				AzureDiskVolumeTypeStandard,
			}, false)
		}
		if v, err := common.SchemaPath(s, "disk_spec", "disk_type", "ebs_volume_type"); err == nil {
			v.ForceNew = true
			// nolint
			v.ValidateFunc = validation.StringInSlice([]string{
				EbsVolumeTypeGeneralPurposeSsd,
				EbsVolumeTypeThroughputOptimizedHdd,
			}, false)
		}
		if v, err := common.SchemaPath(s, "preloaded_docker_image", "url"); err == nil {
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
			if err := common.DataToStructPointer(d, s, &ip); err != nil {
				return err
			}
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
			if err := common.DataToStructPointer(d, s, &ip); err != nil {
				return err
			}
			ip.InstancePoolID = d.Id()
			return NewInstancePoolsAPI(ctx, c).Update(ip)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewInstancePoolsAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
