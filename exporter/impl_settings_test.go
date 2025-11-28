package exporter

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/settingsv2"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceSettingV2Export(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		emptyDestinationNotficationsList,
		{
			Method:   "GET",
			Resource: "/api/2.1/settings-metadata?",
			Response: settingsv2.ListWorkspaceSettingsMetadataResponse{
				SettingsMetadata: []settingsv2.SettingsMetadata{
					{
						Name:        "features.personalCompute.enabled",
						Description: "Personal Compute feature enablement",
					},
					{
						Name:        "features.restrictWorkspaceAdmins.enabled",
						Description: "Restrict workspace admins",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/settings/features.personalCompute.enabled?",
			Response: settingsv2.Setting{
				Name: "features.personalCompute.enabled",
				EffectivePersonalCompute: &settingsv2.PersonalComputeMessage{
					Value: settingsv2.PersonalComputeMessagePersonalComputeMessageEnumOn,
				},
				PersonalCompute: &settingsv2.PersonalComputeMessage{
					Value: settingsv2.PersonalComputeMessagePersonalComputeMessageEnumOn,
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/settings/features.restrictWorkspaceAdmins.enabled?",
			Response: settingsv2.Setting{
				Name: "features.restrictWorkspaceAdmins.enabled",
				EffectiveRestrictWorkspaceAdmins: &settingsv2.RestrictWorkspaceAdminsMessage{
					Status: settingsv2.RestrictWorkspaceAdminsMessageStatusAllowAll,
				},
				RestrictWorkspaceAdmins: &settingsv2.RestrictWorkspaceAdminsMessage{
					Status: settingsv2.RestrictWorkspaceAdminsMessageStatusAllowAll,
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("settings")
		ic.enableServices("settings")

		err := ic.Run()
		assert.NoError(t, err)

		// Verify that the workspace settings were generated in the Terraform code
		content, err := os.ReadFile(tmpDir + "/settings.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))

		// Check that the resources are generated with expected fields
		assert.Contains(t, contentStr, `resource "databricks_workspace_setting_v2" "features_personalcompute_enabled"`)
		assert.Contains(t, contentStr, `name = "features.personalCompute.enabled"`)
		assert.Contains(t, contentStr, `personal_compute`)
		assert.Contains(t, contentStr, `value = "ON"`)

		assert.Contains(t, contentStr, `resource "databricks_workspace_setting_v2" "features_restrictworkspaceadmins_enabled"`)
		assert.Contains(t, contentStr, `name = "features.restrictWorkspaceAdmins.enabled"`)
		assert.Contains(t, contentStr, `restrict_workspace_admins`)
		assert.Contains(t, contentStr, `status = "ALLOW_ALL"`)
	})
}

func TestWorkspaceSettingV2ExportWithBooleanValue(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		emptyDestinationNotficationsList,
		{
			Method:   "GET",
			Resource: "/api/2.1/settings-metadata?",
			Response: settingsv2.ListWorkspaceSettingsMetadataResponse{
				SettingsMetadata: []settingsv2.SettingsMetadata{
					{
						Name:        "example.boolean.setting",
						Description: "Example boolean setting",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/settings/example.boolean.setting?",
			Response: settingsv2.Setting{
				Name: "example.boolean.setting",
				EffectiveBooleanVal: &settingsv2.BooleanMessage{
					Value: true,
				},
				BooleanVal: &settingsv2.BooleanMessage{
					Value: true,
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("settings")
		ic.enableServices("settings")

		err := ic.Run()
		assert.NoError(t, err)

		// Verify that the boolean setting was generated correctly
		content, err := os.ReadFile(tmpDir + "/settings.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))

		// Check that the resource is generated with boolean_val from effective_boolean_val
		assert.Contains(t, contentStr, `resource "databricks_workspace_setting_v2" "example_boolean_setting"`)
		assert.Contains(t, contentStr, `name = "example.boolean.setting"`)
		assert.Contains(t, contentStr, `boolean_val`)
		assert.Contains(t, contentStr, `value = true`)
	})
}

func TestWorkspaceSettingV2ExportWithStringValue(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		emptyDestinationNotficationsList,
		{
			Method:   "GET",
			Resource: "/api/2.1/settings-metadata?",
			Response: settingsv2.ListWorkspaceSettingsMetadataResponse{
				SettingsMetadata: []settingsv2.SettingsMetadata{
					{
						Name:        "example.string.setting",
						Description: "Example string setting",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/settings/example.string.setting?",
			Response: settingsv2.Setting{
				Name: "example.string.setting",
				EffectiveStringVal: &settingsv2.StringMessage{
					Value: "test-value",
				},
				StringVal: &settingsv2.StringMessage{
					Value: "test-value",
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("settings")
		ic.enableServices("settings")

		err := ic.Run()
		assert.NoError(t, err)

		// Verify that the string setting was generated correctly
		content, err := os.ReadFile(tmpDir + "/settings.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))

		// Check that the resource is generated with string_val from effective_string_val
		assert.Contains(t, contentStr, `resource "databricks_workspace_setting_v2" "example_string_setting"`)
		assert.Contains(t, contentStr, `name = "example.string.setting"`)
		assert.Contains(t, contentStr, `string_val`)
		assert.Contains(t, contentStr, `value = "test-value"`)
	})
}

func TestWorkspaceSettingV2ExportWithIntegerValue(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		emptyDestinationNotficationsList,
		{
			Method:   "GET",
			Resource: "/api/2.1/settings-metadata?",
			Response: settingsv2.ListWorkspaceSettingsMetadataResponse{
				SettingsMetadata: []settingsv2.SettingsMetadata{
					{
						Name:        "example.integer.setting",
						Description: "Example integer setting",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/settings/example.integer.setting?",
			Response: settingsv2.Setting{
				Name: "example.integer.setting",
				EffectiveIntegerVal: &settingsv2.IntegerMessage{
					Value: 42,
				},
				IntegerVal: &settingsv2.IntegerMessage{
					Value: 42,
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("settings")
		ic.enableServices("settings")

		err := ic.Run()
		assert.NoError(t, err)

		// Verify that the integer setting was generated correctly
		content, err := os.ReadFile(tmpDir + "/settings.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))

		// Check that the resource is generated with integer_val from effective_integer_val
		assert.Contains(t, contentStr, `resource "databricks_workspace_setting_v2" "example_integer_setting"`)
		assert.Contains(t, contentStr, `name = "example.integer.setting"`)
		assert.Contains(t, contentStr, `integer_val`)
		assert.Contains(t, contentStr, `value = 42`)
	})
}

func TestWorkspaceSettingV2ExportWithAutomaticClusterUpdate(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		emptyDestinationNotficationsList,
		{
			Method:   "GET",
			Resource: "/api/2.1/settings-metadata?",
			Response: settingsv2.ListWorkspaceSettingsMetadataResponse{
				SettingsMetadata: []settingsv2.SettingsMetadata{
					{
						Name:        "automaticClusterUpdate.enabled",
						Description: "Automatic cluster update setting",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/settings/automaticClusterUpdate.enabled?",
			Response: settingsv2.Setting{
				Name: "automaticClusterUpdate.enabled",
				EffectiveAutomaticClusterUpdateWorkspace: &settingsv2.ClusterAutoRestartMessage{
					Enabled:                         true,
					RestartEvenIfNoUpdatesAvailable: false,
					MaintenanceWindow: &settingsv2.ClusterAutoRestartMessageMaintenanceWindow{
						WeekDayBasedSchedule: &settingsv2.ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{
							WindowStartTime: &settingsv2.ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{
								Hours:   2,
								Minutes: 30,
							},
							DayOfWeek: settingsv2.ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSunday,
						},
					},
				},
				AutomaticClusterUpdateWorkspace: &settingsv2.ClusterAutoRestartMessage{
					Enabled:                         true,
					RestartEvenIfNoUpdatesAvailable: false,
					MaintenanceWindow: &settingsv2.ClusterAutoRestartMessageMaintenanceWindow{
						WeekDayBasedSchedule: &settingsv2.ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{
							WindowStartTime: &settingsv2.ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{
								Hours:   2,
								Minutes: 30,
							},
							DayOfWeek: settingsv2.ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSunday,
						},
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("settings")
		ic.enableServices("settings")

		err := ic.Run()
		assert.NoError(t, err)

		// Verify that the automatic cluster update setting was generated correctly
		content, err := os.ReadFile(tmpDir + "/settings.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))

		// Check that the resource is generated with automatic_cluster_update_workspace from effective_*
		assert.Contains(t, contentStr, `resource "databricks_workspace_setting_v2" "automaticclusterupdate_enabled"`)
		assert.Contains(t, contentStr, `name = "automaticClusterUpdate.enabled"`)
		assert.Contains(t, contentStr, `automatic_cluster_update_workspace`)
		assert.Contains(t, contentStr, `enabled = true`)
		// Note: restart_even_if_no_updates_available = false is omitted because false is the zero value
		assert.Contains(t, contentStr, `maintenance_window`)
		assert.Contains(t, contentStr, `week_day_based_schedule`)
		assert.Contains(t, contentStr, `hours = 2`)
		assert.Contains(t, contentStr, `minutes = 30`)
		assert.Contains(t, contentStr, `day_of_week = "SUNDAY"`)
	})
}
