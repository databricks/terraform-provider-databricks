package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Enhanced Security Monitoring Account level setting
var esmEnablementAccountSetting = accountSetting[settings.EsmEnablementAccountSetting]{
	settingStruct: settings.EsmEnablementAccountSetting{},
	readFunc: func(ctx context.Context, w *databricks.AccountClient, etag string) (*settings.EsmEnablementAccountSetting, error) {
		return w.Settings.EsmEnablementAccount().Get(ctx, settings.GetEsmEnablementAccountRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.AccountClient, t settings.EsmEnablementAccountSetting) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.EsmEnablementAccount().Update(ctx, settings.UpdateEsmEnablementAccountSettingRequest{
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
