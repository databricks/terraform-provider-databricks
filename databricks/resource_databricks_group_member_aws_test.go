package databricks

import (
	"errors"
	"fmt"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccAWSGroupMemberResource(t *testing.T) {
	var group model.Group
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest

	groupName := "group test"
	var manuallyCreatedGroup *model.Group
	defer func() {
		client := testAccProvider.Meta().(*service.DBApiClient)
		if client != nil && manuallyCreatedGroup != nil {
			err := client.Groups().Delete(manuallyCreatedGroup.ID)
			assert.NoError(t, err, err)
		}
	}()

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAWSGroupMemberResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				// Creates 2 sub groups and adds as members to parent group
				Config: testAWSGroupMemberResource(groupName),

				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAWSGroupMemberResourceExists("databricks_group.my_group", &group, t),
					// verify remote values
					testAWSGroupMemberValues(t, &group, groupName, 2),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "display_name", groupName),
				),
				Destroy: false,
			},
			{
				PreConfig: func() {
					//	manually create subgroup c
					client := testAccProvider.Meta().(*service.DBApiClient)
					subGroupC, _ := client.Groups().Create("manually-created-group", nil, nil, nil)
					manuallyCreatedGroup = &subGroupC
					//  Add new subgroup to current group
					err := client.Groups().Patch(group.ID, []string{manuallyCreatedGroup.ID}, nil, model.GroupMembersPath)
					assert.NoError(t, err, err)
				},
				Config: testAWSGroupMemberResource(groupName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAWSGroupMemberResourceExists("databricks_group.my_group", &group, t),
					// verify remote values with manually added group
					testAWSGroupMemberValues(t, &group, groupName, 3),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "display_name", groupName),
				),
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testAWSGroupMemberResourceCreateNoMembers(groupName),

				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAWSGroupMemberResourceExists("databricks_group.my_group", &group, t),
					// verify remote values (we added a group manually in the prior test reflecting scim sync)
					testAWSGroupMemberValues(t, &group, groupName, 1),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "display_name", groupName),
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testAWSGroupMemberResourceCreateNoMembers(groupName),

				// Test behavior to expect to attempt to create new role mapping because role is gone
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DBApiClient)
					err := client.Groups().Delete(group.ID)
					assert.NoError(t, err, err)
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testAWSGroupMemberResource(groupName),

				// Lets delete the manually created group
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DBApiClient)
					err := client.Groups().Delete(manuallyCreatedGroup.ID)
					assert.NoError(t, err, err)
					manuallyCreatedGroup = nil
				},
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAWSGroupMemberResourceExists("databricks_group.my_group", &group, t),
					// verify remote values
					testAWSGroupMemberValues(t, &group, groupName, 2),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "display_name", groupName),
				),
				Destroy: false,
			},
		},
	})
}

func testAWSGroupMemberResourceDestroy(s *terraform.State) error {
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

func testAWSGroupMemberValues(t *testing.T, group *model.Group, displayName string, memberCount int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, group.DisplayName == displayName)
		assert.Equal(t, memberCount, len(group.Members), "member count is not matching")
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testAWSGroupMemberResourceExists(n string, group *model.Group, t *testing.T) resource.TestCheckFunc {
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

func testAWSGroupMemberResourceCreateNoMembers(groupName string) string {
	return fmt.Sprintf(`
								resource "databricks_group" "my_group" {
								  display_name = "%s"
								}
								`, groupName)
}

func testAWSGroupMemberResource(groupName string) string {
	return fmt.Sprintf(`
								resource "databricks_group" "my_group" {
								  display_name = "%[1]s"
								}
								resource "databricks_group" "my_sub_group_a" {
								  display_name = "sub_a_%[1]s"
								}
								resource "databricks_group" "my_sub_group_b" {
								  display_name = "sub_b_%[1]s"
								}
								resource "databricks_group_member" "my_member_a" {
								 group_id = databricks_group.my_group.id
								 member_id = databricks_group.my_sub_group_a.id
								}
								resource "databricks_group_member" "my_member_b" {
								 group_id = databricks_group.my_group.id
								 member_id = databricks_group.my_sub_group_b.id
								}
								`, groupName)
}
