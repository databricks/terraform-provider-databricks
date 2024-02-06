package pools

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getPool(poolsAPI InstancePoolsAPI, name string) (*InstancePoolAndStats, error) {
	poolList, err := poolsAPI.List()
	if err != nil {
		return nil, err
	}
	for _, pool := range poolList.InstancePools {
		if pool.InstancePoolName == name {
			return &pool, nil
		}
	}

	return nil, fmt.Errorf("instance pool '%s' doesn't exist", name)
}

// DataSourceInstancePool returns information about instance pool specified by name
func DataSourceInstancePool() common.Resource {
	type poolDetails struct {
		Name       string                `json:"name"`
		Attributes *InstancePoolAndStats `json:"pool_info,omitempty" tf:"computed"`
	}
	s := common.StructToSchema(poolDetails{}, nil)
	return common.Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			name := d.Get("name").(string)
			poolsAPI := NewInstancePoolsAPI(ctx, m)
			pool, err := getPool(poolsAPI, name)
			if err != nil {
				return err
			}
			d.SetId(pool.InstancePoolID)
			err = common.StructToData(poolDetails{Name: name, Attributes: pool}, s, d)
			return err
		},
	}
}
