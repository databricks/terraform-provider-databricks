package access

import (
	"fmt"
	"log"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewSecretScopesAPI creates SecretScopesAPI instance from provider meta
func NewSecretScopesAPI(m interface{}) SecretScopesAPI {
	return SecretScopesAPI{C: m.(*common.DatabricksClient)}
}

// SecretScopesAPI exposes the Secret Scopes API
type SecretScopesAPI struct {
	C *common.DatabricksClient
}

// Create creates a new secret scope
func (a SecretScopesAPI) Create(scope string, initialManagePrincipal string) error {
	paramsMap := map[string]string{
		"scope": scope,
	}
	if len(initialManagePrincipal) > 0 {
		paramsMap["initial_manage_principal"] = initialManagePrincipal
	}

	return a.C.Post("/secrets/scopes/create", paramsMap, nil)
}

// Delete deletes a secret scope
func (a SecretScopesAPI) Delete(scope string) error {
	return a.C.Post("/secrets/scopes/delete", map[string]string{
		"scope": scope,
	}, nil)
}

// List lists all secret scopes available in the workspace
func (a SecretScopesAPI) List() ([]SecretScope, error) {
	var listSecretScopesResponse SecretScopeList
	err := a.C.Get("/secrets/scopes/list", nil, &listSecretScopesResponse)
	return listSecretScopesResponse.Scopes, err
}

// Read will return the metadata for the secret scope
func (a SecretScopesAPI) Read(scopeName string) (SecretScope, error) {
	var secretScope SecretScope
	scopes, err := a.List()
	if err != nil {
		return secretScope, err
	}
	for _, scope := range scopes {
		if scope.Name == scopeName {
			return scope, nil
		}
	}
	return secretScope, common.APIError{
		ErrorCode:  "NOT_FOUND",
		Message:    fmt.Sprintf("no Secret Scope found with scope name %s", scopeName),
		Resource:   "/api/2.0/secrets/scopes/list",
		StatusCode: http.StatusNotFound,
	}
}

func ResourceSecretScope() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretScopeCreate,
		Read:   resourceSecretScopeRead,
		Delete: resourceSecretScopeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
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
			},
			"backend_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSecretScopeCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	scopeName := d.Get("name").(string)
	initialManagePrincipal := d.Get("initial_manage_principal").(string)
	err := NewSecretScopesAPI(client).Create(scopeName, initialManagePrincipal)
	if err != nil {
		return err
	}
	d.SetId(scopeName)
	return resourceSecretScopeRead(d, m)
}

func resourceSecretScopeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	id := d.Id()
	scope, err := NewSecretScopesAPI(client).Read(id)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("[INFO] missing resource due to error: %v\n", e)
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
	client := m.(*common.DatabricksClient)
	id := d.Id()
	err := NewSecretScopesAPI(client).Delete(id)
	return err
}
