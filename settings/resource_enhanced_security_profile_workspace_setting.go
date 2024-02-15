package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Enhanced Security Monitoring Workspace level setting
var esmEnablementSetting = workspaceSetting[settings.ESMEnablement]{
	settingStruct: settings.ESMEnablement{},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.ESMEnablement, error) {
		return w.Settings.GetESMEnablement(ctx, settings.GetESMEnablementRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.ESMEnablement) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.UpdateESMEnablement(ctx, settings.UpdateESMEnablementRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    "namespace.value",
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	}
}
