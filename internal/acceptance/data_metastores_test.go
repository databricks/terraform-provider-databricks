package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestUcAccDataSourceMetastores(t *testing.T) {
	accountLevel(t, step{
		Template: `
		data "databricks_metastores" "this" {
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_metastores.this"]
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
