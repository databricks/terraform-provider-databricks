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

var testRestrictWsAdminsSetting = AllSettingsResources()["restrict_workspace_admins"]

func TestQueryCreateRestrictWsAdminsSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSettingsAPI().EXPECT()
			e.UpdateRestrictWorkspaceAdminsSetting(mock.Anything, settings.UpdateRestrictWorkspaceAdminsSettingRequest{
				AllowMissing: true,
				FieldMask:    "restrict_workspace_admins.status",
				Setting: settings.RestrictWorkspaceAdminsSetting{
					Etag: "",
					RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
						Status: "RESTRICT_TOKENS_AND_JOB_RUN_AS",
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
			e.UpdateRestrictWorkspaceAdminsSetting(mock.Anything, settings.UpdateRestrictWorkspaceAdminsSettingRequest{
				AllowMissing: true,
				FieldMask:    "restrict_workspace_admins.status",
				Setting: settings.RestrictWorkspaceAdminsSetting{
					Etag: "etag1",
					RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
						Status: "RESTRICT_TOKENS_AND_JOB_RUN_AS",
					},
					SettingName: "default",
				},
			}).Return(&settings.RestrictWorkspaceAdminsSetting{
				Etag: "etag2",
				RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
					Status: "RESTRICT_TOKENS_AND_JOB_RUN_AS",
				},
				SettingName: "default",
			}, nil)
			e.GetRestrictWorkspaceAdminsSetting(mock.Anything, settings.GetRestrictWorkspaceAdminRequest{
				Etag: "etag2",
			}).Return(&settings.RestrictWorkspaceAdminsSetting{
				Etag: "etag2",
				RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
					Status: "RESTRICT_TOKENS_AND_JOB_RUN_AS",
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testRestrictWsAdminsSetting,
		Create:   true,
		HCL: `
			restrict_workspace_admins {
				status = "RESTRICT_TOKENS_AND_JOB_RUN_AS"
			}
		`,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, "RESTRICT_TOKENS_AND_JOB_RUN_AS", d.Get("restrict_workspace_admins.0.status"))
}

func TestQueryReadRestrictWsAdminsSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockSettingsAPI().EXPECT().GetRestrictWorkspaceAdminsSetting(mock.Anything, settings.GetRestrictWorkspaceAdminRequest{
				Etag: "etag1",
			}).Return(&settings.RestrictWorkspaceAdminsSetting{
				Etag: "etag2",
				RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
					Status: "RESTRICT_TOKENS_AND_JOB_RUN_AS",
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testRestrictWsAdminsSetting,
		Read:     true,
		HCL: `
			restrict_workspace_admins {
				status = "RESTRICT_TOKENS_AND_JOB_RUN_AS"
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	res := d.Get("restrict_workspace_admins").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "RESTRICT_TOKENS_AND_JOB_RUN_AS", res["status"])
}

func TestQueryUpdateRestrictWsAdminsSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSettingsAPI().EXPECT()
			e.UpdateRestrictWorkspaceAdminsSetting(mock.Anything, settings.UpdateRestrictWorkspaceAdminsSettingRequest{
				AllowMissing: true,
				FieldMask:    "restrict_workspace_admins.status",
				Setting: settings.RestrictWorkspaceAdminsSetting{
					Etag: "etag1",
					RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
						Status: "ALLOW_ALL",
					},
					SettingName: "default",
				},
			}).Return(&settings.RestrictWorkspaceAdminsSetting{
				Etag: "etag2",
				RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
					Status: "ALLOW_ALL",
				},
				SettingName: "default",
			}, nil)
			e.GetRestrictWorkspaceAdminsSetting(mock.Anything, settings.GetRestrictWorkspaceAdminRequest{
				Etag: "etag2",
			}).Return(&settings.RestrictWorkspaceAdminsSetting{
				Etag: "etag2",
				RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
					Status: "ALLOW_ALL",
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testRestrictWsAdminsSetting,
		Update:   true,
		HCL: `
			restrict_workspace_admins {
				status = "ALLOW_ALL"
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	res := d.Get("restrict_workspace_admins").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "ALLOW_ALL", res["status"])
}

func TestQueryUpdateRestrictWsAdminsSettingWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSettingsAPI().EXPECT()
			e.UpdateRestrictWorkspaceAdminsSetting(mock.Anything, settings.UpdateRestrictWorkspaceAdminsSettingRequest{
				AllowMissing: true,
				FieldMask:    "restrict_workspace_admins.status",
				Setting: settings.RestrictWorkspaceAdminsSetting{
					Etag: "etag1",
					RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
						Status: "RESTRICT_TOKENS_AND_JOB_RUN_AS",
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
			e.UpdateRestrictWorkspaceAdminsSetting(mock.Anything, settings.UpdateRestrictWorkspaceAdminsSettingRequest{
				AllowMissing: true,
				FieldMask:    "restrict_workspace_admins.status",
				Setting: settings.RestrictWorkspaceAdminsSetting{
					Etag: "etag2",
					RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
						Status: "RESTRICT_TOKENS_AND_JOB_RUN_AS",
					},
					SettingName: "default",
				},
			}).Return(&settings.RestrictWorkspaceAdminsSetting{
				Etag: "etag3",
				RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
					Status: "RESTRICT_TOKENS_AND_JOB_RUN_AS",
				},
				SettingName: "default",
			}, nil)
			e.GetRestrictWorkspaceAdminsSetting(mock.Anything, settings.GetRestrictWorkspaceAdminRequest{
				Etag: "etag3",
			}).Return(&settings.RestrictWorkspaceAdminsSetting{
				Etag: "etag3",
				RestrictWorkspaceAdmins: settings.RestrictWorkspaceAdminsMessage{
					Status: "RESTRICT_TOKENS_AND_JOB_RUN_AS",
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testRestrictWsAdminsSetting,
		Update:   true,
		HCL: `
			restrict_workspace_admins {
				status = "RESTRICT_TOKENS_AND_JOB_RUN_AS"
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag3", d.Get(etagAttrName).(string))
	res := d.Get("restrict_workspace_admins").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "RESTRICT_TOKENS_AND_JOB_RUN_AS", res["status"])
}

func TestQueryDeleteRestrictWsAdminsSetting(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockSettingsAPI().EXPECT().DeleteRestrictWorkspaceAdminsSetting(mock.Anything, settings.DeleteRestrictWorkspaceAdminRequest{
				Etag: "etag1",
			}).Return(&settings.DeleteRestrictWorkspaceAdminsSettingResponse{
				Etag: "etag2",
			}, nil)
		},
		Resource: testRestrictWsAdminsSetting,
		Delete:   true,
		HCL: `
		restrict_workspace_admins {
			status = "RESTRICT_TOKENS_AND_JOB_RUN_AS"
		}
		etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         defaultSettingId,
		etagAttrName: "etag2",
	})
}

func TestQueryDeleteRestrictWsAdminsSettingWithConflict(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockSettingsAPI().EXPECT().DeleteRestrictWorkspaceAdminsSetting(mock.Anything, settings.DeleteRestrictWorkspaceAdminRequest{
				Etag: "etag1",
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
			w.GetMockSettingsAPI().EXPECT().DeleteRestrictWorkspaceAdminsSetting(mock.Anything, settings.DeleteRestrictWorkspaceAdminRequest{
				Etag: "etag2",
			}).Return(&settings.DeleteRestrictWorkspaceAdminsSettingResponse{
				Etag: "etag3",
			}, nil)
		},
		Resource: testRestrictWsAdminsSetting,
		HCL: `
		restrict_workspace_admins {
			status = "RESTRICT_TOKENS_AND_JOB_RUN_AS"
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
