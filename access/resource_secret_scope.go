package access

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewSecretScopesAPI creates SecretScopesAPI instance from provider meta
func NewSecretScopesAPI(m interface{}) SecretScopesAPI {
	return SecretScopesAPI{client: m.(*common.DatabricksClient)}
}

// SecretScopesAPI exposes the Secret Scopes API
type SecretScopesAPI struct {
	client *common.DatabricksClient
}

// SecretScopeList holds list of secret scopes
type SecretScopeList struct {
	Scopes []SecretScope `json:"scopes,omitempty"`
}

// SecretScope is a struct that encapsulates the secret scope
type SecretScope struct {
	Name                   string            `json:"name"`
	BackendType            string            `json:"backend_type,omitempty" tf:"computed"`
	InitialManagePrincipal string            `json:"initial_manage_principal,omitempty"`
	KeyvaultMetadata       *KeyvaultMetadata `json:"keyvault_metadata,omitempty"`
}

// KeyvaultMetadata Azure Key Vault metadata wrapper
type KeyvaultMetadata struct {
	// /subscriptions/.../resourceGroups/.../providers/Microsoft.KeyVault/vaults/my-azure-kv
	ResourceID string `json:"resource_id"`
	// https://my-azure-kv.vault.azure.net/
	DNSName string `json:"dns_name"`
}

type secretScopeRequest struct {
	Scope string `json:"scope,omitempty"`
	// validate
	// BackendType string `json:"backend_type,omitempty"`
	BackendType string `json:"scope_backend_type,omitempty"`

	// todo: validate conflicting
	InitialManagePrincipal string            `json:"initial_manage_principal,omitempty"`
	BackendAzureKeyvault   *KeyvaultMetadata `json:"backend_azure_keyvault,omitempty"`
}

// Create creates a new secret scope
func (a SecretScopesAPI) Create(s SecretScope) error {
	req := secretScopeRequest{
		Scope:                  s.Name,
		InitialManagePrincipal: s.InitialManagePrincipal,
		BackendType:            "DATABRICKS",
	}
	if s.KeyvaultMetadata != nil {
		if !a.client.IsAzure() {
			return fmt.Errorf("Azure KeyVault is not available")
		}
		if a.client.AzureAuth.IsClientSecretSet() {
			return fmt.Errorf("Azure KeyVault cannot yet be configured for Service Principal authorization")
		}
		req.BackendType = "AZURE_KEYVAULT"
		req.BackendAzureKeyvault = s.KeyvaultMetadata
	}
	return a.client.Post("/secrets/scopes/create", req, nil)
}

// Delete deletes a secret scope
func (a SecretScopesAPI) Delete(scope string) error {
	return a.client.Post("/secrets/scopes/delete", map[string]string{
		"scope": scope,
	}, nil)
}

// List lists all secret scopes available in the workspace
func (a SecretScopesAPI) List() ([]SecretScope, error) {
	var listSecretScopesResponse SecretScopeList
	err := a.client.Get("/secrets/scopes/list", nil, &listSecretScopesResponse)
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

// ResourceSecretScope manages secret scopes
func ResourceSecretScope() *schema.Resource {
	s := internal.StructToSchema(SecretScope{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// TODO: CustomDiffFunc for initial_manage_principal & importing
		s["name"].ForceNew = true
		s["initial_manage_principal"].ForceNew = true
		s["keyvault_metadata"].ForceNew = true
		return s
	})
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		scope, err := NewSecretScopesAPI(m).Read(d.Id())
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			d.SetId("")
			return nil
		}
		if err != nil {
			return diag.FromErr(err)
		}
		err = internal.StructToData(scope, s, d)
		if err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var scope SecretScope
			err := internal.DataToStructPointer(d, s, &scope)
			if err != nil {
				return diag.FromErr(err)
			}
			err = NewSecretScopesAPI(m).Create(scope)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(scope.Name)
			return readContext(ctx, d, m)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			err := NewSecretScopesAPI(m).Delete(d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		ReadContext:   readContext,
		SchemaVersion: 2,
		Schema:        s,
	}
}
