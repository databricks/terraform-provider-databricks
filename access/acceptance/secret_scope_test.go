package acceptance

import (
	"errors"
	"fmt"
	"os"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/access"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccSecretScopeResource(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	var secretScope SecretScope

	scope := "terraform_acc_test_scope"

	acceptance.AccTest(t, resource.TestCase{
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
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", string(ScopeBackendTypeDatabricks)),
				),
			},
			{
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewSecretScopesAPI(client).Delete(scope)
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
					resource.TestCheckResourceAttr("databricks_secret_scope.my_scope", "backend_type", string(ScopeBackendTypeDatabricks)),
				),
			},
		},
	})
}

func testSecretScopeResourceDestroy(s *terraform.State) error {
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_secret_scope" {
			continue
		}
		_, err := NewSecretScopesAPI(client).Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource token is not cleaned up")
	}
	return nil
}

func testSecretScopeValues(t *testing.T, secretScope *SecretScope, scope string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, secretScope.Name == scope)
		assert.True(t, secretScope.BackendType == ScopeBackendTypeDatabricks)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testSecretScopeResourceExists(n string, secretScope *SecretScope, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := common.CommonEnvironmentClient()
		resp, err := NewSecretScopesAPI(conn).Read(rs.Primary.ID)
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
