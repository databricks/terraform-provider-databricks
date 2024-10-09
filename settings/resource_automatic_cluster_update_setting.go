package settings

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Automatic Cluster Update setting
var automaticClusterUpdateFieldMask = strings.Join([]string{
	"automatic_cluster_update_workspace.enabled",
	"automatic_cluster_update_workspace.restart_even_if_no_updates_available",
	"automatic_cluster_update_workspace.maintenance_window.week_day_based_schedule.day_of_week",
	"automatic_cluster_update_workspace.maintenance_window.week_day_based_schedule.frequency",
	"automatic_cluster_update_workspace.maintenance_window.week_day_based_schedule.window_start_time.hours",
	"automatic_cluster_update_workspace.maintenance_window.week_day_based_schedule.window_start_time.minutes",
}, ",")

var automaticClusterUpdateSetting = workspaceSetting[settings.AutomaticClusterUpdateSetting]{
	settingStruct: settings.AutomaticClusterUpdateSetting{},
	customizeSchemaFunc: func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "automatic_cluster_update_workspace", "enablement_details").SetReadOnly()
		common.CustomizeSchemaPath(s, "automatic_cluster_update_workspace", "enabled").SetRequired()
		common.CustomizeSchemaPath(s, "automatic_cluster_update_workspace", "maintenance_window",
			"week_day_based_schedule", "window_start_time", "hours").SetRequired()
		common.CustomizeSchemaPath(s, "automatic_cluster_update_workspace", "maintenance_window",
			"week_day_based_schedule", "day_of_week").SetRequired()
		common.CustomizeSchemaPath(s, "automatic_cluster_update_workspace", "maintenance_window",
			"week_day_based_schedule", "frequency").SetRequired()
		common.CustomizeSchemaPath(s, "automatic_cluster_update_workspace", "maintenance_window",
			"week_day_based_schedule", "window_start_time", "minutes").SetRequired()
		return s
	},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.AutomaticClusterUpdateSetting, error) {
		return w.Settings.AutomaticClusterUpdate().Get(ctx, settings.GetAutomaticClusterUpdateSettingRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.AutomaticClusterUpdateSetting) (string, error) {
		t.SettingName = "default"
		t.AutomaticClusterUpdateWorkspace.ForceSendFields = []string{"Enabled", "RestartEvenIfNoUpdatesAvailable"}
		if t.AutomaticClusterUpdateWorkspace.MaintenanceWindow != nil &&
			t.AutomaticClusterUpdateWorkspace.MaintenanceWindow.WeekDayBasedSchedule != nil &&
			t.AutomaticClusterUpdateWorkspace.MaintenanceWindow.WeekDayBasedSchedule.WindowStartTime != nil {
			t.AutomaticClusterUpdateWorkspace.MaintenanceWindow.WeekDayBasedSchedule.WindowStartTime.ForceSendFields = []string{"Hours", "Minutes"}
		}
		res, err := w.Settings.AutomaticClusterUpdate().Update(ctx, settings.UpdateAutomaticClusterUpdateSettingRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    automaticClusterUpdateFieldMask,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
	deleteFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error) {
		res, err := w.Settings.AutomaticClusterUpdate().Update(ctx, settings.UpdateAutomaticClusterUpdateSettingRequest{
			AllowMissing: true,
			Setting: settings.AutomaticClusterUpdateSetting{
				Etag:        etag,
				SettingName: "default",
				AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
					Enabled:         false,
					ForceSendFields: []string{"Enabled"},
				},
			},
			FieldMask: automaticClusterUpdateFieldMask,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
