package dashboards_test

import (
	"strconv"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
)

func CheckDataSourceDashboardsPopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.Modules[0].Resources["data.databricks_dashboards.this"]
		require.True(t, ok, "data.databricks_dashboards.this has to be there")

		dashboardsCount := ds.Primary.Attributes["databricks_dashboards.#"]
		dashboardsCountInt, err := strconv.Atoi(dashboardsCount)
		require.NoError(t, err, "dashboards count is not a number")
		require.NotEqual(t, 0, dashboardsCountInt, "dashboard list is empty")

		return nil
	}
}

func TestAccDashboardsCreation(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
			resource "databricks_dashboard" "dashboard" {
				display_name         = "New Dashboard"
				warehouse_id         = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
				serialized_dashboard = "{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
				embed_credentials    = false // Optional
				parent_path          = "/Shared/provider-test"
			}

			data "databricks_dashboards" "this" {
				depends_on = [databricks_dashboard.dashboard]
			}
		`,
		Check: CheckDataSourceDashboardsPopulated(t),
	})
}
