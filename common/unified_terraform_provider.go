package common

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type Namespace struct {
	ProviderConfig *ProviderConfig `json:"provider_config,omitempty"`
}

type ProviderConfig struct {
	WorkspaceID string `json:"workspace_id"`
}

func NamespaceCustomizeSchema(s *CustomizableSchema) {
	s.SchemaPath("provider_config", "workspace_id").SetValidateFunc(validation.StringIsNotEmpty)
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

func NamespaceCustomizeSchemaMap(m map[string]*schema.Schema) map[string]*schema.Schema {
	if providerConfig, ok := m["provider_config"]; ok {
		if elem, ok := providerConfig.Elem.(*schema.Resource); ok {
			if workspaceID, ok := elem.Schema["workspace_id"]; ok {
				workspaceID.ValidateFunc = validation.StringIsNotEmpty
			}
		}
	}
	return m
}
