package acceptance

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
)

func TestAccGlobalInitScriptResource_Create(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_global_init_script" "this" {
				name = "init-{var.RANDOM}"
				enabled = false
				content_base64 = "ZWNobyBoZWxsbw=="
			}`,
		},
	})
}
