package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceVolumes() *schema.Resource {
	s := common.StructToSchema(catalog.CreateVolumeRequestContent{}, // This has to be done manually because catalog.UpdateVolumeRequestContent doesn't contain all the necessary fields
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var createVolumeRequestContent catalog.CreateVolumeRequestContent
			common.DataToStructPointer(d, s, &createVolumeRequestContent)
			v, err := w.Volumes.Create(ctx, createVolumeRequestContent)
			if err != nil {
				return err
			}
			d.SetId(v.FullName)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			// Full name = catalog_name . schema_name . name
			v, err := w.Volumes.ReadByFullNameArg(ctx, d.Id())
			if err != nil {
				return err
			}
			common.StructToData(v, s, d)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updateVolumeRequestContent catalog.UpdateVolumeRequestContent
			common.DataToStructPointer(d, s, &updateVolumeRequestContent)
			updateVolumeRequestContent.FullNameArg = d.Id() // to check
			v, err := w.Volumes.Update(ctx, updateVolumeRequestContent)
			_ = v
			if err != nil {
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Volumes.DeleteByFullNameArg(ctx, d.Id())
		},
	}.ToResource()
}
