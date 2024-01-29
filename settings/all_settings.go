package settings

import (
	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Instructions for adding a new setting:
//
//  1. Create a new file named resource_<SETTING_NAME>.go in this directory.
//  2. In that file, create an instance of either the workspaceSettingDefinition or accountSettingDefinition interface for your setting.
//     If the setting name is user-settable, it will be provided in the third argument to the updateFunc method. If not, you must set the
//     SettingName field appropriately. You must also set AllowMissing: true and the field mask to the field to update.
//  3. Add a new entry to the AllSettingsResources map below. The final resource name will be "databricks_<SETTING_NAME>_setting".
func AllSettingsResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"default_namespace": makeSettingResource[settings.DefaultNamespaceSetting, *databricks.WorkspaceClient](defaultNamespaceSetting),
	}
}
