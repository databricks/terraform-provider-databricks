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
	var network model.MWSNetwork

	// cannot use subnets between network registrations...
	networkResourceConfig := EnvironmentTemplate(t, `
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
	name := FirstKeyValue(t, networkResourceConfig, "network_name")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
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
					conn := service.CommonEnvironmentClient()
					err := conn.MWSNetworks().Delete(network.AccountID, network.NetworkID)
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
	client := service.CommonEnvironmentClient()
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
		return errors.New("resource is not cleaned up")
	}
	return nil
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testMWSNetworkResourceExists(n string, network *model.MWSNetwork, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := service.CommonEnvironmentClient()
		packagedMWSIds, err := unpackMWSAccountID(rs.Primary.ID)
		if err != nil {
			return err
		}
		resp, err := conn.MWSNetworks().Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*network = resp
		return nil
	}
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
			Resource: "/api/2.0/accounts/abc/networks/nid",
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
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceMWSNetworksRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/networks/nid",
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
			Resource: "/api/2.0/accounts/abc/networks/nid",
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
			Resource: "/api/2.0/accounts/abc/networks/nid",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSNetworks, nil, actionWithID("abc/nid", resourceMWSNetworksRead))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/nid", d.Id(), "Id should not be empty for error reads")
}

func TestResourceMWSNetworksDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "DELETE",
			Resource: "/api/2.0/accounts/abc/networks/nid",
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/networks/nid",
			Response: model.MWSNetwork{
				NetworkID:   "nid",
				NetworkName: "Open Workers",
				VPCID:       "five",
				VPCStatus:   "SOMETHING",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/networks/nid",
			Response: service.APIErrorBody{
				ErrorCode: "NOT_FOUND",
				Message:   "Yes, it's not found",
			},
			Status: 404,
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
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/nid", d.Id())
}
