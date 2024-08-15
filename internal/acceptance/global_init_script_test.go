package acceptance

import (
	"testing"
)

func TestAccGlobalInitScriptResource_Create(t *testing.T) {
	workspaceLevel(t, LegacyStep{
		Template: `
		resource "databricks_global_init_script" "this" {
			name = "init-{var.RANDOM}"
			enabled = true
			content_base64 = "ZWNobyBoZWxsbw=="
		}`,
	})
}
