package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Enhanced Security Monitoring Account level setting
var esmEnablementAccountSetting = accountSetting[settings.ESMEnablementAccountSetting]{
	settingStruct: settings.ESMEnablementAccountSetting{},
	readFunc: func(ctx context.Context, w *databricks.AccountClient, etag string) (*settings.ESMEnablementAccountSetting, error) {
		return w.Settings.GetESMEnablementAccountSetting(ctx, settings.GetESMEnablementAccountSettingRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.AccountClient, t settings.ESMEnablementAccountSetting) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.UpdateESMEnablementAccountSetting(ctx, settings.UpdateESMEnablementAccountSettingRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    "esm_enablement_account.value",
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
