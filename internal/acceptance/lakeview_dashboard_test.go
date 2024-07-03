package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccLakeviewDashboard(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			serialized_dashboard	=	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			// require.NotEqual(t, m.LastModified, createdTime)

			// raw, err := w.Files.DownloadByFilePath(ctx, id)
			// require.NoError(t, err)
			// contents, err := io.ReadAll(raw.Contents)
			// require.NoError(t, err)
			// // Check that we updated the file
			// assert.Equal(t, "abc\n", string(contents))
			return nil
		}),
	})
}

func TestAccDashboardWithSerializedJSON(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			serialized_dashboard	=	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
			embed_credentials		=	false
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			// require.NotEqual(t, m.LastModified, createdTime)
			return nil
		}),
	}, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			serialized_dashboard	=	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page Modified\"}]}"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			require.NotEqual(t, dashboard.UpdateTime, dashboard.CreateTime)
			return nil
		}),
	})
}

func TestAccDashboardWithFilePath(t *testing.T) {
	workspaceLevel(t, step{
		PreConfig: func() {
			tmpDir := "/tmp/LakeviewDashboardTest"
			fileName := tmpDir + "/Dashboard.json"
			os.Mkdir(tmpDir, 0755)
			os.WriteFile(fileName, []byte("{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"), 0644)
		},
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			file_path				=	"/tmp/LakeviewDashboardTest/Dashboard.json"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			return nil
		}),
	}, step{
		PreConfig: func() {
			os.WriteFile("/tmp/LakeviewDashboardTest/Dashboard.json", []byte("{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page Modified\"}]}"), 0644)
		},
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			file_path				=	"/tmp/LakeviewDashboardTest/Dashboard.json"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			require.NotEqual(t, dashboard.UpdateTime, dashboard.CreateTime)
			os.Remove("/tmp/LakeviewDashboardTest/Dashboard.json")
			os.Remove("/tmp/LakeviewDashboardTest")
			return nil
		}),
	})
}

func TestAccDashboardWithNoChange(t *testing.T) {
	initial_update_time := ""
	workspaceLevel(t, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			serialized_dashboard	=	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			initial_update_time = dashboard.UpdateTime
			return nil
		}),
	}, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			serialized_dashboard	=	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
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
	workspaceLevel(t, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			serialized_dashboard	=	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			dashboard_id = dashboard.DashboardId
			display_name = dashboard.DisplayName
			warehouse_id = dashboard.WarehouseId
			etag = dashboard.Etag
			return nil
		}),
	}, step{
		PreConfig: func() {
			w, err := databricks.NewWorkspaceClient(&databricks.Config{})
			if err != nil {
				fmt.Println(err)
			}
			_, err = w.Lakeview.Update(context.Background(), dashboards.UpdateDashboardRequest{
				DashboardId:         dashboard_id,
				DisplayName:         display_name,
				Etag:                etag,
				WarehouseId:         warehouse_id,
				SerializedDashboard: "{\"pages\":[{\"name\":\"b532570b\",\"displayName\":\"New Page Modified\"}]}",
			})
			if err != nil {
				fmt.Println(err)
			}
		},
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			serialized_dashboard	=	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			require.NotEqual(t, dashboard.UpdateTime, dashboard.CreateTime)
			return nil
		}),
	})
}

func TestAccDashboardTestAll(t *testing.T) {
	dashboard_id := ""
	display_name := ""
	warehouse_id := ""
	etag := ""
	workspaceLevel(t, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			serialized_dashboard	=	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			return nil
		}),
	}, step{
		PreConfig: func() {
			tmpDir := "/tmp/LakeviewDashboardTest"
			fileName := tmpDir + "/Dashboard.json"
			os.Mkdir(tmpDir, 0755)
			os.WriteFile(fileName, []byte("{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page in file\"}]}"), 0644)
		},
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			file_path				=	"/tmp/LakeviewDashboardTest/Dashboard.json"
			embed_credentials		=	false
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			require.NotEqual(t, dashboard.UpdateTime, dashboard.CreateTime)
			require.Equal(t, publish_dash.EmbedCredentials, false)
			return nil
		}),
	}, step{
		PreConfig: func() {
			os.WriteFile("/tmp/LakeviewDashboardTest/Dashboard.json", []byte("{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page Modified\"}]}"), 0644)
		},
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			file_path				=	"/tmp/LakeviewDashboardTest/Dashboard.json"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			dashboard_id = dashboard.DashboardId
			display_name = dashboard.DisplayName
			warehouse_id = dashboard.WarehouseId
			etag = dashboard.Etag
			return nil
		}),
	}, step{
		PreConfig: func() {
			w, err := databricks.NewWorkspaceClient(&databricks.Config{})
			if err != nil {
				fmt.Println(err)
			}
			_, err = w.Lakeview.Update(context.Background(), dashboards.UpdateDashboardRequest{
				DashboardId:         dashboard_id,
				DisplayName:         display_name,
				Etag:                etag,
				WarehouseId:         warehouse_id,
				SerializedDashboard: "{\"pages\":[{\"name\":\"b532570b\",\"displayName\":\"New Page Modified Remote\"}]}",
			})
			if err != nil {
				fmt.Println(err)
			}
		},
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			file_path				=	"/tmp/LakeviewDashboardTest/Dashboard.json"
			parent_path				= 	"/Shared/provider-test"
			embed_credentials		=	false
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			return nil
		}),
	}, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			file_path				=	"/tmp/LakeviewDashboardTest/Dashboard.json"
			parent_path				= 	"/Shared/Teams"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			return nil
		}),
	}, step{
		PreConfig: func() {
			os.WriteFile("/tmp/LakeviewDashboardTest/Dashboard.json", []byte("{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page Modified again\"}]}"), 0644)
		},
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			file_path				=	"/tmp/LakeviewDashboardTest/Dashboard.json"
			parent_path				= 	"/Shared/Teams"
		}
		`,
		Check: resourceCheck("databricks_lakeview_dashboard.d1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
			assert.Equal(t, "Monthly Traffic Report", dashboard.DisplayName)
			fmt.Println(dashboard.SerializedDashboard)
			require.NoError(t, err)
			os.Remove("/tmp/LakeviewDashboardTest/Dashboard.json")
			os.Remove("/tmp/LakeviewDashboardTest")
			return nil
		}),
	})
}
