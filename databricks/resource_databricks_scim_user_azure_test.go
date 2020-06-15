package databricks

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccAzureScimUserResource(t *testing.T) {
	//var secretScope model.Secre
	var scimUser model.User
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	userName := "terraform-test-scim-user@databricks.com"
	displayName := "terraform scim-testuser"
	expectEntitlements := []model.EntitlementsListItem{{Value: model.AllowClusterCreateEntitlement}}

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAzureScimUserResourceDestroy,

		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testAzureScimUserResourceCreate(userName, displayName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAzureScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testAzureScimUserValues(t, &scimUser, userName, displayName, expectEntitlements, "0"),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "user_name", userName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "entitlements.#", "1"),
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testAzureScimUserResourceUpdate(userName, displayName),

				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAzureScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testAzureScimUserValues(t, &scimUser, userName, displayName, nil, "1"),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "user_name", userName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "entitlements.#", "0"),
				),
				Destroy: false,
			},
			{
				// Recreate the user with roles and entitlements again to see if the user gets updated
				Config: testAzureScimUserResourceCreate(userName, displayName),

				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAzureScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testAzureScimUserValues(t, &scimUser, userName, displayName, expectEntitlements, "2"),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "user_name", userName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "entitlements.#", "1"),
				),
				Destroy: false,
			},
			{
				PreConfig: func() {
					err := testAccProvider.Meta().(*service.DBApiClient).Users().Delete(scimUser.ID)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testAzureScimUserResourceUpdate(userName, displayName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAzureScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testAzureScimUserValues(t, &scimUser, userName, displayName, nil, "3"),
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
				Config: testAzureScimUserResourceSetAdmin(userName, displayName, true),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAzureScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testAzureScimUserValues(t, &scimUser, userName, displayName, nil, "4"),
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
				Config: testAzureScimUserResourceSetAdmin(userName, displayName, false),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAzureScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testAzureScimUserValues(t, &scimUser, userName, displayName, nil, "5"),
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
				Config: testAzureScimUserResourceSetAdmin(userName, displayName, true),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAzureScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testAzureScimUserValues(t, &scimUser, userName, displayName, nil, "6"),
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

func testAzureScimUserResourceDestroy(s *terraform.State) error {
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

func testAzureScimUserValues(t *testing.T, user *model.User, userName, displayName string, expectEntitlements []model.EntitlementsListItem, step string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var errorMsg bytes.Buffer

		//ok := assert.True(t, user.UserName == userName)
		ok := reflect.DeepEqual(user.UserName, userName)
		if !ok {
			errorMsg.WriteString(fmt.Sprintf("failed username equality on step: %v;", step))
		}

		ok = reflect.DeepEqual(user.DisplayName, displayName)
		//ok = assert.True(t, user.DisplayName == displayName, failedMsg)
		if !ok {
			errorMsg.WriteString(fmt.Sprintf("failed displayname equality on step: %v;", step))
		}
		ok = reflect.DeepEqual(user.Entitlements, expectEntitlements)
		//ok = assert.EqualValues(t, user.Entitlements, expectEntitlements, failedMsg)
		if !ok {
			errorMsg.WriteString(fmt.Sprintf("failed entitlements equality on step: %v;", step))
		}
		if errorMsg.String() != "" {
			return errors.New(errorMsg.String())
		}
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testAzureScimUserResourceExists(n string, user *model.User, t *testing.T) resource.TestCheckFunc {
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
		log.Print(user)
		return nil
	}
}

func testAzureScimUserResourceCreate(username, displayName string) string {
	return fmt.Sprintf(`
								resource "databricks_scim_user" "my_scim_user" {
								  user_name = "%s"
								  display_name = "%s"
								  default_roles = []
								  entitlements = [
									"allow-cluster-create",
								  ]
								}
								`, username, displayName)
}

func testAzureScimUserResourceUpdate(username, displayName string) string {
	return fmt.Sprintf(`
								resource "databricks_scim_user" "my_scim_user" {
								  user_name = "%s"
								  default_roles = []
								  display_name = "%s"
								}
								`, username, displayName)
}

func testAzureScimUserResourceSetAdmin(username, displayName string, setAdmin bool) string {
	return fmt.Sprintf(`
								resource "databricks_scim_user" "my_scim_user" {
								  user_name = "%s"
								  default_roles = []
								  display_name = "%s"
                                  set_admin = %v
								}
								`, username, displayName, setAdmin)
}
