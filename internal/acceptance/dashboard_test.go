package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type templateStruct struct {
	DisplayName         string
	WarehouseId         string
	ParentPath          string
	SerializedDashboard string
	FilePath            string
	EmbedCredentials    string
}

func makeTemplate(template templateStruct) string {
	templateString := fmt.Sprintf(`
	resource "databricks_dashboard" "d1" {
		display_name			= 	"%s"
		warehouse_id			=	"%s"
		parent_path				= 	"%s"
	`, template.DisplayName, template.WarehouseId, template.ParentPath)
	if template.SerializedDashboard != "" {
		templateString += fmt.Sprintf(`	serialized_dashboard	=	"%s"
	`, template.SerializedDashboard)
	}
	if template.FilePath != "" {
		templateString += fmt.Sprintf(`	file_path				=	"%s"
	`, template.FilePath)
	}
	if template.EmbedCredentials != "" {
		templateString += fmt.Sprintf(`	embed_credentials		=	"%s"
	`, template.EmbedCredentials)
	}
	templateString += `}

resource "databricks_permissions" "dashboard_usage" {
    dashboard_id = databricks_dashboard.d1.id

	access_control {
        group_name       = "users"
        permission_level = "CAN_READ"
    }
}
`

	return templateString
}

// Although EmbedCredentials is an optional field, please specify its value if you want to modify it.
func (t *templateStruct) SetAttributes(mapper map[string]string) templateStruct {
	// Switch case for each attribute. If it is set in the mapper, set it in the struct
	if val, ok := mapper["display_name"]; ok {
		t.DisplayName = val
	}
	if val, ok := mapper["warehouse_id"]; ok {
		t.WarehouseId = val
	}
	if val, ok := mapper["parent_path"]; ok {
		t.ParentPath = val
	}
	if val, ok := mapper["serialized_dashboard"]; ok {
		t.SerializedDashboard = val
	}
	if val, ok := mapper["file_path"]; ok {
		t.FilePath = val
	}
	if val, ok := mapper["embed_credentials"]; ok {
		if val == "true" {
			t.EmbedCredentials = ""
		} else {
			t.EmbedCredentials = val
		}
	}
	return *t
}

func TestAccBasicDashboard(t *testing.T) {
	var template templateStruct
	displayName := fmt.Sprintf("Test Dashboard - %s", qa.RandomName())
	WorkspaceLevel(t, Step{
		Template: makeTemplate(template.SetAttributes(map[string]string{
			"display_name":         displayName,
			"warehouse_id":         "{env.TEST_DEFAULT_WAREHOUSE_ID}",
			"parent_path":          "/Shared/provider-test",
			"serialized_dashboard": `{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}`,
		})),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			// As the format of the serialized dashboard is not fixed, we can only check if it is not empty
			assert.NotEqual(t, "", dashboard.SerializedDashboard)
			assert.Equal(t, displayName, dashboard.DisplayName)
			require.NoError(t, err)
			return nil
		}),
	})
}

func TestAccDashboardWithSerializedJSON(t *testing.T) {
	var template templateStruct
	displayName := fmt.Sprintf("Test Dashboard - %s", qa.RandomName())
	WorkspaceLevel(t, Step{
		Template: makeTemplate(template.SetAttributes(map[string]string{
			"display_name":         displayName,
			"warehouse_id":         "{env.TEST_DEFAULT_WAREHOUSE_ID}",
			"parent_path":          "/Shared/provider-test",
			"serialized_dashboard": `{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}`,
			"embed_credentials":    "false",
		})),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, displayName, dashboard.DisplayName)
			require.NoError(t, err)
			return nil
		}),
	}, Step{
		Template: makeTemplate(template.SetAttributes(map[string]string{
			"serialized_dashboard": `{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page Modified\"}]}`,
			"embed_credentials":    "true",
		})),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, displayName, dashboard.DisplayName)
			require.NoError(t, err)
			require.NotEqual(t, dashboard.UpdateTime, dashboard.CreateTime)
			// As the format of the serialized dashboard is not fixed, we can only check if it is not empty
			assert.NotEqual(t, "", dashboard.SerializedDashboard)
			resource.TestCheckResourceAttr("databricks_dashboard.d1", "etag", dashboard.Etag)
			return nil
		}),
	})
}

func TestAccDashboardWithFilePath(t *testing.T) {
	tmpDir := fmt.Sprintf("/tmp/Dashboard-%s", qa.RandomName())
	fileName := tmpDir + "/Dashboard.json"
	var template templateStruct
	displayName := fmt.Sprintf("Test Dashboard - %s", qa.RandomName())
	WorkspaceLevel(t, Step{
		PreConfig: func() {
			os.Mkdir(tmpDir, 0o755)
			os.WriteFile(fileName, []byte("{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"), 0o644)
		},
		Template: makeTemplate(template.SetAttributes(map[string]string{
			"display_name": displayName,
			"warehouse_id": "{env.TEST_DEFAULT_WAREHOUSE_ID}",
			"file_path":    fileName,
			"parent_path":  "/Shared/provider-test",
		})),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, displayName, dashboard.DisplayName)
			require.NoError(t, err)
			return nil
		}),
	}, Step{
		PreConfig: func() {
			os.WriteFile(fileName, []byte("{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page Modified\"}]}"), 0o644)
		},
		Template: makeTemplate(template),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, displayName, dashboard.DisplayName)
			require.NoError(t, err)
			require.NotEqual(t, dashboard.UpdateTime, dashboard.CreateTime)
			resource.TestCheckResourceAttr("databricks_dashboard.d1", "etag", dashboard.Etag)
			// As the format of the serialized dashboard is not fixed, we can only check if it is not empty
			assert.NotEqual(t, "", dashboard.SerializedDashboard)
			return nil
		}),
	})
}

func TestAccDashboardWithNoChange(t *testing.T) {
	initial_update_time := ""
	var template templateStruct
	displayName := fmt.Sprintf("Test Dashboard - %s", qa.RandomName())
	WorkspaceLevel(t, Step{
		Template: makeTemplate(template.SetAttributes(map[string]string{
			"display_name":         displayName,
			"warehouse_id":         "{env.TEST_DEFAULT_WAREHOUSE_ID}",
			"parent_path":          "/Shared/provider-test",
			"serialized_dashboard": `{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}`,
		})),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, displayName, dashboard.DisplayName)
			require.NoError(t, err)
			initial_update_time = dashboard.UpdateTime
			return nil
		}),
	}, Step{
		Template: makeTemplate(template),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, displayName, dashboard.DisplayName)
			// As the format of the serialized dashboard is not fixed, we can only check if it is not empty
			assert.NotEqual(t, "", dashboard.SerializedDashboard)
			require.NoError(t, err)
			require.Equal(t, dashboard.UpdateTime, initial_update_time)
			return nil
		}),
	})
}

func TestAccDashboardWithRemoteChange(t *testing.T) {
	dashboard_id := ""
	display_name := ""
	warehouse_id := ""
	etag := ""
	var template templateStruct
	displayName := fmt.Sprintf("Test Dashboard - %s", qa.RandomName())
	WorkspaceLevel(t, Step{
		Template: makeTemplate(template.SetAttributes(map[string]string{
			"display_name":         displayName,
			"warehouse_id":         "{env.TEST_DEFAULT_WAREHOUSE_ID}",
			"parent_path":          "/Shared/provider-test",
			"serialized_dashboard": `{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}`,
		})),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, displayName, dashboard.DisplayName)
			require.NoError(t, err)
			dashboard_id = dashboard.DashboardId
			display_name = dashboard.DisplayName
			warehouse_id = dashboard.WarehouseId
			etag = dashboard.Etag
			return nil
		}),
	}, Step{
		PreConfig: func() {
			w, err := databricks.NewWorkspaceClient(&databricks.Config{})
			require.NoError(t, err)
			_, err = w.Lakeview.Update(context.Background(), dashboards.UpdateDashboardRequest{
				DashboardId: dashboard_id,
				Dashboard: &dashboards.Dashboard{
					DashboardId:         dashboard_id,
					DisplayName:         display_name,
					Etag:                etag,
					WarehouseId:         warehouse_id,
					SerializedDashboard: "{\"pages\":[{\"name\":\"b532570b\",\"displayName\":\"New Page Modified Remote\"}]}",
				},
			})
			require.NoError(t, err)
		},
		Template: makeTemplate(template),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, displayName, dashboard.DisplayName)
			require.NoError(t, err)
			// As the format of the serialized dashboard is not fixed, we can only check if it is not empty
			assert.NotEqual(t, "", dashboard.SerializedDashboard)
			require.NotEqual(t, dashboard.UpdateTime, dashboard.CreateTime)
			resource.TestCheckResourceAttr("databricks_dashboard.d1", "etag", dashboard.Etag)
			return nil
		}),
	})
}

func TestAccDashboardTestAll(t *testing.T) {
	dashboard_id := ""
	display_name := ""
	warehouse_id := ""
	etag := ""
	tmpDir := fmt.Sprintf("/tmp/Dashboard-%s", qa.RandomName())
	fileName := tmpDir + "/Dashboard.json"
	var template templateStruct
	displayName := fmt.Sprintf("Test Dashboard - %s", qa.RandomName())
	WorkspaceLevel(t, Step{
		PreConfig: func() {
			os.Mkdir(tmpDir, 0o755)
			os.WriteFile(fileName, []byte("{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page in file\"}]}"), 0o644)
		},
		Template: makeTemplate(template.SetAttributes(map[string]string{
			"display_name":      displayName,
			"warehouse_id":      "{env.TEST_DEFAULT_WAREHOUSE_ID}",
			"parent_path":       "/Shared/provider-test",
			"file_path":         fileName,
			"embed_credentials": "false",
		})),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			publish_dash, err := w.Lakeview.GetPublished(ctx, dashboards.GetPublishedDashboardRequest{
				DashboardId: id,
			})
			assert.Equal(t, displayName, dashboard.DisplayName)
			require.NoError(t, err)
			require.NotEqual(t, dashboard.UpdateTime, dashboard.CreateTime)
			require.Equal(t, publish_dash.EmbedCredentials, false)
			return nil
		}),
	}, Step{
		PreConfig: func() {
			os.WriteFile(fileName, []byte("{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page Modified\"}]}"), 0o644)
		},
		Template: makeTemplate(template),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, displayName, dashboard.DisplayName)
			require.NoError(t, err)
			dashboard_id = dashboard.DashboardId
			display_name = dashboard.DisplayName
			warehouse_id = dashboard.WarehouseId
			etag = dashboard.Etag
			// As the format of the serialized dashboard is not fixed, we can only check if it is not empty
			assert.NotEqual(t, "", dashboard.SerializedDashboard)
			return nil
		}),
	}, Step{
		PreConfig: func() {
			w, err := databricks.NewWorkspaceClient(&databricks.Config{})
			require.NoError(t, err)
			_, err = w.Lakeview.Update(context.Background(), dashboards.UpdateDashboardRequest{
				DashboardId: dashboard_id,
				Dashboard: &dashboards.Dashboard{
					DashboardId:         dashboard_id,
					DisplayName:         display_name,
					Etag:                etag,
					WarehouseId:         warehouse_id,
					SerializedDashboard: "{\"pages\":[{\"name\":\"b532570b\",\"displayName\":\"New Page Modified Remote\"}]}",
				},
			})
			require.NoError(t, err)
		},
		Template: makeTemplate(template),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, displayName, dashboard.DisplayName)
			resource.TestCheckResourceAttr("databricks_dashboard.d1", "etag", dashboard.Etag)
			require.NoError(t, err)
			return nil
		}),
	}, Step{
		Template: makeTemplate(template.SetAttributes(map[string]string{
			"embed_credentials": "true",
			"parent_path":       "/Shared/Teams",
		})),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, displayName, dashboard.DisplayName)
			require.NoError(t, err)
			// As the format of the serialized dashboard is not fixed, we can only check if it is not empty
			assert.NotEqual(t, "", dashboard.SerializedDashboard)
			return nil
		}),
	}, Step{
		PreConfig: func() {
			os.WriteFile(fileName, []byte("{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page Modified again\"}]}"), 0o644)
		},
		Template: makeTemplate(template),
		Check: resourceCheck("databricks_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard, err := w.Lakeview.Get(ctx, dashboards.GetDashboardRequest{
				DashboardId: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, displayName, dashboard.DisplayName)
			require.NoError(t, err)
			resource.TestCheckResourceAttr("databricks_dashboard.d1", "etag", dashboard.Etag)
			// As the format of the serialized dashboard is not fixed, we can only check if it is not empty
			assert.NotEqual(t, "", dashboard.SerializedDashboard)
			return nil
		}),
	})
}

func TestAccDashboardWithWorkspacePrefix(t *testing.T) {
	var template templateStruct

	// Test that the dashboard can use a /Workspace prefix on the parent path and not trigger recreation.
	// If this does NOT work, the test fails with an error that the non-refresh plan is non-empty.

	WorkspaceLevel(t, Step{
		Template: makeTemplate(template.SetAttributes(map[string]string{
			"display_name":         fmt.Sprintf("Test Dashboard - %s", qa.RandomName()),
			"warehouse_id":         "{env.TEST_DEFAULT_WAREHOUSE_ID}",
			"parent_path":          "/Workspace/Shared/provider-test",
			"serialized_dashboard": `{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}`,
		})),
	})
}
