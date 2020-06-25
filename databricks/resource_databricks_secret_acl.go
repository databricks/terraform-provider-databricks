package databricks

import (
	"log"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSecretACL() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretACLCreate,
		Read:   resourceSecretACLRead,
		Delete: resourceSecretACLDelete,

		Schema: map[string]*schema.Schema{
			"scope": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"principal": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"permission": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func getSecretACLID(scope string, key string) (string, error) {
	return scope + "|||" + key, nil
}

func getScopeAndKeyFromSecretACLID(secretACLIDString string) (string, string, error) {
	return strings.Split(secretACLIDString, "|||")[0], strings.Split(secretACLIDString, "|||")[1], nil
}

func resourceSecretACLCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	scopeName := d.Get("scope").(string)
	principal := d.Get("principal").(string)
	permission := model.ACLPermission(d.Get("permission").(string))
	err := client.SecretAcls().Create(scopeName, principal, permission)
	if err != nil {
		return err
	}
	id, err := getSecretACLID(scopeName, principal)
	if err != nil {
		return err
	}
	d.SetId(id)
	return resourceSecretACLRead(d, m)
}

func resourceSecretACLRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	scope, principal, err := getScopeAndKeyFromSecretACLID(id)
	if err != nil {
		return err
	}
	client := m.(*service.DBApiClient)
	secretACL, err := client.SecretAcls().Read(scope, principal)
	if err != nil {
		if e, ok := err.(service.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}
	err = d.Set("scope", scope)
	if err != nil {
		return err
	}
	err = d.Set("principal", principal)
	if err != nil {
		return err
	}
	err = d.Set("permission", secretACL.Permission)
	return err
}

func resourceSecretACLDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	id := d.Id()
	scope, key, err := getScopeAndKeyFromSecretACLID(id)
	if err != nil {
		return err
	}
	err = client.SecretAcls().Delete(scope, key)
	return err
}
