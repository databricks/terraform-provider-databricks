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

var testDisableLegacyDbfs = AllSettingsResources()["disable_legacy_dbfs"]

func TestCreateDisableLegacyDbfs(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockDisableLegacyDbfsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDisableLegacyDbfsRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_dbfs.value",
				Setting: settings.DisableLegacyDbfs{
					Etag: "",
					DisableLegacyDbfs: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_dbfs",
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
			e.Update(mock.Anything, settings.UpdateDisableLegacyDbfsRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_dbfs.value",
				Setting: settings.DisableLegacyDbfs{
					Etag: "etag1",
					DisableLegacyDbfs: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_dbfs",
				},
			}).Return(&settings.DisableLegacyDbfs{
				Etag: "etag2",
				DisableLegacyDbfs: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_dbfs",
			}, nil)
			e.Get(mock.Anything, settings.GetDisableLegacyDbfsRequest{
				Etag: "etag2",
			}).Return(&settings.DisableLegacyDbfs{
				Etag: "etag2",
				DisableLegacyDbfs: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_dbfs",
			}, nil)
		},
		Resource: testDisableLegacyDbfs,
		Create:   true,
		HCL: `
			disable_legacy_dbfs {
    			value = "true"
  			}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("disable_legacy_dbfs.0.value"))
}

func TestReadDisableLegacyDbfs(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDisableLegacyDbfsAPI().EXPECT().Get(mock.Anything, settings.GetDisableLegacyDbfsRequest{
				Etag: "etag1",
			}).Return(&settings.DisableLegacyDbfs{
				Etag: "etag2",
				DisableLegacyDbfs: settings.BooleanMessage{
					Value: false,
				},
				SettingName: "disable_legacy_dbfs",
			}, nil)
		},
		Resource: testDisableLegacyDbfs,
		Read:     true,
		HCL: `
			disable_legacy_dbfs {
    			value = "false"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, false, d.Get("disable_legacy_dbfs.0.value"))
}

func TestUpdateDisableLegacyDbfs(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockDisableLegacyDbfsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDisableLegacyDbfsRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_dbfs.value",
				Setting: settings.DisableLegacyDbfs{
					Etag: "etag1",
					DisableLegacyDbfs: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_dbfs",
				},
			}).Return(&settings.DisableLegacyDbfs{
				Etag: "etag2",
				DisableLegacyDbfs: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_dbfs",
			}, nil)
			e.Get(mock.Anything, settings.GetDisableLegacyDbfsRequest{
				Etag: "etag2",
			}).Return(&settings.DisableLegacyDbfs{
				Etag: "etag2",
				DisableLegacyDbfs: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_dbfs",
			}, nil)
		},
		Resource: testDisableLegacyDbfs,
		Update:   true,
		HCL: `
			disable_legacy_dbfs {
    			value = "true"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("disable_legacy_dbfs.0.value"))
}

func TestUpdateDisableLegacyDbfsWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockDisableLegacyDbfsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDisableLegacyDbfsRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_dbfs.value",
				Setting: settings.DisableLegacyDbfs{
					Etag: "etag1",
					DisableLegacyDbfs: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_dbfs",
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
			e.Update(mock.Anything, settings.UpdateDisableLegacyDbfsRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_dbfs.value",
				Setting: settings.DisableLegacyDbfs{
					Etag: "etag2",
					DisableLegacyDbfs: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_dbfs",
				},
			}).Return(&settings.DisableLegacyDbfs{
				Etag: "etag3",
				DisableLegacyDbfs: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_dbfs",
			}, nil)
			e.Get(mock.Anything, settings.GetDisableLegacyDbfsRequest{
				Etag: "etag3",
			}).Return(&settings.DisableLegacyDbfs{
				Etag: "etag3",
				DisableLegacyDbfs: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_dbfs",
			}, nil)
		},
		Resource: testDisableLegacyDbfs,
		Update:   true,
		HCL: `
			disable_legacy_dbfs {
    			value = "true"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag3", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("disable_legacy_dbfs.0.value"))
}

func TestDeleteDisableLegacyDbfs(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDisableLegacyDbfsAPI().EXPECT().Delete(mock.Anything,
				settings.DeleteDisableLegacyDbfsRequest{
					Etag: "etag1",
				}).Return(&settings.DeleteDisableLegacyDbfsResponse{
				Etag: "etag2",
			}, nil)
		},
		Resource: testDisableLegacyDbfs,
		Delete:   true,
		HCL: `
			disable_legacy_dbfs {
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

func TestDeleteDisableLegacyDbfsWithConflict(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDisableLegacyDbfsAPI().EXPECT().Delete(mock.Anything, settings.DeleteDisableLegacyDbfsRequest{
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
			w.GetMockDisableLegacyDbfsAPI().EXPECT().Delete(mock.Anything, settings.DeleteDisableLegacyDbfsRequest{
				Etag: "etag2",
			}).Return(&settings.DeleteDisableLegacyDbfsResponse{
				Etag: "etag3",
			}, nil)
		},
		Resource: testDisableLegacyDbfs,
		HCL: `
			disable_legacy_dbfs {
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
