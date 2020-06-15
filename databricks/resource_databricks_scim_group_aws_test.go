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

func TestAccAwsScimGroupResource(t *testing.T) {
	//var secretScope model.Secre
	var ScimGroup model.Group
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	userName := "scimgroup-test@databricks.com"
	displayName := "scimgroup test"
	groupName := "scimgroup test"
	role := "arn:aws:iam::999999999999:instance-profile/terraform-scim-group-test"
	entitlement := "allow-cluster-create"
	expectEntitlements := []model.EntitlementsListItem{{Value: model.AllowClusterCreateEntitlement}}

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAwsScimGroupResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testAwsScimGroupResourceCreate(userName, displayName, groupName, role, entitlement),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAwsScimGroupResourceExists("databricks_scim_group.my_scim_group", &ScimGroup, t),
					// verify remote values
					testAwsScimGroupValues(t, &ScimGroup, displayName, expectEntitlements,
						[]model.RoleListItem{{Value: role}}, true),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "entitlements.#", "1"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "roles.#", "1"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "members.#", "1"),
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testAwsScimGroupResourceUpdate(groupName),

				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAwsScimGroupResourceExists("databricks_scim_group.my_scim_group", &ScimGroup, t),
					// verify remote values
					testAwsScimGroupValues(t, &ScimGroup, displayName, nil, nil, false),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "entitlements.#", "0"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "roles.#", "0"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "members.#", "0"),
				),
				Destroy: false,
			},
			{
				// Recreate the group with roles and entitlements again to see if the group gets updated
				Config: testAwsScimGroupResourceCreate(userName, displayName, groupName, role, entitlement),

				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAwsScimGroupResourceExists("databricks_scim_group.my_scim_group", &ScimGroup, t),
					// verify remote values
					testAwsScimGroupValues(t, &ScimGroup, displayName, expectEntitlements,
						[]model.RoleListItem{{Value: role}}, true),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "entitlements.#", "1"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "roles.#", "1"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "members.#", "1"),
				),
				Destroy: false,
			},
			{
				PreConfig: func() {
					err := testAccProvider.Meta().(*service.DBApiClient).Groups().Delete(ScimGroup.ID)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testAwsScimGroupResourceUpdate(displayName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAwsScimGroupResourceExists("databricks_scim_group.my_scim_group", &ScimGroup, t),
					// verify remote values
					testAwsScimGroupValues(t, &ScimGroup, displayName, nil, nil, false),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "entitlements.#", "0"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "roles.#", "0"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "members.#", "0"),
				),
				Destroy: false,
			},
			{
				// Recreate the group with roles and entitlements again to see if the group gets updated
				Config: testAwsScimGroupResourceInheritedRole(userName, displayName, groupName, role, entitlement),

				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAwsScimGroupResourceExists("databricks_scim_group.my_scim_group", &ScimGroup, t),
					// verify remote values
					testAwsScimGroupValues(t, &ScimGroup, displayName, expectEntitlements,
						[]model.RoleListItem{{Value: role}}, true),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "entitlements.#", "1"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "roles.#", "0"),
					resource.TestCheckResourceAttr("databricks_scim_group.my_scim_group", "members.#", "1"),
				),
				Destroy: false,
			},
		},
	})
}

func testAwsScimGroupResourceDestroy(s *terraform.State) error {
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

func testAwsScimGroupValues(t *testing.T, group *model.Group, displayName string, expectEntitlements []model.EntitlementsListItem, expectRoles []model.RoleListItem, verifyMembers bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, group.DisplayName == displayName)
		assert.EqualValues(t, group.Entitlements, expectEntitlements)
		assert.EqualValues(t, group.Roles, expectRoles)
		assert.True(t, (verifyMembers && len(group.Members) == 1) == verifyMembers)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testAwsScimGroupResourceExists(n string, group *model.Group, t *testing.T) resource.TestCheckFunc {
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

func testAwsScimGroupResourceCreate(username, displayName, groupName, role, entitlement string) string {
	return fmt.Sprintf(`
								resource "databricks_instance_profile" "instance_profile" {
								  instance_profile_arn = "%s"
								  skip_validation = true
								}
								data "databricks_default_user_roles" "default_roles" {
								  default_username = "terraform-all-user-roles@databricks.com"
								}
								resource "databricks_scim_user" "my_scim_user" {
								  user_name = "%s"
								  default_roles = data.databricks_default_user_roles.default_roles.roles
								  display_name = "%s"
								  entitlements = [
									"allow-cluster-create",
								  ]
								  roles = [
									databricks_instance_profile.instance_profile.id,
								  ]
								}
								resource "databricks_scim_group" "my_scim_group" {
								  display_name = "%s"
								  members = [databricks_scim_user.my_scim_user.id]
								  entitlements = [
									"%s",
								  ]
								  roles = [
									databricks_instance_profile.instance_profile.id
								  ]
								}
								`, role, username, displayName, groupName, entitlement)
}

func testAwsScimGroupResourceUpdate(groupName string) string {
	return fmt.Sprintf(`
								resource "databricks_scim_group" "my_scim_group" {
								  display_name = "%s"
								}
								`, groupName)
}

func testAwsScimGroupResourceInheritedRole(username, displayName, groupName, role, entitlement string) string {
	return fmt.Sprintf(`
								resource "databricks_instance_profile" "instance_profile" {
								  instance_profile_arn = "%s"
								  skip_validation = true
								}
								data "databricks_default_user_roles" "default_roles" {
								  default_username = "terraform-all-user-roles@databricks.com"
								}
								resource "databricks_scim_user" "my_scim_user" {
								  user_name = "%s"
								  default_roles = data.databricks_default_user_roles.default_roles.roles
								  display_name = "%s"
								  entitlements = [
									"allow-cluster-create",
								  ]
								  roles = [
									databricks_instance_profile.instance_profile.id,
								  ]
								}
								resource "databricks_scim_group" "my_scim_group_parent" {
								  display_name = "%s_parent"
								  members = [databricks_scim_group.my_scim_group.id]
								  roles = [
									databricks_instance_profile.instance_profile.id
								  ]
								}
								resource "databricks_scim_group" "my_scim_group" {
								  display_name = "%s"
								  members = [databricks_scim_user.my_scim_user.id]
								  entitlements = [
									"%s",
								  ]
								}
								`, role, username, displayName, groupName, groupName, entitlement)
}
