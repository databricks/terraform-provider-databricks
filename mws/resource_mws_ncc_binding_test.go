package mws

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestResourceNccBindingCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/abc/workspaces/123456789",
				ExpectedRequest: provisioning.UpdateWorkspaceRequest{
					NetworkConnectivityConfigId: "ncc_id",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/123456789?",
				Response: Workspace{
					WorkspaceStatus: WorkspaceStatusRunning,
					WorkspaceID:     123456789,
				},
			},
		},
		Resource:  ResourceMwsNccBinding(),
		AccountID: "abc",
		HCL: `
		workspace_id                   = 123456789
		network_connectivity_config_id = "ncc_id"
		`,
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{"id": "123456789/ncc_id"})
}

func TestResourceNccBindingUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/abc/workspaces/123456789",
				ExpectedRequest: provisioning.UpdateWorkspaceRequest{
					NetworkConnectivityConfigId: "ncc_id",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/accounts/abc/workspaces/123456789?",
				Response: Workspace{
					WorkspaceStatus: WorkspaceStatusRunning,
					WorkspaceID:     123456789,
				},
			},
		},
		Resource:  ResourceMwsNccBinding(),
		AccountID: "abc",
		InstanceState: map[string]string{
			"network_connectivity_config_id": "ncc_prev_id",
			"workspace_id":                   "123456789",
		},
		Update: true,
		ID:     "123456789/ncc_prev_id",
		HCL: `
		workspace_id                   = 123456789
		network_connectivity_config_id = "ncc_id"
		`,
	}.ApplyAndExpectData(t, map[string]any{"id": "123456789/ncc_id"})
}
