package catalog_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestUcAccDataSourceMetastore(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		data "databricks_metastore" "this" {
			metastore_id = "{env.TEST_METASTORE_ID}"
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_metastore.this"]
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
