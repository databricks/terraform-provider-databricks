package app_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccAppsDataSource(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("not available on GCP")
	}
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_apps" "this" { }
		`,
	})
}
