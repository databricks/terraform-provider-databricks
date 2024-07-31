package workspace

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceConfCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/workspace-conf",
				ExpectedRequest: map[string]string{
					"enableIpAccessLists": "true",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace-conf?keys=enableIpAccessLists",
				Response: map[string]any{
					"enableIpAccessLists": "true",
				},
			},
		},
		Resource: ResourceWorkspaceConf(),
		HCL: `custom_config {
			enableIpAccessLists = "true"
		}`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "_", d.Id())
	assert.Equal(t, "true", d.Get("custom_config.enableIpAccessLists"))
}

func TestWorkspaceConfCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/workspace-conf",
				ExpectedRequest: map[string]string{
					"enableIpAccessLists": "true",
				},
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceWorkspaceConf(),
		HCL: `custom_config {
			enableIpAccessLists = "true"
		}`,
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestWorkspaceConfUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/workspace-conf",
				ExpectedRequest: map[string]string{
					"enableIpAccessLists": "true",
					"enforceSomething":    "1",
					"enableSomething":     "false",
					"someProperty":        "",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace-conf?keys=enableIpAccessLists%2CenforceSomething",
				Response: map[string]string{
					"enableIpAccessLists": "true",
					"enforceSomething":    "true",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace-conf?keys=enforceSomething%2CenableIpAccessLists",
				Response: map[string]string{
					"enableIpAccessLists": "true",
					"enforceSomething":    "true",
				},
			},
		},
		Resource: ResourceWorkspaceConf(),
		InstanceState: map[string]string{
			"custom_config.enableSomething": "true",
			"custom_config.someProperty":    "thing",
		},
		HCL: `custom_config {
			enableIpAccessLists = "true"
			enforceSomething = true
		}`,
		Update: true,
		ID:     "_",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "_", d.Id())
	assert.Equal(t, "true", d.Get("custom_config.enableIpAccessLists"))
}

func TestWorkspaceConfUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/workspace-conf",
				ExpectedRequest: map[string]string{
					"enableIpAccessLists": "true",
				},
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceWorkspaceConf(),
		HCL: `custom_config {
			enableIpAccessLists = "true"
		}`,
		Update: true,
		ID:     "_",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
}

func TestWorkspaceConfRead(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace-conf?",
				Response: map[string]string{},
			},
		},
		Resource: ResourceWorkspaceConf(),
		Read:     true,
		ID:       "_",
	}.Apply(t)
	assert.NoError(t, err)
}

func TestWorkspaceConfRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace-conf?",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceWorkspaceConf(),
		Read:     true,
		ID:       "_",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "_", d.Id(), "Id should not be empty for error reads")
}

func TestWorkspaceConfDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/workspace-conf",
				ExpectedRequest: map[string]string{
					"enableFancyThing":     "false",
					"enableSomething":      "false",
					"enforceSomethingElse": "false",
					"someProperty":         "",
				},
			},
		},
		HCL: `custom_config {
			enableSomething = true
			enforceSomethingElse = "true"
			enableFancyThing = "false"
			someProperty = "thing"
		}`,
		Resource: ResourceWorkspaceConf(),
		Delete:   true,
		ID:       "_",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "_", d.Id())
}

func TestWorkspaceConfDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/workspace-conf",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceWorkspaceConf(),
		Delete:   true,
		ID:       "_",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "_", d.Id())
}

func TestWorkspaceConfUpdateOnInvalidConf(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/workspace-conf",
				Status:   400,
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "some-invalid-conf is an invalid config key",
				},
				ExpectedRequest: map[string]string{
					"some-invalid-conf": "foo",
					"some-valid-conf":   "",
				},
			},
		},
		ID: "_",
		InstanceState: map[string]string{
			"custom_config.some-valid-conf": "bar",
		},
		Resource: ResourceWorkspaceConf(),
		HCL: `custom_config {
			some-invalid-conf = "foo"
		}`,
		Update: true,
	}.Apply(t)

	// Expect error returned during update
	assert.ErrorContains(t, err, "some-invalid-conf is an invalid config key")

	// Expect previous state of the resource to be preserved
	config := d.Get("custom_config")
	assert.Equal(t, map[string]any{
		"some-valid-conf": "bar",
	}, config)
}
