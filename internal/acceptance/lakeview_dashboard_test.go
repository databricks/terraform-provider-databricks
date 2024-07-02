package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccLakeviewDashboard(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			file_path				=	"/Users/divyansh.vijayvergia/terraform-provider-databricks/Experimental-Divyansh/json_file"
			parent_path				= 	"/Users/divyansh.vijayvergia@databricks.com"
		}
		`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
		),
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
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "embed_credentials", "false"),
		),
	}, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			serialized_dashboard	=	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page Modified\"}]}"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "embed_credentials", "true"),
		),
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
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
		),
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
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
		),
	})
}

func TestAccDashboardWithNoChange(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			serialized_dashboard	=	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
		),
	}, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			serialized_dashboard	=	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
			parent_path				= 	"/Shared/provider-test"
		}
		`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
		),
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
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
			func(s *terraform.State) error {
				dashboard, ok := s.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				dashboard_id = dashboard.Primary.ID
				assert.NotEmpty(t, dashboard_id)
				display_name = dashboard.Primary.Attributes["display_name"]
				assert.NotEmpty(t, display_name)
				warehouse_id = dashboard.Primary.Attributes["warehouse_id"]
				assert.NotEmpty(t, warehouse_id)
				etag = dashboard.Primary.Attributes["etag"]
				assert.NotEmpty(t, etag)
				return nil
			},
		),
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
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
		),
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
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
		),
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
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "embed_credentials", "false"),
		),
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
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
			func(s *terraform.State) error {
				dashboard, ok := s.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				dashboard_id = dashboard.Primary.ID
				assert.NotEmpty(t, dashboard_id)
				display_name = dashboard.Primary.Attributes["display_name"]
				assert.NotEmpty(t, display_name)
				warehouse_id = dashboard.Primary.Attributes["warehouse_id"]
				assert.NotEmpty(t, warehouse_id)
				etag = dashboard.Primary.Attributes["etag"]
				assert.NotEmpty(t, etag)
				return nil
			},
		),
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
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
		),
	}, step{
		Template: `
		resource "databricks_lakeview_dashboard" "d1" {
			display_name			= 	"Monthly Traffic Report"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			file_path				=	"/tmp/LakeviewDashboardTest/Dashboard.json"
			parent_path				= 	"/Shared/Teams"
		}
		`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_lakeview_dashboard.d1", "display_name", "Monthly Traffic Report"),
			func(state *terraform.State) error {
				dashboard, ok := state.RootModule().Resources["databricks_lakeview_dashboard.d1"]
				if !ok {
					return fmt.Errorf("dashboard resource not found")
				}
				fmt.Println(dashboard.Primary.Attributes["serialized_dashboard"])
				return nil
			},
		),
	})
}
