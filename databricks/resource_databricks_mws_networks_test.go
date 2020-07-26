package databricks

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestMwsAccNetworks(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Cannot run test on non-MWS environment")
	}
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

func TestResourceMWSNetworksCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/abc/networks",
			ExpectedRequest: model.MWSNetwork{
				SecurityGroupIds: []string{"one", "two"},
				NetworkName:      "Open Workers",
				VPCID:            "five",
				SubnetIds:        []string{"four", "three"},
			},
			Response: model.MWSNetwork{
				NetworkID: "nid",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/networks/nid?",
			Response: model.MWSNetwork{
				NetworkID:        "nid",
				SecurityGroupIds: []string{"one", "two"},
				NetworkName:      "Open Workers",
				VPCID:            "five",
				SubnetIds:        []string{"four", "three"},
			},
		},
	}, resourceMWSNetworks, map[string]interface{}{
		"account_id":         "abc",
		"network_name":       "Open Workers",
		"security_group_ids": []interface{}{"one", "two"},
		"subnet_ids":         []interface{}{"three", "four"},
		"vpc_id":             "five",
	}, resourceMWSNetworksCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/nid", d.Id())
}

func TestResourceMWSNetworksCreate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/abc/networks",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSNetworks, map[string]interface{}{
		"account_id":         "abc",
		"network_name":       "Open Workers",
		"security_group_ids": []interface{}{"one", "two"},
		"subnet_ids":         []interface{}{"three", "four"},
		"vpc_id":             "five",
	}, resourceMWSNetworksCreate)
	assert.Errorf(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceMWSNetworksRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/networks/nid?",
			Response: model.MWSNetwork{
				NetworkID:        "nid",
				SecurityGroupIds: []string{"one", "two"},
				NetworkName:      "Open Workers",
				VPCID:            "five",
				SubnetIds:        []string{"four", "three"},
				WorkspaceID:      789,
			},
		},
	}, resourceMWSNetworks, nil, actionWithID("abc/nid", resourceMWSNetworksRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/nid", d.Id(), "Id should not be empty")
	assert.Equal(t, 0, d.Get("creation_time"))
	assert.Equal(t, "nid", d.Get("network_id"))
	assert.Equal(t, "Open Workers", d.Get("network_name"))
	assert.Equal(t, "five", d.Get("vpc_id"))
	assert.Equal(t, 789, d.Get("workspace_id"))
}

func TestResourceMWSNetworksRead_NotFound(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/networks/nid?",
			Response: service.APIErrorBody{
				ErrorCode: "NOT_FOUND",
				Message:   "Item not found",
			},
			Status: 404,
		},
	}, resourceMWSNetworks, nil, actionWithID("abc/nid", resourceMWSNetworksRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceMWSNetworksRead_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/networks/nid?",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSNetworks, nil, actionWithID("abc/nid", resourceMWSNetworksRead))
	assert.Errorf(t, err, "Internal error happened")
	assert.Equal(t, "abc/nid", d.Id(), "Id should not be empty for error reads")
}

func TestResourceMWSNetworksDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "DELETE",
			Resource: "/api/2.0/accounts/abc/networks/nid",
		},
	}, resourceMWSNetworks, nil, actionWithID("abc/nid", resourceMWSNetworksDelete))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/nid", d.Id())
}

func TestResourceMWSNetworksDelete_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "DELETE",
			Resource: "/api/2.0/accounts/abc/networks/nid",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSNetworks, nil, actionWithID("abc/nid", resourceMWSNetworksDelete))
	assert.Errorf(t, err, "Internal error happened")
	assert.Equal(t, "abc/nid", d.Id())
}
