package catalog_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestUcAccResourceSystemSchema(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_system_schema" "access" {
			schema = "access"
		}
		resource "databricks_system_schema" "billing" {
			schema = "billing"
		}`,
	})
}
