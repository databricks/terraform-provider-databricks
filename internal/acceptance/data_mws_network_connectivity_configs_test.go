package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataSourceMwsNetworkConnectivityConfigsTest(t *testing.T) {
	LoadWorkspaceEnv(t)
	if IsGcp(t) {
		Skipf(t)("GCP not supported")
	}
	var region string
	if IsAzure(t) {
		region = "eastus2"
	} else if IsAws(t) {
		region = "us-east-2"
	}
	AccountLevel(t,
		Step{
			Template: fmt.Sprintf(`
			resource "databricks_mws_network_connectivity_configs" "this" {
				name = "tf-{var.RANDOM}"
				region = "%s"
			}	

			data "databricks_mws_network_connectivity_configs" "this" {
			  depends_on = [databricks_mws_network_connectivity_config.this]
			  region = databricks_mws_network_connectivity_config.this.region
			}`, region),
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
