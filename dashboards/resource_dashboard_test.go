package dashboards

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDashboardCreate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockLakeviewAPI().EXPECT()
			e.Create(mock.Anything, dashboards.CreateDashboardRequest{
				Dashboard: &dashboards.Dashboard{
					DisplayName:         "Dashboard name",
					WarehouseId:         "abc",
					ParentPath:          "/path",
					SerializedDashboard: "serialized_json",
				},
			}).Return(&dashboards.Dashboard{
				DashboardId:         "xyz",
				DisplayName:         "Dashboard name",
				SerializedDashboard: "serialized_json_2",
				WarehouseId:         "abc",
				UpdateTime:          "2125678",
			}, nil)
			e.Publish(mock.Anything, dashboards.PublishRequest{
				EmbedCredentials: true,
				WarehouseId:      "abc",
				DashboardId:      "xyz",
				ForceSendFields:  []string{"EmbedCredentials"},
			}).Return(&dashboards.PublishedDashboard{
				EmbedCredentials:   true,
				WarehouseId:        "abc",
				DisplayName:        "Dashboard name",
				RevisionCreateTime: "823828",
			}, nil)
			e.Get(mock.Anything, dashboards.GetDashboardRequest{
				DashboardId: "xyz",
			}).Return(&dashboards.Dashboard{
				DashboardId:         "xyz",
				DisplayName:         "Dashboard name",
				SerializedDashboard: "serialized_json_2",
				WarehouseId:         "abc",
				UpdateTime:          "2125678",
			}, nil)
		},
		Resource: ResourceDashboard(),
		Create:   true,
		HCL: `
			display_name = "Dashboard name"
			warehouse_id = "abc"
			parent_path = "/path"
			serialized_dashboard = "serialized_json"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           "xyz",
		"display_name": "Dashboard name",
	})
}

func TestDashboardCreate_NoParent(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			lv := w.GetMockLakeviewAPI().EXPECT()
			lv.Create(mock.Anything, dashboards.CreateDashboardRequest{
				Dashboard: &dashboards.Dashboard{
					DisplayName:         "Dashboard name",
					WarehouseId:         "abc",
					ParentPath:          "/path",
					SerializedDashboard: "serialized_json",
				},
			}).Return(nil, fmt.Errorf("Path (/path) doesn't exist.")).Once()
			w.GetMockWorkspaceAPI().EXPECT().MkdirsByPath(mock.Anything, "/path").Return(nil)
			lv.Create(mock.Anything, dashboards.CreateDashboardRequest{
				Dashboard: &dashboards.Dashboard{
					DisplayName:         "Dashboard name",
					WarehouseId:         "abc",
					ParentPath:          "/path",
					SerializedDashboard: "serialized_json",
				},
			}).Return(&dashboards.Dashboard{
				DashboardId:         "xyz",
				DisplayName:         "Dashboard name",
				SerializedDashboard: "serialized_json_2",
				WarehouseId:         "abc",
				UpdateTime:          "2125678",
			}, nil)
			lv.Publish(mock.Anything, dashboards.PublishRequest{
				EmbedCredentials: true,
				WarehouseId:      "abc",
				DashboardId:      "xyz",
				ForceSendFields:  []string{"EmbedCredentials"},
			}).Return(&dashboards.PublishedDashboard{
				EmbedCredentials:   true,
				WarehouseId:        "abc",
				DisplayName:        "Dashboard name",
				RevisionCreateTime: "823828",
			}, nil)
			lv.Get(mock.Anything, dashboards.GetDashboardRequest{
				DashboardId: "xyz",
			}).Return(&dashboards.Dashboard{
				DashboardId:         "xyz",
				DisplayName:         "Dashboard name",
				SerializedDashboard: "serialized_json_2",
				WarehouseId:         "abc",
				UpdateTime:          "2125678",
			}, nil)
		},
		Resource: ResourceDashboard(),
		Create:   true,
		HCL: `
			display_name = "Dashboard name"
			warehouse_id = "abc"
			parent_path = "/path"
			serialized_dashboard = "serialized_json"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           "xyz",
		"display_name": "Dashboard name",
	})
}

func TestDashboardRead(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockLakeviewAPI().EXPECT().Get(mock.Anything, dashboards.GetDashboardRequest{
				DashboardId: "xyz",
			}).Return(&dashboards.Dashboard{
				DashboardId:         "xyz",
				WarehouseId:         "abc",
				SerializedDashboard: "serialized_json",
				CreateTime:          "12345678",
				UpdateTime:          "2125678",
			}, nil)
		},
		Resource: ResourceDashboard(),
		Read:     true,
		ID:       "xyz",
		HCL: `
			display_name = "Dashboard name"
			warehouse_id = "abc"
			parent_path = "/path"
			serialized_dashboard = "serialized_json"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           "xyz",
		"warehouse_id": "abc",
		"create_time":  "12345678",
	})
}

func TestDashboardUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockLakeviewAPI().EXPECT()
			e.Update(mock.Anything, dashboards.UpdateDashboardRequest{
				DashboardId: "xyz",
				Dashboard: &dashboards.Dashboard{
					DashboardId:         "xyz",
					DisplayName:         "Dashboard name",
					WarehouseId:         "abc",
					SerializedDashboard: "serialized_dashboard_updated",
					ParentPath:          "/path",
				},
			}).Return(&dashboards.Dashboard{
				DashboardId:         "xyz",
				DisplayName:         "Dashboard name",
				SerializedDashboard: "serialized_dashboard_updated_2",
				WarehouseId:         "abc",
				UpdateTime:          "2125678",
				ParentPath:          "/path",
			}, nil)
			e.Publish(mock.Anything, dashboards.PublishRequest{
				EmbedCredentials: true,
				WarehouseId:      "abc",
				DashboardId:      "xyz",
				ForceSendFields:  []string{"EmbedCredentials"},
			}).Return(&dashboards.PublishedDashboard{
				EmbedCredentials:   true,
				WarehouseId:        "abc",
				DisplayName:        "Dashboard name",
				RevisionCreateTime: "823828",
			}, nil)
			e.Get(mock.Anything, dashboards.GetDashboardRequest{
				DashboardId: "xyz",
			}).Return(&dashboards.Dashboard{
				DashboardId:         "xyz",
				DisplayName:         "Dashboard name",
				SerializedDashboard: "serialized_dashboard_updated_2",
				WarehouseId:         "abc",
				UpdateTime:          "2125679",
				ParentPath:          "/path",
			}, nil)
		},
		Resource: ResourceDashboard(),
		Update:   true,
		ID:       "xyz",
		HCL: `
			display_name = "Dashboard name"
			warehouse_id = "abc"
			parent_path = "/path"
			serialized_dashboard = "serialized_dashboard_updated"
		`,
		InstanceState: map[string]string{
			"display_name":         "Dashboard name",
			"warehouse_id":         "abc",
			"parent_path":          "/path",
			"serialized_dashboard": "serialized_dashboard",
		},
	}.ApplyAndExpectData(t, map[string]any{
		"id":                   "xyz",
		"display_name":         "Dashboard name",
		"warehouse_id":         "abc",
		"serialized_dashboard": "serialized_dashboard_updated_2",
	})
}

func TestDashboardDelete(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockLakeviewAPI().EXPECT()

			// Expect the dashboard to be trashed.
			e.Trash(mock.Anything, dashboards.TrashDashboardRequest{
				DashboardId: "xyz",
			}).Return(nil)
		},
		Resource: ResourceDashboard(),
		Delete:   true,
		ID:       "xyz",
		HCL: `
			display_name = "Dashboard name"
			warehouse_id = "abc"
			parent_path = "/path"
			serialized_dashboard = "serialized_json"
		`,
	}.Apply(t)

	// Expect this to succeed.
	assert.NoError(t, err)
}

func TestDashboardDeletePermissionDenied(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockLakeviewAPI().EXPECT()

			// First, expect the dashboard to be trashed.
			e.Trash(mock.Anything, dashboards.TrashDashboardRequest{
				DashboardId: "xyz",
			}).Return(apierr.ErrPermissionDenied)

			// Then, expect to get the dashboard to confirm if it was already trashed.
			// We confirm below that the below error is ignored and the original error is returned.
			e.Get(mock.Anything, dashboards.GetDashboardRequest{
				DashboardId: "xyz",
			}).Return(nil, fmt.Errorf("some other error"))
		},
		Resource: ResourceDashboard(),
		Delete:   true,
		ID:       "xyz",
		HCL: `
			display_name = "Dashboard name"
			warehouse_id = "abc"
			parent_path = "/path"
			serialized_dashboard = "serialized_json"
		`,
	}.Apply(t)

	// We expect the original error to be returned.
	assert.Equal(t, err, apierr.ErrPermissionDenied)
}

func TestDashboardDeleteAlreadyTrashed(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockLakeviewAPI().EXPECT()

			// First, expect the dashboard to be trashed.
			e.Trash(mock.Anything, dashboards.TrashDashboardRequest{
				DashboardId: "xyz",
			}).Return(apierr.ErrPermissionDenied)

			// Then, expect to get the dashboard to confirm if it was already trashed.
			e.Get(mock.Anything, dashboards.GetDashboardRequest{
				DashboardId: "xyz",
			}).Return(&dashboards.Dashboard{
				DashboardId:    "xyz",
				ParentPath:     "/path",
				LifecycleState: dashboards.LifecycleStateTrashed,
			}, nil)
		},
		Resource: ResourceDashboard(),
		Delete:   true,
		ID:       "xyz",
		HCL: `
			display_name = "Dashboard name"
			warehouse_id = "abc"
			parent_path = "/path"
			serialized_dashboard = "serialized_json"
		`,
	}.Apply(t)

	// Expect this to succeed.
	assert.NoError(t, err)
}
