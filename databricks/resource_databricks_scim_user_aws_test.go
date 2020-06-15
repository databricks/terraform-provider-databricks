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

func TestAccAwsScimUserResource(t *testing.T) {
	//var secretScope model.Secre
	var scimUser model.User
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	userName := "terraform-testuser@databricks.com"
	displayName := "terraform testuser"
	expectEntitlements := []model.EntitlementsListItem{{Value: model.AllowClusterCreateEntitlement}}

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testScimUserResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testScimUserResourceCreate(userName, displayName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testScimUserValues(t, &scimUser, userName, displayName, expectEntitlements),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "user_name", userName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "entitlements.#", "1"),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "roles.#", "1"),
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testScimUserResourceUpdate(userName, displayName),

				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testScimUserValues(t, &scimUser, userName, displayName, nil),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "user_name", userName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "entitlements.#", "0"),
				),
				Destroy: false,
			},
			{
				// Recreate the user with roles and entitlements again to see if the user gets updated
				Config: testScimUserResourceCreate(userName, displayName),

				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testScimUserValues(t, &scimUser, userName, displayName, expectEntitlements),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "user_name", userName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "entitlements.#", "1"),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "roles.#", "1"),
				),
				Destroy: false,
			},
			{
				PreConfig: func() {
					err := testAccProvider.Meta().(*service.DBApiClient).Users().Delete(scimUser.ID)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testScimUserResourceUpdate(userName, displayName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testScimUserValues(t, &scimUser, userName, displayName, nil),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "user_name", userName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "entitlements.#", "0"),
				),
				Destroy: false,
			},
			{
				//Create a new user
				PreConfig: func() {
					err := testAccProvider.Meta().(*service.DBApiClient).Users().Delete(scimUser.ID)
					assert.NoError(t, err, err)
				},
				// Create new admin user
				Config: testScimUserResourceSetAdmin(userName, displayName, true),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testScimUserValues(t, &scimUser, userName, displayName, nil),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "user_name", userName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "entitlements.#", "0"),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "set_admin", "true"),
				),
				Destroy: false,
			},
			{
				// Update admin to false
				Config: testScimUserResourceSetAdmin(userName, displayName, false),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testScimUserValues(t, &scimUser, userName, displayName, nil),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "user_name", userName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "entitlements.#", "0"),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "set_admin", "false"),
				),
				Destroy: false,
			},
			{
				// Update admin back to true
				Config: testScimUserResourceSetAdmin(userName, displayName, true),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testScimUserValues(t, &scimUser, userName, displayName, nil),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "user_name", userName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "entitlements.#", "0"),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "set_admin", "true"),
				),
				Destroy: false,
			},
		},
	})
}

func testScimUserResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*service.DBApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_scim_user" {
			continue
		}
		_, err := client.Users().Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource Scim User is not cleaned up")
	}
	return nil
}

func testScimUserValues(t *testing.T, user *model.User, userName, displayName string, expectEntitlements []model.EntitlementsListItem) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, user.UserName == userName)
		assert.True(t, user.DisplayName == displayName)
		assert.EqualValues(t, user.Entitlements, expectEntitlements)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testScimUserResourceExists(n string, user *model.User, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DBApiClient)
		resp, err := conn.Users().Read(rs.Primary.ID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*user = resp
		return nil
	}
}

func testScimUserResourceCreate(username, displayName string) string {
	return fmt.Sprintf(`
								data "databricks_default_user_roles" "default_roles" {
								  default_username = "terraform-all-user-roles@databricks.com"
								}
								resource "databricks_instance_profile" "instance_profile" {
								  instance_profile_arn = "arn:aws:iam::999999999999:instance-profile/terraform-scim-user-test"
								  skip_validation = true
								}
								resource "databricks_scim_user" "my_scim_user" {
								  user_name = "%s"
								  display_name = "%s"
								  default_roles = data.databricks_default_user_roles.default_roles.roles
								  entitlements = [
									"allow-cluster-create",
								  ]
								  roles = [
									databricks_instance_profile.instance_profile.id,
								  ]
								}
								`, username, displayName)
}

func testScimUserResourceUpdate(username, displayName string) string {
	return fmt.Sprintf(`
								data "databricks_default_user_roles" "default_roles" {
								  default_username = "terraform-all-user-roles@databricks.com"
								}
								resource "databricks_scim_user" "my_scim_user" {
								  user_name = "%s"
								  default_roles = data.databricks_default_user_roles.default_roles.roles
								  display_name = "%s"
								}
								`, username, displayName)
}

func testScimUserResourceSetAdmin(username, displayName string, setAdmin bool) string {
	return fmt.Sprintf(`
								data "databricks_default_user_roles" "default_roles" {
								  default_username = "terraform-all-user-roles@databricks.com"
								}
								resource "databricks_scim_user" "my_scim_user" {
								  user_name = "%s"
								  default_roles = data.databricks_default_user_roles.default_roles.roles
								  display_name = "%s"
                                  set_admin = %v
								}
								`, username, displayName, setAdmin)
}
