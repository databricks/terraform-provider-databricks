package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Compliance Security Profile Account level setting
var cspEnablementAccountSetting = accountSetting[settings.CSPEnablementAccountSetting]{
	settingStruct: settings.CSPEnablementAccountSetting{},
	readFunc: func(ctx context.Context, w *databricks.AccountClient, etag string) (*settings.CSPEnablementAccountSetting, error) {
		return w.Settings.GetCSPEnablementAccountSetting(ctx, settings.GetCSPEnablementAccountRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.AccountClient, t settings.CSPEnablementAccountSetting) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.UpdateCSPEnablementAccountSetting(ctx, settings.UpdateCSPEnablementAccountRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    "csp_enablement_account.value",
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
