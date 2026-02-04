package dashboards_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAccDashboardFileContentChangeDetection tests that changing the content of a file
// referenced by file_path triggers an update, even when the file_path string itself doesn't change.
//
// This test reproduces ES-1687403: when only the file CONTENT changes (not the file_path string),
// Terraform should detect the change and trigger an update.
//
// The bug: DiffSuppressFunc is only called when there's an attribute diff to suppress.
// When file_path string is unchanged and serialized_dashboard is removed from state by StructToData,
// there's no diff, so DiffSuppressFunc is never called, and the file content change is not detected.
func TestAccDashboardFileContentChangeDetection(t *testing.T) {
	tmpDir := fmt.Sprintf("/tmp/Dashboard-FileContent-%s", qa.RandomName())
	fileName := tmpDir + "/Dashboard.json"
	displayName := fmt.Sprintf("Test Dashboard - %s", qa.RandomName())

	initialContent := `{"pages":[{"name":"page1","displayName":"Initial Page"}]}`
	modifiedContent := `{"pages":[{"name":"page1","displayName":"Modified Page"}]}`

	// Capture update times to verify the dashboard was actually updated
	var initialUpdateTime, afterContentChangeUpdateTime string

	acceptance.WorkspaceLevel(t,
		// Step 1: Create dashboard with file_path pointing to initial content
		acceptance.Step{
			PreConfig: func() {
				os.MkdirAll(tmpDir, 0755)
				os.WriteFile(fileName, []byte(initialContent), 0644)
			},
			Template: fmt.Sprintf(`
				resource "databricks_dashboard" "test" {
					display_name = "%s"
					warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
					parent_path  = "/Shared/provider-test"
					file_path    = "%s"
				}
			`, displayName, fileName),
			Check: acceptance.ResourceCheck("databricks_dashboard.test",
				func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
					assert.Contains(t, dashboard.SerializedDashboard, "Initial Page")
					initialUpdateTime = dashboard.UpdateTime
					return nil
				}),
		},
		// Step 2: Change file content WITHOUT changing file_path string
		// This is the critical test: the file_path in HCL is identical, only the file content differs.
		// Bug: Without the fix, this step detects no changes and the dashboard is NOT updated.
		acceptance.Step{
			PreConfig: func() {
				// Modify the file content - this is the only change
				os.WriteFile(fileName, []byte(modifiedContent), 0644)
			},
			// IMPORTANT: Template is EXACTLY the same as Step 1 - file_path string unchanged
			Template: fmt.Sprintf(`
				resource "databricks_dashboard" "test" {
					display_name = "%s"
					warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
					parent_path  = "/Shared/provider-test"
					file_path    = "%s"
				}
			`, displayName, fileName),
			Check: acceptance.ResourceCheck("databricks_dashboard.test",
				func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
					afterContentChangeUpdateTime = dashboard.UpdateTime

					// CRITICAL ASSERTIONS:
					// 1. The dashboard content should be updated to reflect the new file content
					assert.Contains(t, dashboard.SerializedDashboard, "Modified Page",
						"Dashboard content should be updated when file content changes")

					// 2. The update time should have changed (proving an update occurred)
					require.NotEqual(t, initialUpdateTime, afterContentChangeUpdateTime,
						"Dashboard should have been updated when file content changed - "+
							"this failure indicates ES-1687403: file content changes are not detected")

					return nil
				}),
		},
	)
}
