package acceptance

import (
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

func TestAccScimGroupResource(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	//var secretScope Secre
	var ScimGroup Group
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	userName := fmt.Sprintf("scimgroup-test-%s@example.com", randomName)
	displayName := fmt.Sprintf("scimgroup %s", randomName)
	groupName := fmt.Sprintf("scimgroup %s", randomName)
	role := "arn:aws:iam::999999999999:instance-profile/terraform-scim-group-test"
	entitlement := "allow-cluster-create"
	expectEntitlements := []EntitlementsListItem{{Value: AllowClusterCreateEntitlement}}

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testScimGroupResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testScimGroupResourceCreate(userName, displayName, groupName, role, entitlement),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimGroupResourceExists("databricks_scim_group.my_scim_group", &ScimGroup, t),
					// verify remote values
					testScimGroupValues(t, &ScimGroup, displayName, expectEntitlements,
						[]RoleListItem{{Value: role}}, true),
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
				Config: testScimGroupResourceUpdate(groupName),

				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimGroupResourceExists("databricks_scim_group.my_scim_group", &ScimGroup, t),
					// verify remote values
					testScimGroupValues(t, &ScimGroup, displayName, nil, nil, false),
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
				Config: testScimGroupResourceCreate(userName, displayName, groupName, role, entitlement),

				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimGroupResourceExists("databricks_scim_group.my_scim_group", &ScimGroup, t),
					// verify remote values
					testScimGroupValues(t, &ScimGroup, displayName, expectEntitlements,
						[]RoleListItem{{Value: role}}, true),
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
					err := NewGroupsAPI(common.CommonEnvironmentClient()).Delete(ScimGroup.ID)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testScimGroupResourceUpdate(displayName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimGroupResourceExists("databricks_scim_group.my_scim_group", &ScimGroup, t),
					// verify remote values
					testScimGroupValues(t, &ScimGroup, displayName, nil, nil, false),
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
				Config: testScimGroupResourceInheritedRole(userName, displayName, groupName, role, entitlement),

				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimGroupResourceExists("databricks_scim_group.my_scim_group", &ScimGroup, t),
					// verify remote values
					testScimGroupValues(t, &ScimGroup, displayName, expectEntitlements,
						[]RoleListItem{{Value: role}}, true),
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

func testScimGroupResourceDestroy(s *terraform.State) error {
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_scim_group" {
			continue
		}
		_, err := NewUsersAPI(client).Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource Scim Group is not cleaned up")
	}
	return nil
}

func testScimGroupValues(t *testing.T, group *Group, displayName string, expectEntitlements []EntitlementsListItem, expectRoles []RoleListItem, verifyMembers bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, group.DisplayName == displayName)
		assert.EqualValues(t, group.Entitlements, expectEntitlements)
		assert.EqualValues(t, group.Roles, expectRoles)
		assert.True(t, (verifyMembers && len(group.Members) == 1) == verifyMembers)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testScimGroupResourceExists(n string, group *Group, t *testing.T) resource.TestCheckFunc {
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

func testScimGroupResourceCreate(username, displayName, groupName, role, entitlement string) string {
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

func testScimGroupResourceUpdate(groupName string) string {
	return fmt.Sprintf(`
		resource "databricks_scim_group" "my_scim_group" {
			display_name = "%s"
		}
		`, groupName)
}

func testScimGroupResourceInheritedRole(username, displayName, groupName, role, entitlement string) string {
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
