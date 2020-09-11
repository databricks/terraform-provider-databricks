package acceptance

import (
	"errors"
	"fmt"
	"os"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/identity"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccScimUserResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	//var secretScope Secre
	var scimUser User
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	userName := fmt.Sprintf("terraform.test+%s@example.com", randomName)
	displayName := fmt.Sprintf("Terra %s", randomName)
	expectEntitlements := []EntitlementsListItem{{Value: AllowClusterCreateEntitlement}}
	userNameAndDisplay := map[string]string{
		"USER_NAME":    userName,
		"DISPLAY_NAME": displayName,
	}

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testScimUserResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: qa.EnvironmentTemplate(t, `
				data "databricks_default_user_roles" "default_roles" {
					default_username = "terraform-all-user-roles+{var.RANDOM}@example.com"
				}
				resource "databricks_instance_profile" "instance_profile" {
					instance_profile_arn = "arn:aws:iam::999999999999:instance-profile/terraform-scim-user-{var.RANDOM}"
					skip_validation = true
				}
				resource "databricks_scim_user" "my_scim_user" {
					user_name = "{var.USER_NAME}"
					display_name = "{var.DISPLAY_NAME}"
					default_roles = data.databricks_default_user_roles.default_roles.roles
					entitlements = ["allow-cluster-create"]
					roles = [databricks_instance_profile.instance_profile.id]
				}`, userNameAndDisplay),
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
				Config: qa.EnvironmentTemplate(t, `
				data "databricks_default_user_roles" "default_roles" {
					default_username = "terraform-all-user-roles+{var.RANDOM}@example.com"
				}
				resource "databricks_scim_user" "my_scim_user" {
					user_name = "{var.USER_NAME}"
					display_name = "{var.DISPLAY_NAME}"
					default_roles = data.databricks_default_user_roles.default_roles.roles
				}`, userNameAndDisplay),
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
				Config: qa.EnvironmentTemplate(t, `
				data "databricks_default_user_roles" "default_roles" {
					default_username = "terraform-all-user-roles+{var.RANDOM}@example.com"
				}
				resource "databricks_instance_profile" "instance_profile" {
					instance_profile_arn = "arn:aws:iam::999999999999:instance-profile/terraform-scim-user-{var.RANDOM}"
					skip_validation = true
				}
				resource "databricks_scim_user" "my_scim_user" {
					user_name = "{var.USER_NAME}"
					display_name = "{var.DISPLAY_NAME}"
					default_roles = data.databricks_default_user_roles.default_roles.roles
					entitlements = ["allow-cluster-create"]
					roles = [databricks_instance_profile.instance_profile.id]
				}`, userNameAndDisplay),

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
					err := UsersAPI{C: common.CommonEnvironmentClient()}.Delete(scimUser.ID)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: qa.EnvironmentTemplate(t, `
				data "databricks_default_user_roles" "default_roles" {
					default_username = "terraform-all-user-roles+{var.RANDOM}@example.com"
				}
				resource "databricks_scim_user" "my_scim_user" {
					user_name = "{var.USER_NAME}"
					display_name = "{var.DISPLAY_NAME}"
					default_roles = data.databricks_default_user_roles.default_roles.roles
				}`, userNameAndDisplay),
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
					err := UsersAPI{C: common.CommonEnvironmentClient()}.Delete(scimUser.ID)
					assert.NoError(t, err, err)
				},
				// Create new admin user
				Config: qa.EnvironmentTemplate(t, `
				data "databricks_default_user_roles" "default_roles" {
					default_username = "terraform-all-user-roles+{var.RANDOM}@example.com"
				}
				resource "databricks_scim_user" "my_scim_user" {
					user_name = "{var.USER_NAME}"
					display_name = "{var.DISPLAY_NAME}"
					default_roles = data.databricks_default_user_roles.default_roles.roles
					set_admin = {var.ADMIN}
				}`, map[string]string{
					"USER_NAME":    userName,
					"DISPLAY_NAME": displayName,
					"ADMIN":        "true",
				}),
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
				Config: qa.EnvironmentTemplate(t, `
				data "databricks_default_user_roles" "default_roles" {
					default_username = "terraform-all-user-roles+{var.RANDOM}@example.com"
				}
				resource "databricks_scim_user" "my_scim_user" {
					user_name = "{var.USER_NAME}"
					display_name = "{var.DISPLAY_NAME}"
					default_roles = data.databricks_default_user_roles.default_roles.roles
					set_admin = {var.ADMIN}
				}`, map[string]string{
					"USER_NAME":    userName,
					"DISPLAY_NAME": displayName,
					"ADMIN":        "false",
				}),
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
				Config: qa.EnvironmentTemplate(t, `
				data "databricks_default_user_roles" "default_roles" {
					default_username = "terraform-all-user-roles+{var.RANDOM}@example.com"
				}
				resource "databricks_scim_user" "my_scim_user" {
					user_name = "{var.USER_NAME}"
					display_name = "{var.DISPLAY_NAME}"
					default_roles = data.databricks_default_user_roles.default_roles.roles
					set_admin = {var.ADMIN}
				}`, map[string]string{
					"USER_NAME":    userName,
					"DISPLAY_NAME": displayName,
					"ADMIN":        "true",
				}),
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
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_scim_user" {
			continue
		}
		_, err := NewUsersAPI(client).Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource Scim User is not cleaned up")
	}
	return nil
}

func testScimUserValues(t *testing.T, user *User, userName, displayName string, expectEntitlements []EntitlementsListItem) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, user.UserName == userName)
		assert.True(t, user.DisplayName == displayName)
		assert.EqualValues(t, user.Entitlements, expectEntitlements)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testScimUserResourceExists(n string, user *User, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := common.CommonEnvironmentClient()
		resp, err := NewUsersAPI(conn).Read(rs.Primary.ID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*user = resp
		return nil
	}
}
