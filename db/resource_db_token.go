package db

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/databrickslabs/databricks-terraform/client/service"
)

func resourceToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceTokenCreate,
		Read:   resourceTokenRead,
		Delete: resourceTokenDelete,

		Schema: map[string]*schema.Schema{
			"lifetime_seconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  0,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",
			},
			"creation_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"token_value": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"expiry_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceTokenCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	lifeTimeSeconds := d.Get("lifetime_seconds").(int)
	comment := d.Get("comment").(string)
	tokenResp, err := client.Tokens().Create(int32(lifeTimeSeconds), comment)
	if err != nil {
		return err
	}
	d.SetId(tokenResp.TokenInfo.TokenID)
	err = d.Set("token_value", tokenResp.TokenValue)
	if err != nil {
		return err
	}
	return resourceTokenRead(d, m)
}
func resourceTokenRead(d *schema.ResourceData, m interface{}) error {
	tokenId := d.Id()
	client := m.(service.DBApiClient)
	token, err := client.Tokens().Read(tokenId)
	if err != nil {
		return err
	}
	err = d.Set("creation_time", token.CreationTime)
	if err != nil {
		return err
	}
	err = d.Set("expiry_time", token.ExpiryTime)
	return err
}
func resourceTokenDelete(d *schema.ResourceData, m interface{}) error {
	tokenId := d.Id()
	client := m.(service.DBApiClient)
	err := client.Tokens().Delete(tokenId)
	return err
}
