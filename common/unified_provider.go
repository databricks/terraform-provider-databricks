package common

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type Namespace struct {
	ProviderConfig *ProviderConfig `json:"provider_config,omitempty"`
}

type ProviderConfig struct {
	WorkspaceID string `json:"workspace_id"`
}

func workspaceIDValidateFunc() func(interface{}, string) ([]string, []error) {
	return validation.StringIsNotEmpty
}

func NamespaceCustomizeSchema(s *CustomizableSchema) {
	s.SchemaPath("provider_config", "workspace_id").SetValidateFunc(workspaceIDValidateFunc())
}

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

func NamespaceCustomizeDiff(d *schema.ResourceDiff) error {
	workspaceIDKey := "provider_config.0.workspace_id"
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
func WorkspaceClientUnifiedProvider(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) (*databricks.WorkspaceClient, error) {
	workspaceIDFromSchema := d.Get("provider_config.0.workspace_id")
	// workspace_id does not exist in the schema
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
