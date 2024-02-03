package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

// Restrict Workspace Admins setting
var restrictWsAdminsSetting = workspaceSetting[settings.RestrictWorkspaceAdminsSetting]{
	settingStruct: settings.RestrictWorkspaceAdminsSetting{},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.RestrictWorkspaceAdminsSetting, error) {
		return w.Settings.ReadRestrictWorkspaceAdmins(ctx, settings.ReadRestrictWorkspaceAdminsRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.RestrictWorkspaceAdminsSetting) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.UpdateRestrictWorkspaceAdmins(ctx, settings.UpdateDefaultWorkspaceNamespaceRequest{
			AllowMissing: true,
			Setting:      &t,
			FieldMask:    "restrict_workspace_admins",
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
	deleteFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error) {
		res, err := w.Settings.DeleteRestrictWorkspaceAdmins(ctx, settings.DeleteDefaultWorkspaceNamespaceRequest{
			Etag: etag,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
