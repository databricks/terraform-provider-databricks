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
	return strings.Split(secretIDString, "|||")[0], strings.Split(secretIDString, "|||")[1], nil
}

func resourceSecretCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
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
	client := m.(*service.DBApiClient)
	scope, key, err := getScopeAndKeyFromSecretID(id)
	if err != nil {
		return err
	}
	secretMetaData, err := client.Secrets().Read(scope, key)
	if err != nil {
		if isErrorRecoverable(err, scope, key) {
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
	client := m.(*service.DBApiClient)
	id := d.Id()
	scope, key, err := getScopeAndKeyFromSecretID(id)
	if err != nil {
		return err
	}
	err = client.Secrets().Delete(scope, key)
	return err
}

func isErrorRecoverable(err error, scope string, key string) bool {
	if isSecretMissing(err.Error(), scope, key) {
		log.Printf("Missing secret with id: %s in scope with id: %s.", scope, key)
		return true
	}
	if isScopeMissing(err.Error(), scope) {
		log.Printf("Missing scope with id: %s; secret %s cannot exist without scope", scope, key)
		return true
	}

	return false
}

func isSecretMissing(errorMsg, scope string, key string) bool {
	return strings.Contains(errorMsg, fmt.Sprintf("no Secret Scope found with secret metadata scope name: %s and key: %s", scope, key))
}

func isScopeMissing(errorMsg, scope string) bool {
	return strings.Contains(errorMsg, fmt.Sprintf("Scope %s does not exist!", scope))
}
