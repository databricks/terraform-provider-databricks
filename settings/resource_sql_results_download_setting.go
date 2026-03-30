package settings

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var sqlResultsDownloadFieldMask = strings.Join([]string{
	"boolean_val.value",
}, ",")

var sqlResultsDownloadSetting = workspaceSetting[settings.SqlResultsDownload]{
	settingStruct: settings.SqlResultsDownload{},
	customizeSchemaFunc: func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "boolean_val", "value").SetRequired()
		return s
	},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.SqlResultsDownload, error) {
		return w.Settings.SqlResultsDownload().Get(ctx, settings.GetSqlResultsDownloadRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.SqlResultsDownload) (string, error) {
		t.SettingName = "default"
		t.BooleanVal.ForceSendFields = []string{"Value"}
		res, err := w.Settings.SqlResultsDownload().Update(ctx, settings.UpdateSqlResultsDownloadRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    sqlResultsDownloadFieldMask,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
	deleteFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error) {
        res, err := w.Settings.SqlResultsDownload().Delete(ctx, settings.DeleteSqlResultsDownloadRequest{
            Etag: etag,
        })
        if err != nil {
            return "", err
        }
        return res.Etag, err
    },
}