package acceptance

import (
	"testing"
)

func TestUcAccResourceSystemSchema(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_system_schema" "this" {
			system_schema = ["access", "billing"]
			}`,
	})
}
