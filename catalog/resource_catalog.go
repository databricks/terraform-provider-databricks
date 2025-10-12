package catalog

import (
	"context"
	"fmt"
	"log"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/catalog/bindings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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

func ResourceCatalog() common.Resource {
	catalogSchema := common.StructToSchema(catalog.CatalogInfo{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			s["force_destroy"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			}
			// mark all computed values
			for _, v := range []string{"owner", "isolation_mode", "metastore_id", "enable_predictive_optimization"} {
				common.CustomizeSchemaPath(s, v).SetOptional().SetComputed()
			}
			// case sensitive suppress diff
			for _, v := range []string{"name", "connection_name", "share_name", "provider_name"} {
				common.CustomizeSchemaPath(s, v).SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
			}
			// can only have one of provider_name + share_name, connection_name
			common.CustomizeSchemaPath(s, "connection_name").SetConflictsWith([]string{"provider_name", "share_name"}).SetForceNew()
			for _, v := range []string{"provider_name", "share_name"} {
				common.CustomizeSchemaPath(s, v).SetConflictsWith([]string{"connection_name", "storage_root"}).SetForceNew()
			}
			common.CustomizeSchemaPath(s, "storage_root").SetCustomSuppressDiff(ucDirectoryPathSlashOnlySuppressDiff).SetForceNew()
			common.CustomizeSchemaPath(s, "enable_predictive_optimization").SetValidateFunc(
				validation.StringInSlice([]string{"DISABLE", "ENABLE", "INHERIT"}, false),
			)
			for _, v := range []string{"catalog_type", "created_at", "created_by",
				"updated_at", "updated_by", "securable_type", "full_name", "storage_location"} {
				common.CustomizeSchemaPath(s, v).SetReadOnly()
			}
			common.CustomizeSchemaPath(s, "effective_predictive_optimization_flag").SetComputed().SetSuppressDiff()
			return s
		})
	return common.Resource{
		Schema: catalogSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}

			err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
			if err != nil {
				return err
			}

			var createCatalogRequest catalog.CreateCatalog
			common.DataToStructPointer(d, catalogSchema, &createCatalogRequest)
			ci, err := w.Catalogs.Create(ctx, createCatalogRequest)
			if err != nil {
				return err
			}
			// only remove catalog default schema for standard catalog (e.g. non-Delta Sharing, non-foreign)
			if ci.ShareName == "" && ci.ConnectionName == "" {
				if err := w.Schemas.DeleteByFullName(ctx, ci.Name+".default"); err != nil {
					return fmt.Errorf("cannot remove new catalog default schema: %w", err)
				}
			}

			d.SetId(ci.Name)

			// Update owner, isolation mode or predictive optimization if it is provided
			if !updateRequired(d, []string{"owner", "isolation_mode", "enable_predictive_optimization"}) {
				return nil
			}

			var updateCatalogRequest catalog.UpdateCatalog
			common.DataToStructPointer(d, catalogSchema, &updateCatalogRequest)
			updateCatalogRequest.Name = d.Id()
			// Options must be set in the create request only (aside from HMS-backed catalogs).
			updateCatalogRequest.Options = nil
			_, err = w.Catalogs.Update(ctx, updateCatalogRequest)
			if err != nil {
				return err
			}

			// Bind the current workspace if the catalog is isolated, otherwise the read will fail
			return bindings.AddCurrentWorkspaceBindings(ctx, d, w, ci.Name, bindings.BindingsSecurableTypeCatalog)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
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
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}

			err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
			if err != nil {
				return err
			}
			if !d.HasChangeExcept("force_destroy") {
				return nil
			}

			var updateCatalogRequest catalog.UpdateCatalog
			common.DataToStructPointer(d, catalogSchema, &updateCatalogRequest)
			updateCatalogRequest.Name = d.Id()

			if d.HasChange("owner") {
				_, err = w.Catalogs.Update(ctx, catalog.UpdateCatalog{
					Name:  updateCatalogRequest.Name,
					Owner: updateCatalogRequest.Owner,
				})
				if err != nil {
					return err
				}
			}

			if !d.HasChangeExcept("owner") {
				return nil
			}

			if d.HasChange("comment") && updateCatalogRequest.Comment == "" {
				updateCatalogRequest.ForceSendFields = append(updateCatalogRequest.ForceSendFields, "Comment")
			}

			updateCatalogRequest.Owner = ""
			// The only option allowed in update is "authorized_paths". All other options must be removed.
			if opts := updateCatalogRequest.Options; opts != nil {
				if v, ok := opts["authorized_paths"]; ok {
					updateCatalogRequest.Options = map[string]string{"authorized_paths": v}
				} else {
					updateCatalogRequest.Options = nil
				}
			}
			ci, err := w.Catalogs.Update(ctx, updateCatalogRequest)

			if err != nil {
				if d.HasChange("owner") {
					// Rollback
					old, new := d.GetChange("owner")
					_, rollbackErr := w.Catalogs.Update(ctx, catalog.UpdateCatalog{
						Name:  updateCatalogRequest.Name,
						Owner: old.(string),
					})
					if rollbackErr != nil {
						return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
					}
				}
				return err
			}

			// We need to update the resource data because Name is updatable
			// So if we don't update the field then the requests would be made to old Name which doesn't exists.
			d.SetId(ci.Name)

			// Bind the current workspace if the catalog is isolated, otherwise the read will fail
			return bindings.AddCurrentWorkspaceBindings(ctx, d, w, ci.Name, bindings.BindingsSecurableTypeCatalog)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}

			err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
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
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			// The only scenario in which we can update options is for the `authorized_paths` key. Any
			// other changes to the options field will result in an error.
			if d.HasChange("options") {
				old, new := d.GetChange("options")
				oldMap := old.(map[string]interface{})
				newMap := new.(map[string]interface{})
				delete(oldMap, "authorized_paths")
				delete(newMap, "authorized_paths")
				// If any attribute other than `authorized_paths` is removed, the resource should be recreated.
				for k := range oldMap {
					if _, ok := newMap[k]; !ok {
						if err := d.ForceNew("options"); err != nil {
							return err
						}
					}
				}
				// If any attribute other than `authorized_paths` is added or changed, the resource should be recreated.
				for k, v := range newMap {
					if oldV, ok := oldMap[k]; !ok || oldV != v {
						if err := d.ForceNew("options"); err != nil {
							return err
						}
					}
				}
			}
			return nil
		},
	}
}
