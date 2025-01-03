package settings

import (
	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
)

// Instructions for adding a new setting:
//
//  1. Create a new file named resource_<SETTING_NAME>.go in this directory.
//  2. In that file, create an instance of either the workspaceSettingDefinition or accountSettingDefinition interface for your setting.
//     If the setting name is user-settable, it will be provided in the third argument to the updateFunc method. If not, you must set the
//     SettingName field appropriately. You must also set AllowMissing: true and the field mask to the field to update.
//  3. Add a new entry to the AllSettingsResources map below. The final resource name will be "databricks_<SETTING_NAME>_setting".
func AllSettingsResources() map[string]common.Resource {
	return map[string]common.Resource{
		"default_namespace":                         makeSettingResource[settings.DefaultNamespaceSetting, *databricks.WorkspaceClient](defaultNamespaceSetting),
		"restrict_workspace_admins":                 makeSettingResource[settings.RestrictWorkspaceAdminsSetting, *databricks.WorkspaceClient](restrictWsAdminsSetting),
		"compliance_security_profile_workspace":     makeSettingResource[settings.ComplianceSecurityProfileSetting, *databricks.WorkspaceClient](complianceSecurityProfileSetting),
		"enhanced_security_monitoring_workspace":    makeSettingResource[settings.EnhancedSecurityMonitoringSetting, *databricks.WorkspaceClient](enhancedSecurityMonitoringSetting),
		"automatic_cluster_update_workspace":        makeSettingResource[settings.AutomaticClusterUpdateSetting, *databricks.WorkspaceClient](automaticClusterUpdateSetting),
		"aibi_dashboard_embedding_access_policy":    makeSettingResource[settings.AibiDashboardEmbeddingAccessPolicySetting, *databricks.WorkspaceClient](aibiDashboardEmbeddingAccessPolicySetting),
		"aibi_dashboard_embedding_approved_domains": makeSettingResource[settings.AibiDashboardEmbeddingApprovedDomainsSetting, *databricks.WorkspaceClient](aibiDashboardEmbeddingApprovedDomainsSetting),
	}
}
