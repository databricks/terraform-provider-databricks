package catalog

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// This structure contains the fields of catalog.UpdateVolumeRequestContent and catalog.CreateVolumeRequestContent
// We need to create this because we need Owner, FullNameArg, SchemaName and CatalogName which aren't present in a single of them.
// We also need to annotate tf:"computed" for the Owner field.
type VolumeInfo struct {
	// The name of the catalog where the schema and the volume are
	CatalogName string `json:"catalog_name" tf:"force_new"`
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`
	// The name of the schema where the volume is
	SchemaName  string `json:"schema_name" tf:"force_new"`
	FullNameArg string `json:"-" url:"-"`
	// The name of the volume
	Name string `json:"name"`
	// The identifier of the user who owns the volume
	Owner string `json:"owner,omitempty" tf:"computed"`
	// The storage location on the cloud
	StorageLocation string             `json:"storage_location,omitempty" tf:"force_new"`
	VolumeType      catalog.VolumeType `json:"volume_type" tf:"force_new"`
}

func getNameFromId(id string) (string, error) {
	split := strings.Split(id, ".")
	if len(split) != 3 {
		return "", fmt.Errorf("invalid id <%s>: id should be in the format catalog.schema.volume", id)
	}
	return split[2], nil
}

func ResourceVolume() common.Resource {
	s := common.StructToSchema(VolumeInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			caseInsensitiveFields := []string{"name", "catalog_name", "schema_name"}
			for _, field := range caseInsensitiveFields {
				m[field].DiffSuppressFunc = common.EqualFoldDiffSuppress
			}
			m["storage_location"].DiffSuppressFunc = ucDirectoryPathSlashAndEmptySuppressDiff
			m["volume_path"] = &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			}
			m["volume_type"].ValidateDiagFunc = func(i interface{}, p cty.Path) diag.Diagnostics {
				s, ok := i.(string)
				if !ok {
					return diag.Errorf("expected string, got %s", reflect.TypeOf(i))
				}

				v := catalog.VolumeType("")
				err := v.Set(s)
				if err != nil {
					return diag.Errorf("invalid volume type %s: %s", s, err)
				}
				return nil
			}
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
			updateVolumeRequestContent.Name = d.Id()
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
			v, err := w.Volumes.ReadByName(ctx, d.Id())
			if err != nil {
				return err
			}
			err = common.StructToData(v, s, d)
			if err != nil {
				return err
			}
			return d.Set("volume_path", "/Volumes/"+strings.ReplaceAll(v.FullName, ".", "/"))
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updateVolumeRequestContent catalog.UpdateVolumeRequestContent
			common.DataToStructPointer(d, s, &updateVolumeRequestContent)
			updateVolumeRequestContent.Name = d.Id()
			userProvidedName := d.Get("name").(string)
			storedName, err := getNameFromId(d.Id())
			if err != nil {
				return err
			}
			if storedName != userProvidedName {
				updateVolumeRequestContent.NewName = userProvidedName
			}

			if d.HasChange("owner") {
				_, err := w.Volumes.Update(ctx, catalog.UpdateVolumeRequestContent{
					Name:  updateVolumeRequestContent.Name,
					Owner: updateVolumeRequestContent.Owner,
				})
				if err != nil {
					return err
				}
			}

			if !d.HasChangeExcept("owner") {
				return nil
			}

			if d.HasChange("comment") && updateVolumeRequestContent.Comment == "" {
				updateVolumeRequestContent.ForceSendFields = append(updateVolumeRequestContent.ForceSendFields, "Comment")
			}

			updateVolumeRequestContent.Owner = ""
			v, err := w.Volumes.Update(ctx, updateVolumeRequestContent)
			if err != nil {
				if d.HasChange("owner") {
					// Rollback
					old, new := d.GetChange("owner")
					_, rollbackErr := w.Volumes.Update(ctx, catalog.UpdateVolumeRequestContent{
						Name:  updateVolumeRequestContent.Name,
						Owner: old.(string),
					})
					if rollbackErr != nil {
						return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
					}
				}
				return err
			}

			// We need to update the resource Id because Name is updatable and FullName consists of Name,
			// So if we don't update the field then the requests would be made to old FullName which doesn't exists.
			d.SetId(v.FullName)
			d.Set("volume_path", "/Volumes/"+strings.ReplaceAll(v.FullName, ".", "/"))
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Volumes.DeleteByName(ctx, d.Id())
		},
	}
}
