package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// This structure contains the fields of catalog.UpdateVolumeRequestContent and catalog.CreateVolumeRequestContent
// We need to create this because we need Owner, FullNameArg, SchemaName and CatalogName which aren't present in a single of them.
// We also need to annotate tf:"computed" for the Owner field.
type VolumeInfo struct {
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
	s := common.StructToSchema(VolumeInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["storage_location"].DiffSuppressFunc = ucDirectoryPathSlashAndEmptySuppressDiff
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

			// Don't update owner if it is not provided
			if d.Get("owner") == "" {
				return nil
			}

			var updateVolumeRequestContent catalog.UpdateVolumeRequestContent
			common.DataToStructPointer(d, s, &updateVolumeRequestContent)
			updateVolumeRequestContent.FullNameArg = d.Id()
			_, err = w.Volumes.Update(ctx, updateVolumeRequestContent)
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
			if err != nil {
				return err
			}
			// We need to update the resource Id because Name is updatable and FullName consists of Name,
			// So if we don't update the field then the requests would be made to old FullName which doesn't exists.
			d.SetId(v.FullName)
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
