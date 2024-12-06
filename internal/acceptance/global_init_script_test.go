package acceptance

import (
	"testing"
)

func TestAccGlobalInitScriptResource_Create(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `
		resource "databricks_global_init_script" "this" {
			name = "init-{var.RANDOM}"
			enabled = true
			content_base64 = "ZWNobyBoZWxsbw=="
		}`,
	})
}
