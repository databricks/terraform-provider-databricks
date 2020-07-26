package databricks

import (
	"fmt"
	"log"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretCreate,
		Read:   resourceSecretRead,
		Delete: resourceSecretDelete,

		Schema: map[string]*schema.Schema{
			"string_value": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"last_updated_timestamp": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func getSecretID(scope string, key string) (string, error) {
	return scope + "|||" + key, nil
}

func getScopeAndKeyFromSecretID(secretIDString string) (string, string, error) {
	split := strings.Split(secretIDString, "|||")
	if len(split) != 2 {
		return "", "", fmt.Errorf("Malformed secret id: %s", secretIDString)
	}
	return strings.Split(secretIDString, "|||")[0], strings.Split(secretIDString, "|||")[1], nil
}

func resourceSecretCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DatabricksClient)
	scopeName := d.Get("scope").(string)
	key := d.Get("key").(string)
	secretValue := d.Get("string_value").(string)
	err := client.Secrets().Create(secretValue, scopeName, key)
	if err != nil {
		return err
	}
	id, err := getSecretID(scopeName, key)
	if err != nil {
		return err
	}
	d.SetId(id)
	return resourceSecretRead(d, m)
}

func resourceSecretRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DatabricksClient)
	scope, key, err := getScopeAndKeyFromSecretID(id)
	if err != nil {
		return err
	}
	secretMetaData, err := client.Secrets().Read(scope, key)
	if err != nil {
		if e, ok := err.(service.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}

	err = d.Set("last_updated_timestamp", secretMetaData.LastUpdatedTimestamp)
	if err != nil {
		return err
	}

	err = d.Set("scope", scope)
	if err != nil {
		return err
	}
	err = d.Set("key", secretMetaData.Key)
	if err != nil {
		return err
	}
	d.SetId(id)
	return nil
}

func resourceSecretDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DatabricksClient)
	id := d.Id()
	scope, key, err := getScopeAndKeyFromSecretID(id)
	if err != nil {
		return err
	}
	err = client.Secrets().Delete(scope, key)
	return err
}
