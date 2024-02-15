package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Enhanced Security Monitoring Account level setting
var esmEnablementAccountSetting = accountSetting[settings.ESMEnablementAccount]{
	settingStruct: settings.ESMEnablementAccount{},
	readFunc: func(ctx context.Context, w *databricks.AccountClient, etag string) (*settings.ESMEnablementAccount, error) {
		return w.Settings.GetESMEnablementAccount(ctx, settings.GetESMEnablementAccountRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.AccountClient, t settings.ESMEnablementAccount) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.UpdateESMEnablementAccount(ctx, settings.UpdateESMEnablementAccountRequest{
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
