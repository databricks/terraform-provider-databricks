package workspace

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
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
				Response: map[string]interface{}{
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
	assert.NoError(t, err, err)
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
				Response: common.APIErrorBody{
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
					"enableSomething":     "false",
					"someProperty":        "",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace-conf?keys=enableIpAccessLists",
				Response: map[string]string{
					"enableIpAccessLists": "true",
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
		}`,
		Update: true,
		ID:     "_",
	}.Apply(t)
	assert.NoError(t, err, err)
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
				Response: common.APIErrorBody{
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
	assert.NoError(t, err, err)
}

func TestWorkspaceConfRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace-conf?",
				Response: common.APIErrorBody{
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
			enableSomething = "true"
			enforceSomethingElse = "true"
			enableFancyThing = "false"
			someProperty = "thing"
		}`,
		Resource: ResourceWorkspaceConf(),
		Delete:   true,
		ID:       "_",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "_", d.Id())
}

func TestWorkspaceConfDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/workspace-conf",
				Response: common.APIErrorBody{
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
