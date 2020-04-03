package db

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stikkireddy/databricks-tf-provider/client/model"
	"github.com/stikkireddy/databricks-tf-provider/client/service"
	"strings"
)

func resourceSecretAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretAclCreate,
		Read:   resourceSecretAclRead,
		Delete: resourceSecretAclDelete,

		Schema: map[string]*schema.Schema{
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"principal": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"permission": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

type SecretAclId map[string]string

func getSecretAclId(scope string, key string) (string, error) {
	return scope + "|||" + key, nil
}

func getScopeAndKeyFromSecretAclId(SecretAclIdString string) (string, string, error) {
	return strings.Split(SecretAclIdString, "|||")[0], strings.Split(SecretAclIdString, "|||")[1], nil
}

func resourceSecretAclCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	scopeName := d.Get("scope").(string)
	principal := d.Get("principal").(string)
	permission := model.AclPermission(d.Get("permission").(string))
	err := client.SecretAcls().Create(scopeName, principal, permission)
	if err != nil {
		return err
	}
	id, err := getSecretAclId(scopeName, principal)
	if err != nil {
		return err
	}
	d.SetId(id)
	return resourceSecretAclRead(d, m)
}

func resourceSecretAclRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	scope, principal, err := getScopeAndKeyFromSecretAclId(id)
	if err != nil {
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

	client := m.(service.DBApiClient)
	secretAcl, err := client.SecretAcls().Read(scope, principal)
	if err != nil {
		return err
	}
	err = d.Set("permission", secretAcl.Permission)
	return err
}

func resourceSecretAclDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	id := d.Id()
	scope, key, err := getScopeAndKeyFromSecretAclId(id)
	if err != nil {
		return err
	}
	err = client.SecretAcls().Delete(scope, key)
	return err
}
