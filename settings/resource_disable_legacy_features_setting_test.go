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

var testDisableLegacyFeatures = AllSettingsResources()["disable_legacy_features"]

func TestCreateDisableLegacyFeatures(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockDisableLegacyFeaturesAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDisableLegacyFeaturesRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_features.value",
				Setting: settings.DisableLegacyFeatures{
					Etag: "",
					DisableLegacyFeatures: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_features",
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
			e.Update(mock.Anything, settings.UpdateDisableLegacyFeaturesRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_features.value",
				Setting: settings.DisableLegacyFeatures{
					Etag: "etag1",
					DisableLegacyFeatures: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_features",
				},
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag2",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_features",
			}, nil)
			e.Get(mock.Anything, settings.GetDisableLegacyFeaturesRequest{
				Etag: "etag2",
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag2",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_features",
			}, nil)
		},
		Resource: testDisableLegacyFeatures,
		Create:   true,
		HCL: `
			disable_legacy_features {
    			value = "true"
  			}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("disable_legacy_features.0.value"))
}

func TestReadDisableLegacyFeatures(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			w.GetMockDisableLegacyFeaturesAPI().EXPECT().Get(mock.Anything, settings.GetDisableLegacyFeaturesRequest{
				Etag: "etag1",
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag2",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: false,
				},
				SettingName: "disable_legacy_features",
			}, nil)
		},
		Resource: testDisableLegacyFeatures,
		Read:     true,
		HCL: `
			disable_legacy_features {
    			value = "false"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, false, d.Get("disable_legacy_features.0.value"))
}

func TestUpdateDisableLegacyFeatures(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockDisableLegacyFeaturesAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDisableLegacyFeaturesRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_features.value",
				Setting: settings.DisableLegacyFeatures{
					Etag: "etag1",
					DisableLegacyFeatures: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_features",
				},
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag2",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_features",
			}, nil)
			e.Get(mock.Anything, settings.GetDisableLegacyFeaturesRequest{
				Etag: "etag2",
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag2",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_features",
			}, nil)
		},
		Resource: testDisableLegacyFeatures,
		Update:   true,
		HCL: `
			disable_legacy_features {
    			value = "true"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("disable_legacy_features.0.value"))
}

func TestUpdateDisableLegacyFeaturesWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockDisableLegacyFeaturesAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDisableLegacyFeaturesRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_features.value",
				Setting: settings.DisableLegacyFeatures{
					Etag: "etag1",
					DisableLegacyFeatures: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_features",
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
			e.Update(mock.Anything, settings.UpdateDisableLegacyFeaturesRequest{
				AllowMissing: true,
				FieldMask:    "disable_legacy_features.value",
				Setting: settings.DisableLegacyFeatures{
					Etag: "etag2",
					DisableLegacyFeatures: settings.BooleanMessage{
						Value: true,
					},
					SettingName: "disable_legacy_features",
				},
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag3",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_features",
			}, nil)
			e.Get(mock.Anything, settings.GetDisableLegacyFeaturesRequest{
				Etag: "etag3",
			}).Return(&settings.DisableLegacyFeatures{
				Etag: "etag3",
				DisableLegacyFeatures: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "disable_legacy_features",
			}, nil)
		},
		Resource: testDisableLegacyFeatures,
		Update:   true,
		HCL: `
			disable_legacy_features {
    			value = "true"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag3", d.Get(etagAttrName).(string))
	assert.Equal(t, true, d.Get("disable_legacy_features.0.value"))
}

func TestDeleteDisableLegacyFeatures(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			w.GetMockDisableLegacyFeaturesAPI().EXPECT().Delete(mock.Anything,
				settings.DeleteDisableLegacyFeaturesRequest{
					Etag: "etag1",
				}).Return(&settings.DeleteDisableLegacyFeaturesResponse{
				Etag: "etag2",
			}, nil)
		},
		Resource: testDisableLegacyFeatures,
		Delete:   true,
		HCL: `
			disable_legacy_features {
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

func TestDeleteDisableLegacyFeaturesWithConflict(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			w.GetMockDisableLegacyFeaturesAPI().EXPECT().Delete(mock.Anything, settings.DeleteDisableLegacyFeaturesRequest{
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
			w.GetMockDisableLegacyFeaturesAPI().EXPECT().Delete(mock.Anything, settings.DeleteDisableLegacyFeaturesRequest{
				Etag: "etag2",
			}).Return(&settings.DeleteDisableLegacyFeaturesResponse{
				Etag: "etag3",
			}, nil)
		},
		Resource: testDisableLegacyFeatures,
		HCL: `
			disable_legacy_features {
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
