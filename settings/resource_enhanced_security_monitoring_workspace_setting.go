package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Enhanced Security Monitoring Workspace level setting
var esmEnablementWorkspaceSetting = workspaceSetting[settings.ESMEnablementSetting]{
	settingStruct: settings.ESMEnablementSetting{},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.ESMEnablementSetting, error) {
		return w.Settings.GetESMEnablementSetting(ctx, settings.GetESMEnablementSettingRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.ESMEnablementSetting) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.UpdateESMEnablementSetting(ctx, settings.UpdateESMEnablementSettingRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    "esm_enablement_workspace.value",
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
