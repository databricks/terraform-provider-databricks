package settings

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestQueryCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default?",
				Status:   404,
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
					Etag: "etag3",
					Namespace: settings.StringMessage{
						Value: "namespace_value",
					},
					SettingName: "default",
				},
			},
		},
		Resource: ResourceNamespaceSettings(),
		Create:   true,
		HCL: `
			setting_name = "default"
			namespace {
				value = "namespace_value"
			}
		`,
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, "etag3", d.Id())
	assert.Equal(t, "default", d.Get("setting_name"))
	assert.Equal(t, "namespace_value", d.Get("namespace.0.value"))
}

func TestQueryRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default?etag=etag1",
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
		Resource: ResourceNamespaceSettings(),
		Read:     true,
		HCL: `
			setting_name = "default"
			namespace {
				value = "namespace_value"
			}
		`,
		ID: "etag1",
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, "etag2", d.Id())
	assert.Equal(t, "default", d.Get("setting_name"))
	res := d.Get("namespace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "namespace_value", res["value"])
}

func TestQueryUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default?etag=etag1",
				Status:   200,
				Response: &settings.DefaultNamespaceSetting{
					Etag: "etag2",
					Namespace: settings.StringMessage{
						Value: "namespace_value",
					},
					SettingName: "default",
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
					Etag: "etag4",
					Namespace: settings.StringMessage{
						Value: "new_namespace_value",
					},
					SettingName: "default",
				},
			},
		},
		Resource: ResourceNamespaceSettings(),
		Update:   true,
		HCL: `
			setting_name = "default"
			namespace {
				value = "new_namespace_value"
			}
		`,
		ID: "etag1",
	}.Apply(t)

	assert.NoError(t, err)

	assert.Equal(t, "etag4", d.Id())
	assert.Equal(t, "default", d.Get("setting_name"))
	res := d.Get("namespace").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "new_namespace_value", res["value"])
}

func TestQueryDelete(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default?etag=etag1",
				Status:   200,
				Response: &settings.DefaultNamespaceSetting{
					Etag: "etag2",
					Namespace: settings.StringMessage{
						Value: "namespace_value",
					},
					SettingName: "default",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.0/settings/types/default_namespace_ws/names/default?etag=etag2",
				Status:   200,
			},
		},
		Resource: ResourceNamespaceSettings(),
		Delete:   true,
		ID:       "etag1",
	}.Apply(t)

	assert.NoError(t, err)
}
