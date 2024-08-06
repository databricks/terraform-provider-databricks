package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataSourceMwsCredentials(t *testing.T) {
	accountLevel(t, step{
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
			return nil
		},
	})
}
