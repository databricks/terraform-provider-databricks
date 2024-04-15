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
type SecretScope workspace.CreateScope

func (s SecretScope) CustomizeSchema(m map[string]*schema.Schema, path []string) map[string]*schema.Schema {
	common.CustomizeSchemaPath(m, "name").SetValidateFunc(validScope)
	common.CustomizeSchemaPath(m, "backend_type").SetComputed()
	return m
}

func (s SecretScope) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"secrets.SecretScope": {
			"scope":                  "name",
			"scope_backend_type":     "backend_type",
			"backend_azure_keyvault": "keyvault_metadata",
		},
	}
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
			secretScope.Scope = scope.Name
			secretScope.ScopeBackendType = scope.BackendType
			secretScope.BackendAzureKeyvault = scope.KeyvaultMetadata
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
	s := common.StructToSchema(SecretScope{}, nil)
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
			if scope.BackendAzureKeyvault != nil {
				scope.ScopeBackendType = "AZURE_KEYVAULT"
			} else {
				scope.ScopeBackendType = "DATABRICKS"
			}
			if err := w.Secrets.CreateScope(ctx, workspace.CreateScope(scope)); err != nil {
				return err
			}
			d.SetId(scope.Scope)
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
