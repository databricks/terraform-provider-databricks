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

func TestAccAwsSecretScopeResource(t *testing.T) {
	var secretScope model.SecretScope

	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	scope := "terraform_acc_test_scope"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAwsSecretScopeResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testAwsSecretScopeResource(scope),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAwsSecretScopeResourceExists("databricks_secret_scope.my_scope", &secretScope, t),
					// verify remote values
					testAwsSecretScopeValues(t, &secretScope, scope),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "name", scope),
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", string(model.ScopeBackendTypeDatabricks)),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DBApiClient)
					err := client.SecretScopes().Delete(scope)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testAwsSecretScopeResource(scope),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAwsSecretScopeResourceExists("databricks_secret_scope.my_scope", &secretScope, t),
					// verify remote values
					testAwsSecretScopeValues(t, &secretScope, scope),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "name", scope),
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", string(model.ScopeBackendTypeDatabricks)),
				),
			},
		},
	})
}

func testAwsSecretScopeResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*service.DBApiClient)
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

func testAwsSecretScopeValues(t *testing.T, secretScope *model.SecretScope, scope string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, secretScope.Name == scope)
		assert.True(t, secretScope.BackendType == model.ScopeBackendTypeDatabricks)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testAwsSecretScopeResourceExists(n string, secretScope *model.SecretScope, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DBApiClient)
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
func testAwsSecretScopeResource(scopeName string) string {
	return fmt.Sprintf(`
								resource "databricks_secret_scope" "my_scope" {
								  name = "%s"
								}
								`, scopeName)
}
