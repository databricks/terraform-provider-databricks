package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestUcAccDataSourceAwsCrossaccountPolicy(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET") // marker for AWS test env
	accountLevel(t, step{
		Template: `
		data "databricks_aws_crossaccount_policy" "this" {
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_aws_crossaccount_policy.this"]
			if !ok {
				return fmt.Errorf("data not found in state")
			}
			policy := r.Primary.Attributes["json"]
			if policy == "" {
				return fmt.Errorf("CrossAccount Policy JSON is empty: %v", r.Primary.Attributes)
			}
			return nil
		},
	})
}
