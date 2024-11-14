package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataSourceMwsNetworkConnectivityConfigTest(t *testing.T) {
	loadWorkspaceEnv(t)
	if isGcp(t) {
		skipf(t)("GCP not supported")
	}
	AccountLevel(t,
		Step{
			Template: `
			resource "databricks_mws_network_connectivity_config" "this" {
				name = "tf-{var.RANDOM}"
				region = "eastus2"
			}	

			data "databricks_mws_network_connectivity_config" "this" {
			  depends_on = [databricks_mws_network_connectivity_config.this]
			  name = databricks_mws_network_connectivity_config.this.name
			}`,
			Check: func(s *terraform.State) error {
				r, ok := s.RootModule().Resources["data.databricks_mws_network_connectivity_config.this"]
				if !ok {
					return fmt.Errorf("data not found in state")
				}
				name := r.Primary.Attributes["name"]
				if name == "" {
					return fmt.Errorf("name is empty: %v", r.Primary.Attributes)
				}
				expect := "eastus2"
				region := r.Primary.Attributes["region"]
				if region != expect {
					return fmt.Errorf("incorrect region. expected: %v, received: %v",
						expect, region)
				}
				return nil
			},
		})
}
