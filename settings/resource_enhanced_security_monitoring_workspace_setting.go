package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Enhanced Security Monitoring Workspace level setting
var esmEnablementWorkspaceSetting = workspaceSetting[settings.EsmEnablementSetting]{
	settingStruct: settings.EsmEnablementSetting{},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.EsmEnablementSetting, error) {
		return w.Settings.EsmEnablement().Get(ctx, settings.GetEsmEnablementRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.EsmEnablementSetting) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.EsmEnablement().Update(ctx, settings.UpdateEsmEnablementSettingRequest{
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
