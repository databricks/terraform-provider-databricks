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
<<<<<<< HEAD
		return w.Settings.GetRestrictWorkspaceAdminsSetting(ctx, settings.GetRestrictWorkspaceAdminsSettingRequest{
=======
		return w.Settings.ReadRestrictWorkspaceAdmins(ctx, settings.ReadRestrictWorkspaceAdminsRequest{
>>>>>>> 0932ea2084ac64919cbe4929aba74b681a16be26
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.RestrictWorkspaceAdminsSetting) (string, error) {
		t.SettingName = "default"
<<<<<<< HEAD
		res, err := w.Settings.UpdateRestrictWorkspaceAdminsSetting(ctx, settings.UpdateRestrictWorkspaceAdminsSettingRequest{
			AllowMissing: true,
			Setting:      t,
=======
		res, err := w.Settings.UpdateRestrictWorkspaceAdmins(ctx, settings.UpdateDefaultWorkspaceNamespaceRequest{
			AllowMissing: true,
			Setting:      &t,
>>>>>>> 0932ea2084ac64919cbe4929aba74b681a16be26
			FieldMask:    "restrict_workspace_admins",
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
	deleteFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error) {
<<<<<<< HEAD
		res, err := w.Settings.DeleteRestrictWorkspaceAdminsSetting(ctx, settings.DeleteRestrictWorkspaceAdminsSettingRequest{
=======
		res, err := w.Settings.DeleteRestrictWorkspaceAdmins(ctx, settings.DeleteDefaultWorkspaceNamespaceRequest{
>>>>>>> 0932ea2084ac64919cbe4929aba74b681a16be26
			Etag: etag,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
