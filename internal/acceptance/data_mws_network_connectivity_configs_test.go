package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataSourceMwsNetworkConnectivityConfigsTest(t *testing.T) {
	loadWorkspaceEnv(t)
	if isGcp(t) {
		skipf(t)("GCP not supported")
	}
	AccountLevel(t,
		Step{
			Template: `
			resource "databricks_mws_network_connectivity_configs" "this" {
				name = "tf-{var.RANDOM}"
				region = "eastus2"
			}	

			data "databricks_mws_network_connectivity_configs" "this" {
			  depends_on = [databricks_mws_network_connectivity_config.this]
			  region = databricks_mws_network_connectivity_config.this.region
			}`,
			Check: func(s *terraform.State) error {
				r, ok := s.RootModule().Resources["data.databricks_mws_network_connectivity_configs.this"]
				if !ok {
					return fmt.Errorf("data not found in state")
				}
				names := r.Primary.Attributes["names"]
				if names == "" {
					return fmt.Errorf("names is empty: %v", r.Primary.Attributes)
				}
				return nil
			},
		})
}
