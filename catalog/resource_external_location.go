package catalog

import (
	"context"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// This structure contains the fields of both catalog.UpdateExternalLocation and catalog.CreateExternalLocation
type ExternalLocationInfo struct {
	Name           string                     `json:"name" tf:"force_new"`
	URL            string                     `json:"url"`
	CredentialName string                     `json:"credential_name"`
	Comment        string                     `json:"comment,omitempty"`
	SkipValidation bool                       `json:"skip_validation,omitempty"`
	Owner          string                     `json:"owner,omitempty" tf:"computed"`
	MetastoreID    string                     `json:"metastore_id,omitempty" tf:"computed"`
	ReadOnly       bool                       `json:"read_only,omitempty"`
	AccessPoint    string                     `json:"access_point,omitempty"`
	EncDetails     *catalog.EncryptionDetails `json:"encryption_details,omitempty"`
}

func ResourceExternalLocation() *schema.Resource {
	s := common.StructToSchema(ExternalLocationInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["force_destroy"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["force_update"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["skip_validation"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				return old == "false" && new == "true"
			}
			m["url"].DiffSuppressFunc = ucDirectoryPathSlashOnlySuppressDiff
			return m
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var createExternalLocationRequest catalog.CreateExternalLocation
			common.DataToStructPointer(d, s, &createExternalLocationRequest)
			el, err := w.ExternalLocations.Create(ctx, createExternalLocationRequest)
			if err != nil {
				return err
			}
			d.SetId(el.Name)

			// Don't update owner if it is not provided
			if d.Get("owner") == "" {
				return nil
			}

			var updateExternalLocationRequest catalog.UpdateExternalLocation
			common.DataToStructPointer(d, s, &updateExternalLocationRequest)
			_, err = w.ExternalLocations.Update(ctx, updateExternalLocationRequest)
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
			el, err := w.ExternalLocations.GetByName(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(el, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_update").(bool)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updateExternalLocationRequest catalog.UpdateExternalLocation
			common.DataToStructPointer(d, s, &updateExternalLocationRequest)
			updateExternalLocationRequest.Force = force

			ownerInRequest := updateExternalLocationRequest.Owner
			updateExternalLocationRequest.Owner = ""
			_, err = w.ExternalLocations.Update(ctx, updateExternalLocationRequest)
			if err != nil {
				return err
			}
			if ownerInRequest != "" && d.HasChange("owner") {
				var updateOwnerExternalLocationRequest catalog.UpdateExternalLocation
				updateOwnerExternalLocationRequest.Owner = ownerInRequest
				updateOwnerExternalLocationRequest.Name = updateExternalLocationRequest.Name
				updateOwnerExternalLocationRequest.Force = updateExternalLocationRequest.Force
				_, err = w.ExternalLocations.Update(ctx, updateExternalLocationRequest)
				if err != nil {
					// roll back the previous changes
					var updateRollbackExternalLocationRequest catalog.UpdateExternalLocation
					values := reflect.ValueOf(updateRollbackExternalLocationRequest)
					for i := 0; i < values.NumField(); i++ {
						field := values.Type().Field(i)
						jsonFull := field.Tag.Get("json")
						jsonName := strings.Split(jsonFull, ",")[0]
						if jsonName == "owner" {
							continue
						}
						if d.HasChange(jsonName) {
							old, _ := d.GetChange(jsonName)
							fieldValue := values.Elem().FieldByName(field.Type.String())
							fieldValue.Set(reflect.ValueOf(old))
						}
					}
					return err
				}
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_destroy").(bool)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.ExternalLocations.Delete(ctx, catalog.DeleteExternalLocationRequest{
				Name:  d.Id(),
				Force: force,
			})
		},
	}.ToResource()
}
