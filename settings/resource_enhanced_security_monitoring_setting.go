package settings

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Enhanced Security Monitoring setting
var enhancedSecurityMonitoringFieldMask = strings.Join([]string{
	"enhanced_security_monitoring_workspace.is_enabled",
}, ",")

var enhancedSecurityMonitoringSetting = workspaceSetting[settings.EnhancedSecurityMonitoringSetting]{
	settingStruct: settings.EnhancedSecurityMonitoringSetting{},
	customizeSchemaFunc: func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "enhanced_security_monitoring_workspace", "is_enabled").SetRequired()
		return s
	},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.EnhancedSecurityMonitoringSetting, error) {
		return w.Settings.EnhancedSecurityMonitoring().Get(ctx, settings.GetEnhancedSecurityMonitoringSettingRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.EnhancedSecurityMonitoringSetting) (string, error) {
		t.SettingName = "default"
		t.EnhancedSecurityMonitoringWorkspace.ForceSendFields = []string{"IsEnabled"}
		res, err := w.Settings.EnhancedSecurityMonitoring().Update(ctx, settings.UpdateEnhancedSecurityMonitoringSettingRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    enhancedSecurityMonitoringFieldMask,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
	deleteFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error) {
		res, err := w.Settings.EnhancedSecurityMonitoring().Update(ctx, settings.UpdateEnhancedSecurityMonitoringSettingRequest{
			AllowMissing: true,
			Setting: settings.EnhancedSecurityMonitoringSetting{
				Etag:        etag,
				SettingName: "default",
				EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
					IsEnabled:       false,
					ForceSendFields: []string{"IsEnabled"},
				},
			},
			FieldMask: enhancedSecurityMonitoringFieldMask,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
