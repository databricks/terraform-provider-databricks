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

var testAiBiEmbeddingAccessPolicySetting = AllSettingsResources()["aibi_dashboard_embedding_access_policy"]

func TestCreateAiBiEmbeddingAccessPolicySetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAibiDashboardEmbeddingAccessPolicyAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest{
				AllowMissing: true,
				FieldMask:    "aibi_dashboard_embedding_access_policy.access_policy_type",
				Setting: settings.AibiDashboardEmbeddingAccessPolicySetting{
					Etag: "",
					AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
						AccessPolicyType: "ALLOW_APPROVED_DOMAINS",
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
			e.Update(mock.Anything, settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest{
				AllowMissing: true,
				FieldMask:    "aibi_dashboard_embedding_access_policy.access_policy_type",
				Setting: settings.AibiDashboardEmbeddingAccessPolicySetting{
					Etag: "etag1",
					AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
						AccessPolicyType: "ALLOW_APPROVED_DOMAINS",
					},
					SettingName: "default",
				},
			}).Return(&settings.AibiDashboardEmbeddingAccessPolicySetting{
				Etag: "etag2",
				AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
					AccessPolicyType: "ALLOW_APPROVED_DOMAINS",
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest{
				Etag: "etag2",
			}).Return(&settings.AibiDashboardEmbeddingAccessPolicySetting{
				Etag: "etag2",
				AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
					AccessPolicyType: "ALLOW_APPROVED_DOMAINS",
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testAiBiEmbeddingAccessPolicySetting,
		Create:   true,
		HCL: `
			aibi_dashboard_embedding_access_policy {
    			access_policy_type = "ALLOW_APPROVED_DOMAINS"
  			}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, "ALLOW_APPROVED_DOMAINS", d.Get("aibi_dashboard_embedding_access_policy.0.access_policy_type"))
}

func TestReadAiBiEmbeddingAccessPolicySetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockAibiDashboardEmbeddingAccessPolicyAPI().EXPECT().Get(mock.Anything, settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest{
				Etag: "etag1",
			}).Return(&settings.AibiDashboardEmbeddingAccessPolicySetting{
				Etag: "etag2",
				AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
					AccessPolicyType: "ALLOW_APPROVED_DOMAINS",
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testAiBiEmbeddingAccessPolicySetting,
		Read:     true,
		HCL: `
			aibi_dashboard_embedding_access_policy {
    			access_policy_type = "ALLOW_APPROVED_DOMAINS"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, "ALLOW_APPROVED_DOMAINS", d.Get("aibi_dashboard_embedding_access_policy.0.access_policy_type"))
}

func TestUpdateAiBiEmbeddingAccessPolicySetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAibiDashboardEmbeddingAccessPolicyAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest{
				AllowMissing: true,
				FieldMask:    "aibi_dashboard_embedding_access_policy.access_policy_type",
				Setting: settings.AibiDashboardEmbeddingAccessPolicySetting{
					Etag: "etag1",
					AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
						AccessPolicyType: "ALLOW_APPROVED_DOMAINS",
					},
					SettingName: "default",
				},
			}).Return(&settings.AibiDashboardEmbeddingAccessPolicySetting{
				Etag: "etag2",
				AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
					AccessPolicyType: "ALLOW_APPROVED_DOMAINS",
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest{
				Etag: "etag2",
			}).Return(&settings.AibiDashboardEmbeddingAccessPolicySetting{
				Etag: "etag2",
				AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
					AccessPolicyType: "ALLOW_APPROVED_DOMAINS",
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testAiBiEmbeddingAccessPolicySetting,
		Update:   true,
		HCL: `
			aibi_dashboard_embedding_access_policy {
    			access_policy_type = "ALLOW_APPROVED_DOMAINS"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, "ALLOW_APPROVED_DOMAINS", d.Get("aibi_dashboard_embedding_access_policy.0.access_policy_type"))
}

func TestUpdateAiBiEmbeddingAccessPolicySettingWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAibiDashboardEmbeddingAccessPolicyAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest{
				AllowMissing: true,
				FieldMask:    "aibi_dashboard_embedding_access_policy.access_policy_type",
				Setting: settings.AibiDashboardEmbeddingAccessPolicySetting{
					Etag: "etag1",
					AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
						AccessPolicyType: "ALLOW_APPROVED_DOMAINS",
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
			e.Update(mock.Anything, settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest{
				AllowMissing: true,
				FieldMask:    "aibi_dashboard_embedding_access_policy.access_policy_type",
				Setting: settings.AibiDashboardEmbeddingAccessPolicySetting{
					Etag: "etag2",
					AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
						AccessPolicyType: "ALLOW_APPROVED_DOMAINS",
					},
					SettingName: "default",
				},
			}).Return(&settings.AibiDashboardEmbeddingAccessPolicySetting{
				Etag: "etag3",
				AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
					AccessPolicyType: "ALLOW_APPROVED_DOMAINS",
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest{
				Etag: "etag3",
			}).Return(&settings.AibiDashboardEmbeddingAccessPolicySetting{
				Etag: "etag3",
				AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
					AccessPolicyType: "ALLOW_APPROVED_DOMAINS",
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testAiBiEmbeddingAccessPolicySetting,
		Update:   true,
		HCL: `
			aibi_dashboard_embedding_access_policy {
    			access_policy_type = "ALLOW_APPROVED_DOMAINS"
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag3", d.Get(etagAttrName).(string))
	assert.Equal(t, "ALLOW_APPROVED_DOMAINS", d.Get("aibi_dashboard_embedding_access_policy.0.access_policy_type"))
}

func TestDeleteAiBiEmbeddingAccessPolicySetting(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockAibiDashboardEmbeddingAccessPolicyAPI().EXPECT().Delete(mock.Anything,
				settings.DeleteAibiDashboardEmbeddingAccessPolicySettingRequest{
					Etag: "etag1",
				}).Return(&settings.DeleteAibiDashboardEmbeddingAccessPolicySettingResponse{
				Etag: "etag2",
			}, nil)
		},
		Resource: testAiBiEmbeddingAccessPolicySetting,
		Delete:   true,
		HCL: `
		aibi_dashboard_embedding_access_policy {
    		access_policy_type = "ALLOW_APPROVED_DOMAINS"
  		}
		etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         defaultSettingId,
		etagAttrName: "etag2",
	})
}

func TestDeleteAiBiEmbeddingAccessPolicySettingWithConflict(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockAibiDashboardEmbeddingAccessPolicyAPI().EXPECT().Delete(mock.Anything, settings.DeleteAibiDashboardEmbeddingAccessPolicySettingRequest{
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
			w.GetMockAibiDashboardEmbeddingAccessPolicyAPI().EXPECT().Delete(mock.Anything, settings.DeleteAibiDashboardEmbeddingAccessPolicySettingRequest{
				Etag: "etag2",
			}).Return(&settings.DeleteAibiDashboardEmbeddingAccessPolicySettingResponse{
				Etag: "etag3",
			}, nil)
		},
		Resource: testAiBiEmbeddingAccessPolicySetting,
		HCL: `
		aibi_dashboard_embedding_access_policy {
    		access_policy_type = "ALLOW_APPROVED_DOMAINS"
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
