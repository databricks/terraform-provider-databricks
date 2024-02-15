package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Compliance Security Profile Account level setting
var cspEnablementAccountSetting = accountSetting[settings.CSPEnablementAccount]{
	settingStruct: settings.CSPEnablementAccount{},
	readFunc: func(ctx context.Context, w *databricks.AccountClient, etag string) (*settings.CSPEnablementAccount, error) {
		return w.Settings.GetCSPEnablementAccount(ctx, settings.GetCSPEnablementAccountRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.AccountClient, t settings.CSPEnablementAccount) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.UpdateCSPEnablementAccount(ctx, settings.UpdateCSPEnablementAccountRequest{
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
