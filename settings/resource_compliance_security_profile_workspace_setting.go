package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Compliance Security Profile Workspace level setting
var cspEnablementWorkspaceSetting = workspaceSetting[settings.CSPEnablement]{
	settingStruct: settings.CSPEnablement{},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.CSPEnablement, error) {
		return w.Settings.GetCSPEnablement(ctx, settings.GetCSPEnablementRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.CSPEnablement) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.UpdateCSPEnablement(ctx, settings.UpdateCSPEnablementRequest{
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
