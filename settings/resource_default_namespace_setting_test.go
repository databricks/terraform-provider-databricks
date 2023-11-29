package settings

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestQueryCreateDefaultNameSetting(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default",
				Status:   404,
				ExpectedRequest: settings.UpdateDefaultWorkspaceNamespaceRequest{
					AllowMissing: true,
					FieldMask:    "namespace.value",
					Setting: &settings.DefaultNamespaceSetting{
						Etag: "",
						Namespace: settings.StringMessage{
							Value: "namespace_value",
						},
						SettingName: "default",
					},
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "SomeMessage",
					Details: []apierr.ErrorDetail{{
						Type: "type.googleapis.com/google.rpc.ErrorInfo",
						Metadata: map[string]string{
							"etag": "etag1",
						},
					}},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default",
				Status:   200,
				ExpectedRequest: settings.UpdateDefaultWorkspaceNamespaceRequest{
					AllowMissing: true,
					FieldMask:    "namespace.value",
					Setting: &settings.DefaultNamespaceSetting{
						Etag: "etag1",
						Namespace: settings.StringMessage{
							Value: "namespace_value",
						},
						SettingName: "default",
					},
				},
				Response: &settings.DefaultNamespaceSetting{
					Etag: "etag2",
					Namespace: settings.StringMessage{
						Value: "namespace_value",
					},
					SettingName: "default",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default?etag=etag2",
				Status:   200,
				Response: &settings.DefaultNamespaceSetting{
					Etag: "etag2",
					Namespace: settings.StringMessage{
						Value: "namespace_value",
					},
					SettingName: "default",
				},
			},
		},
		Resource: ResourceDefaultNamespaceSetting(),
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default?etag=etag1",
				Status:   200,
				Response: &settings.DefaultNamespaceSetting{
					// This simulates that the Setting has been changed externally. Thus the different etag.
					Etag: "etag2",
					Namespace: settings.StringMessage{
						Value: "namespace_value",
					},
					SettingName: "default",
				},
			},
		},
		Resource: ResourceDefaultNamespaceSetting(),
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default",
				Status:   200,
				ExpectedRequest: settings.UpdateDefaultWorkspaceNamespaceRequest{
					AllowMissing: true,
					FieldMask:    "namespace.value",
					Setting: &settings.DefaultNamespaceSetting{
						Etag: "etag1",
						Namespace: settings.StringMessage{
							Value: "new_namespace_value",
						},
						SettingName: "default",
					},
				},
				Response: &settings.DefaultNamespaceSetting{
					Etag: "etag2",
					Namespace: settings.StringMessage{
						Value: "new_namespace_value",
					},
					SettingName: "default",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default?etag=etag2",
				Status:   200,
				Response: &settings.DefaultNamespaceSetting{
					Etag: "etag2",
					Namespace: settings.StringMessage{
						Value: "new_namespace_value",
					},
					SettingName: "default",
				},
			},
		},
		Resource: ResourceDefaultNamespaceSetting(),
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default",
				Status:   409,
				ExpectedRequest: settings.UpdateDefaultWorkspaceNamespaceRequest{
					AllowMissing: true,
					FieldMask:    "namespace.value",
					Setting: &settings.DefaultNamespaceSetting{
						Etag: "etag1",
						Namespace: settings.StringMessage{
							Value: "new_namespace_value",
						},
						SettingName: "default",
					},
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "RESOURCE_CONFLICT",
					Message:   "SomeMessage",
					Details: []apierr.ErrorDetail{{
						Type: "type.googleapis.com/google.rpc.ErrorInfo",
						Metadata: map[string]string{
							"etag": "etag2",
						},
					}},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default",
				Status:   200,
				ExpectedRequest: settings.UpdateDefaultWorkspaceNamespaceRequest{
					AllowMissing: true,
					FieldMask:    "namespace.value",
					Setting: &settings.DefaultNamespaceSetting{
						Etag: "etag2",
						Namespace: settings.StringMessage{
							Value: "new_namespace_value",
						},
						SettingName: "default",
					},
				},
				Response: &settings.DefaultNamespaceSetting{
					Etag: "etag3",
					Namespace: settings.StringMessage{
						Value: "new_namespace_value",
					},
					SettingName: "default",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default?etag=etag3",
				Status:   200,
				Response: &settings.DefaultNamespaceSetting{
					Etag: "etag3",
					Namespace: settings.StringMessage{
						Value: "new_namespace_value",
					},
					SettingName: "default",
				},
			},
		},
		Resource: ResourceDefaultNamespaceSetting(),
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default?etag=etag1",
				Status:   200,
				Response: &settings.DeleteDefaultWorkspaceNamespaceResponse{
					Etag: "etag2",
				},
			},
		},
		Resource: ResourceDefaultNamespaceSetting(),
		Delete:   true,
		ID:       "etag1",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "etag2", d.Id())
}

func TestQueryDeleteDefaultNameSettingWithConflict(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default?etag=etag1",
				Status:   409,
				Response: apierr.APIErrorBody{
					ErrorCode: "RESOURCE_CONFLICT",
					Message:   "SomeMessage",
					Details: []apierr.ErrorDetail{{
						Type: "type.googleapis.com/google.rpc.ErrorInfo",
						Metadata: map[string]string{
							"etag": "etag2",
						},
					}},
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default?etag=etag2",
				Status:   200,
				Response: &settings.DeleteDefaultWorkspaceNamespaceResponse{
					Etag: "etag3",
				},
			},
		},
		Resource: ResourceDefaultNamespaceSetting(),
		Delete:   true,
		ID:       "etag1",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "etag3", d.Id())
}
