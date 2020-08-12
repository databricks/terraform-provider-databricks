package acceptance

import (
	"errors"
	"fmt"
	"os"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/access"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccSecretAclResource(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	//var secretScope Secre
	var secretACL ACLItem
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	scope := "terraform_acc_test_acl"
	principal := "users"
	permission := "READ"

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testSecretACLResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testSecretACLResource(scope, principal, permission),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testSecretACLResourceExists("databricks_secret_acl.my_secret_acl", &secretACL, t),
					// verify remote values
					testSecretACLValues(t, &secretACL, permission, principal),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "scope", scope),
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "principal", principal),
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "permission", permission),
				),
			},
			{
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewSecretAclsAPI(client).Delete(scope, principal)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testSecretACLResource(scope, principal, permission),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testSecretACLResourceExists("databricks_secret_acl.my_secret_acl", &secretACL, t),
					// verify remote values
					testSecretACLValues(t, &secretACL, permission, principal),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "scope", scope),
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "principal", principal),
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "permission", permission),
				),
			},
		},
	})
}

func testSecretACLResourceDestroy(s *terraform.State) error {
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_secret" && rs.Type != "databricks_secret_scope" {
			continue
		}
		_, err := NewSecretAclsAPI(client).Read(rs.Primary.Attributes["scope"], rs.Primary.Attributes["principal"])
		if err == nil {
			return errors.New("resource secret acl is not cleaned up")
		}
		_, err = NewSecretScopesAPI(client).Read(rs.Primary.Attributes["scope"])
		if err == nil {
			return errors.New("resource secret is not cleaned up")
		}
	}
	return nil
}

func testSecretACLValues(t *testing.T, acl *ACLItem, permission, principal string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, acl.Permission == ACLPermissionRead)
		assert.True(t, acl.Principal == principal)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testSecretACLResourceExists(n string, aclItem *ACLItem, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := common.CommonEnvironmentClient()
		resp, err := NewSecretAclsAPI(conn).Read(rs.Primary.Attributes["scope"], rs.Primary.Attributes["principal"])
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
func testSecretACLResource(scopeName, principal, permission string) string {
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
