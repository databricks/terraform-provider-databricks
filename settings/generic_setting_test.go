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

// Choose an arbitrary setting to test.
var testSetting = AllSettingsResources()["default_namespace"]

func TestQueryCreateDefaultNameSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSettingsAPI().EXPECT()
			e.UpdateDefaultWorkspaceNamespace(mock.Anything, settings.UpdateDefaultWorkspaceNamespaceRequest{
				AllowMissing: true,
				FieldMask:    "namespace.value",
				Setting: &settings.DefaultNamespaceSetting{
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
						"etag": "etag1",
					},
				}},
			})
			e.UpdateDefaultWorkspaceNamespace(mock.Anything, settings.UpdateDefaultWorkspaceNamespaceRequest{
				AllowMissing: true,
				FieldMask:    "namespace.value",
				Setting: &settings.DefaultNamespaceSetting{
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
			e.ReadDefaultWorkspaceNamespace(mock.Anything, settings.ReadDefaultWorkspaceNamespaceRequest{
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
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, "etag2", d.Id())
	assert.Equal(t, "namespace_value", d.Get("namespace.0.value"))
}

func TestQueryReadDefaultNameSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockSettingsAPI().EXPECT().ReadDefaultWorkspaceNamespace(mock.Anything, settings.ReadDefaultWorkspaceNamespaceRequest{
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
		`,
		ID: "etag1",
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, "etag2", d.Id())
	res := d.Get("namespace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "namespace_value", res["value"])
}

func TestQueryUpdateDefaultNameSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSettingsAPI().EXPECT()
			e.UpdateDefaultWorkspaceNamespace(mock.Anything, settings.UpdateDefaultWorkspaceNamespaceRequest{
				AllowMissing: true,
				FieldMask:    "namespace.value",
				Setting: &settings.DefaultNamespaceSetting{
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
			e.ReadDefaultWorkspaceNamespace(mock.Anything, settings.ReadDefaultWorkspaceNamespaceRequest{
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
		`,
		ID: "etag1",
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, "etag2", d.Id())
	res := d.Get("namespace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "new_namespace_value", res["value"])
}

func TestQueryUpdateDefaultNameSettingWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSettingsAPI().EXPECT()
			e.UpdateDefaultWorkspaceNamespace(mock.Anything, settings.UpdateDefaultWorkspaceNamespaceRequest{
				AllowMissing: true,
				FieldMask:    "namespace.value",
				Setting: &settings.DefaultNamespaceSetting{
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
						"etag": "etag2",
					},
				}},
			})
			e.UpdateDefaultWorkspaceNamespace(mock.Anything, settings.UpdateDefaultWorkspaceNamespaceRequest{
				AllowMissing: true,
				FieldMask:    "namespace.value",
				Setting: &settings.DefaultNamespaceSetting{
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
			e.ReadDefaultWorkspaceNamespace(mock.Anything, settings.ReadDefaultWorkspaceNamespaceRequest{
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
		`,
		ID: "etag1",
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, "etag3", d.Id())
	res := d.Get("namespace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "new_namespace_value", res["value"])
}

func TestQueryDeleteDefaultNameSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockSettingsAPI().EXPECT().DeleteDefaultWorkspaceNamespace(mock.Anything, settings.DeleteDefaultWorkspaceNamespaceRequest{
				Etag: "etag1",
			}).Return(&settings.DeleteDefaultWorkspaceNamespaceResponse{
				Etag: "etag2",
			}, nil)
		},
		Resource: testSetting,
		Delete:   true,
		ID:       "etag1",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "etag2", d.Id())
}

func TestQueryDeleteDefaultNameSettingWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockSettingsAPI().EXPECT().DeleteDefaultWorkspaceNamespace(mock.Anything, settings.DeleteDefaultWorkspaceNamespaceRequest{
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
			w.GetMockSettingsAPI().EXPECT().DeleteDefaultWorkspaceNamespace(mock.Anything, settings.DeleteDefaultWorkspaceNamespaceRequest{
				Etag: "etag2",
			}).Return(&settings.DeleteDefaultWorkspaceNamespaceResponse{
				Etag: "etag3",
			}, nil)
		},
		Resource: testSetting,
		Delete:   true,
		ID:       "etag1",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "etag3", d.Id())
}
