package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Compliance Security Profile Workspace level setting
var cspEnablementWorkspaceSetting = workspaceSetting[settings.CspEnablementSetting]{
	settingStruct: settings.CspEnablementSetting{},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.CspEnablementSetting, error) {
		return w.Settings.CspEnablement().Get(ctx, settings.GetCspEnablementRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.CspEnablementSetting) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.CspEnablement().Update(ctx, settings.UpdateCspEnablementSettingRequest{
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
