package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Disable Legacy DBFS setting
var disableLegacyDbfs = workspaceSetting[settings.DisableLegacyDbfs]{
	settingStruct: settings.DisableLegacyDbfs{},
	customizeSchemaFunc: func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "disable_legacy_access", "value").SetRequired()
		return s
	},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.DisableLegacyDbfs, error) {
		return w.Settings.DisableLegacyDbfs().Get(ctx, settings.GetDisableLegacyDbfsRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.DisableLegacyDbfs) (string, error) {
		t.SettingName = "disable_legacy_dbfs"
		res, err := w.Settings.DisableLegacyDbfs().Update(ctx, settings.UpdateDisableLegacyDbfsRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    "disable_legacy_dbfs.value",
		})
		if err != nil {
			return "", err
		}
		return res.Etag, nil
	},
	deleteFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error) {
		res, err := w.Settings.DisableLegacyDbfs().Delete(ctx, settings.DeleteDisableLegacyDbfsRequest{
			Etag: etag,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
