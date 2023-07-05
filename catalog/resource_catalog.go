package catalog

import (
	"context"
	"fmt"
	"log"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ucDirectoryPathSlashOnlySuppressDiff(k, old, new string, d *schema.ResourceData) bool {
	if (new == (old + "/")) || (old == (new + "/")) {
		log.Printf("[DEBUG] Ignoring configuration drift from %s to %s", old, new)
		return true
	}
	return false
}

func ucDirectoryPathSlashAndEmptySuppressDiff(k, old, new string, d *schema.ResourceData) bool {
	if (new == (old + "/")) || (old == (new + "/")) || (new == "" && old != "") {
		log.Printf("[DEBUG] Ignoring configuration drift from %s to %s", old, new)
		return true
	}
	return false
}

type CatalogInfo struct {
	Name          string            `json:"name"`
	Comment       string            `json:"comment,omitempty"`
	StorageRoot   string            `json:"storage_root,omitempty" tf:"force_new"`
	ProviderName  string            `json:"provider_name,omitempty" tf:"force_new,conflicts:storage_root"`
	ShareName     string            `json:"share_name,omitempty" tf:"force_new,conflicts:storage_root"`
	Properties    map[string]string `json:"properties,omitempty"`
	Owner         string            `json:"owner,omitempty" tf:"computed"`
	IsolationMode string            `json:"isolation_mode,omitempty" tf:"computed"`
	MetastoreID   string            `json:"metastore_id,omitempty" tf:"computed"`
}

func ResourceCatalog() *schema.Resource {
	catalogSchema := common.StructToSchema(CatalogInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["force_destroy"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			}
			m["storage_root"].DiffSuppressFunc = ucDirectoryPathSlashOnlySuppressDiff
			return m
		})
	return common.Resource{
		Schema: catalogSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			var createCatalogRequest catalog.CreateCatalog
			common.DataToStructPointer(d, catalogSchema, &createCatalogRequest)
			ci, err := w.Catalogs.Create(ctx, createCatalogRequest)
			if err != nil {
				return err
			}
			// only remove catalog default schema for non-Delta Sharing catalog
			if ci.ShareName == "" {
				if err := w.Schemas.DeleteByFullName(ctx, ci.Name+".default"); err != nil {
					return fmt.Errorf("cannot remove new catalog default schema: %w", err)
				}
			}

			d.SetId(ci.Name)

			// Update owner or isolation mode if it is provided
			if d.Get("owner") == "" && d.Get("isolation_mode") == "" {
				return nil
			}
			var updateCatalogRequest catalog.UpdateCatalog
			common.DataToStructPointer(d, catalogSchema, &updateCatalogRequest)
			_, err = w.Catalogs.Update(ctx, updateCatalogRequest)
			if err != nil {
				return err
			}

			if d.Get("isolation_mode") != "ISOLATED" {
				return nil
			}
			// Bind the current workspace if the catalog is isolated, otherwise the read will fail
			currentMetastoreAssignment, err := w.Metastores.Current(ctx)
			if err != nil {
				return err
			}
			_, err = w.WorkspaceBindings.Update(ctx, catalog.UpdateWorkspaceBindings{
				Name:             ci.Name,
				AssignWorkspaces: []int64{currentMetastoreAssignment.WorkspaceId},
			})
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

			ci, err := w.Catalogs.GetByName(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(ci, catalogSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updateCatalogRequest catalog.UpdateCatalog
			common.DataToStructPointer(d, catalogSchema, &updateCatalogRequest)
			ci, err := w.Catalogs.Update(ctx, updateCatalogRequest)
			if err != nil {
				return err
			}

			// We need to update the resource data because Name is updatable
			// So if we don't update the field then the requests would be made to old Name which doesn't exists.
			d.SetId(ci.Name)

			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			force := d.Get("force_destroy").(bool)
			// If the workspace has isolation mode ISOLATED, we need to add the current workspace to its
			// bindings before deleting.
			if d.Get("isolation_mode").(string) == "ISOLATED" {
				currentMetastoreAssignment, err := w.Metastores.Current(ctx)
				if err != nil {
					return err
				}
				_, err = w.WorkspaceBindings.Update(ctx, catalog.UpdateWorkspaceBindings{
					Name:             d.Id(),
					AssignWorkspaces: []int64{currentMetastoreAssignment.WorkspaceId},
				})
				if err != nil {
					return err
				}
			}
			return w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{Force: force, Name: d.Id()})
		},
	}.ToResource()
}
