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

func TestAccAwsSecretAclResource(t *testing.T) {
	//var secretScope model.Secre
	var secretACL model.ACLItem
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	scope := "terraform_acc_test_acl"
	principal := "users"
	permission := "READ"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAwsSecretACLResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testAwsSecretACLResource(scope, principal, permission),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAwsSecretACLResourceExists("databricks_secret_acl.my_secret_acl", &secretACL, t),
					// verify remote values
					testAwsSecretACLValues(t, &secretACL, permission, principal),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "scope", scope),
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "principal", principal),
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "permission", permission),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DBApiClient)
					err := client.SecretAcls().Delete(scope, principal)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testAwsSecretACLResource(scope, principal, permission),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAwsSecretACLResourceExists("databricks_secret_acl.my_secret_acl", &secretACL, t),
					// verify remote values
					testAwsSecretACLValues(t, &secretACL, permission, principal),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "scope", scope),
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "principal", principal),
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "permission", permission),
				),
			},
		},
	})
}

func testAwsSecretACLResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*service.DBApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_secret" && rs.Type != "databricks_secret_scope" {
			continue
		}
		_, err := client.SecretAcls().Read(rs.Primary.Attributes["scope"], rs.Primary.Attributes["principal"])
		if err == nil {
			return errors.New("resource secret acl is not cleaned up")
		}
		_, err = client.SecretScopes().Read(rs.Primary.Attributes["scope"])
		if err == nil {
			return errors.New("resource secret is not cleaned up")
		}
	}
	return nil
}

func testAwsSecretACLValues(t *testing.T, acl *model.ACLItem, permission, principal string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, acl.Permission == model.ACLPermissionRead)
		assert.True(t, acl.Principal == principal)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testAwsSecretACLResourceExists(n string, aclItem *model.ACLItem, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DBApiClient)
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
func testAwsSecretACLResource(scopeName, principal, permission string) string {
	return fmt.Sprintf(`
								resource "databricks_secret_scope" "my_scope" {
								  name = "%s"
								}
								resource "databricks_secret_acl" "my_secret_acl" {
								  principal = "%s"
								  permission = "%s"
								  scope = databricks_secret_scope.my_scope.name
								}
								`, scopeName, principal, permission)
}
