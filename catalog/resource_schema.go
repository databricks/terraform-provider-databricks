package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceSchema() common.Resource {
	s := common.StructToSchema(catalog.SchemaInfo{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			s["force_destroy"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			}

			// Mark fields as force_new
			for _, v := range []string{"name", "catalog_name", "storage_root"} {
				common.CustomizeSchemaPath(s, v).SetForceNew()
			}

			// Mark optional and computed fields (user can set, but has default from backend)
			for _, v := range []string{"metastore_id", "enable_predictive_optimization", "owner"} {
				common.CustomizeSchemaPath(s, v).SetComputed()
			}

			// Set read-only fields (returned by backend, not settable by user)
			// Note: SetReadOnly() implies SetComputed(), so no need to call both
			for _, v := range []string{"schema_id", "browse_only", "catalog_type",
				"created_at", "created_by", "storage_location", "updated_at", "updated_by",
				"effective_predictive_optimization_flag", "full_name"} {
				common.CustomizeSchemaPath(s, v).SetReadOnly()
			}

			common.CustomizeSchemaPath(s, "storage_root").SetCustomSuppressDiff(ucDirectoryPathSlashOnlySuppressDiff)
			s["storage_root"].DiffSuppressFunc = ucDirectoryPathSlashOnlySuppressDiff
			common.CustomizeSchemaPath(s, "name").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
			common.CustomizeSchemaPath(s, "catalog_name").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
			common.CustomizeSchemaPath(s, "enable_predictive_optimization").SetValidateFunc(
				validation.StringInSlice([]string{"DISABLE", "ENABLE", "INHERIT"}, false),
			)

			return s
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
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

			// Update owner or predictive optimization if it is provided
			updateRequired := false
			for _, key := range []string{"owner", "enable_predictive_optimization"} {
				if d.Get(key) != "" {
					updateRequired = true
					break
				}
			}
			if !updateRequired {
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
			schema, err := w.Schemas.GetByFullName(ctx, d.Id())
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
			err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
			if err != nil {
				return err
			}
			if !d.HasChangeExcept("force_destroy") {
				return nil
			}
			var updateSchemaRequest catalog.UpdateSchema
			common.DataToStructPointer(d, s, &updateSchemaRequest)
			updateSchemaRequest.FullName = d.Id()

			if d.HasChange("owner") {
				_, err := w.Schemas.Update(ctx, catalog.UpdateSchema{
					FullName: updateSchemaRequest.FullName,
					Owner:    updateSchemaRequest.Owner,
				})
				if err != nil {
					return err
				}
			}

			if !d.HasChangeExcept("owner") {
				return nil
			}

			if d.HasChange("comment") && updateSchemaRequest.Comment == "" {
				updateSchemaRequest.ForceSendFields = append(updateSchemaRequest.ForceSendFields, "Comment")
			}

			updateSchemaRequest.Owner = ""
			schema, err := w.Schemas.Update(ctx, updateSchemaRequest)
			if err != nil {
				if d.HasChange("owner") {
					// Rollback
					old, new := d.GetChange("owner")
					_, rollbackErr := w.Schemas.Update(ctx, catalog.UpdateSchema{
						FullName: updateSchemaRequest.FullName,
						Owner:    old.(string),
					})
					if rollbackErr != nil {
						return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
					}
				}
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
			err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
			if err != nil {
				return err
			}
			return w.Schemas.Delete(ctx, catalog.DeleteSchemaRequest{FullName: name, Force: force})
		},
	}
}
