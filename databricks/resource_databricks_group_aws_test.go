package databricks

import (
	"errors"
	"fmt"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccAWSGroupResource(t *testing.T) {
	var Group model.Group
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	randomStr := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	displayName := fmt.Sprintf("tf group test %s", randomStr)
	newDisplayName := fmt.Sprintf("new tf group test %s", randomStr)
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAWSGroupResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testAWSDatabricksGroup(displayName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAWSGroupResourceExists("databricks_group.my_group", &Group, t),
					// verify remote values
					testAWSGroupValues(t, &Group, displayName),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "display_name", displayName),
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testAWSDatabricksGroup(newDisplayName),
				// test to see if new resource is attempted to be planned
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
			{
				ResourceName:      "databricks_group.my_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAWSGroupResource_verify_entitlements(t *testing.T) {
	var Group model.Group
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	randomStr := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	displayName := fmt.Sprintf("tf group test %s", randomStr)
	newDisplayName := fmt.Sprintf("new tf group test %s", randomStr)
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAWSGroupResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testAWSDatabricksGroupEntitlements(displayName, "true", "true"),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAWSGroupResourceExists("databricks_group.my_group", &Group, t),
					// verify remote values
					testAWSGroupValues(t, &Group, displayName),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "allow_cluster_create", "true"),
					resource.TestCheckResourceAttr("databricks_group.my_group", "allow_instance_pool_create", "true"),
				),
				Destroy: false,
			},
			// Remove entitlements and expect a non empty plan
			{
				// use a dynamic configuration with the random name from above
				Config: testAWSDatabricksGroup(newDisplayName),
				// test to see if new resource is attempted to be planned
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
		},
	})
}

func testAWSGroupResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*service.DBApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_group" {
			continue
		}
		_, err := client.Users().Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource Group is not cleaned up")
	}
	return nil
}

func testAWSGroupValues(t *testing.T, group *model.Group, displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, group.DisplayName == displayName)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testAWSGroupResourceExists(n string, group *model.Group, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DBApiClient)
		resp, err := conn.Groups().Read(rs.Primary.ID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*group = resp
		return nil
	}
}

func testAWSDatabricksGroup(groupName string) string {
	return fmt.Sprintf(`
								resource "databricks_group" "my_group" {
								  display_name = "%s"
								}
								`, groupName)
}

func testAWSDatabricksGroupEntitlements(groupName, allowClusterCreate, allowPoolCreate string) string {
	return fmt.Sprintf(`
								resource "databricks_group" "my_group" {
								  display_name = "%s"
								  allow_cluster_create = %s
								  allow_instance_pool_create = %s
								}
								`, groupName, allowClusterCreate, allowPoolCreate)
}
