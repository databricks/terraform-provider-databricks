package databricks

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccSecretAclResource(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
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
					client := testAccProvider.Meta().(*service.DatabricksClient)
					err := client.SecretAcls().Delete(scope, principal)
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
	client := testAccProvider.Meta().(*service.DatabricksClient)
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

func testSecretACLValues(t *testing.T, acl *model.ACLItem, permission, principal string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, acl.Permission == model.ACLPermissionRead)
		assert.True(t, acl.Principal == principal)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testSecretACLResourceExists(n string, aclItem *model.ACLItem, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DatabricksClient)
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

func TestResourceSecretACLRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
			Response: model.ACLItem{
				Permission: "CAN_MANAGE",
			},
		},
	}, resourceSecretACL, nil, actionWithID("global|||something", resourceSecretACLRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "global|||something", d.Id(), "Id should not be empty")
	assert.Equal(t, "CAN_MANAGE", d.Get("permission"))
	assert.Equal(t, "something", d.Get("principal"))
	assert.Equal(t, "global", d.Get("scope"))
}

func TestResourceSecretACLRead_NotFound(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
			Response: service.APIErrorBody{
				ErrorCode: "NOT_FOUND",
				Message:   "Item not found",
			},
			Status: 404,
		},
	}, resourceSecretACL, nil, actionWithID("global|||something", resourceSecretACLRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceSecretACLRead_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceSecretACL, nil, actionWithID("global|||something", resourceSecretACLRead))
	assert.Errorf(t, err, "Internal error happened")
	assert.Equal(t, "global|||something", d.Id(), "Id should not be empty for error reads")
}

func TestResourceSecretACLCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/secrets/acls/put",
			ExpectedRequest: model.SecretACLRequest{
				Principal:  "something",
				Permission: "CAN_MANAGE",
				Scope:      "global",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/acls/get?principal=something&scope=global",
			Response: model.ACLItem{
				Permission: "CAN_MANAGE",
			},
		},
	}, resourceSecretACL, map[string]interface{}{
		"permission": "CAN_MANAGE",
		"principal":  "something",
		"scope":      "global",
	}, resourceSecretACLCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "global|||something", d.Id())
}

func TestResourceSecretACLCreate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{ // read log output for better stub url...
			Method:   "POST",
			Resource: "/api/2.0/secrets/acls/put",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceSecretACL, map[string]interface{}{
		"permission": "CAN_MANAGE",
		"principal":  "something",
		"scope":      "global",
	}, resourceSecretACLCreate)
	assert.Errorf(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceSecretACLDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/secrets/acls/delete",
			ExpectedRequest: map[string]string{
				"scope":     "global",
				"principal": "something",
			},
		},
	}, resourceSecretACL, nil, actionWithID("global|||something", resourceSecretACLDelete))
	assert.NoError(t, err, err)
	assert.Equal(t, "global|||something", d.Id())
}

func TestResourceSecretACLDelete_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/secrets/acls/delete",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceSecretACL, nil, actionWithID("global|||something", resourceSecretACLDelete))
	assert.Errorf(t, err, "Internal error happened")
	assert.Equal(t, "global|||something", d.Id())
}
