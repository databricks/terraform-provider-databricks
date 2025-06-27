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

var testEnhancedSecurityMonitoringSetting = AllSettingsResources()["enhanced_security_monitoring_workspace"]

func TestQueryCreateEnhancedSecurityMonitoringSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockEnhancedSecurityMonitoringAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateEnhancedSecurityMonitoringSettingRequest{
				AllowMissing: true,
				FieldMask:    enhancedSecurityMonitoringFieldMask,
				Setting: settings.EnhancedSecurityMonitoringSetting{
					Etag: "",
					EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
						IsEnabled:       true,
						ForceSendFields: []string{"IsEnabled"},
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
			e.Update(mock.Anything, settings.UpdateEnhancedSecurityMonitoringSettingRequest{
				AllowMissing: true,
				FieldMask:    enhancedSecurityMonitoringFieldMask,
				Setting: settings.EnhancedSecurityMonitoringSetting{
					Etag: "etag1",
					EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
						IsEnabled:       true,
						ForceSendFields: []string{"IsEnabled"},
					},
					SettingName: "default",
				},
			}).Return(&settings.EnhancedSecurityMonitoringSetting{
				Etag: "etag2",
				EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
					IsEnabled: true,
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetEnhancedSecurityMonitoringSettingRequest{
				Etag: "etag2",
			}).Return(&settings.EnhancedSecurityMonitoringSetting{
				Etag: "etag2",
				EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
					IsEnabled: true,
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testEnhancedSecurityMonitoringSetting,
		Create:   true,
		HCL: `
			enhanced_security_monitoring_workspace {
				is_enabled = true
			}
		`,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("enhanced_security_monitoring_workspace.0.is_enabled"))
}

func TestQueryReadEnhancedSecurityMonitoringSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockEnhancedSecurityMonitoringAPI().EXPECT().Get(mock.Anything, settings.GetEnhancedSecurityMonitoringSettingRequest{
				Etag: "etag1",
			}).Return(&settings.EnhancedSecurityMonitoringSetting{
				Etag: "etag2",
				EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
					IsEnabled: true,
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testEnhancedSecurityMonitoringSetting,
		Read:     true,
		HCL: `
			enhanced_security_monitoring_workspace {
				is_enabled = true
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	res := d.Get("enhanced_security_monitoring_workspace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["is_enabled"])
}

func TestQueryUpdateEnhancedSecurityMonitoringSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockEnhancedSecurityMonitoringAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateEnhancedSecurityMonitoringSettingRequest{
				AllowMissing: true,
				FieldMask:    enhancedSecurityMonitoringFieldMask,
				Setting: settings.EnhancedSecurityMonitoringSetting{
					Etag: "etag1",
					EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
						IsEnabled:       true,
						ForceSendFields: []string{"IsEnabled"},
					},
					SettingName: "default",
				},
			}).Return(&settings.EnhancedSecurityMonitoringSetting{
				Etag: "etag2",
				EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
					IsEnabled:       true,
					ForceSendFields: []string{"IsEnabled"},
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetEnhancedSecurityMonitoringSettingRequest{
				Etag: "etag2",
			}).Return(&settings.EnhancedSecurityMonitoringSetting{
				Etag: "etag2",
				EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
					IsEnabled: true,
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testEnhancedSecurityMonitoringSetting,
		Update:   true,
		HCL: `
			enhanced_security_monitoring_workspace {
				is_enabled = true
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	res := d.Get("enhanced_security_monitoring_workspace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["is_enabled"])
}

func TestQueryUpdateEnhancedSecurityMonitoringSettingWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockEnhancedSecurityMonitoringAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateEnhancedSecurityMonitoringSettingRequest{
				AllowMissing: true,
				FieldMask:    enhancedSecurityMonitoringFieldMask,
				Setting: settings.EnhancedSecurityMonitoringSetting{
					Etag: "etag1",
					EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
						IsEnabled:       true,
						ForceSendFields: []string{"IsEnabled"},
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
			e.Update(mock.Anything, settings.UpdateEnhancedSecurityMonitoringSettingRequest{
				AllowMissing: true,
				FieldMask:    enhancedSecurityMonitoringFieldMask,
				Setting: settings.EnhancedSecurityMonitoringSetting{
					Etag: "etag2",
					EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
						IsEnabled:       true,
						ForceSendFields: []string{"IsEnabled"},
					},
					SettingName: "default",
				},
			}).Return(&settings.EnhancedSecurityMonitoringSetting{
				Etag: "etag3",
				EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
					IsEnabled: true,
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetEnhancedSecurityMonitoringSettingRequest{
				Etag: "etag3",
			}).Return(&settings.EnhancedSecurityMonitoringSetting{
				Etag: "etag3",
				EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
					IsEnabled: true,
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testEnhancedSecurityMonitoringSetting,
		Update:   true,
		HCL: `
			enhanced_security_monitoring_workspace {
				is_enabled = true
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag3", d.Get(etagAttrName).(string))
	res := d.Get("enhanced_security_monitoring_workspace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["is_enabled"])
}

func TestQueryDeleteEnhancedSecurityMonitoringSetting(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockEnhancedSecurityMonitoringAPI().EXPECT().Update(mock.Anything, settings.UpdateEnhancedSecurityMonitoringSettingRequest{
				AllowMissing: true,
				FieldMask:    enhancedSecurityMonitoringFieldMask,
				Setting: settings.EnhancedSecurityMonitoringSetting{
					Etag:        "etag1",
					SettingName: "default",
					EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
						IsEnabled:       false,
						ForceSendFields: []string{"IsEnabled"},
					},
				},
			}).Return(&settings.EnhancedSecurityMonitoringSetting{
				EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
					IsEnabled: false,
				},
				Etag: "etag2",
			}, nil)
		},
		Resource: testEnhancedSecurityMonitoringSetting,
		Delete:   true,
		HCL: `
			enhanced_security_monitoring_workspace {
				is_enabled = false
			}
		etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         defaultSettingId,
		etagAttrName: "etag2",
	})
}

func TestQueryDeleteEnhancedSecurityMonitoringSettingWithConflict(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockEnhancedSecurityMonitoringAPI().EXPECT().Update(mock.Anything, settings.UpdateEnhancedSecurityMonitoringSettingRequest{
				AllowMissing: true,
				FieldMask:    enhancedSecurityMonitoringFieldMask,
				Setting: settings.EnhancedSecurityMonitoringSetting{
					Etag:        "etag1",
					SettingName: "default",
					EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
						IsEnabled:       false,
						ForceSendFields: []string{"IsEnabled"},
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
			w.GetMockEnhancedSecurityMonitoringAPI().EXPECT().Update(mock.Anything, settings.UpdateEnhancedSecurityMonitoringSettingRequest{
				AllowMissing: true,
				FieldMask:    enhancedSecurityMonitoringFieldMask,
				Setting: settings.EnhancedSecurityMonitoringSetting{
					Etag:        "etag2",
					SettingName: "default",
					EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
						IsEnabled:       false,
						ForceSendFields: []string{"IsEnabled"},
					},
				},
			}).Return(&settings.EnhancedSecurityMonitoringSetting{
				EnhancedSecurityMonitoringWorkspace: settings.EnhancedSecurityMonitoring{
					IsEnabled: false,
				},
				Etag: "etag3",
			}, nil)
		},
		Resource: testEnhancedSecurityMonitoringSetting,
		HCL: `
			enhanced_security_monitoring_workspace {
				is_enabled = true
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
