// +build azure

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

func TestAccScimGroupResource(t *testing.T) {
	//var secretScope model.Secre
	var ScimGroup model.Group
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	userName := "scimgroup-test2@databricks.com"
	displayName := "scimgroup test2"
	groupName := "scimgroup test2"
	entitlement := "allow-cluster-create"
	expectEntitlements := []model.EntitlementsListItem{{Value: model.AllowClusterCreateEntitlement}}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testScimGroupResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testScimGroupResourceCreate(userName, displayName, groupName, entitlement),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimGroupResourceExists("databricks_scim_group.my_scim_group_test", &ScimGroup, t),
					// verify remote values
					testScimGroupValues(t, &ScimGroup, displayName, expectEntitlements, true),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "entitlements.#", "1"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "members.#", "1"),
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testScimGroupResourceUpdate(groupName),

				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimGroupResourceExists("databricks_scim_group.my_scim_group_test", &ScimGroup, t),
					// verify remote values
					testScimGroupValues(t, &ScimGroup, displayName, nil, false),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "entitlements.#", "0"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "members.#", "0"),
				),
				Destroy: false,
			},
			{
				// Recreate the group with roles and entitlements again to see if the group gets updated
				Config: testScimGroupResourceCreate(userName, displayName, groupName, entitlement),

				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimGroupResourceExists("databricks_scim_group.my_scim_group_test", &ScimGroup, t),
					// verify remote values
					testScimGroupValues(t, &ScimGroup, displayName, expectEntitlements, true),
					// verify local values

					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "entitlements.#", "1"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "members.#", "1"),
				),
				Destroy: false,
			},
			{
				PreConfig: func() {
					err := testAccProvider.Meta().(*service.DBApiClient).Groups().Delete(ScimGroup.ID)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testScimGroupResourceUpdate(displayName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimGroupResourceExists("databricks_scim_group.my_scim_group_test", &ScimGroup, t),
					// verify remote values
					testScimGroupValues(t, &ScimGroup, displayName, nil, false),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "entitlements.#", "0"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "members.#", "0"),
				),
				Destroy: false,
			},
			{
				// Recreate the group with roles and entitlements again to see if the group gets updated
				Config: testScimGroupResourceInheritedRole(userName, displayName, groupName, entitlement),

				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimGroupResourceExists("databricks_scim_group.my_scim_group_test", &ScimGroup, t),
					// verify remote values
					testScimGroupValues(t, &ScimGroup, displayName, expectEntitlements, true),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "entitlements.#", "1"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group_test", "members.#", "1"),
				),
				Destroy: false,
			},
		},
	})
}

func testScimGroupResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*service.DBApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_scim_group" {
			continue
		}
		_, err := client.Users().Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource Scim Group is not cleaned up")
	}
	return nil
}

func testScimGroupValues(t *testing.T, group *model.Group, displayName string, expectEntitlements []model.EntitlementsListItem, verifyMembers bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, group.DisplayName == displayName)
		assert.EqualValues(t, group.Entitlements, expectEntitlements)
		assert.True(t, (verifyMembers && len(group.Members) == 1) == verifyMembers)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testScimGroupResourceExists(n string, group *model.Group, t *testing.T) resource.TestCheckFunc {
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

func testScimGroupResourceCreate(username, displayName, groupName, entitlement string) string {
	return fmt.Sprintf(`
								resource "databricks_scim_user" "my_scim_group_test_user" {
								  user_name = "%s"
								  default_roles = []
								  display_name = "%s"
								  entitlements = [
									"allow-cluster-create",
								  ]
								}
								resource "databricks_scim_group" "my_scim_group_test" {
								  display_name = "%s"
								  members = [databricks_scim_user.my_scim_group_test_user.id]
								  entitlements = [
									"%s",
								  ]
								}
								`, username, displayName, groupName, entitlement)
}

func testScimGroupResourceUpdate(groupName string) string {
	return fmt.Sprintf(`

								resource "databricks_scim_group" "my_scim_group_test" {
								  display_name = "%s"
								}
								`, groupName)
}

func testScimGroupResourceInheritedRole(username, displayName, groupName, entitlement string) string {
	return fmt.Sprintf(`
								resource "databricks_scim_user" "my_scim_group_test_user" {
								  user_name = "%s"
								  default_roles = []
								  display_name = "%s"
								  entitlements = [
									"allow-cluster-create",
								  ]
								}
								resource "databricks_scim_group" "my_scim_group_test_parent" {
								  display_name = "%s_parent"
								  members = [databricks_scim_group.my_scim_group_test.id]
								}
								resource "databricks_scim_group" "my_scim_group_test" {
								  display_name = "%s"
								  members = [databricks_scim_user.my_scim_group_test_user.id]
								  entitlements = [
									"%s",
								  ]
								}
								`, username, displayName, groupName, groupName, entitlement)
}
