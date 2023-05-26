package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type UpdateVolumeRequestContent struct {
	// The name of the catalog where the schema and the volume are
	CatalogName string `json:"catalog_name"`
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`
	// The name of the schema where the volume is
	SchemaName  string `json:"schema_name"`
	FullNameArg string `json:"-" url:"-"`
	// The name of the volume
	Name string `json:"name"`
	// The identifier of the user who owns the volume
	Owner string `json:"owner,omitempty" tf:"computed"`
	// The storage location on the cloud
	StorageLocation string             `json:"storage_location,omitempty"`
	VolumeType      catalog.VolumeType `json:"volume_type"`
}

func ResourceVolume() *schema.Resource {
	// We cannot use catalog.UpdateVolumeRequestContent because it doesn't contain all the necessary fields, example - SchemaName, CatalogName,
	// We also cannot use catalog.CreateVolumeRequestContent because it doesn't contain Owner and FullNameArg
	// We also need to do tf:"computed" for the Owner field hence cannot use a struct in Go SDK.
	s := common.StructToSchema(UpdateVolumeRequestContent{},
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

			// Update owner if it is provided
			if d.Get("owner") != "" {
				var updateVolumeRequestContent catalog.UpdateVolumeRequestContent
				common.DataToStructPointer(d, s, &updateVolumeRequestContent)
				updateVolumeRequestContent.FullNameArg = d.Id()
				_, err = w.Volumes.Update(ctx, updateVolumeRequestContent)
				if err != nil {
					return err
				}
			}

			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			v, err := w.Volumes.ReadByFullNameArg(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(v, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updateVolumeRequestContent catalog.UpdateVolumeRequestContent
			common.DataToStructPointer(d, s, &updateVolumeRequestContent)
			updateVolumeRequestContent.FullNameArg = d.Id()
			v, err := w.Volumes.Update(ctx, updateVolumeRequestContent)
			d.SetId(v.FullName)
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
