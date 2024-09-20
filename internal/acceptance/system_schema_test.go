package acceptance

import (
	"os"
	"testing"
)

func TestUcAccResourceSystemSchema(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		t.Skipf("databricks_system_schema resource not available on GCP")
	}
	UnityWorkspaceLevel(t, Step{
		Template: `
		resource "databricks_system_schema" "this" {
			schema = "access"
		}`,
	})
}
