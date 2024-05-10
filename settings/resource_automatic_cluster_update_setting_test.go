package settings

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testAutomaticClusterUpdateSetting = AllSettingsResources()["automatic_cluster_update_workspace"]

func TestQueryCreateAutomaticClusterUpdateSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAutomaticClusterUpdateAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateAutomaticClusterUpdateSettingRequest{
				AllowMissing: true,
				FieldMask:    automaticClusterUpdateFieldMask,
				Setting: settings.AutomaticClusterUpdateSetting{
					Etag: "",
					AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
						Enabled: true,
					},
					SettingName: "default",
				},
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "SomeMessage",
				Details: []apierr.ErrorDetail{{
					Type: "type.googleapis.com/google.rpc.ErrorInfo",
					Metadata: map[string]string{
						etagAttrName: "etag1",
					},
				}},
			})
			e.Update(mock.Anything, settings.UpdateAutomaticClusterUpdateSettingRequest{
				AllowMissing: true,
				FieldMask:    automaticClusterUpdateFieldMask,
				Setting: settings.AutomaticClusterUpdateSetting{
					Etag: "etag1",
					AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
						Enabled: true,
					},
					SettingName: "default",
				},
			}).Return(&settings.AutomaticClusterUpdateSetting{
				Etag: "etag2",
				AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
					Enabled: true,
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetAutomaticClusterUpdateSettingRequest{
				Etag: "etag2",
			}).Return(&settings.AutomaticClusterUpdateSetting{
				Etag: "etag2",
				AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
					Enabled: true,
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testAutomaticClusterUpdateSetting,
		Create:   true,
		HCL: `
			automatic_cluster_update_workspace {
				enabled = true
			}
		`,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("automatic_cluster_update_workspace.0.enabled"))
}

func TestQueryReadAutomaticClusterUpdateSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockAutomaticClusterUpdateAPI().EXPECT().Get(mock.Anything, settings.GetAutomaticClusterUpdateSettingRequest{
				Etag: "etag1",
			}).Return(&settings.AutomaticClusterUpdateSetting{
				Etag: "etag2",
				AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
					Enabled:                         true,
					RestartEvenIfNoUpdatesAvailable: true,
					MaintenanceWindow: &settings.ClusterAutoRestartMessageMaintenanceWindow{
						WeekDayBasedSchedule: &settings.ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{
							DayOfWeek: "MONDAY",
							Frequency: "EVERY_WEEK",
							WindowStartTime: &settings.ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{
								Hours:   1,
								Minutes: 0,
							},
						},
					},
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testAutomaticClusterUpdateSetting,
		Read:     true,
		HCL: `
			automatic_cluster_update_workspace {
				enabled = true
				restart_even_if_no_updates_available = true
				maintenance_window {
					week_day_based_schedule {
						day_of_week = "MONDAY"
						frequency = "EVERY_WEEK"
						window_start_time {
							hours = 1
							minutes = 0
						}
					}
				}
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	res := d.Get("automatic_cluster_update_workspace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["restart_even_if_no_updates_available"])
	week_day_based_schedule := res["maintenance_window"].([]interface{})[0].(map[string]interface{})["week_day_based_schedule"].([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "MONDAY", week_day_based_schedule["day_of_week"])
	window_start_time := week_day_based_schedule["window_start_time"].([]interface{})[0].(map[string]interface{})
	assert.Equal(t, 1, window_start_time["hours"])
}

func TestQueryUpdateAutomaticClusterUpdateSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAutomaticClusterUpdateAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateAutomaticClusterUpdateSettingRequest{
				AllowMissing: true,
				FieldMask:    automaticClusterUpdateFieldMask,
				Setting: settings.AutomaticClusterUpdateSetting{
					Etag: "etag1",
					AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
						Enabled:                         true,
						RestartEvenIfNoUpdatesAvailable: true,
						MaintenanceWindow: &settings.ClusterAutoRestartMessageMaintenanceWindow{
							WeekDayBasedSchedule: &settings.ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{
								DayOfWeek: "MONDAY",
								Frequency: "EVERY_WEEK",
								WindowStartTime: &settings.ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{
									Hours:   1,
									Minutes: 30,
								},
							},
						},
					},
					SettingName: "default",
				},
			}).Return(&settings.AutomaticClusterUpdateSetting{
				Etag: "etag2",
				AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
					Enabled: true,
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetAutomaticClusterUpdateSettingRequest{
				Etag: "etag2",
			}).Return(&settings.AutomaticClusterUpdateSetting{
				Etag: "etag2",
				AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
					Enabled:                         true,
					RestartEvenIfNoUpdatesAvailable: true,
					MaintenanceWindow: &settings.ClusterAutoRestartMessageMaintenanceWindow{
						WeekDayBasedSchedule: &settings.ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{
							DayOfWeek: "MONDAY",
							Frequency: "EVERY_WEEK",
							WindowStartTime: &settings.ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{
								Hours:   1,
								Minutes: 30,
							},
						},
					},
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testAutomaticClusterUpdateSetting,
		Update:   true,
		HCL: `
			automatic_cluster_update_workspace {
				enabled = true
				restart_even_if_no_updates_available = true
				maintenance_window {
					week_day_based_schedule {
						day_of_week = "MONDAY"
						frequency = "EVERY_WEEK"
						window_start_time {
							hours = 1
							minutes = 30
						}
					}
				}
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	res := d.Get("automatic_cluster_update_workspace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["restart_even_if_no_updates_available"])
	week_day_based_schedule := res["maintenance_window"].([]interface{})[0].(map[string]interface{})["week_day_based_schedule"].([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "MONDAY", week_day_based_schedule["day_of_week"])
	window_start_time := week_day_based_schedule["window_start_time"].([]interface{})[0].(map[string]interface{})
	assert.Equal(t, 1, window_start_time["hours"])
}

func TestQueryUpdateAutomaticClusterUpdateSettingWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAutomaticClusterUpdateAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateAutomaticClusterUpdateSettingRequest{
				AllowMissing: true,
				FieldMask:    automaticClusterUpdateFieldMask,
				Setting: settings.AutomaticClusterUpdateSetting{
					Etag: "etag1",
					AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
						Enabled:                         true,
						RestartEvenIfNoUpdatesAvailable: true,
					},
					SettingName: "default",
				},
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "RESOURCE_CONFLICT",
				StatusCode: 409,
				Message:    "SomeMessage",
				Details: []apierr.ErrorDetail{{
					Type: "type.googleapis.com/google.rpc.ErrorInfo",
					Metadata: map[string]string{
						etagAttrName: "etag2",
					},
				}},
			})
			e.Update(mock.Anything, settings.UpdateAutomaticClusterUpdateSettingRequest{
				AllowMissing: true,
				FieldMask:    automaticClusterUpdateFieldMask,
				Setting: settings.AutomaticClusterUpdateSetting{
					Etag: "etag2",
					AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
						Enabled:                         true,
						RestartEvenIfNoUpdatesAvailable: true,
					},
					SettingName: "default",
				},
			}).Return(&settings.AutomaticClusterUpdateSetting{
				Etag: "etag3",
				AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
					Enabled:                         true,
					RestartEvenIfNoUpdatesAvailable: true,
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetAutomaticClusterUpdateSettingRequest{
				Etag: "etag3",
			}).Return(&settings.AutomaticClusterUpdateSetting{
				Etag: "etag3",
				AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
					Enabled:                         true,
					RestartEvenIfNoUpdatesAvailable: true,
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testAutomaticClusterUpdateSetting,
		Update:   true,
		HCL: `
			automatic_cluster_update_workspace {
				enabled = true
				restart_even_if_no_updates_available = true
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag3", d.Get(etagAttrName).(string))
	res := d.Get("automatic_cluster_update_workspace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["restart_even_if_no_updates_available"])
}

func TestQueryDeleteAutomaticClusterUpdateSetting(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockAutomaticClusterUpdateAPI().EXPECT().Update(mock.Anything, settings.UpdateAutomaticClusterUpdateSettingRequest{
				AllowMissing: true,
				FieldMask:    automaticClusterUpdateFieldMask,
				Setting: settings.AutomaticClusterUpdateSetting{
					Etag:        "etag1",
					SettingName: "default",
					AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
						Enabled: false,
					},
				},
			}).Return(&settings.AutomaticClusterUpdateSetting{
				AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
					Enabled: false,
				},
				Etag: "etag2",
			}, nil)
		},
		Resource: testAutomaticClusterUpdateSetting,
		Delete:   true,
		HCL: `
			automatic_cluster_update_workspace {
				enabled = true
			}
		etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         defaultSettingId,
		etagAttrName: "etag2",
	})
}

func TestQueryDeleteAutomaticClusterUpdateSettingWithConflict(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockAutomaticClusterUpdateAPI().EXPECT().Update(mock.Anything, settings.UpdateAutomaticClusterUpdateSettingRequest{
				AllowMissing: true,
				FieldMask:    automaticClusterUpdateFieldMask,
				Setting: settings.AutomaticClusterUpdateSetting{
					Etag:        "etag1",
					SettingName: "default",
					AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
						Enabled: false,
					},
				},
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "RESOURCE_CONFLICT",
				StatusCode: 409,
				Message:    "SomeMessage",
				Details: []apierr.ErrorDetail{{
					Type: "type.googleapis.com/google.rpc.ErrorInfo",
					Metadata: map[string]string{
						etagAttrName: "etag2",
					},
				}},
			})
			w.GetMockAutomaticClusterUpdateAPI().EXPECT().Update(mock.Anything, settings.UpdateAutomaticClusterUpdateSettingRequest{
				AllowMissing: true,
				FieldMask:    automaticClusterUpdateFieldMask,
				Setting: settings.AutomaticClusterUpdateSetting{
					Etag:        "etag2",
					SettingName: "default",
					AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
						Enabled: false,
					},
				},
			}).Return(&settings.AutomaticClusterUpdateSetting{
				AutomaticClusterUpdateWorkspace: settings.ClusterAutoRestartMessage{
					Enabled: false,
				},
				Etag: "etag3",
			}, nil)
		},
		Resource: testAutomaticClusterUpdateSetting,
		HCL: `
			automatic_cluster_update_workspace {
				enabled = true
			}
		etag = "etag1"
		`,
		Delete: true,
		ID:     defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         defaultSettingId,
		etagAttrName: "etag3",
	})
}