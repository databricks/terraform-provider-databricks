package databricks

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccSecretScopeResource(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	var secretScope model.SecretScope

	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	scope := "terraform_acc_test_scope"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testSecretScopeResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testSecretScopeResource(scope),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testSecretScopeResourceExists("databricks_secret_scope.my_scope", &secretScope, t),
					// verify remote values
					testSecretScopeValues(t, &secretScope, scope),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "name", scope),
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", string(model.ScopeBackendTypeDatabricks)),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DatabricksClient)
					err := client.SecretScopes().Delete(scope)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testSecretScopeResource(scope),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testSecretScopeResourceExists("databricks_secret_scope.my_scope", &secretScope, t),
					// verify remote values
					testSecretScopeValues(t, &secretScope, scope),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "name", scope),
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", string(model.ScopeBackendTypeDatabricks)),
				),
			},
		},
	})
}

func testSecretScopeResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*service.DatabricksClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_secret_scope" {
			continue
		}
		_, err := client.Tokens().Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource token is not cleaned up")
	}
	return nil
}

func testSecretScopeValues(t *testing.T, secretScope *model.SecretScope, scope string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, secretScope.Name == scope)
		assert.True(t, secretScope.BackendType == model.ScopeBackendTypeDatabricks)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testSecretScopeResourceExists(n string, secretScope *model.SecretScope, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DatabricksClient)
		resp, err := conn.SecretScopes().Read(rs.Primary.ID)
		//t.Log(resp)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*secretScope = resp
		return nil
		//return fmt.Errorf("Token (%s) not found", rs.Primary.ID)
	}
}

// testAccTokenResource returns an configuration for an Example Widget with the provided name
func testSecretScopeResource(scopeName string) string {
	return fmt.Sprintf(`
		resource "databricks_secret_scope" "my_scope" {
			name = "%s"
		}
		`, scopeName)
}

func TestResourceSecretScopeRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/secrets/scopes/list?",
			Response: model.SecretScopeList{
				Scopes: []model.SecretScope{
					{
						Name:        "abc",
						BackendType: "DATABRICKS",
					},
				},
			},
			Status: 200,
		},
	}, resourceSecretScope, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId("abc")
		return resourceSecretScopeRead(d, c)
	})
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "DATABRICKS", d.Get("backend_type"))
	assert.Equal(t, "users", d.Get("initial_manage_principal"))
	assert.Equal(t, "abc", d.Get("name"))
}

func TestResourceSecretScopeRead_NotFound(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/secrets/scopes/list?",
			Response: model.SecretScopeList{
				Scopes: []model.SecretScope{
					{
						Name:        "bcd",
						BackendType: "DATABRICKS",
					},
				},
			},
			Status: 200,
		},
	}, resourceSecretScope, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId("abc")
		return resourceSecretScopeRead(d, c)
	})
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id())
}

func TestResourceSecretScopeRead_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/scopes/list?",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceSecretScope, nil, actionWithID("abc", resourceSecretScopeRead))
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceSecretScopeCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/secrets/scopes/create",
			ExpectedRequest: map[string]string{
				"scope":                    "Boom",
				"initial_manage_principal": "groups",
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/secrets/scopes/list?",
			Response: model.SecretScopeList{
				Scopes: []model.SecretScope{
					{
						Name:        "Boom",
						BackendType: "DATABRICKS",
					},
				},
			},
			Status: 200,
		},
	}, resourceSecretScope, map[string]interface{}{
		"initial_manage_principal": "groups",
		"name":                     "Boom",
	}, resourceSecretScopeCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "Boom", d.Id())
}

func TestResourceSecretScopeCreate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/secrets/scopes/create",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceSecretScope, map[string]interface{}{
		"initial_manage_principal": "groups",
		"name":                     "Boom",
	}, resourceSecretScopeCreate)
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceSecretScopeDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{ // read log output for better stub url...
			Method:   "POST",
			Resource: "/api/2.0/secrets/scopes/delete",
			ExpectedRequest: map[string]string{
				"scope": "abc",
			},
		},
	}, resourceSecretScope, nil, actionWithID("abc", resourceSecretScopeDelete))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceSecretScopeDelete_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/secrets/scopes/delete",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceSecretScope, nil, actionWithID("abc", resourceSecretScopeDelete))
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
