package common

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type ProviderConfig struct {
	WorkspaceID string `json:"workspace_id"`
}

func ProviderConfigCustomizeSchema(m map[string]*schema.Schema) map[string]*schema.Schema {
	// Debug: print all keys in the schema map
	fmt.Printf("DEBUG: Schema keys: ")
	for k := range m {
		fmt.Printf("%s, ", k)
	}
	fmt.Printf("\n")

	if providerConfig, ok := m["provider_config"]; ok {
		fmt.Printf("DEBUG: Found provider_config in schema\n")
		if elem, ok := providerConfig.Elem.(*schema.Resource); ok {
			fmt.Printf("DEBUG: provider_config.Elem is a Resource\n")
			if workspaceID, ok := elem.Schema["workspace_id"]; ok {
				fmt.Printf("DEBUG: Found workspace_id in provider_config schema\n")
				workspaceID.ValidateFunc = validation.StringIsNotEmpty
			} else {
				fmt.Printf("DEBUG: workspace_id NOT found in provider_config schema\n")
			}
		} else {
			fmt.Printf("DEBUG: provider_config.Elem is NOT a Resource, type: %T\n", providerConfig.Elem)
		}
	} else {
		fmt.Printf("DEBUG: provider_config NOT found in schema\n")
	}
	return m
}

func ProviderConfigCustomizeDiff(d *schema.ResourceDiff) error {
	workspaceIDKey := "provider_config.0.workspace_id"
	oldWorkspaceID, newWorkspaceID := d.GetChange(workspaceIDKey)
	if oldWorkspaceID != "" && newWorkspaceID != "" && oldWorkspaceID != newWorkspaceID {
		if err := d.ForceNew(workspaceIDKey); err != nil {
			return err
		}
	}
	return nil
}
