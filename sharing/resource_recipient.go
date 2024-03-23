package sharing

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceRecipient() common.Resource {
	recipientSchema := common.StructToSchema(sharing.RecipientInfo{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "authentication_type").SetForceNew().SetRequired().SetValidateFunc(validation.StringInSlice([]string{"TOKEN", "DATABRICKS"}, false))
		common.CustomizeSchemaPath(s, "sharing_code").SetSuppressDiff().SetForceNew().SetSensitive()
		common.CustomizeSchemaPath(s, "authentication_type").SetForceNew()
		common.CustomizeSchemaPath(s, "name").SetForceNew().SetRequired()
		common.CustomizeSchemaPath(s, "owner").SetSuppressDiff()
		common.CustomizeSchemaPath(s, "properties_kvpairs").SetSuppressDiff()
		common.CustomizeSchemaPath(s, "properties_kvpairs", "properties").SetSuppressDiff()
		common.CustomizeSchemaPath(s, "data_recipient_global_metastore_id").SetForceNew().SetConflictsWith([]string{"ip_access_list"})
		common.CustomizeSchemaPath(s, "ip_access_list").SetConflictsWith([]string{"data_recipient_global_metastore_id"})

		common.CustomizeSchemaPath(s, "created_at").SetReadOnly()
		common.CustomizeSchemaPath(s, "created_by").SetReadOnly()
		common.CustomizeSchemaPath(s, "updated_at").SetReadOnly()
		common.CustomizeSchemaPath(s, "updated_by").SetReadOnly()
		common.CustomizeSchemaPath(s, "metastore_id").SetReadOnly()
		common.CustomizeSchemaPath(s, "region").SetReadOnly()
		common.CustomizeSchemaPath(s, "cloud").SetReadOnly()
		common.CustomizeSchemaPath(s, "activated").SetReadOnly()
		common.CustomizeSchemaPath(s, "activation_url").SetReadOnly()
		common.CustomizeSchemaPath(s, "tokens").SetReadOnly()
		common.CustomizeSchemaPath(s, "tokens", "id").SetReadOnly()
		common.CustomizeSchemaPath(s, "tokens", "created_at").SetReadOnly()
		common.CustomizeSchemaPath(s, "tokens", "created_by").SetReadOnly()
		common.CustomizeSchemaPath(s, "tokens", "activation_url").SetReadOnly()
		common.CustomizeSchemaPath(s, "tokens", "expiration_time").SetReadOnly()
		common.CustomizeSchemaPath(s, "tokens", "updated_at").SetReadOnly()
		common.CustomizeSchemaPath(s, "tokens", "updated_by").SetReadOnly()

		return s
	})
	return common.Resource{
		Schema: recipientSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var createRecipientRequest sharing.CreateRecipient
			common.DataToStructPointer(d, recipientSchema, &createRecipientRequest)
			ri, err := w.Recipients.Create(ctx, createRecipientRequest)
			if err != nil {
				return err
			}
			d.SetId(ri.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			ri, err := w.Recipients.GetByName(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(ri, recipientSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updateRecipientRequest sharing.UpdateRecipient
			common.DataToStructPointer(d, recipientSchema, &updateRecipientRequest)
			updateRecipientRequest.Name = d.Id()

			if d.HasChange("owner") {
				err = w.Recipients.Update(ctx, sharing.UpdateRecipient{
					Name:  updateRecipientRequest.Name,
					Owner: updateRecipientRequest.Owner,
				})
				if err != nil {
					return err
				}
			}

			if !d.HasChangeExcept("owner") {
				return nil
			}

			updateRecipientRequest.Owner = ""
			err = w.Recipients.Update(ctx, updateRecipientRequest)
			if err != nil {
				if d.HasChange("owner") {
					// Rollback
					old, new := d.GetChange("owner")
					rollbackErr := w.Recipients.Update(ctx, sharing.UpdateRecipient{
						Name:  updateRecipientRequest.Name,
						Owner: old.(string),
					})
					if rollbackErr != nil {
						return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
					}
				}
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Recipients.DeleteByName(ctx, d.Id())
		},
	}
}
