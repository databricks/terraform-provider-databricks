package catalog

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type MetastoresAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewMetastoresAPI(ctx context.Context, m interface{}) MetastoresAPI {
	return MetastoresAPI{m.(*common.DatabricksClient), ctx}
}

type MetastoreInfo struct {
	Name                      string  `json:"name"`
	StorageRoot               string  `json:"storage_root"`
	DefaultDataAccessConfigID string  `json:"default_data_access_config_id,omitempty"`
	Owner                     string  `json:"owner,omitempty" tf:"computed"`
	DataSharingEnabled        bool    `json:"data_sharing_enabled,omitempty"`
	MetastoreID               string  `json:"metastore_id,omitempty" tf:"computed"`
	WorkspaceIDs              []int64 `json:"workspace_ids,omitempty" tf:"computed"`
}

func (a MetastoresAPI) listMetastores() (mis []MetastoreInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/metastores", nil, &mis)
	return
}

func (a MetastoresAPI) createMetastore(mi *MetastoreInfo) error {
	return a.client.Post(a.context, "/unity-catalog/metastores", mi, mi)
}

func (a MetastoresAPI) getMetastore(id string) (mi MetastoreInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/metastores/"+id, nil, &mi)
	return
}

func (a MetastoresAPI) updateMetastore(mi MetastoreInfo) error {
	return a.client.Patch(a.context, "/unity-catalog/metastores/"+mi.MetastoreID, mi)
}

func (a MetastoresAPI) deleteMetastore(id string, force bool) error {
	return a.client.Delete(a.context, "/unity-catalog/metastores/"+id, map[string]interface{}{
		"force": force,
	})
}

func ResourceMetastore() *schema.Resource {
	s := common.StructToSchema(MetastoreInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			delete(m, "metastore_id")
			delete(m, "workspace_ids") // todo: bring it back when it works
			m["force_destroy"] = &schema.Schema{
				Type: schema.TypeBool,
				// Default:  false,
				Optional: true,
			}
			return m
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var mi MetastoreInfo
			if err := common.DataToStructPointer(d, s, &mi); err != nil {
				return err
			}
			if err := NewMetastoresAPI(ctx, c).createMetastore(&mi); err != nil {
				return err
			}
			d.SetId(mi.MetastoreID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			mi, err := NewMetastoresAPI(ctx, c).getMetastore(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(mi, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var mi MetastoreInfo
			if err := common.DataToStructPointer(d, s, &mi); err != nil {
				return err
			}
			mi.MetastoreID = d.Id()
			return NewMetastoresAPI(ctx, c).updateMetastore(mi)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_destroy").(bool)
			return NewMetastoresAPI(ctx, c).deleteMetastore(d.Id(), force)
		},
	}.ToResource()
}
