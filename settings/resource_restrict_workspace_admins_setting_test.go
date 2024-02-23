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
						"etag": "etag1",
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
			e.GetRestrictWorkspaceAdminsSetting(mock.Anything, settings.GetRestrictWorkspaceAdminsSettingRequest{
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

	assert.Equal(t, "etag2", d.Id())
	assert.Equal(t, "RESTRICT_TOKENS_AND_JOB_RUN_AS", d.Get("restrict_workspace_admins.0.status"))
}

func TestQueryReadRestrictWsAdminsSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockSettingsAPI().EXPECT().GetRestrictWorkspaceAdminsSetting(mock.Anything, settings.GetRestrictWorkspaceAdminsSettingRequest{
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
		`,
		ID: "etag1",
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, "etag2", d.Id())
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
			e.GetRestrictWorkspaceAdminsSetting(mock.Anything, settings.GetRestrictWorkspaceAdminsSettingRequest{
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
		`,
		ID: "etag1",
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, "etag2", d.Id())
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
						"etag": "etag2",
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
			e.GetRestrictWorkspaceAdminsSetting(mock.Anything, settings.GetRestrictWorkspaceAdminsSettingRequest{
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
		`,
		ID: "etag1",
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, "etag3", d.Id())
	res := d.Get("restrict_workspace_admins").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "RESTRICT_TOKENS_AND_JOB_RUN_AS", res["status"])
}

func TestQueryDeleteRestrictWsAdminsSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockSettingsAPI().EXPECT().DeleteRestrictWorkspaceAdminsSetting(mock.Anything, settings.DeleteRestrictWorkspaceAdminsSettingRequest{
				Etag: "etag1",
			}).Return(&settings.DeleteRestrictWorkspaceAdminsSettingResponse{
				Etag: "etag2",
			}, nil)
		},
		Resource: testRestrictWsAdminsSetting,
		Delete:   true,
		ID:       "etag1",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "etag2", d.Id())
}

func TestQueryDeleteRestrictWsAdminsSettingWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockSettingsAPI().EXPECT().DeleteRestrictWorkspaceAdminsSetting(mock.Anything, settings.DeleteRestrictWorkspaceAdminsSettingRequest{
				Etag: "etag1",
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "RESOURCE_CONFLICT",
				StatusCode: 409,
				Message:    "SomeMessage",
				Details: []apierr.ErrorDetail{{
					Type: "type.googleapis.com/google.rpc.ErrorInfo",
					Metadata: map[string]string{
						"etag": "etag2",
					},
				}},
			})
			w.GetMockSettingsAPI().EXPECT().DeleteRestrictWorkspaceAdminsSetting(mock.Anything, settings.DeleteRestrictWorkspaceAdminsSettingRequest{
				Etag: "etag2",
			}).Return(&settings.DeleteRestrictWorkspaceAdminsSettingResponse{
				Etag: "etag3",
			}, nil)
		},
		Resource: testRestrictWsAdminsSetting,
		Delete:   true,
		ID:       "etag1",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "etag3", d.Id())
}
