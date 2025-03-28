package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Disable Legacy Access setting
var disableLegacyAccess = workspaceSetting[settings.DisableLegacyAccess]{
	settingStruct: settings.DisableLegacyAccess{},
	customizeSchemaFunc: func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "disable_legacy_access", "value").SetRequired()
		return s
	},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.DisableLegacyAccess, error) {
		return w.Settings.DisableLegacyAccess().Get(ctx, settings.GetDisableLegacyAccessRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.DisableLegacyAccess) (string, error) {
		t.SettingName = "disable_legacy_access"
		res, err := w.Settings.DisableLegacyAccess().Update(ctx, settings.UpdateDisableLegacyAccessRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    "disable_legacy_access.value",
		})
		if err != nil {
			return "", err
		}
		return res.Etag, nil
	},
	deleteFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error) {
		res, err := w.Settings.DisableLegacyAccess().Delete(ctx, settings.DeleteDisableLegacyAccessRequest{
			Etag: etag,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
