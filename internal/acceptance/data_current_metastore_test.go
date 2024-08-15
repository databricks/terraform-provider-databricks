package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestUcAccDataSourceCurrentMetastore(t *testing.T) {
	workspaceLevel(t, LegacyStep{
		Template: `
		data "databricks_current_metastore" "this" {
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_current_metastore.this"]
			if !ok {
				return fmt.Errorf("data not found in state")
			}
			metastore_info := r.Primary.Attributes["metastore_info.0.%"]
			if metastore_info == "" {
				return fmt.Errorf("MetastoreInfo is empty: %v", r.Primary.Attributes)
			}
			return nil
		},
	})
}
