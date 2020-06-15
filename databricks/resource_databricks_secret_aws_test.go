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

func TestAccAwsSecretResource(t *testing.T) {
	//var secretScope model.Secre
	var secret model.SecretMetadata
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	scope := "terraform_acc_test_secret"
	key := "my_cool_key"
	stringValue := "my super secret key"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAwsSecretResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testAwsSecretResource(scope, key, stringValue),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAwsSecretResourceExists("databricks_secret.my_secret", &secret, t),
					// verify remote values
					testAwsSecretValues(t, &secret, key),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "scope", scope),
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "key", key),
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "string_value", stringValue),
				),
			},
			{
				//Deleting and recreating the secret
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DBApiClient)
					err := client.Secrets().Delete(scope, secret.Key)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testAwsSecretResource(scope, key, stringValue),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAwsSecretResourceExists("databricks_secret.my_secret", &secret, t),
					// verify remote values
					testAwsSecretValues(t, &secret, key),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "scope", scope),
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "key", key),
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "string_value", stringValue),
				),
			},
			{
				//Deleting the scope should recreate the secret
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DBApiClient)
					err := client.SecretScopes().Delete(scope)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testAwsSecretResource(scope, key, stringValue),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAwsSecretResourceExists("databricks_secret.my_secret", &secret, t),
					// verify remote values
					testAwsSecretValues(t, &secret, key),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "scope", scope),
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "key", key),
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "string_value", stringValue),
				),
			},
		},
	})
}

func testAwsSecretResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*service.DBApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_secret" && rs.Type != "databricks_secret_scope" {
			continue
		}
		_, err := client.Secrets().Read(rs.Primary.Attributes["scope"], rs.Primary.Attributes["key"])
		if err == nil {
			return errors.New("resource secret is not cleaned up")
		}
		_, err = client.SecretScopes().Read(rs.Primary.Attributes["scope"])
		if err == nil {
			return errors.New("resource secret is not cleaned up")
		}
	}
	return nil
}

func testAwsSecretValues(t *testing.T, secret *model.SecretMetadata, key string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, secret.Key == key)
		assert.True(t, secret.LastUpdatedTimestamp > 0)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testAwsSecretResourceExists(n string, secret *model.SecretMetadata, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DBApiClient)
		resp, err := conn.Secrets().Read(rs.Primary.Attributes["scope"], rs.Primary.Attributes["key"])
		//t.Log(resp)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*secret = resp
		return nil
		//return fmt.Errorf("Token (%s) not found", rs.Primary.ID)
	}
}

// testAccTokenResource returns an configuration for an Example Widget with the provided name
func testAwsSecretResource(scopeName, key, value string) string {
	return fmt.Sprintf(`
								resource "databricks_secret_scope" "my_scope" {
								  name = "%s"
								}
								resource "databricks_secret" "my_secret" {
								  key = "%s"
								  string_value = "%s"
								  scope = databricks_secret_scope.my_scope.name
								}
								`, scopeName, key, value)
}
