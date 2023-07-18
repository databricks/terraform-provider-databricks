package catalog

import (
	"context"
	"log"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type MetastoresAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewMetastoresAPI(ctx context.Context, m any) MetastoresAPI {
	return MetastoresAPI{m.(*common.DatabricksClient), context.WithValue(ctx, common.Api, common.API_2_1)}
}

type MetastoreInfo struct {
	Name                                        string  `json:"name"`
	StorageRoot                                 string  `json:"storage_root" tf:"force_new"`
	DefaultDacID                                string  `json:"default_data_access_config_id,omitempty" tf:"suppress_diff"`
	Owner                                       string  `json:"owner,omitempty" tf:"computed"`
	MetastoreID                                 string  `json:"metastore_id,omitempty" tf:"computed"`
	WorkspaceIDs                                []int64 `json:"workspace_ids,omitempty" tf:"computed"`
	Region                                      string  `json:"region,omitempty" tf:"computed"`
	Cloud                                       string  `json:"cloud,omitempty" tf:"computed"`
	GlobalMetastoreId                           string  `json:"global_metastore_id,omitempty" tf:"computed"`
	CreatedAt                                   int64   `json:"created_at,omitempty" tf:"computed"`
	CreatedBy                                   string  `json:"created_by,omitempty" tf:"computed"`
	UpdatedAt                                   int64   `json:"updated_at,omitempty" tf:"computed"`
	UpdatedBy                                   string  `json:"updated_by,omitempty" tf:"computed"`
	DeltaSharingScope                           string  `json:"delta_sharing_scope,omitempty" tf:"suppress_diff"`
	DeltaSharingRecipientTokenLifetimeInSeconds int64   `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	DeltaSharingOrganizationName                string  `json:"delta_sharing_organization_name,omitempty"`
}

type CreateMetastore struct {
	Name        string `json:"name"`
	StorageRoot string `json:"storage_root"`
}

func (a MetastoresAPI) createMetastore(cm CreateMetastore) (mi MetastoreInfo, err error) {
	err = a.client.Post(a.context, "/unity-catalog/metastores", cm, &mi)
	return
}

func (a MetastoresAPI) getMetastore(id string) (mi MetastoreInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/metastores/"+id, nil, &mi)
	return
}

func (a MetastoresAPI) updateMetastore(metastoreID string, update map[string]any) error {
	return a.client.Patch(a.context, "/unity-catalog/metastores/"+metastoreID, update)
}

func (a MetastoresAPI) deleteMetastore(id string, force bool) error {
	return a.client.Delete(a.context, "/unity-catalog/metastores/"+id, map[string]any{
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
			m["delta_sharing_scope"].RequiredWith = []string{"delta_sharing_recipient_token_lifetime_in_seconds"}
			m["delta_sharing_scope"].ValidateFunc = validation.StringInSlice([]string{"INTERNAL", "INTERNAL_AND_EXTERNAL"}, false)
			m["delta_sharing_recipient_token_lifetime_in_seconds"].RequiredWith = []string{"delta_sharing_scope"}
			m["storage_root"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				if strings.HasPrefix(old, new) {
					log.Printf("[DEBUG] Ignoring configuration drift from %s to %s", old, new)
					return true
				}
				return false
			}
			return m
		})
	update := updateFunctionFactory("/unity-catalog/metastores", []string{"owner", "name", "delta_sharing_scope",
		"delta_sharing_recipient_token_lifetime_in_seconds", "delta_sharing_organization_name"})

	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var create CreateMetastore
			common.DataToStructPointer(d, s, &create)
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
