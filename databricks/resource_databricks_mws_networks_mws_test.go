package databricks

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccMWSNetworks(t *testing.T) {
	var MWSNetwork model.MWSNetwork
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	mwsAcctID := os.Getenv("DATABRICKS_MWS_ACCT_ID")
	mwsHost := os.Getenv("DATABRICKS_MWS_HOST")
	networkName := "test-mws-network-tf"

	vpc := "vpc-11111111"
	subnet1 := "subnet-11111111"
	subnet2 := "subnet-99999999"
	sg1 := "sg-11111111"
	sg2 := "sg-99999999"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testMWSNetworkResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testMWSNetworkCreate(mwsAcctID, mwsHost, networkName, vpc, subnet1, subnet2, sg1, sg2),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testMWSNetworkResourceExists("databricks_mws_networks.my_network", &MWSNetwork, t),
					// verify local values
					resource.TestCheckResourceAttr("databricks_mws_networks.my_network", "account_id", mwsAcctID),
					resource.TestCheckResourceAttr("databricks_mws_networks.my_network", "network_name", networkName),
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testMWSNetworkCreate(mwsAcctID, mwsHost, networkName, vpc, subnet1, subnet2, sg1, sg2),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testMWSNetworkResourceExists("databricks_mws_networks.my_network", &MWSNetwork, t),
					// verify local values
					resource.TestCheckResourceAttr("databricks_mws_networks.my_network", "account_id", mwsAcctID),
					resource.TestCheckResourceAttr("databricks_mws_networks.my_network", "network_name", networkName),
				),
				ExpectNonEmptyPlan: false,
				Destroy:            false,
			},
			{
				PreConfig: func() {
					conn := getMWSClient()
					err := conn.MWSNetworks().Delete(MWSNetwork.AccountID, MWSNetwork.NetworkID)
					if err != nil {
						panic(err)
					}
				},
				// use a dynamic configuration with the random name from above
				Config: testMWSNetworkCreate(mwsAcctID, mwsHost, networkName, vpc, subnet1, subnet2, sg1, sg2),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testMWSNetworkResourceExists("databricks_mws_networks.my_network", &MWSNetwork, t),
					// verify local values
					resource.TestCheckResourceAttr("databricks_mws_networks.my_network", "account_id", mwsAcctID),
					resource.TestCheckResourceAttr("databricks_mws_networks.my_network", "network_name", networkName),
				),
				ExpectNonEmptyPlan: false,
				Destroy:            false,
			},
		},
	})
}

func testMWSNetworkResourceDestroy(s *terraform.State) error {
	client := getMWSClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_mws_storage_configurations" {
			continue
		}
		packagedMWSIds, err := unpackMWSAccountID(rs.Primary.ID)
		if err != nil {
			return err
		}
		_, err = client.MWSNetworks().Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		if err != nil {
			return nil
		}
		return errors.New("resource Scim Group is not cleaned up")
	}
	return nil
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testMWSNetworkResourceExists(n string, mwsCreds *model.MWSNetwork, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := getMWSClient()
		packagedMWSIds, err := unpackMWSAccountID(rs.Primary.ID)
		if err != nil {
			return err
		}
		resp, err := conn.MWSNetworks().Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*mwsCreds = resp
		return nil
	}
}

func testMWSNetworkCreate(mwsAcctID, mwsHost, networkName, vpcID, subnetID1, subnetID2, sgID1, sgID2 string) string {
	return fmt.Sprintf(`
								provider "databricks" {
								  host = "%s"
								  basic_auth {}
								}
								resource "databricks_mws_networks" "my_network" {
								  account_id = "%s"
								  network_name = "%s"
								  vpc_id = "%s"
								  subnet_ids = [
									"%s",
									"%s",
								  ]
								  security_group_ids = [
									"%s",
									"%s",
								  ]
								}

								`, mwsHost, mwsAcctID, networkName, vpcID, subnetID1, subnetID2, sgID1, sgID2)
}
