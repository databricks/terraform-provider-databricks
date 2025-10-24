package common

import (
	"context"
	"fmt"
	"regexp"

	"github.com/databricks/databricks-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// workspaceIDSchemaKey is the key for the workspace ID in schema
const workspaceIDSchemaKey = "provider_config.0.workspace_id"

type Namespace struct {
	ProviderConfig *ProviderConfig `json:"provider_config,omitempty"`
}

// ProviderConfig is used to store the provider configurations for unified terraform provider
// across resources onboarded to SDKv2.
type ProviderConfig struct {
	WorkspaceID string `json:"workspace_id"`
}

// workspaceIDValidateFunc is used to validate the workspace ID for the provider configuration
func workspaceIDValidateFunc() func(interface{}, string) ([]string, []error) {
	return validation.All(
		validation.StringIsNotEmpty,
		validation.StringMatch(regexp.MustCompile(`^\d+$`), "workspace_id must be a valid integer"),
	)
}

// NamespaceCustomizeSchema is used to customize the schema for the provider configuration
// for a single schema.
func NamespaceCustomizeSchema(s *CustomizableSchema) {
	s.SchemaPath("provider_config", "workspace_id").SetValidateFunc(workspaceIDValidateFunc())
}

// NamespaceCustomizeSchemaMap is used to customize the schema for the provider configuration
// in a map of schemas.
func NamespaceCustomizeSchemaMap(m map[string]*schema.Schema) map[string]*schema.Schema {
	if providerConfig, ok := m["provider_config"]; ok {
		if elem, ok := providerConfig.Elem.(*schema.Resource); ok {
			if workspaceID, ok := elem.Schema["workspace_id"]; ok {
				workspaceID.ValidateFunc = workspaceIDValidateFunc()
			}
		}
	}
	return m
}

// NamespaceCustomizeDiff is used to customize the diff for the provider configuration
// in a resource diff.
func NamespaceCustomizeDiff(d *schema.ResourceDiff) error {
	// Force New
	workspaceIDKey := workspaceIDSchemaKey
	oldWorkspaceID, newWorkspaceID := d.GetChange(workspaceIDKey)
	if oldWorkspaceID != "" && newWorkspaceID != "" && oldWorkspaceID != newWorkspaceID {
		if err := d.ForceNew(workspaceIDKey); err != nil {
			return err
		}
	}

	return nil
}

// WorkspaceClientUnifiedProvider returns the WorkspaceClient for the workspace ID from the resource data
// This is used by resources and data sources that are developed over SDKv2.
func (c *DatabricksClient) WorkspaceClientUnifiedProvider(ctx context.Context, d *schema.ResourceData) (*databricks.WorkspaceClient, error) {
	workspaceIDFromSchema := d.Get(workspaceIDSchemaKey)
	// workspace_id does not exist in the resource data
	if workspaceIDFromSchema == nil {
		return c.GetWorkspaceClientForUnifiedProvider(ctx, "")
	}
	var workspaceID string
	workspaceID, ok := workspaceIDFromSchema.(string)
	if !ok {
		return nil, fmt.Errorf("workspace_id must be a string")
	}
	return c.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
}
