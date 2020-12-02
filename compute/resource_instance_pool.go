package compute

import (
	"context"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/util"
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
func (a InstancePoolsAPI) Update(instancePoolInfo InstancePool) error {
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

// ResourceInstancePool ...
func ResourceInstancePool() *schema.Resource {
	s := internal.StructToSchema(InstancePool{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["aws_attributes"].ForceNew = true
		s["node_type_id"].ForceNew = true
		s["custom_tags"].ForceNew = true
		s["enable_elastic_disk"].ForceNew = true
		s["enable_elastic_disk"].Default = true
		// TODO: check if it's really force new...
		if v, err := internal.SchemaPath(s, "aws_attributes", "availability"); err != nil {
			v.ForceNew = true
		}
		if v, err := internal.SchemaPath(s, "aws_attributes", "zone_id"); err != nil {
			v.ForceNew = true
		}
		if v, err := internal.SchemaPath(s, "aws_attributes", "spot_bid_price_percent"); err != nil {
			v.ForceNew = true
		}
		if v, err := internal.SchemaPath(s, "disk_spec", "disk_type", "azure_disk_volume_type"); err != nil {
			v.ForceNew = true
			// nolint
			v.ValidateFunc = validation.StringInSlice([]string{
				AzureDiskVolumeTypePremium,
				AzureDiskVolumeTypeStandard,
			}, false)
		}
		if v, err := internal.SchemaPath(s, "disk_spec", "disk_type", "ebs_volume_type"); err != nil {
			v.ForceNew = true
			// nolint
			v.ValidateFunc = validation.StringInSlice([]string{
				EbsVolumeTypeGeneralPurposeSsd,
				EbsVolumeTypeThroughputOptimizedHdd,
			}, false)
		}
		return s
	})
	return util.CommonResource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ip InstancePool
			if err := internal.DataToStructPointer(d, s, &ip); err != nil {
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
			return internal.StructToData(ip, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ip InstancePool
			if err := internal.DataToStructPointer(d, s, &ip); err != nil {
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
