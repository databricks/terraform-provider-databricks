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

var testAiBiEmbeddingAllowedDomainsSetting = AllSettingsResources()["aibi_dashboard_embedding_approved_domains"]

func TestCreateAiBiEmbeddingAllowedDomains(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAibiDashboardEmbeddingApprovedDomainsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest{
				AllowMissing: true,
				FieldMask:    "aibi_dashboard_embedding_approved_domains.approved_domains",
				Setting: settings.AibiDashboardEmbeddingApprovedDomainsSetting{
					Etag: "",
					AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
						ApprovedDomains: []string{"test1.com", "test2.com"},
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
			e.Update(mock.Anything, settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest{
				AllowMissing: true,
				FieldMask:    "aibi_dashboard_embedding_approved_domains.approved_domains",
				Setting: settings.AibiDashboardEmbeddingApprovedDomainsSetting{
					Etag: "etag1",
					AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
						ApprovedDomains: []string{"test1.com", "test2.com"},
					},
					SettingName: "default",
				},
			}).Return(&settings.AibiDashboardEmbeddingApprovedDomainsSetting{
				Etag: "etag2",
				AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
					ApprovedDomains: []string{"test1.com", "test2.com"},
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest{
				Etag: "etag2",
			}).Return(&settings.AibiDashboardEmbeddingApprovedDomainsSetting{
				Etag: "etag2",
				AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
					ApprovedDomains: []string{"test1.com", "test2.com"},
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testAiBiEmbeddingAllowedDomainsSetting,
		Create:   true,
		HCL: `
			aibi_dashboard_embedding_approved_domains {
    			approved_domains = ["test1.com", "test2.com"]
  			}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, []any{"test1.com", "test2.com"}, d.Get("aibi_dashboard_embedding_approved_domains.0.approved_domains"))
}

func TestReadRestrictAiBiEmbeddingAllowedDomainsSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockAibiDashboardEmbeddingApprovedDomainsAPI().EXPECT().Get(mock.Anything, settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest{
				Etag: "etag1",
			}).Return(&settings.AibiDashboardEmbeddingApprovedDomainsSetting{
				Etag: "etag2",
				AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
					ApprovedDomains: []string{"test1.com", "test2.com"},
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testAiBiEmbeddingAllowedDomainsSetting,
		Read:     true,
		HCL: `
			aibi_dashboard_embedding_approved_domains {
    			approved_domains = ["test1.com", "test2.com"]
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, []any{"test1.com", "test2.com"}, d.Get("aibi_dashboard_embedding_approved_domains.0.approved_domains"))
}

func TestUpdateRestrictAiBiEmbeddingAllowedDomainsSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAibiDashboardEmbeddingApprovedDomainsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest{
				AllowMissing: true,
				FieldMask:    "aibi_dashboard_embedding_approved_domains.approved_domains",
				Setting: settings.AibiDashboardEmbeddingApprovedDomainsSetting{
					Etag: "etag1",
					AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
						ApprovedDomains: []string{"test1.com", "test2.com"},
					},
					SettingName: "default",
				},
			}).Return(&settings.AibiDashboardEmbeddingApprovedDomainsSetting{
				Etag: "etag2",
				AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
					ApprovedDomains: []string{"test1.com", "test2.com"},
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest{
				Etag: "etag2",
			}).Return(&settings.AibiDashboardEmbeddingApprovedDomainsSetting{
				Etag: "etag2",
				AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
					ApprovedDomains: []string{"test1.com", "test2.com"},
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testAiBiEmbeddingAllowedDomainsSetting,
		Update:   true,
		HCL: `
			aibi_dashboard_embedding_approved_domains {
    			approved_domains = ["test1.com", "test2.com"]
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
	assert.Equal(t, []any{"test1.com", "test2.com"}, d.Get("aibi_dashboard_embedding_approved_domains.0.approved_domains"))
}

func TestUpdateRestrictAiBiEmbeddingAllowedDomainsSettingWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAibiDashboardEmbeddingApprovedDomainsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest{
				AllowMissing: true,
				FieldMask:    "aibi_dashboard_embedding_approved_domains.approved_domains",
				Setting: settings.AibiDashboardEmbeddingApprovedDomainsSetting{
					Etag: "etag1",
					AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
						ApprovedDomains: []string{"test1.com", "test2.com"},
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
			e.Update(mock.Anything, settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest{
				AllowMissing: true,
				FieldMask:    "aibi_dashboard_embedding_approved_domains.approved_domains",
				Setting: settings.AibiDashboardEmbeddingApprovedDomainsSetting{
					Etag: "etag2",
					AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
						ApprovedDomains: []string{"test1.com", "test2.com"},
					},
					SettingName: "default",
				},
			}).Return(&settings.AibiDashboardEmbeddingApprovedDomainsSetting{
				Etag: "etag3",
				AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
					ApprovedDomains: []string{"test1.com", "test2.com"},
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest{
				Etag: "etag3",
			}).Return(&settings.AibiDashboardEmbeddingApprovedDomainsSetting{
				Etag: "etag3",
				AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
					ApprovedDomains: []string{"test1.com", "test2.com"},
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testAiBiEmbeddingAllowedDomainsSetting,
		Update:   true,
		HCL: `
			aibi_dashboard_embedding_approved_domains {
    			approved_domains = ["test1.com", "test2.com"]
  			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, defaultSettingId, d.Id())
	assert.Equal(t, "etag3", d.Get(etagAttrName).(string))
	assert.Equal(t, []any{"test1.com", "test2.com"}, d.Get("aibi_dashboard_embedding_approved_domains.0.approved_domains"))
}

func TestDeleteRestrictAiBiEmbeddingAllowedDomainsSetting(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockAibiDashboardEmbeddingApprovedDomainsAPI().EXPECT().Delete(mock.Anything, settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest{
				Etag: "etag1",
			}).Return(&settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse{
				Etag: "etag2",
			}, nil)
		},
		Resource: testAiBiEmbeddingAllowedDomainsSetting,
		Delete:   true,
		HCL: `
		aibi_dashboard_embedding_approved_domains {
    		approved_domains = ["test1.com", "test2.com"]
		}
		etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         defaultSettingId,
		etagAttrName: "etag2",
	})
}

// TestAiBiEmbeddingApprovedDomainsLifecycle mirrors TestAccAiBiEmbeddings
// (in internal/acceptance/aibi_embedding_settings_test.go) at the unit level
// for the approved-domains resource: the acceptance test applies ["test.com"]
// then updates to ["test.com", "test2.com"]; this exercises the same two
// transitions against mocked SDK calls. Use this as the per-resource
// expected-behavior check while the acceptance test is skipped.
// The companion access-policy resource is covered by
// TestAiBiEmbeddingAccessPolicySettingLifecycle in its own _test.go.
func TestAiBiEmbeddingApprovedDomainsLifecycle(t *testing.T) {
	t.Run("create with single domain persists into state", func(t *testing.T) {
		d, err := qa.ResourceFixture{
			MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
				e := w.GetMockAibiDashboardEmbeddingApprovedDomainsAPI().EXPECT()
				e.Update(mock.Anything, settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest{
					AllowMissing: true,
					FieldMask:    "aibi_dashboard_embedding_approved_domains.approved_domains",
					Setting: settings.AibiDashboardEmbeddingApprovedDomainsSetting{
						AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
							ApprovedDomains: []string{"test.com"},
						},
						SettingName: "default",
					},
				}).Return(&settings.AibiDashboardEmbeddingApprovedDomainsSetting{
					Etag: "etag-after-create",
					AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
						ApprovedDomains: []string{"test.com"},
					},
					SettingName: "default",
				}, nil)
				e.Get(mock.Anything, settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest{
					Etag: "etag-after-create",
				}).Return(&settings.AibiDashboardEmbeddingApprovedDomainsSetting{
					Etag: "etag-after-create",
					AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
						ApprovedDomains: []string{"test.com"},
					},
					SettingName: "default",
				}, nil)
			},
			Resource: testAiBiEmbeddingAllowedDomainsSetting,
			Create:   true,
			HCL: `
				aibi_dashboard_embedding_approved_domains {
					approved_domains = ["test.com"]
				}
			`,
		}.Apply(t)
		assert.NoError(t, err)
		assert.Equal(t, defaultSettingId, d.Id())
		assert.Equal(t, "etag-after-create", d.Get(etagAttrName).(string))
		assert.Equal(t, []any{"test.com"}, d.Get("aibi_dashboard_embedding_approved_domains.0.approved_domains"))
	})

	t.Run("update appends a second domain", func(t *testing.T) {
		d, err := qa.ResourceFixture{
			MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
				e := w.GetMockAibiDashboardEmbeddingApprovedDomainsAPI().EXPECT()
				e.Update(mock.Anything, settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest{
					AllowMissing: true,
					FieldMask:    "aibi_dashboard_embedding_approved_domains.approved_domains",
					Setting: settings.AibiDashboardEmbeddingApprovedDomainsSetting{
						Etag: "etag-after-create",
						AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
							ApprovedDomains: []string{"test.com", "test2.com"},
						},
						SettingName: "default",
					},
				}).Return(&settings.AibiDashboardEmbeddingApprovedDomainsSetting{
					Etag: "etag-after-update",
					AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
						ApprovedDomains: []string{"test.com", "test2.com"},
					},
					SettingName: "default",
				}, nil)
				e.Get(mock.Anything, settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest{
					Etag: "etag-after-update",
				}).Return(&settings.AibiDashboardEmbeddingApprovedDomainsSetting{
					Etag: "etag-after-update",
					AibiDashboardEmbeddingApprovedDomains: settings.AibiDashboardEmbeddingApprovedDomains{
						ApprovedDomains: []string{"test.com", "test2.com"},
					},
					SettingName: "default",
				}, nil)
			},
			Resource: testAiBiEmbeddingAllowedDomainsSetting,
			Update:   true,
			HCL: `
				aibi_dashboard_embedding_approved_domains {
					approved_domains = ["test.com", "test2.com"]
				}
				etag = "etag-after-create"
			`,
			ID: defaultSettingId,
		}.Apply(t)
		assert.NoError(t, err)
		assert.Equal(t, defaultSettingId, d.Id())
		assert.Equal(t, "etag-after-update", d.Get(etagAttrName).(string))
		assert.Equal(t, []any{"test.com", "test2.com"}, d.Get("aibi_dashboard_embedding_approved_domains.0.approved_domains"))
	})
}

func TestDeleteRestrictAiBiEmbeddingAllowedDomainsSettingWithConflict(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockAibiDashboardEmbeddingApprovedDomainsAPI().EXPECT().Delete(mock.Anything, settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest{
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
			w.GetMockAibiDashboardEmbeddingApprovedDomainsAPI().EXPECT().Delete(mock.Anything, settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest{
				Etag: "etag2",
			}).Return(&settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse{
				Etag: "etag3",
			}, nil)
		},
		Resource: testAiBiEmbeddingAllowedDomainsSetting,
		HCL: `
		aibi_dashboard_embedding_approved_domains {
    		approved_domains = ["test1.com", "test2.com"]
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
