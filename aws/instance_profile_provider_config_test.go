package aws_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func instanceProfileProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_instance_profile" "this" {
		instance_profile_arn = "arn:aws:iam::999999999999:instance-profile/fake-profile"
		skip_validation = true
		%s
	}
	`, providerConfig)
}

func TestAccInstanceProfile_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: instanceProfileProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
