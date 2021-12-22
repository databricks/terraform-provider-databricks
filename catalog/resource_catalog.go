package catalog

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CatalogsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewCatalogsAPI(ctx context.Context, m interface{}) CatalogsAPI {
	return CatalogsAPI{m.(*common.DatabricksClient), ctx}
}

type CatalogInfo struct {
	Name        string            `json:"name" tf:"force_new"`
	Comment     string            `json:"comment,omitempty"`
	Properties  map[string]string `json:"properties,omitempty"`
	Owner       string            `json:"owner,omitempty" tf:"computed"`
	MetastoreID string            `json:"metastore_id,omitempty" tf:"computed"`
}

func (a CatalogsAPI) createCatalog(ci *CatalogInfo) error {
	return a.client.Post(a.context, "/unity-catalog/catalogs", ci, ci)
}

func (a CatalogsAPI) getCatalog(name string) (ci CatalogInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/catalogs/"+name, nil, &ci)
	return
}

func (a CatalogsAPI) updateCatalog(ci CatalogInfo) error {
	return a.client.Patch(a.context, "/unity-catalog/catalogs/"+ci.Name, ci)
}

func (a CatalogsAPI) deleteCatalog(name string) error {
	// TODO: force_destroy attribute
	return a.client.Delete(a.context, "/unity-catalog/catalogs/"+name+"?force=true", nil)
}

func ResourceCatalog() *schema.Resource {
	catalogSchema := common.StructToSchema(CatalogInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})
	return common.Resource{
		Schema: catalogSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ci CatalogInfo
			common.DataToStructPointer(d, catalogSchema, &ci)
			if err := NewCatalogsAPI(ctx, c).createCatalog(&ci); err != nil {
				return err
			}
			if err := NewSchemasAPI(ctx, c).deleteSchema(ci.Name + ".default"); err != nil {
				return fmt.Errorf("cannot remove new catalog default schema: %w", err)
			}
			d.SetId(ci.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ci, err := NewCatalogsAPI(ctx, c).getCatalog(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(ci, catalogSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ci CatalogInfo
			common.DataToStructPointer(d, catalogSchema, &ci)
			return NewCatalogsAPI(ctx, c).updateCatalog(ci)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewCatalogsAPI(ctx, c).deleteCatalog(d.Id())
		},
	}.ToResource()
}
