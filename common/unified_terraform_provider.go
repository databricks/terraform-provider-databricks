package common

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ProviderConfig struct {
	WorkspaceID string `json:"workspace_id"`
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
