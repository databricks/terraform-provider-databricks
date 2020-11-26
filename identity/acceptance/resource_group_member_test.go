package acceptance

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/identity"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccGroupMemberResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	var group ScimGroup
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	groupName := "Terraform Integration " + randomName
	var manuallyCreatedGroup *ScimGroup
	defer func() {
		client := common.CommonEnvironmentClient()
		if client != nil && manuallyCreatedGroup != nil {
			err := NewGroupsAPI(client).Delete(manuallyCreatedGroup.ID)
			assert.NoError(t, err, err)
		}
	}()

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testGroupMemberResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				// Creates 2 sub groups and adds as members to parent group
				Config: testGroupMemberResource(groupName),

				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testGroupMemberResourceExists("databricks_group.my_group", &group, t),
					// verify remote values
					testGroupMemberValues(t, &group, groupName, 2),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "display_name", groupName),
				),
				Destroy: false,
			},
			{
				PreConfig: func() {
					//	manually create subgroup c
					client := common.CommonEnvironmentClient()
					subGroupC, _ := NewGroupsAPI(client).Create("manually-created-group-"+randomName, nil, nil, nil)
					manuallyCreatedGroup = &subGroupC
					//  Add new subgroup to current group
					err := NewGroupsAPI(client).Patch(group.ID, []string{manuallyCreatedGroup.ID}, nil, GroupMembersPath)
					assert.NoError(t, err, err)
				},
				Config: testGroupMemberResource(groupName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testGroupMemberResourceExists("databricks_group.my_group", &group, t),
					// verify remote values with manually added group
					testGroupMemberValues(t, &group, groupName, 3),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "display_name", groupName),
				),
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testGroupMemberResourceCreateNoMembers(groupName),

				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testGroupMemberResourceExists("databricks_group.my_group", &group, t),
					// verify remote values (we added a group manually in the prior test reflecting scim sync)
					testGroupMemberValues(t, &group, groupName, 1),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "display_name", groupName),
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testGroupMemberResourceCreateNoMembers(groupName),

				// Test behavior to expect to attempt to create new role mapping because role is gone
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewGroupsAPI(client).Delete(group.ID)
					assert.NoError(t, err, err)
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testGroupMemberResource(groupName),

				// Lets delete the manually created group
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewGroupsAPI(client).Delete(manuallyCreatedGroup.ID)
					assert.NoError(t, err, err)
					manuallyCreatedGroup = nil
				},
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testGroupMemberResourceExists("databricks_group.my_group", &group, t),
					// verify remote values
					testGroupMemberValues(t, &group, groupName, 2),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "display_name", groupName),
				),
				Destroy: false,
			},
		},
	})
}

func testGroupMemberResourceDestroy(s *terraform.State) error {
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_group" {
			continue
		}
		ctx := context.Background()
		usersAPI := NewUsersAPI(ctx, client)
		_, err := usersAPI.Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource Group is not cleaned up")
	}
	return nil
}

func testGroupMemberValues(t *testing.T, group *ScimGroup, displayName string, memberCount int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, group.DisplayName == displayName)
		assert.Equal(t, memberCount, len(group.Members), "member count is not matching")
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testGroupMemberResourceExists(n string, group *ScimGroup, t *testing.T) resource.TestCheckFunc {
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

func testGroupMemberResourceCreateNoMembers(groupName string) string {
	return fmt.Sprintf(`
		resource "databricks_group" "my_group" {
			display_name = "%s"
		}
		`, groupName)
}

func testGroupMemberResource(groupName string) string {
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
