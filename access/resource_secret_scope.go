package access

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// NewSecretScopesAPI creates SecretScopesAPI instance from provider meta
func NewSecretScopesAPI(ctx context.Context, m interface{}) SecretScopesAPI {
	return SecretScopesAPI{m.(*common.DatabricksClient), ctx}
}

// SecretScopesAPI exposes the Secret Scopes API
type SecretScopesAPI struct {
	client  *common.DatabricksClient
	context context.Context
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
	Scope                  string            `json:"scope,omitempty"`
	BackendType            string            `json:"scope_backend_type,omitempty"`
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
		if err := a.client.Authenticate(); err != nil {
			return err
		}
		if !a.client.IsAzure() {
			return fmt.Errorf("Azure KeyVault is not available")
		}
		if a.client.AzureAuth.IsClientSecretSet() {
			return fmt.Errorf("Azure KeyVault cannot yet be configured for Service Principal authorization")
		}
		req.BackendType = "AZURE_KEYVAULT"
		req.BackendAzureKeyvault = s.KeyvaultMetadata
	}
	return a.client.Post(a.context, "/secrets/scopes/create", req, nil)
}

// Delete deletes a secret scope
func (a SecretScopesAPI) Delete(scope string) error {
	return a.client.Post(a.context, "/secrets/scopes/delete", map[string]string{
		"scope": scope,
	}, nil)
}

// List lists all secret scopes available in the workspace
func (a SecretScopesAPI) List() ([]SecretScope, error) {
	var listSecretScopesResponse SecretScopeList
	err := a.client.Get(a.context, "/secrets/scopes/list", nil, &listSecretScopesResponse)
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

var validScope = validation.StringMatch(regexp.MustCompile(`^[\w\.@_-]{1,128}$`),
	"Must consist of alphanumeric characters, dashes, underscores, and periods, "+
		"and may not exceed 128 characters.")

// ResourceSecretScope manages secret scopes
func ResourceSecretScope() *schema.Resource {
	s := internal.StructToSchema(SecretScope{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// TODO: DiffSuppressFunc for initial_manage_principal & importing
		s["name"].ForceNew = true
		// nolint
		s["name"].ValidateFunc = validScope
		s["initial_manage_principal"].ForceNew = true
		s["keyvault_metadata"].ForceNew = true
		return s
	})
	return util.CommonResource{
		Schema:        s,
		SchemaVersion: 2,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var scope SecretScope
			if err := internal.DataToStructPointer(d, s, &scope); err != nil {
				return err
			}
			if err := NewSecretScopesAPI(ctx, c).Create(scope); err != nil {
				return err
			}
			d.SetId(scope.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			scope, err := NewSecretScopesAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			return internal.StructToData(scope, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewSecretScopesAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
