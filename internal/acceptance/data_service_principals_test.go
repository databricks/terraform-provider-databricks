package acceptance

import (
	"testing"
)

func TestAccDataSourceSPNs(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		data databricks_service_principals "this" {
			display_name_contains = ""
		}`,
	})
}
