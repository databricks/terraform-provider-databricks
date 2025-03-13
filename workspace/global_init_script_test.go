package workspace_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccGlobalInitScriptResource_Create(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_global_init_script" "this" {
			name = "init-{var.RANDOM}"
			enabled = true
			content_base64 = "ZWNobyBoZWxsbw=="
		}`,
	})
}
