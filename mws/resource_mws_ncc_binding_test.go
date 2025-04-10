package mws

import (
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

var mockWorkspace = &provisioning.Workspace{
	WorkspaceId:     123456789,
	WorkspaceStatus: provisioning.WorkspaceStatusRunning,
}

var mockWaiter = &provisioning.WaitGetWorkspaceRunning[struct{}]{
	WorkspaceId: 123456789,
	Poll: func(d time.Duration, f func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
		return mockWorkspace, nil
	},
}

func TestResourceNccBindingCreate(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Update(mock.Anything, provisioning.UpdateWorkspaceRequest{
				WorkspaceId:                 123456789,
				NetworkConnectivityConfigId: "ncc_id",
			}).Return(mockWaiter, nil)
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 123456789,
			}).Return(mockWorkspace, nil)
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().Update(mock.Anything, provisioning.UpdateWorkspaceRequest{
				WorkspaceId:                 123456789,
				NetworkConnectivityConfigId: "new_ncc_id",
			}).Return(mockWaiter, nil)
			a.GetMockWorkspacesAPI().EXPECT().Get(mock.Anything, provisioning.GetWorkspaceRequest{
				WorkspaceId: 123456789,
			}).Return(mockWorkspace, nil)
		},
		Resource:  ResourceMwsNccBinding(),
		AccountID: "abc",
		Update:    true,
		InstanceState: map[string]string{
			"id":                             "123456789/old_ncc_id",
			"workspace_id":                   "123456789",
			"network_connectivity_config_id": "old_ncc_id",
		},
		HCL: `
		workspace_id                   = 123456789
		network_connectivity_config_id = "new_ncc_id"
		`,
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{"id": "123456789/new_ncc_id"})
}
