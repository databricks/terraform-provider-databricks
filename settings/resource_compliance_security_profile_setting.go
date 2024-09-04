package settings

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Enhanced Security Monitoring setting
var complianceSecurityProfileFieldMask = strings.Join([]string{
	"compliance_security_profile_workspace.is_enabled",
	"compliance_security_profile_workspace.compliance_standards",
}, ",")
var complianceSecurityProfileSetting = workspaceSetting[settings.ComplianceSecurityProfileSetting]{
	settingStruct: settings.ComplianceSecurityProfileSetting{},
	customizeSchemaFunc: func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "compliance_security_profile_workspace", "compliance_standards").SetRequired()
		common.CustomizeSchemaPath(s, "compliance_security_profile_workspace", "is_enabled").SetRequired()
		return s
	},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.ComplianceSecurityProfileSetting, error) {
		return w.Settings.ComplianceSecurityProfile().Get(ctx, settings.GetComplianceSecurityProfileSettingRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.ComplianceSecurityProfileSetting) (string, error) {
		t.SettingName = "default"
		t.ComplianceSecurityProfileWorkspace.ForceSendFields = []string{"IsEnabled"}
		res, err := w.Settings.ComplianceSecurityProfile().Update(ctx, settings.UpdateComplianceSecurityProfileSettingRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    complianceSecurityProfileFieldMask,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
	deleteFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error) {
		tflog.Warn(ctx, "databricks_compliance_security_profile_workspace_setting couldn't be disabled!")
		return etag, nil
	},
}
