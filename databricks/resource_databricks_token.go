package databricks

import (
	"log"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceTokenCreate,
		Read:   resourceTokenRead,
		Delete: resourceTokenDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"lifetime_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  0,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"token_value": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"expiry_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceTokenCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DatabricksClient)
	lifeTimeSeconds := d.Get("lifetime_seconds").(int)
	comment := d.Get("comment").(string)
	tokenDuration := time.Duration(lifeTimeSeconds) * time.Second
	tokenResp, err := client.Tokens().Create(tokenDuration, comment)
	if err != nil {
		return err
	}
	d.SetId(tokenResp.TokenInfo.TokenID)
	err = d.Set("token_value", tokenResp.TokenValue)
	if err != nil {
		return err
	}
	// TODO: (loprio) check if we can omit read
	return resourceTokenRead(d, m)
}

func resourceTokenRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DatabricksClient)
	token, err := client.Tokens().Read(id)
	if err != nil {
		if e, ok := err.(service.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}
	err = d.Set("creation_time", token.CreationTime)
	if err != nil {
		return err
	}
	err = d.Set("comment", token.Comment)
	if err != nil {
		return err
	}
	err = d.Set("expiry_time", token.ExpiryTime)
	return err
}

func resourceTokenDelete(d *schema.ResourceData, m interface{}) error {
	tokenID := d.Id()
	client := m.(*service.DatabricksClient)
	err := client.Tokens().Delete(tokenID)
	return err
}
