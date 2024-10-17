package acceptance

import (
	"fmt"
	"testing"
)

var (
	appTemplate = `
	resource "databricks_workspace_file" "this" {
		source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
		path = "/Shared/%s/xx_{var.RANDOM}"
	}
	resource "databricks_app" "this" {
		name = "my-custom-app"
		description = "%s"
		source_code_path = databricks_workspace_file.this.workspace_path
		mode = "SNAPSHOT"
	}`
)

func TestAccAppCreate(t *testing.T) {
	loadWorkspaceEnv(t)
	if isGcp(t) {
		skipf(t)("not available on GCP")
	}
	WorkspaceLevel(t, Step{
		Template: fmt.Sprintf(appTemplate, "app", "My app"),
	})
}

func TestAccAppUpdate(t *testing.T) {
	loadWorkspaceEnv(t)
	if isGcp(t) {
		skipf(t)("not available on GCP")
	}
	WorkspaceLevel(t, Step{
		Template: fmt.Sprintf(budgetTemplate, "app", "My app"),
	}, Step{
		Template: fmt.Sprintf(budgetTemplate, "app", "My new app"),
	})
}
