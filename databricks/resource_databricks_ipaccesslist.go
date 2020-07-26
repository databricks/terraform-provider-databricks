package databricks

// Preview feature: https://docs.databricks.com/security/network/ip-access-list.html
// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"log"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIPAccessList() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPACLCreate,
		Read:   resourceIPACLRead,
		Update: resourceIPACLUpdate,
		Delete: resourceIPACLDelete,

		Schema: map[string]*schema.Schema{
			"label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"list_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						string(model.WhiteList),
						string(model.BlackList),
					}, false),
			},
			"ip_addresses": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.Any(validation.IsIPv4Address, validation.IsCIDR),
				},
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceIPACLCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	label := d.Get("label").(string)
	ipAddresses := d.Get("ip_addresses").([]interface{})
	listType := d.Get("list_type").(string)

	log.Println("IPACLLists: Calling IP ACL Create")
	status, err := client.IPAccessLists().Create(
		convertListInterfaceToString(ipAddresses),
		label,
		model.IPAccessListType(listType))
	log.Printf("IPACLLists: Created as  %v\n", status)
	if err != nil {
		log.Printf("IPACLLists:  Creation error %v\n", err)
		return err
	}

	d.SetId(status.ListID)

	return resourceIPACLRead(d, m)
}

func resourceIPACLRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	status, err := client.IPAccessLists().Read(d.Id())
	if err != nil {
		// check 404 (missing) and set id to empty string for tf
		if e, ok := err.(service.APIError); ok && e.IsMissing() {
			log.Printf("[IPACLLists:  missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}

	return updateFromStatus(d, status)
}

func resourceIPACLUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	ipAddresses := convertListInterfaceToString(d.Get("ip_addresses").([]interface{}))
	_, err := client.IPAccessLists().Update(
		d.Id(),
		d.Get("label").(string),
		model.IPAccessListType(d.Get("list_type").(string)),
		ipAddresses,
		d.Get("enabled").(bool),
	)
	if err != nil {
		return err
	}
	return resourceIPACLRead(d, m)
}

func resourceIPACLDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	return client.IPAccessLists().Delete(d.Id())
}

func updateFromStatus(d *schema.ResourceData, status model.IPAccessListStatus) error {
	err := d.Set("label", status.Label)
	if err != nil {
		return err
	}
	err = d.Set("list_type", string(status.ListType))
	if err != nil {
		return err
	}
	err = d.Set("ip_addresses", status.IPAddresses)
	if err != nil {
		return err
	}
	err = d.Set("enabled", status.Enabled)

	return err
}
