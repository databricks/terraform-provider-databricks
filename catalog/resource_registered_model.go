package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceRegisteredModel() common.Resource {
	s := common.StructToSchema(
		catalog.CreateRegisteredModelRequest{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			caseInsensitiveFields := []string{"name", "catalog_name", "schema_name"}
			for _, field := range caseInsensitiveFields {
				m[field].DiffSuppressFunc = common.EqualFoldDiffSuppress
			}
			m["name"].ForceNew = true
			m["catalog_name"].ForceNew = true
			m["schema_name"].ForceNew = true
			m["storage_location"].ForceNew = true
			m["storage_location"].Computed = true
			m["owner"] = &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			}

			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var m catalog.CreateRegisteredModelRequest
			common.DataToStructPointer(d, s, &m)
			model, err := w.RegisteredModels.Create(ctx, m)
			if err != nil {
				return err
			}
			d.SetId(model.FullName)
			// Don't update owner if it is not provided
			if d.Get("owner") == "" {
				return nil
			}

			var update catalog.UpdateRegisteredModelRequest
			common.DataToStructPointer(d, s, &update)
			update.FullName = d.Id()
			_, err = w.RegisteredModels.Update(ctx, update)
			if err != nil {
				return err
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			model, err := w.RegisteredModels.GetByFullName(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(*model, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}

			var u catalog.UpdateRegisteredModelRequest
			common.DataToStructPointer(d, s, &u)
			u.FullName = d.Id()

			if d.HasChange("owner") {
				_, err := w.RegisteredModels.Update(ctx, catalog.UpdateRegisteredModelRequest{
					FullName: u.FullName,
					Owner:    u.Owner,
				})
				if err != nil {
					return err
				}
			}

			if !d.HasChangeExcept("owner") {
				return nil
			}

			if d.HasChange("comment") && u.Comment == "" {
				u.ForceSendFields = append(u.ForceSendFields, "Comment")
			}
			u.Owner = ""
			_, err = w.RegisteredModels.Update(ctx, u)
			if err != nil {
				if d.HasChange("owner") {
					// Rollback
					old, new := d.GetChange("owner")
					_, rollbackErr := w.RegisteredModels.Update(ctx, catalog.UpdateRegisteredModelRequest{
						FullName: u.FullName,
						Owner:    old.(string),
					})
					if rollbackErr != nil {
						return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
					}
				}
				return err
			}
			_, err = w.RegisteredModels.Update(ctx, u)
			return err
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return w.RegisteredModels.DeleteByFullName(ctx, d.Id())
		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
	}
}
