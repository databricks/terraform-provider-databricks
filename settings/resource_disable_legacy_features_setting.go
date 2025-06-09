package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Disable Legacy Features setting
var disableLegacyFeatures = accountSetting[settings.DisableLegacyFeatures]{
	settingStruct: settings.DisableLegacyFeatures{},
	customizeSchemaFunc: func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "disable_legacy_features", "value").SetRequired()
		return s
	},
	readFunc: func(ctx context.Context, w *databricks.AccountClient, etag string) (*settings.DisableLegacyFeatures, error) {
		return w.Settings.DisableLegacyFeatures().Get(ctx, settings.GetDisableLegacyFeaturesRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.AccountClient, t settings.DisableLegacyFeatures) (string, error) {
		t.SettingName = "disable_legacy_features"
		res, err := w.Settings.DisableLegacyFeatures().Update(ctx, settings.UpdateDisableLegacyFeaturesRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    "disable_legacy_features.value",
		})
		if err != nil {
			return "", err
		}
		return res.Etag, nil
	},
	deleteFunc: func(ctx context.Context, w *databricks.AccountClient, etag string) (string, error) {
		res, err := w.Settings.DisableLegacyFeatures().Delete(ctx, settings.DeleteDisableLegacyFeaturesRequest{
			Etag: etag,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
