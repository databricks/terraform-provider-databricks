package databricks

import (
	"errors"
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccAWSGroupRoleResource(t *testing.T) {
	var group model.Group
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest

	groupName := "group test"
	role := "arn:aws:iam::999999999999:instance-profile/terraform-group-test"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAWSGroupRoleResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testAWSGroupRoleResourceCreate(role, groupName),

				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAWSGroupRoleResourceExists("databricks_group.my_group", &group, t),
					// verify remote values
					testAWSGroupRoleValues(t, &group, groupName, role),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "display_name", groupName),
					resource.TestCheckResourceAttr("databricks_group_role.my_group_role", "instance_profile_id", role),

				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testAWSGroupRoleResourceCreate(role, groupName),

				// Test behavior to expect to attempt to create new role mapping because role is gone
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DBApiClient)
					client.InstanceProfiles().Delete(role)
				},
				PlanOnly: true,
				ExpectNonEmptyPlan: true,
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testAWSGroupRoleResourceCreate(role, groupName),

				// Test behavior to expect to attempt to create new role mapping because role is gone
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DBApiClient)
					client.Groups().Delete(group.ID)
				},
				PlanOnly: true,
				ExpectNonEmptyPlan: true,
				Destroy: false,
			},
		},
	})
}

func testAWSGroupRoleResourceDestroy(s *terraform.State) error {
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

func testAWSGroupRoleValues(t *testing.T, group *model.Group, displayName, role string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, group.DisplayName == displayName)
		assert.True(t, iRoleInGroup(role, group), "role is not in group")
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testAWSGroupRoleResourceExists(n string, group *model.Group, t *testing.T) resource.TestCheckFunc {
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

func testAWSGroupRoleResourceCreate(roleArn, groupName string) string {
	return fmt.Sprintf(`
								resource "databricks_instance_profile" "instance_profile" {
								  instance_profile_arn = "%s"
								  skip_validation = true
								}
								resource "databricks_group" "my_group" {
								  display_name = "%s"
								}
								resource "databricks_group_role" "my_group_role" {
								 group_id = databricks_group.my_group.id
								 instance_profile_id = databricks_instance_profile.instance_profile.id
								}
								`, roleArn, groupName)
}
