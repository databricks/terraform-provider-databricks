package catalog

import (
	"context"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type RecipientsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewRecipientsAPI(ctx context.Context, m interface{}) RecipientsAPI {
	return RecipientsAPI{m.(*common.DatabricksClient), ctx}
}

type Token struct {
	Id             string `json:"id,omitempty" tf:"computed"`
	CreatedAt      int64  `json:"created_at,omitempty" tf:"computed"`
	CreatedBy      string `json:"created_by,omitempty" tf:"computed"`
	ActivationUrl  string `json:"activation_url,omitempty" tf:"computed"`
	ExpirationTime int64  `json:"expiration_time,omitempty" tf:"computed"`
	UpdatedAt      int64  `json:"updated_at,omitempty" tf:"computed"`
	UpdatedBy      string `json:"updated_by,omitempty" tf:"computed"`
}

type RecipientInfo struct {
	Name               string  `json:"name" tf:"force_new"`
	Comment            string  `json:"comment,omitempty"`
	SharingCode        string  `json:"sharing_code,omitempty" tf:"sensitive,force_new,suppress_diff"`
	AuthenticationType string  `json:"authentication_type" tf:"force_new"`
	Tokens             []Token `json:"tokens,omitempty" tf:"computed"`
}

type Recipients struct {
	Recipients []RecipientInfo `json:"recipients"`
}

func (a RecipientsAPI) list() (recipients Recipients, err error) {
	err = a.client.Get(a.context, "/unity-catalog/recipients", nil, &recipients)
	return
}

func (a RecipientsAPI) createRecipient(ci *RecipientInfo) error {
	return a.client.Post(a.context, "/unity-catalog/recipients", ci, ci)
}

func (a RecipientsAPI) getRecipient(name string) (ci RecipientInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/recipients/"+name, nil, &ci)
	return
}

func (a RecipientsAPI) deleteRecipient(name string) error {
	return a.client.Delete(a.context, "/unity-catalog/recipients/"+name, nil)
}

func (a RecipientsAPI) updateRecipient(ci *RecipientInfo) error {
	patch := struct {
		Comment string `json:"comment"`
	}{
		Comment: ci.Comment,
	}
	return a.client.Patch(a.context, "/unity-catalog/recipients/"+ci.Name, patch)
}

func ResourceRecipient() *schema.Resource {
	recipientSchema := common.StructToSchema(RecipientInfo{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		return m
	})
	return common.Resource{
		Schema: recipientSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ri RecipientInfo
			common.DataToStructPointer(d, recipientSchema, &ri)
			if err := NewRecipientsAPI(ctx, c).createRecipient(&ri); err != nil {
				return err
			}
			d.SetId(ri.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ri, err := NewRecipientsAPI(ctx, c).getRecipient(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(ri, recipientSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ri RecipientInfo
			common.DataToStructPointer(d, recipientSchema, &ri)
			return NewRecipientsAPI(ctx, c).updateRecipient(&ri)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewRecipientsAPI(ctx, c).deleteRecipient(d.Id())
		},
	}.ToResource()
}
