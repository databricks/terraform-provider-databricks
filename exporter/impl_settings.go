package exporter

import (
	"github.com/databricks/databricks-sdk-go/service/settingsv2"
	account_setting_v2_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/account_setting_v2"
	workspace_setting_v2_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/workspace_setting_v2"
)

func importWorkspaceSettingV2(ic *importContext, r *resource) error {
	// Copy values from effective_* fields to their input counterparts using converter-based approach
	// This works by:
	// 1. Converting TF state to Go SDK struct
	// 2. Copying effective_* fields to input fields using reflection
	// 3. Converting back to TF state
	// This automatically handles all types (simple and complex)
	copyEffectiveFieldsToInputFieldsWithConverters[workspace_setting_v2_resource.Setting](
		ic, r, settingsv2.Setting{})

	return nil
}

func listWorkspaceSettingsV2(ic *importContext) error {
	settings, err := ic.workspaceClient.WorkspaceSettingsV2.ListWorkspaceSettingsMetadataAll(ic.Context, settingsv2.ListWorkspaceSettingsMetadataRequest{})
	if err != nil {
		return err
	}
	for _, setting := range settings {
		ic.Emit(&resource{
			Resource: "databricks_workspace_setting_v2",
			ID:       setting.Name,
			Name:     setting.Name,
		})
	}
	return nil
}

func importAccountSettingV2(ic *importContext, r *resource) error {
	// Copy values from effective_* fields to their input counterparts using converter-based approach
	// This works by:
	// 1. Converting TF state to Go SDK struct
	// 2. Copying effective_* fields to input fields using reflection
	// 3. Converting back to TF state
	// This automatically handles all types (simple and complex)
	copyEffectiveFieldsToInputFieldsWithConverters[account_setting_v2_resource.Setting](
		ic, r, settingsv2.Setting{})

	return nil
}

func listAccountSettingsV2(ic *importContext) error {
	settings, err := ic.accountClient.SettingsV2.ListAccountSettingsMetadataAll(ic.Context, settingsv2.ListAccountSettingsMetadataRequest{})
	if err != nil {
		return err
	}
	for _, setting := range settings {
		ic.Emit(&resource{
			Resource: "databricks_account_setting_v2",
			ID:       setting.Name,
			Name:     setting.Name,
		})
	}
	return nil
}
