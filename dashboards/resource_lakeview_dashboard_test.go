package dashboards

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
)

func TestLakeviewDashboardCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/lakeview/dashboards",
				ExpectedRequest: dashboards.Dashboard{
					DisplayName:         "Dashboard name",
					WarehouseId:         "abc",
					ParentPath:          "/path",
					SerializedDashboard: "serialized_json",
				},
				Response: dashboards.Dashboard{
					DashboardId:         "xyz",
					DisplayName:         "Dashboard name",
					SerializedDashboard: "serialized_json_2",
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/lakeview/dashboards/xyz/published",
				ExpectedRequest: dashboards.PublishRequest{
					EmbedCredentials: true,
					WarehouseId:      "abc",
				},
				Response: dashboards.PublishedDashboard{
					EmbedCredentials:   true,
					WarehouseId:        "abc",
					DisplayName:        "Dashboard name",
					RevisionCreateTime: "823828",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/lakeview/dashboards/xyz?",
				Response: dashboards.Dashboard{
					DashboardId:         "xyz",
					DisplayName:         "Dashboard name",
					SerializedDashboard: "serialized_json_2",
					WarehouseId:         "abc",
					CreateTime:          "12345678",
					UpdateTime:          "2125678",
				},
			},
		},
		Resource: ResourceLakeviewDashboard(),
		Create:   true,
		State: map[string]any{
			"display_name":         "Dashboard name",
			"warehouse_id":         "abc",
			"parent_path":          "/path",
			"serialized_dashboard": "serialized_json",
		},
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "Dashboard name", d.Get("display_name"))
	assert.Equal(t, "12345678", d.Get("create_time"))
}

func TestLakeviewDashboardRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/lakeview/dashboards/xyz?",
				Response: dashboards.Dashboard{
					DashboardId:         "xyz",
					WarehouseId:         "abc",
					SerializedDashboard: "{\"pages\":[{\"name\":\"43622\",\"displayName\":\"New Page\"}]}",
					CreateTime:          "12345678",
					UpdateTime:          "2125678",
				},
			},
		},
		Resource: ResourceLakeviewDashboard(),
		Read:     true,
		ID:       "xyz",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "12345678", d.Get("create_time"))
	assert.Equal(t, "2125678", d.Get("update_time"))
}

func TestLakeviewDashboardUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/lakeview/dashboards/",
				ExpectedRequest: dashboards.UpdateDashboardRequest{
					DisplayName:         "Dashboard name",
					WarehouseId:         "abc",
					SerializedDashboard: "serialized_dashboard_updated",
				},
				Response: dashboards.Dashboard{
					WarehouseId:         "abc",
					DashboardId:         "xyz",
					DisplayName:         "Dashboard name",
					SerializedDashboard: "serialized_dashboard_updated_2",
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/lakeview/dashboards/xyz/published",
				ExpectedRequest: dashboards.PublishRequest{
					EmbedCredentials: true,
					WarehouseId:      "abc",
				},
				Response: dashboards.PublishedDashboard{
					EmbedCredentials:   true,
					WarehouseId:        "abc",
					DisplayName:        "Dashboard name",
					RevisionCreateTime: "823828",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/lakeview/dashboards/xyz?",
				Response: dashboards.Dashboard{
					WarehouseId:         "abc",
					DashboardId:         "xyz",
					DisplayName:         "Dashboard name",
					SerializedDashboard: "serialized_dashboard_updated_2",
				},
			},
		},
		Resource:    ResourceLakeviewDashboard(),
		Update:      true,
		ID:          "xyz",
		RequiresNew: true,
		State: map[string]any{
			"display_name":         "Dashboard name",
			"warehouse_id":         "abc",
			"parent_path":          "/path",
			"serialized_dashboard": "serialized_dashboard_updated",
		},
		InstanceState: map[string]string{
			"display_name":         "Dashboard name",
			"warehouse_id":         "abc",
			"parent_path":          "/path",
			"serialized_dashboard": "serialized_dashboard",
		},
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "Dashboard name", d.Get("display_name"))
	assert.Equal(t, "serialized_dashboard_updated_2", d.Get("serialized_dashboard"))
	assert.Equal(t, "abc", d.Get("warehouse_id"))
}

func TestLakeviewDashboardDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/lakeview/dashboards/xyz?",
			},
		},
		Resource: ResourceLakeviewDashboard(),
		Delete:   true,
		ID:       "xyz",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}
