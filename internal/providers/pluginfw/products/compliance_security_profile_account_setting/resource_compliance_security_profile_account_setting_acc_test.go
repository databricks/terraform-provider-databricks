package compliance_security_profile_account_setting_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccCspAccountSettingResource(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
        resource "databricks_compliance_security_profile_account_setting" "this" {
            csp_enablement_account {
                is_enforced          = true
                compliance_standards = ["NONE"]
            }
        }
        `,
	})
}
