package acceptance

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/access"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccSecretResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	var secret SecretMetadata
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	scope := fmt.Sprintf("tf-scope-%s", randomName)
	key := fmt.Sprintf("tf-key-%s", randomName)
	stringValue := "my super secret key"

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testSecretResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testSecretResource(scope, key, stringValue),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testSecretResourceExists("databricks_secret.my_secret", &secret, t),
					// verify remote values
					testSecretValues(t, &secret, key),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "scope", scope),
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "key", key),
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "string_value", stringValue),
				),
			},
			{
				//Deleting and recreating the secret
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewSecretsAPI(context.Background(), client).Delete(scope, secret.Key)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testSecretResource(scope, key, stringValue),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testSecretResourceExists("databricks_secret.my_secret", &secret, t),
					// verify remote values
					testSecretValues(t, &secret, key),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "scope", scope),
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "key", key),
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "string_value", stringValue),
				),
			},
			{
				//Deleting the scope should recreate the secret
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewSecretScopesAPI(context.Background(), client).Delete(scope)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testSecretResource(scope, key, stringValue),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testSecretResourceExists("databricks_secret.my_secret", &secret, t),
					// verify remote values
					testSecretValues(t, &secret, key),
					// verify local values
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "scope", scope),
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "key", key),
					resource.TestCheckResourceAttr("databricks_secret.my_secret", "string_value", stringValue),
				),
			},
		},
	})
}

func testSecretResourceDestroy(s *terraform.State) error {
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_secret" && rs.Type != "databricks_secret_scope" {
			continue
		}
		ctx := context.Background()
		_, err := NewSecretsAPI(ctx, client).Read(rs.Primary.Attributes["scope"], rs.Primary.Attributes["key"])
		if err == nil {
			return errors.New("resource secret is not cleaned up")
		}
		_, err = NewSecretScopesAPI(ctx, client).Read(rs.Primary.Attributes["scope"])
		if err == nil {
			return errors.New("resource secret is not cleaned up")
		}
	}
	return nil
}

func testSecretValues(t *testing.T, secret *SecretMetadata, key string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, secret.Key == key)
		assert.True(t, secret.LastUpdatedTimestamp > 0)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testSecretResourceExists(n string, secret *SecretMetadata, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := common.CommonEnvironmentClient()
		resp, err := NewSecretsAPI(context.Background(), conn).Read(rs.Primary.Attributes["scope"], rs.Primary.Attributes["key"])
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
func testSecretResource(scopeName, key, value string) string {
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
