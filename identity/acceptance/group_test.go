package acceptance

import (
	"errors"
	"fmt"
	"os"

	"github.com/databrickslabs/databricks-terraform/common"
	. "github.com/databrickslabs/databricks-terraform/identity"
	"github.com/stretchr/testify/assert"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"testing"
)

func TestAccUserMemberships(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	config := qa.EnvironmentTemplate(t, `
	data "databricks_group" "admins" {
		display_name = "admins"
	}
	
	resource "databricks_user" "me" {
		user_name    = "me+{var.RANDOM}@example.com"
	}
	
	resource "databricks_group_member" "my_member_a" {
		group_id = data.databricks_group.admins.id
		member_id = databricks_user.me.id
	}`)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config:             config,
				Destroy:            false,
				ExpectNonEmptyPlan: true,
				// Check: resource.ComposeTestCheckFunc(
				// 	resource.TestCheckResourceAttr("databricks_user.first", "allow_cluster_create", "false"),
				// 	resource.TestCheckResourceAttr("databricks_user.first", "allow_instance_pool_create", "false"),
				// 	resource.TestCheckResourceAttr("databricks_user.second", "allow_cluster_create", "true"),
				// 	resource.TestCheckResourceAttr("databricks_user.second", "allow_instance_pool_create", "false"),
				// 	resource.TestCheckResourceAttr("databricks_user.third", "allow_cluster_create", "false"),
				// 	resource.TestCheckResourceAttr("databricks_user.third", "allow_instance_pool_create", "true"),
				// 	func(s *terraform.State) error {
				// 		r := s.RootModule().Resources
				// 		client := common.CommonEnvironmentClient()
				// 		return NewGroupsAPI(client).Patch(r["databricks_group.first"].Primary.ID, []string{
				// 			r["databricks_user.first"].Primary.ID,
				// 			r["databricks_user.second"].Primary.ID,
				// 		}, nil, GroupMembersPath)
				// 	},
				// ),
			},
		},
	})
}

func TestAccGroupResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	var Group ScimGroup
	randomStr := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	displayName := fmt.Sprintf("tf group test %s", randomStr)
	newDisplayName := fmt.Sprintf("new tf group test %s", randomStr)
	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testGroupResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testDatabricksGroup(displayName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testGroupResourceExists("databricks_group.my_group", &Group, t),
					// verify remote values
					testGroupValues(t, &Group, displayName),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "display_name", displayName),
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testDatabricksGroup(newDisplayName),
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

func TestAccGroupResource_verify_entitlements(t *testing.T) {
	var Group ScimGroup
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	randomStr := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	displayName := fmt.Sprintf("tf group test %s", randomStr)
	newDisplayName := fmt.Sprintf("new tf group test %s", randomStr)
	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testGroupResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testDatabricksGroupEntitlements(displayName, "true", "true"),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testGroupResourceExists("databricks_group.my_group", &Group, t),
					// verify remote values
					testGroupValues(t, &Group, displayName),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "allow_cluster_create", "true"),
					resource.TestCheckResourceAttr("databricks_group.my_group", "allow_instance_pool_create", "true"),
				),
				Destroy: false,
			},
			// Remove entitlements and expect a non empty plan
			{
				// use a dynamic configuration with the random name from above
				Config: testDatabricksGroup(newDisplayName),
				// test to see if new resource is attempted to be planned
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
		},
	})
}

func testGroupResourceDestroy(s *terraform.State) error {
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_group" {
			continue
		}
		_, err := NewUsersAPI(client).Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource Group is not cleaned up")
	}
	return nil
}

func testGroupValues(t *testing.T, group *ScimGroup, displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, group.DisplayName == displayName)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testGroupResourceExists(n string, group *ScimGroup, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := common.CommonEnvironmentClient()
		resp, err := NewGroupsAPI(conn).Read(rs.Primary.ID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*group = resp
		return nil
	}
}

func testDatabricksGroup(groupName string) string {
	return fmt.Sprintf(`
		resource "databricks_group" "my_group" {
			display_name = "%s"
		}
		`, groupName)
}

func testDatabricksGroupEntitlements(groupName, allowClusterCreate, allowPoolCreate string) string {
	return fmt.Sprintf(`
		resource "databricks_group" "my_group" {
			display_name = "%s"
			allow_cluster_create = %s
			allow_instance_pool_create = %s
		}
		`, groupName, allowClusterCreate, allowPoolCreate)
}
