package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Compliance Security Profile Workspace level setting
var cspEnablementWorkspaceSetting = workspaceSetting[settings.CSPEnablementSetting]{
	settingStruct: settings.CSPEnablementSetting{},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.CSPEnablementSetting, error) {
		return w.Settings.GetCSPEnablementSetting(ctx, settings.GetCSPEnablementSettingRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.CSPEnablementSetting) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.UpdateCSPEnablementSetting(ctx, settings.UpdateCSPEnablementSettingRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    "csp_enablement_workspace.value",
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
