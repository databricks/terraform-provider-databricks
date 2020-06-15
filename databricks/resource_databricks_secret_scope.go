package databricks

import (
	"fmt"
	"log"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSecretScope() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretScopeCreate,
		Read:   resourceSecretScopeRead,
		Delete: resourceSecretScopeDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"initial_manage_principal": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "users",
			},
			"backend_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSecretScopeCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
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
	client := m.(*service.DBApiClient)
	id := d.Id()
	scope, err := client.SecretScopes().Read(id)
	if err != nil {
		if isSecretScopeMissing(err.Error(), id) {
			log.Printf("Missing secret scope with name: %s.", id)
			d.SetId("")
			return nil
		}
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
	client := m.(*service.DBApiClient)
	id := d.Id()
	err := client.SecretScopes().Delete(id)
	return err
}

func isSecretScopeMissing(errorMsg, resourceID string) bool {
	return strings.Contains(errorMsg, fmt.Sprintf("no Secret Scope found with scope name %s", resourceID))
}
