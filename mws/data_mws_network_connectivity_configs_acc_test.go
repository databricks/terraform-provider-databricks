package mws_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataSourceMwsNetworkConnectivityConfigsTest(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("GCP not supported")
	}
	var region string
	if acceptance.IsAzure(t) {
		region = "eastus2"
	} else if acceptance.IsAws(t) {
		region = "us-east-2"
	}
	acceptance.AccountLevel(t,
		acceptance.Step{
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
