package sharing

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type Token struct {
	Id             string `json:"id,omitempty" tf:"computed"`
	CreatedAt      int64  `json:"created_at,omitempty" tf:"computed"`
	CreatedBy      string `json:"created_by,omitempty" tf:"computed"`
	ActivationUrl  string `json:"activation_url,omitempty" tf:"computed"`
	ExpirationTime int64  `json:"expiration_time,omitempty" tf:"computed"`
	UpdatedAt      int64  `json:"updated_at,omitempty" tf:"computed"`
	UpdatedBy      string `json:"updated_by,omitempty" tf:"computed"`
}

type IpAccessList struct {
	AllowedIpAddresses []string `json:"allowed_ip_addresses"`
}

type RecipientInfo struct {
	Name                           string        `json:"name" tf:"force_new"`
	Comment                        string        `json:"comment,omitempty"`
	SharingCode                    string        `json:"sharing_code,omitempty" tf:"sensitive,force_new,suppress_diff"`
	AuthenticationType             string        `json:"authentication_type" tf:"force_new"`
	Tokens                         []Token       `json:"tokens,omitempty" tf:"computed"`
	Owner                          string        `json:"owner,omitempty" tf:"suppress_diff"`
	DataRecipientGlobalMetastoreId string        `json:"data_recipient_global_metastore_id,omitempty" tf:"force_new,conflicts:ip_access_list"`
	IpAccessList                   *IpAccessList `json:"ip_access_list,omitempty"`
}

func ResourceRecipient() common.Resource {
	recipientSchema := common.StructToSchema(RecipientInfo{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		m["authentication_type"].ValidateFunc = validation.StringInSlice([]string{"TOKEN", "DATABRICKS"}, false)
		m["name"].DiffSuppressFunc = common.EqualFoldDiffSuppress
		return m
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
