package catalog

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SchemaInfo struct {
	Name        string            `json:"name" tf:"force_new"`
	CatalogName string            `json:"catalog_name" tf:"force_new"`
	StorageRoot string            `json:"storage_root,omitempty" tf:"force_new"`
	Comment     string            `json:"comment,omitempty"`
	Properties  map[string]string `json:"properties,omitempty"`
	Owner       string            `json:"owner,omitempty" tf:"computed"`
	MetastoreID string            `json:"metastore_id,omitempty" tf:"computed"`
	FullName    string            `json:"full_name,omitempty" tf:"computed"`
}

func ResourceSchema() *schema.Resource {
	s := common.StructToSchema(SchemaInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			delete(m, "full_name")
			m["force_destroy"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			}
			m["storage_root"].DiffSuppressFunc = ucDirectoryPathSlashOnlySuppressDiff
			return m
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var createSchemaRequest catalog.CreateSchema
			common.DataToStructPointer(d, s, &createSchemaRequest)
			schema, err := w.Schemas.Create(ctx, createSchemaRequest)
			if err != nil {
				return err
			}
			d.SetId(schema.FullName)

			// Don't update owner if it is not provided
			if d.Get("owner") == "" {
				return nil
			}

			var updateSchemaRequest catalog.UpdateSchema
			common.DataToStructPointer(d, s, &updateSchemaRequest)
			updateSchemaRequest.FullName = d.Id()
			_, err = w.Schemas.Update(ctx, updateSchemaRequest)
			if err != nil {
				return err
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			schema, err := w.Schemas.Get(ctx, catalog.GetSchemaRequest{FullName: d.Id()})
			if err != nil {
				return err
			}
			return common.StructToData(schema, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updateSchemaRequest catalog.UpdateSchema
			common.DataToStructPointer(d, s, &updateSchemaRequest)
			updateSchemaRequest.FullName = d.Id()
			schema, err := w.Schemas.Update(ctx, updateSchemaRequest)
			if err != nil {
				return err
			}
			// We need to update the resource Id because Name is updatable and FullName consists of Name,
			// So if we don't update the field then the requests would be made to old FullName which doesn't exists.
			d.SetId(schema.FullName)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_destroy").(bool)
			name := d.Id()
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			if force {
				tables, err := w.Tables.ListAll(ctx, catalog.ListTablesRequest{
					CatalogName: strings.Split(name, ".")[0],
					SchemaName:  strings.Split(name, ".")[1],
				})
				if err != nil {
					return err
				}
				for _, t := range tables {
					w.Tables.DeleteByFullName(ctx, t.FullName)
				}
			}
			return w.Schemas.DeleteByFullName(ctx, name)
		},
	}.ToResource()
}
