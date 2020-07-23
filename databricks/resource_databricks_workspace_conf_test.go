package databricks

// Preview feature: https://docs.databricks.com/security/network/ip-access-list.html
// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"testing"
)

func TestWorkspaceConfCreateOrUpdate(t *testing.T) {
	// d, err := ResourceTester(t, []HTTPFixture{
	// 	{
	// 		Method:   http.MethodGet,
	// 		Resource: "/api/2.0/preview/workspace-conf",
	// 		Response: map[string]string{
	// 			"enableIpAccessLists": "true",
	// 		},
	// 	},
	// 	{
	// 		Method:   http.MethodPatch,
	// 		Resource: "/api/2.0/preview/workspace-conf",
	// 		Response: map[string]string{
	// 			"enableIpAccessLists": "true",
	// 		},
	// 	},
	// },
	// 	resourceWorkspaceConf,
	// 	map[string]interface{}{
	// 		"enableIpAccessLists": "true",
	// 	},
	// 	resourceWorkspaceConfCreateOrUpdate,
	// )
	// assert.NoError(t, err, err)
	// assert.Equal(t, "workspace_configs", d.Id())
	// assert.Equal(t, "true", d.Get("enableIpAccessLists"))
}

func TestWorkspaceConfRead(t *testing.T) {

}

func TestWorkspaceConfDelete(t *testing.T) {

}
