package db

import (
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSecretScope() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretScopeCreate,
		Read:   resourceSecretScopeRead,
		Delete: resourceSecretScopeDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"initial_manage_principal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  nil,
			},
			"backend_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSecretScopeCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	scopeName := d.Get("name").(string)
	initialManagePrincipal := d.Get("initial_manage_principal").(string)
	err := client.SecretScopes().Create(scopeName, initialManagePrincipal)
	if err != nil {
		return err
	}
	d.SetId(scopeName)
	return resourceSecretScopeRead(d, m)
}

func resourceSecretScopeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	id := d.Id()
	scope, err := client.SecretScopes().Read(id)
	if err != nil {
		return err
	}
	d.SetId(scope.Name)
	err = d.Set("name", scope.Name)
	if err != nil {
		return err
	}
	err = d.Set("backend_type", scope.BackendType)
	return err
}

func resourceSecretScopeDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	id := d.Id()
	err := client.SecretScopes().Delete(id)
	return err
}
