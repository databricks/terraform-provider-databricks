package databricks

// Preview feature: https://docs.databricks.com/security/network/ip-access-list.html
// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceConfCreateOrUpdate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
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
				"enabledIpAccessLists": "true",
			},
		},
	},
		resourceWorkspaceConf,
		map[string]interface{}{
			"enable_ip_access_lists": "true",
		},
		resourceWorkspaceConfCreateOrUpdate,
	)
	assert.NoError(t, err, err)
	assert.Equal(t, "workspace_configs", d.Id())
	assert.Equal(t, true, d.Get("enable_ip_access_lists"))
}

func TestWorkspaceConfRead_true(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/workspace-conf?keys=enableIpAccessLists",
			Response: map[string]string{
				"enableIpAccessLists": "true",
			},
		},
	},
		resourceWorkspaceConf,
		nil,
		resourceWorkspaceConfRead,
	)
	assert.NoError(t, err, err)
	assert.Equal(t, true, d.Get("enable_ip_access_lists"))
}

func TestWorkspaceConfRead_false(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/workspace-conf?keys=enableIpAccessLists",
			Response: map[string]string{
				"enableIpAccessLists": "false",
			},
		},
	},
		resourceWorkspaceConf,
		nil,
		resourceWorkspaceConfRead,
	)
	assert.NoError(t, err, err)
	assert.Equal(t, false, d.Get("enable_ip_access_lists"))
}

func TestWorkspaceConfRead_empty(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/workspace-conf?keys=enableIpAccessLists",
			Response: map[string]string{
				"enableIpAccessLists": "",
			},
		},
	},
		resourceWorkspaceConf,
		nil,
		resourceWorkspaceConfRead,
	)
	assert.NoError(t, err, err)
	assert.Equal(t, false, d.Get("enable_ip_access_lists"))
}

func TestWorkspaceConfRead_error(t *testing.T) {
	_, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/workspace-conf?keys=enableIpAccessLists",
			Response: service.APIErrorBody{
				ErrorCode: "SERVER_ERROR",
				Message:   "Something unexpected happened",
			},
			Status: 500,
		},
	},
		resourceWorkspaceConf,
		nil,
		resourceWorkspaceConfRead,
	)
	assert.Error(t, err)
}

func TestWorkspaceConfDelete(t *testing.T) {
	// interrestingly enough, once you set ip access lists, you cant reset to null--only true/false are valid
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodPatch,
			Resource: "/api/2.0/preview/workspace-conf",
			ExpectedRequest: map[string]string{
				"enabledIpAccessLists": "false",
			},
		},
	},
		resourceWorkspaceConf,
		nil,
		resourceWorkspaceConfDelete,
	)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id())
}
