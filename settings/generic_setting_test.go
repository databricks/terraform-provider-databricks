package settings

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Choose an arbitrary setting to test.
var testSetting = AllSettingsResources()["default_namespace"]

func TestDefaultNamespaceSettingProviderConfigEmptyID(t *testing.T) {
	qa.ResourceFixture{
		Resource: testSetting,
		Create:   true,
		HCL: `
			namespace {
				value = "namespace_value"
			}
			provider_config {
				workspace_id = ""
			}
		`,
	}.ExpectError(t, "invalid config supplied. [provider_config.#.workspace_id] expected provider_config.0.workspace_id to not be an empty string, got ")
}

func TestDefaultNamespaceSettingProviderConfigChangeRequiresRecreate(t *testing.T) {
	resource := testSetting.ToResource()
	databricksClient := &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://test.cloud.databricks.com",
				Token: "test-token",
			},
		},
	}
	databricksClient.SetCachedWorkspaceID(222)
	diff, err := resource.Diff(
		context.Background(),
		&terraform.InstanceState{
			ID: defaultSettingId,
			Attributes: map[string]string{
				"namespace.#":                    "1",
				"namespace.0.value":              "namespace_value",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "111",
			},
		},
		terraform.NewResourceConfigRaw(map[string]any{
			"namespace": []any{map[string]any{"value": "namespace_value"}},
			"provider_config": []any{map[string]any{
				"workspace_id": "222",
			}},
		}),
		databricksClient,
	)
	require.NoError(t, err)
	require.NotNil(t, diff)
	workspaceIDDiff, ok := diff.Attributes["provider_config.0.workspace_id"]
	require.True(t, ok)
	assert.True(t, workspaceIDDiff.RequiresNew)
}

func TestQueryCreateDefaultNameSetting(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockDefaultNamespaceAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDefaultNamespaceSettingRequest{
				AllowMissing: true,
				FieldMask:    "namespace.value",
				Setting: settings.DefaultNamespaceSetting{
					Etag: "",
					Namespace: settings.StringMessage{
						Value: "namespace_value",
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
			e.Update(mock.Anything, settings.UpdateDefaultNamespaceSettingRequest{
				AllowMissing: true,
				FieldMask:    "namespace.value",
				Setting: settings.DefaultNamespaceSetting{
					Etag: "etag1",
					Namespace: settings.StringMessage{
						Value: "namespace_value",
					},
					SettingName: "default",
				},
			}).Return(&settings.DefaultNamespaceSetting{
				Etag: "etag2",
				Namespace: settings.StringMessage{
					Value: "namespace_value",
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetDefaultNamespaceSettingRequest{
				Etag: "etag2",
			}).Return(&settings.DefaultNamespaceSetting{
				Etag: "etag2",
				Namespace: settings.StringMessage{
					Value: "namespace_value",
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testSetting,
		Create:   true,
		HCL: `
			namespace {
				value = "namespace_value"
			}
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":                defaultSettingId,
		etagAttrName:        "etag2",
		"namespace.0.value": "namespace_value",
	})
}

func TestQueryReadDefaultNameSetting(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDefaultNamespaceAPI().EXPECT().Get(mock.Anything, settings.GetDefaultNamespaceSettingRequest{
				Etag: "etag1",
			}).Return(&settings.DefaultNamespaceSetting{
				Etag: "etag2",
				Namespace: settings.StringMessage{
					Value: "namespace_value",
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testSetting,
		Read:     true,
		HCL: `
			namespace {
				value = "namespace_value"
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":                defaultSettingId,
		etagAttrName:        "etag2",
		"namespace.0.value": "namespace_value",
	})
}

func TestQueryUpdateDefaultNameSetting(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockDefaultNamespaceAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDefaultNamespaceSettingRequest{
				AllowMissing: true,
				FieldMask:    "namespace.value",
				Setting: settings.DefaultNamespaceSetting{
					Etag: "etag1",
					Namespace: settings.StringMessage{
						Value: "new_namespace_value",
					},
					SettingName: "default",
				},
			}).Return(&settings.DefaultNamespaceSetting{
				Etag: "etag2",
				Namespace: settings.StringMessage{
					Value: "new_namespace_value",
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetDefaultNamespaceSettingRequest{
				Etag: "etag2",
			}).Return(&settings.DefaultNamespaceSetting{
				Etag: "etag2",
				Namespace: settings.StringMessage{
					Value: "new_namespace_value",
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testSetting,
		Update:   true,
		HCL: `
			namespace {
				value = "new_namespace_value"
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":                defaultSettingId,
		etagAttrName:        "etag2",
		"namespace.0.value": "new_namespace_value",
	})
}

func TestQueryUpdateDefaultNameSettingWithConflict(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockDefaultNamespaceAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateDefaultNamespaceSettingRequest{
				AllowMissing: true,
				FieldMask:    "namespace.value",
				Setting: settings.DefaultNamespaceSetting{
					Etag: "etag1",
					Namespace: settings.StringMessage{
						Value: "new_namespace_value",
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
			e.Update(mock.Anything, settings.UpdateDefaultNamespaceSettingRequest{
				AllowMissing: true,
				FieldMask:    "namespace.value",
				Setting: settings.DefaultNamespaceSetting{
					Etag: "etag2",
					Namespace: settings.StringMessage{
						Value: "new_namespace_value",
					},
					SettingName: "default",
				},
			}).Return(&settings.DefaultNamespaceSetting{
				Etag: "etag3",
				Namespace: settings.StringMessage{
					Value: "new_namespace_value",
				},
				SettingName: "default",
			}, nil)
			e.Get(mock.Anything, settings.GetDefaultNamespaceSettingRequest{
				Etag: "etag3",
			}).Return(&settings.DefaultNamespaceSetting{
				Etag: "etag3",
				Namespace: settings.StringMessage{
					Value: "new_namespace_value",
				},
				SettingName: "default",
			}, nil)
		},
		Resource: testSetting,
		Update:   true,
		HCL: `
			namespace {
				value = "new_namespace_value"
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":                defaultSettingId,
		etagAttrName:        "etag3",
		"namespace.0.value": "new_namespace_value",
	})
}

func TestQueryDeleteDefaultNameSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDefaultNamespaceAPI().EXPECT().Delete(mock.Anything, settings.DeleteDefaultNamespaceSettingRequest{
				Etag: "etag1",
			}).Return(&settings.DeleteDefaultNamespaceSettingResponse{
				Etag: "etag2",
			}, nil)
		},
		Resource: testSetting,
		Delete:   true,
		HCL: `
			namespace {
				value = "new_namespace_value"
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "etag2", d.Get(etagAttrName).(string))
}

func TestQueryDeleteDefaultNameSettingWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockDefaultNamespaceAPI().EXPECT().Delete(mock.Anything, settings.DeleteDefaultNamespaceSettingRequest{
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
			w.GetMockDefaultNamespaceAPI().EXPECT().Delete(mock.Anything, settings.DeleteDefaultNamespaceSettingRequest{
				Etag: "etag2",
			}).Return(&settings.DeleteDefaultNamespaceSettingResponse{
				Etag: "etag3",
			}, nil)
		},
		Resource: testSetting,
		Delete:   true,
		HCL: `
			namespace {
				value = "new_namespace_value"
			}
			etag = "etag1"
		`,
		ID: defaultSettingId,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "etag3", d.Get(etagAttrName).(string))
}
