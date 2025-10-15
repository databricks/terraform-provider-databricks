package app_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccAppsDataSource(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_apps" "this" { }
		`,
	})
}
