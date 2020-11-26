package acceptance

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	. "github.com/databrickslabs/databricks-terraform/mws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestMwsAccNetworks(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Cannot run test on non-MWS environment")
	}
	var network Network

	// cannot use subnets between network registrations...
	networkResourceConfig := qa.EnvironmentTemplate(t, `
	provider "databricks" {
		host     = "{env.DATABRICKS_HOST}"
		username = "{env.DATABRICKS_USERNAME}"
		password = "{env.DATABRICKS_PASSWORD}"
	}
	resource "databricks_mws_networks" "my_network" {
		account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
		network_name = "network-test-{var.RANDOM}"
		vpc_id       = "vpc-11111111"
		subnet_ids   = [
			"subnet-11111111",
			"subnet-99999999"
		]
		security_group_ids = [
			"sg-99999999"
		]
	}`)
	name := qa.FirstKeyValue(t, networkResourceConfig, "network_name")
	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testMWSNetworkResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: networkResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testMWSNetworkResourceExists("databricks_mws_networks.my_network", &network, t),
					resource.TestCheckResourceAttr("databricks_mws_networks.my_network", "network_name", name),
				),
				Destroy: false,
			},
			{
				Config: networkResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testMWSNetworkResourceExists("databricks_mws_networks.my_network", &network, t),
					resource.TestCheckResourceAttr("databricks_mws_networks.my_network", "network_name", name),
				),
				ExpectNonEmptyPlan: false,
				Destroy:            false,
			},
			{
				PreConfig: func() {
					conn := common.CommonEnvironmentClient()
					err := NewNetworksAPI(context.Background(), conn).Delete(network.AccountID, network.NetworkID)
					if err != nil {
						panic(err)
					}
				},
				Config: networkResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testMWSNetworkResourceExists("databricks_mws_networks.my_network", &network, t),
					resource.TestCheckResourceAttr("databricks_mws_networks.my_network", "network_name", name),
				),
				ExpectNonEmptyPlan: false,
				Destroy:            false,
			},
		},
	})
}

func testMWSNetworkResourceDestroy(s *terraform.State) error {
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_mws_networks" {
			continue
		}
		packagedMWSIds, err := UnpackMWSAccountID(rs.Primary.ID)
		if err != nil {
			return err
		}
		_, err = NewNetworksAPI(context.Background(), client).Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		if err != nil {
			return nil
		}
		return errors.New("resource is not cleaned up")
	}
	return nil
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testMWSNetworkResourceExists(n string, network *Network, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := common.CommonEnvironmentClient()
		packagedMWSIds, err := UnpackMWSAccountID(rs.Primary.ID)
		if err != nil {
			return err
		}
		resp, err := NewNetworksAPI(context.Background(), conn).Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*network = resp
		return nil
	}
}
