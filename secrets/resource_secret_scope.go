package secrets

import (
	"context"
	"fmt"
	"regexp"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// SecretScope is a struct that encapsulates the secret scope
type SecretScope struct {
	Name                   string                                      `json:"name" tf:"force_new"`
	BackendType            string                                      `json:"backend_type,omitempty" tf:"computed"`
	InitialManagePrincipal string                                      `json:"initial_manage_principal,omitempty" tf:"force_new"`
	KeyvaultMetadata       *workspace.AzureKeyVaultSecretScopeMetadata `json:"keyvault_metadata,omitempty" tf:"force_new"`
}

// Read will return the metadata for the secret scope
func readSecretScope(ctx context.Context, w *databricks.WorkspaceClient, scopeName string) (SecretScope, error) {
	var secretScope SecretScope
	scopes := w.Secrets.ListScopes(ctx)
	for scopes.HasNext(ctx) {
		scope, err := scopes.Next(ctx)
		if err != nil {
			return secretScope, err
		}
		if scope.Name == scopeName {
			secretScope.Name = scope.Name
			secretScope.BackendType = scope.BackendType.String()
			secretScope.KeyvaultMetadata = scope.KeyvaultMetadata
			return secretScope, nil
		}

	}
	return secretScope, apierr.NotFound(
		fmt.Sprintf("no Secret Scope found with scope name %s", scopeName))
}

var validScope = validation.StringMatch(regexp.MustCompile(`^[\w\.@_/-]{1,128}$`),
	"Must consist of alphanumeric characters, dashes, underscores, and periods, "+
		"and may not exceed 128 characters.")

// ResourceSecretScope manages secret scopes
func ResourceSecretScope() common.Resource {
	s := common.StructToSchema(SecretScope{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// TODO: DiffSuppressFunc for initial_manage_principal & importing
		// nolint
		s["name"].ValidateFunc = validScope

		return s
	})
	return common.Resource{
		Schema:        s,
		SchemaVersion: 2,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var scope SecretScope
			common.DataToStructPointer(d, s, &scope)
			scopeRequest := workspace.CreateScope{
				Scope:                  scope.Name,
				InitialManagePrincipal: scope.InitialManagePrincipal,
				ScopeBackendType:       "DATABRICKS",
			}
			if scope.KeyvaultMetadata != nil {
				scopeRequest.ScopeBackendType = "AZURE_KEYVAULT"
				scopeRequest.BackendAzureKeyvault = scope.KeyvaultMetadata
			}
			if err := w.Secrets.CreateScope(ctx, scopeRequest); err != nil {
				return err
			}
			d.SetId(scope.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			scope, err := readSecretScope(ctx, w, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(scope, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Secrets.DeleteScope(ctx, workspace.DeleteScope{Scope: d.Id()})
		},
	}
}
