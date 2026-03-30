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

var testSqlResultsDownloadSetting = AllSettingsResources()["sql_results_download"]

func TestQueryCreateSqlResultsDownloadSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSqlResultsDownloadAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateSqlResultsDownloadRequest{
				AllowMissing: true,
				FieldMask:    sqlResultsDownloadFieldMask,
				Setting: settings.SqlResultsDownload{
					Etag: "",
					BooleanVal: settings.BooleanMessage{
						Value:           true,
						ForceSendFields: []string{"Value"},
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
			e.Update(mock.Anything, settings.UpdateSqlResultsDownloadRequest{
				AllowMissing: true,
				FieldMask:    sqlResultsDownloadFieldMask,
				Setting: settings.SqlResultsDownload{
					Etag: "etag1",
					BooleanVal: settings.BooleanMessage{
						Value:           true,
						ForceSendFields: []string{"Value"},
					},
					SettingName: "default",
				},
			}).Return(&settings.SqlResultsDownload{
				Etag: "etag2",
				BooleanVal: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetSqlResultsDownloadRequest{
				Etag: "etag2",
			}).Return(&settings.SqlResultsDownload{
				Etag: "etag2",
				BooleanVal: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testSqlResultsDownloadSetting,
		Create:   true,
		HCL: `
			boolean_val {
				value = true
			}
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	res := d.Get("boolean_val").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["value"])
}

func TestQueryReadSqlResultsDownloadSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockSqlResultsDownloadAPI().EXPECT().Get(mock.Anything, settings.GetSqlResultsDownloadRequest{
				Etag: "etag1",
			}).Return(&settings.SqlResultsDownload{
				Etag: "etag2",
				BooleanVal: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testSqlResultsDownloadSetting,
		Read:     true,
		HCL: `
			boolean_val {
				value = true
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	res := d.Get("boolean_val").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["value"])
}

func TestQueryUpdateSqlResultsDownloadSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSqlResultsDownloadAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateSqlResultsDownloadRequest{
				AllowMissing: true,
				FieldMask:    sqlResultsDownloadFieldMask,
				Setting: settings.SqlResultsDownload{
					Etag: "etag1",
					BooleanVal: settings.BooleanMessage{
						Value:           true,
						ForceSendFields: []string{"Value"},
					},
					SettingName: "default",
				},
			}).Return(&settings.SqlResultsDownload{
				Etag: "etag2",
				BooleanVal: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetSqlResultsDownloadRequest{
				Etag: "etag2",
			}).Return(&settings.SqlResultsDownload{
				Etag: "etag2",
				BooleanVal: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testSqlResultsDownloadSetting,
		Update:   true,
		HCL: `
			boolean_val {
				value = true
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	res := d.Get("boolean_val").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["value"])
}

func TestQueryUpdateSqlResultsDownloadSettingWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSqlResultsDownloadAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateSqlResultsDownloadRequest{
				AllowMissing: true,
				FieldMask:    sqlResultsDownloadFieldMask,
				Setting: settings.SqlResultsDownload{
					Etag: "etag1",
					BooleanVal: settings.BooleanMessage{
						Value:           true,
						ForceSendFields: []string{"Value"},
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
			e.Update(mock.Anything, settings.UpdateSqlResultsDownloadRequest{
				AllowMissing: true,
				FieldMask:    sqlResultsDownloadFieldMask,
				Setting: settings.SqlResultsDownload{
					Etag: "etag2",
					BooleanVal: settings.BooleanMessage{
						Value:           true,
						ForceSendFields: []string{"Value"},
					},
					SettingName: "default",
				},
			}).Return(&settings.SqlResultsDownload{
				Etag: "etag3",
				BooleanVal: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetSqlResultsDownloadRequest{
				Etag: "etag3",
			}).Return(&settings.SqlResultsDownload{
				Etag: "etag3",
				BooleanVal: settings.BooleanMessage{
					Value: true,
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testSqlResultsDownloadSetting,
		Update:   true,
		HCL: `
			boolean_val {
				value = true
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag3", d.Get(etagAttrName).(string))
	res := d.Get("boolean_val").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, res["value"])
}

func TestQueryDeleteSqlResultsDownloadSetting(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockSqlResultsDownloadAPI().EXPECT().Delete(mock.Anything, settings.DeleteSqlResultsDownloadRequest{
				Etag: "etag1",
			}).Return(&settings.DeleteSqlResultsDownloadResponse{
				Etag: "etag2",
			}, nil)
		},
		Resource: testSqlResultsDownloadSetting,
		Delete:   true,
		HCL: `
			boolean_val {
				value = true
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         defaultSettingId,
		etagAttrName: "etag2",
	})
}

func TestQueryDeleteSqlResultsDownloadSettingWithConflict(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSqlResultsDownloadAPI().EXPECT()
			e.Delete(mock.Anything, settings.DeleteSqlResultsDownloadRequest{
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
			e.Delete(mock.Anything, settings.DeleteSqlResultsDownloadRequest{
				Etag: "etag2",
			}).Return(&settings.DeleteSqlResultsDownloadResponse{
				Etag: "etag3",
			}, nil)
		},
		Resource: testSqlResultsDownloadSetting,
		Delete:   true,
		HCL: `
			boolean_val {
				value = true
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         defaultSettingId,
		etagAttrName: "etag3",
	})
}