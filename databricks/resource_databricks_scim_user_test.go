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

func TestAccScimUserResource(t *testing.T) {
	//var secretScope model.Secre
	var scimUser model.User
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	userName := "terraform-testuser@databricks.com"
	displayName := "terraform testuser"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testScimUserResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testScimUserResource(userName, displayName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testScimUserResourceExists("databricks_scim_user.my_scim_user", &scimUser, t),
					// verify remote values
					testScimUserValues(t, &scimUser, userName, displayName),
					// verify local values
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "user_name", userName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "display_name", displayName),
					resource.TestCheckResourceAttr("databricks_scim_user.my_scim_user", "entitlements.#", "1"),
				),
			},
		},
	})
}

func testScimUserResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(service.DBApiClient)
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

func testScimUserPreCheck(t *testing.T) {
	return
}

func testScimUserValues(t *testing.T, user *model.User, userName, displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, user.UserName == userName)
		assert.True(t, user.DisplayName == displayName)
		assert.True(t, user.Entitlements[0].Value == model.AllowClusterCreateEntitlement)
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
		conn := testAccProvider.Meta().(service.DBApiClient)
		resp, err := conn.Users().Read(rs.Primary.ID)
		t.Log(resp)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*user = resp
		return nil
	}
}

func testScimUserResource(username, displayName string) string {
	return fmt.Sprintf(`
								resource "databricks_scim_user" "my_scim_user" {
								  user_name = "%s"
								  display_name = "%s"
								  entitlements = [
									"allow-cluster-create",
								  ]
								}
								`, username, displayName)
}
