package mws_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataSourceMwsCredentials(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		data "databricks_mws_credentials" "this" {
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_mws_credentials.this"]
			if !ok {
				return fmt.Errorf("data not found in state")
			}
			ids := r.Primary.Attributes["ids.%"]
			if ids == "" {
				return fmt.Errorf("ids is empty: %v", r.Primary.Attributes)
			}
			// Regression check for https://github.com/databricks/terraform-provider-databricks/issues/5664.
			// databricks_mws_credentials is account-only and must not carry a
			// provider_config block in state — the post-Read hook that populates
			// it is keyed off schema presence, so the absence of any
			// provider_config.* attribute confirms the schema is clean and the
			// hook never ran.
			for k := range r.Primary.Attributes {
				if strings.HasPrefix(k, "provider_config") {
					return fmt.Errorf("unexpected provider_config in state: %s=%q", k, r.Primary.Attributes[k])
				}
			}
			return nil
		},
	})
}
