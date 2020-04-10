package db

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

func TestAccSecretAclResource(t *testing.T) {
	//var secretScope model.Secre
	var secretAcl model.AclItem
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	scope := "terraform_acc_test_scope"
	principal := "USERS"
	permission := "READ"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testSecretAclResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testSecretAclResource(scope, principal, permission),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testSecretAclResourceExists("db_secret_acl.my_secret_acl", &secretAcl, t),
					// verify remote values
					testSecretAclValues(t, &secretAcl, permission, principal),
					// verify local values
					resource.TestCheckResourceAttr("db_secret_acl.my_secret_acl", "scope", scope),
					resource.TestCheckResourceAttr("db_secret_acl.my_secret_acl", "principal", principal),
					resource.TestCheckResourceAttr("db_secret_acl.my_secret_acl", "permission", permission),
				),
			},
		},
	})
}

func testSecretAclResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(service.DBApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "db_secret_acl" {
			continue
		}
		_, err := client.SecretAcls().Read(rs.Primary.Attributes["scope"], rs.Primary.Attributes["principal"])
		if err != nil {
			return nil
		}
		return errors.New("Resource secret acl is not cleaned up!")
	}
	return nil
}

func testSecretAclPreCheck(t *testing.T) {
	return
}

func testSecretAclValues(t *testing.T, acl *model.AclItem, permission, principal string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, acl.Permission == model.AclPermissionRead)
		assert.True(t, acl.Principal == principal)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testSecretAclResourceExists(n string, aclItem *model.AclItem, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(service.DBApiClient)
		resp, err := conn.SecretAcls().Read(rs.Primary.Attributes["scope"], rs.Primary.Attributes["principal"])
		//t.Log(resp)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*aclItem = resp
		return nil
		//return fmt.Errorf("Token (%s) not found", rs.Primary.ID)
	}
}

// testAccTokenResource returns an configuration for an Example Widget with the provided name
func testSecretAclResource(scopeName, principal, permission string) string {
	return fmt.Sprintf(`
								resource "db_secret_scope" "my_scope" {
								  name = "%s"
								}
								resource "db_secret_acl" "my_secret_acl" {
								  principal = "%s"
								  permission = "%s"
								  scope = db_secret_scope.my_scope.name
								}
								`, scopeName, principal, permission)
}
