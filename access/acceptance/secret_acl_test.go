package acceptance

import (
	"errors"
	"fmt"
	"os"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/identity"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccSecretAclResource(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	//var secretScope Secre
	var secretACL ACLItem
	scope := fmt.Sprintf("tf-scope-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	principal := "users"
	permission := "READ"
	client := common.CommonEnvironmentClient()
	me, _ := identity.NewUsersAPI(client).Me()
	userName := me.UserName

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testSecretACLResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testSecretACLResource(scope, principal, permission),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// test scope permissions - it should be current user
					testSecretScopeHasPrincipal(t, scope, userName, "MANAGE"),
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

// this test checks that any user has access when initial principal is set to 'users'
func TestAccSecretAclResourceDefaultPrincipal(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	scope := fmt.Sprintf("tf-scope-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	client := common.CommonEnvironmentClient()
	me, _ := identity.NewUsersAPI(client).Me()
	userName := me.UserName
	userPermission := "READ"
	initialPrincipal := "users"
	initialPermission := "MANAGE"
	var secretACL ACLItem

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testSecretACLResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testSecretACLResourceWithDefaultPrincipal(scope, initialPrincipal, userName, userPermission),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// test scope permissions - it should be users
					testSecretScopeHasPrincipal(t, scope, initialPrincipal, initialPermission),
					// query the API to retrieve the tokenInfo object
					testSecretACLResourceExists("databricks_secret_acl.my_secret_acl", &secretACL, t),
					// verify remote values
					testSecretACLValues(t, &secretACL, userPermission, userName),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "scope", scope),
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "principal", userName),
					resource.TestCheckResourceAttr("databricks_secret_acl.my_secret_acl", "permission", userPermission),
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
		assert.EqualValues(t, permission, acl.Permission)
		assert.EqualValues(t, principal, acl.Principal)
		return nil
	}
}

func testSecretScopeHasPrincipal(t *testing.T, scope, principal, permission string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var acl ACLItem
		err := getSecretACLResourceExistsForScopeAndPrincipal(scope, principal, &acl)
		if err != nil {
			return err
		}
		assert.EqualValues(t, permission, acl.Permission)
		assert.EqualValues(t, principal, acl.Principal)
		return nil
	}
}

func getSecretACLResourceExistsForScopeAndPrincipal(scope, principal string, aclItem *ACLItem) error {
	// retrieve the configured client from the test setup
	conn := common.CommonEnvironmentClient()
	resp, err := NewSecretAclsAPI(conn).Read(scope, principal)
	if err != nil {
		return err
	}
	// If no error, assign the response Widget attribute to the widget pointer
	*aclItem = resp
	return nil
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testSecretACLResourceExists(n string, aclItem *ACLItem, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		return getSecretACLResourceExistsForScopeAndPrincipal(rs.Primary.Attributes["scope"],
			rs.Primary.Attributes["principal"], aclItem)
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

// testAccTokenResource returns an configuration for an Example Widget with the provided name
func testSecretACLResourceWithDefaultPrincipal(scopeName, defaultPrincipal, principal, permission string) string {
	return fmt.Sprintf(`
        resource "databricks_secret_scope" "my_scope" {
            name = "%s"
            initial_manage_principal = "%s"
        }
        resource "databricks_secret_acl" "my_secret_acl" {
            principal = "%s"
            permission = "%s"
            scope = databricks_secret_scope.my_scope.name
        }
        `, scopeName, defaultPrincipal, principal, permission)
}
