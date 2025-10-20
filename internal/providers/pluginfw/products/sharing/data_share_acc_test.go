package sharing_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func dataShareTemplate(provider_config string) string {
	return fmt.Sprintf(`
	resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-share-config"
			object {
				name = databricks_schema.schema1.id
				data_object_type = "SCHEMA"
			}
	}
	data "databricks_share" "this" {
		name = databricks_share.myshare.name
		%s
	}
`, provider_config)
}

func TestAccDataShare_ProviderConfig_Invalid(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + dataShareTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`(?s)failed to get workspace client.*failed to parse workspace_id.*valid integer`),
	})
}
