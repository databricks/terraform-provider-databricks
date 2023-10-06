package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceRegisteredModel() *schema.Resource {
	s := common.StructToSchema(
		catalog.CreateRegisteredModelRequest{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["name"].ForceNew = true
			m["catalog_name"].ForceNew = true
			m["schema_name"].ForceNew = true
			m["storage_location"].ForceNew = true
			m["storage_location"].Computed = true

			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
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
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			model, err := w.RegisteredModels.GetByFullName(ctx, d.Id())
			if err != nil {
				return err
			}
			err = common.StructToData(*model, s, d)
			if err != nil {
				return err
			}
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var m catalog.CreateRegisteredModelRequest
			var u catalog.UpdateRegisteredModelRequest
			common.DataToStructPointer(d, s, &m)
			u.FullName = d.Id()
			u.Comment = m.Comment
			_, err = w.RegisteredModels.Update(ctx, u)
			return err
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.RegisteredModels.DeleteByFullName(ctx, d.Id())
		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
	}.ToResource()
}
