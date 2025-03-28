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

var testDisableLegacyAccess = AllSettingsResources()["disable_legacy_access"]

func TestCreateDisableLegacyAccess(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockDisableLegacyAccessAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDisableLegacyAccessRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_access.value",
				Setting: settings.DisableLegacyAccess{
					Etag: "",
					DisableLegacyAccess: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_access",
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
			e.Update(mock.Anything, settings.UpdateDisableLegacyAccessRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_access.value",
				Setting: settings.DisableLegacyAccess{
					Etag: "etag1",
					DisableLegacyAccess: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_access",
				},
			}).Return(&settings.DisableLegacyAccess{
				Etag: "etag2",
				DisableLegacyAccess: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_access",
			}, nil)
			e.Get(mock.Anything, settings.GetDisableLegacyAccessRequest{
				Etag: "etag2",
			}).Return(&settings.DisableLegacyAccess{
				Etag: "etag2",
				DisableLegacyAccess: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_access",
			}, nil)
		},
		Resource: testDisableLegacyAccess,
		Create:   true,
		HCL: `
			disable_legacy_access {
    			value = "true"
  			}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("disable_legacy_access.0.value"))
}

func TestReadDisableLegacyAccess(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDisableLegacyAccessAPI().EXPECT().Get(mock.Anything, settings.GetDisableLegacyAccessRequest{
				Etag: "etag1",
			}).Return(&settings.DisableLegacyAccess{
				Etag: "etag2",
				DisableLegacyAccess: settings.BooleanMessage{
					Value: false,
				},
				SettingName: "disable_legacy_access",
			}, nil)
		},
		Resource: testDisableLegacyAccess,
		Read:     true,
		HCL: `
			disable_legacy_access {
    			value = "false"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, false, d.Get("disable_legacy_access.0.value"))
}

func TestUpdateDisableLegacyAccess(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockDisableLegacyAccessAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDisableLegacyAccessRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_access.value",
				Setting: settings.DisableLegacyAccess{
					Etag: "etag1",
					DisableLegacyAccess: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_access",
				},
			}).Return(&settings.DisableLegacyAccess{
				Etag: "etag2",
				DisableLegacyAccess: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_access",
			}, nil)
			e.Get(mock.Anything, settings.GetDisableLegacyAccessRequest{
				Etag: "etag2",
			}).Return(&settings.DisableLegacyAccess{
				Etag: "etag2",
				DisableLegacyAccess: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_access",
			}, nil)
		},
		Resource: testDisableLegacyAccess,
		Update:   true,
		HCL: `
			disable_legacy_access {
    			value = "true"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("disable_legacy_access.0.value"))
}

func TestUpdateDisableLegacyAccessWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockDisableLegacyAccessAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDisableLegacyAccessRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_access.value",
				Setting: settings.DisableLegacyAccess{
					Etag: "etag1",
					DisableLegacyAccess: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_access",
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
			e.Update(mock.Anything, settings.UpdateDisableLegacyAccessRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_access.value",
				Setting: settings.DisableLegacyAccess{
					Etag: "etag2",
					DisableLegacyAccess: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_access",
				},
			}).Return(&settings.DisableLegacyAccess{
				Etag: "etag3",
				DisableLegacyAccess: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_access",
			}, nil)
			e.Get(mock.Anything, settings.GetDisableLegacyAccessRequest{
				Etag: "etag3",
			}).Return(&settings.DisableLegacyAccess{
				Etag: "etag3",
				DisableLegacyAccess: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_access",
			}, nil)
		},
		Resource: testDisableLegacyAccess,
		Update:   true,
		HCL: `
			disable_legacy_access {
    			value = "true"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag3", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("disable_legacy_access.0.value"))
}

func TestDeleteDisableLegacyAccess(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDisableLegacyAccessAPI().EXPECT().Delete(mock.Anything,
				settings.DeleteDisableLegacyAccessRequest{
					Etag: "etag1",
				}).Return(&settings.DeleteDisableLegacyAccessResponse{
				Etag: "etag2",
			}, nil)
		},
		Resource: testDisableLegacyAccess,
		Delete:   true,
		HCL: `
			disable_legacy_access {
    			value = "true"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         defaultSettingId,
		etagAttrName: "etag2",
	})
}

func TestDeleteDisableLegacyAccessWithConflict(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDisableLegacyAccessAPI().EXPECT().Delete(mock.Anything, settings.DeleteDisableLegacyAccessRequest{
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
			w.GetMockDisableLegacyAccessAPI().EXPECT().Delete(mock.Anything, settings.DeleteDisableLegacyAccessRequest{
				Etag: "etag2",
			}).Return(&settings.DeleteDisableLegacyAccessResponse{
				Etag: "etag3",
			}, nil)
		},
		Resource: testDisableLegacyAccess,
		HCL: `
			disable_legacy_access {
    			value = "true"
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
