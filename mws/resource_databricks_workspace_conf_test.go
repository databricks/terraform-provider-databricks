package mws

// Preview feature: https://docs.databricks.com/security/network/ip-access-list.html
// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceConfCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/workspace-conf?keys=enableIpAccessLists",
				Response: map[string]string{
					"enableIpAccessLists": "true",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/preview/workspace-conf",
				ExpectedRequest: map[string]string{
					"enableIpAccessLists": "true",
				},
			},
		},
		Resource: ResourceWorkspaceConf(),
		State: map[string]interface{}{
			"enable_ip_access_lists": "true",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "workspace_configs", d.Id())
	assert.Equal(t, true, d.Get("enable_ip_access_lists"))
}

func TestWorkspaceConfCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/workspace-conf?keys=enableIpAccessLists",
				Response: map[string]string{
					"enableIpAccessLists": "true",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/preview/workspace-conf",
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
		State: map[string]interface{}{
			"enable_ip_access_lists": "true",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestWorkspaceConfUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/workspace-conf?keys=enableIpAccessLists",
				Response: map[string]string{
					"enableIpAccessLists": "true",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/preview/workspace-conf",
				ExpectedRequest: map[string]string{
					"enableIpAccessLists": "true",
				},
			},
		},
		Resource: ResourceWorkspaceConf(),
		State: map[string]interface{}{
			"enable_ip_access_lists": "true",
		},
		Update: true,
		ID:     "workspace_configs",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "workspace_configs", d.Id())
	assert.Equal(t, true, d.Get("enable_ip_access_lists"))
}

func TestWorkspaceConfUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/workspace-conf?keys=enableIpAccessLists",
				Response: map[string]string{
					"enableIpAccessLists": "true",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/preview/workspace-conf",
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
		State: map[string]interface{}{
			"enable_ip_access_lists": "true",
		},
		Update: true,
		ID:     "workspace_configs",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
}

func TestWorkspaceConfRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/workspace-conf?keys=enableIpAccessLists",
				Response: map[string]string{
					"enableIpAccessLists": "true",
				},
			},
		},
		Resource: ResourceWorkspaceConf(),
		Read:     true,
		ID:       "workspace_configs",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, true, d.Get("enable_ip_access_lists"))
}

func TestWorkspaceConfRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/workspace-conf?keys=enableIpAccessLists",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceWorkspaceConf(),
		Read:     true,
		ID:       "workspace_configs",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestWorkspaceConfRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/preview/workspace-conf?keys=enableIpAccessLists",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceWorkspaceConf(),
		Read:     true,
		ID:       "workspace_configs",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "workspace_configs", d.Id(), "Id should not be empty for error reads")
}

func TestWorkspaceConfDelete(t *testing.T) {
	// interrestingly enough, once you set ip access lists, you cant reset to null--only true/false are valid
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/preview/workspace-conf",
				ExpectedRequest: map[string]string{
					"enableIpAccessLists": "false",
				},
			},
		},
		Resource: ResourceWorkspaceConf(),
		Delete:   true,
		ID:       "workspace_configs",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "workspace_configs", d.Id())
}

func TestWorkspaceConfDelete_Error(t *testing.T) {
	// interrestingly enough, once you set ip access lists, you cant reset to null--only true/false are valid
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/preview/workspace-conf",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceWorkspaceConf(),
		Delete:   true,
		ID:       "workspace_configs",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "workspace_configs", d.Id())
}
