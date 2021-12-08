package catalog

import (
	"context"
	"log"
	"strings"

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
	Name         string  `json:"name"`
	StorageRoot  string  `json:"storage_root" tf:"force_new"`
	DefaultDacID string  `json:"default_data_access_config_id,omitempty"`
	Owner        string  `json:"owner,omitempty" tf:"computed"`
	MetastoreID  string  `json:"metastore_id,omitempty" tf:"computed"`
	WorkspaceIDs []int64 `json:"workspace_ids,omitempty" tf:"computed"`
}

type CreateMetastore struct {
	Name        string `json:"name"`
	StorageRoot string `json:"storage_root"`
}

// func (a MetastoresAPI) listMetastores() (mis []MetastoreInfo, err error) {
// 	err = a.client.Get(a.context, "/unity-catalog/metastores", nil, &mis)
// 	return
// }

func (a MetastoresAPI) createMetastore(cm CreateMetastore) (mi MetastoreInfo, err error) {
	err = a.client.Post(a.context, "/unity-catalog/metastores", cm, &mi)
	return
}

func (a MetastoresAPI) getMetastore(id string) (mi MetastoreInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/metastores/"+id, nil, &mi)
	return
}

func (a MetastoresAPI) updateMetastore(metastoreID string, update map[string]interface{}) error {
	return a.client.Patch(a.context, "/unity-catalog/metastores/"+metastoreID, update)
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
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["storage_root"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				if strings.HasPrefix(old, new) {
					log.Printf("[DEBUG] Ignoring configuration drift from %s to %s", old, new)
					return true
				}
				return false
			}
			return m
		})
	update := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		// other fields to come later
		updatable := []string{"owner", "name"}
		patch := map[string]interface{}{}
		for _, field := range updatable {
			old, new := d.GetChange(field)
			if old == new {
				continue
			}
			if field == "name" && old == "" {
				continue
			}
			patch[field] = new
		}
		if len(patch) == 0 {
			return nil
		}
		return NewMetastoresAPI(ctx, c).updateMetastore(d.Id(), patch)
	}
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var create CreateMetastore
			if err := common.DataToStructPointer(d, s, &create); err != nil {
				return err
			}
			mi, err := NewMetastoresAPI(ctx, c).createMetastore(create)
			if err != nil {
				return err
			}
			d.SetId(mi.MetastoreID)
			return update(ctx, d, c)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			mi, err := NewMetastoresAPI(ctx, c).getMetastore(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(mi, s, d)
		},
		Update: update,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_destroy").(bool)
			return NewMetastoresAPI(ctx, c).deleteMetastore(d.Id(), force)
		},
	}.ToResource()
}
