package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Compliance Security Profile Account level setting
var cspEnablementAccountSetting = accountSetting[settings.CspEnablementAccountSetting]{
	settingStruct: settings.CspEnablementAccountSetting{},
	readFunc: func(ctx context.Context, w *databricks.AccountClient, etag string) (*settings.CspEnablementAccountSetting, error) {
		return w.Settings.CspEnablementAccount().Get(ctx, settings.GetCspEnablementAccountRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.AccountClient, t settings.CspEnablementAccountSetting) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.CspEnablementAccount().Update(ctx, settings.UpdateCspEnablementAccountSettingRequest{
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
