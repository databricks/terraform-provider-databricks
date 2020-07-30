package databricks

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccSecretResource(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
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
					client := testAccProvider.Meta().(*service.DatabricksClient)
					err := client.Secrets().Delete(scope, secret.Key)
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
					client := testAccProvider.Meta().(*service.DatabricksClient)
					err := client.SecretScopes().Delete(scope)
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
	client := testAccProvider.Meta().(*service.DatabricksClient)
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

func testSecretValues(t *testing.T, secret *model.SecretMetadata, key string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, secret.Key == key)
		assert.True(t, secret.LastUpdatedTimestamp > 0)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testSecretResourceExists(n string, secret *model.SecretMetadata, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DatabricksClient)
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

func TestResourceSecretRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/list?scope=foo",
			Response: model.SecretsList{
				Secrets: []model.SecretMetadata{
					{
						Key:                  "bar",
						LastUpdatedTimestamp: 12345678,
					},
				},
			},
		},
	}, resourceSecret, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId("foo|||bar")
		return resourceSecretRead(d, c)
	})
	assert.NoError(t, err, err)
	assert.Equal(t, "foo|||bar", d.Id())
	assert.Equal(t, "bar", d.Get("key"))
	assert.Equal(t, 12345678, d.Get("last_updated_timestamp"))
	assert.Equal(t, "foo", d.Get("scope"))
	assert.Equal(t, "", d.Get("string_value"))
}

func TestResourceSecretRead_NotFound(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/list?scope=foo",
			Response: model.SecretsList{
				Secrets: []model.SecretMetadata{
					{
						Key:                  "bar",
						LastUpdatedTimestamp: 12345678,
					},
				},
			},
		},
	}, resourceSecret, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId("foo|||missing")
		return resourceSecretRead(d, c)
	})
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id())
}

func TestResourceSecretRead_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/list?scope=foo",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceSecret, nil, actionWithID("foo|||bar", resourceSecretRead))
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "foo|||bar", d.Id(), "Id should not be empty for error reads")
}

func TestResourceSecretCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/secrets/put",
			ExpectedRequest: model.SecretsRequest{
				StringValue: "SparkIsTh3Be$t",
				Scope:       "foo",
				Key:         "bar",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/list?scope=foo",
			Response: model.SecretsList{
				Secrets: []model.SecretMetadata{
					{
						Key:                  "bar",
						LastUpdatedTimestamp: 12345678,
					},
				},
			},
		},
	}, resourceSecret, map[string]interface{}{
		"scope":        "foo",
		"key":          "bar",
		"string_value": "SparkIsTh3Be$t",
	}, resourceSecretCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "foo|||bar", d.Id())
}

func TestResourceSecretCreate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/secrets/put",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceSecret, map[string]interface{}{
		"key":          "...",
		"scope":        "...",
		"string_value": "...",
	}, resourceSecretCreate)
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceSecretDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{ // read log output for better stub url...
			Method:   "POST",
			Resource: "/api/2.0/secrets/delete",
			ExpectedRequest: map[string]string{
				"scope": "foo",
				"key":   "bar",
			},
		},
	}, resourceSecret, nil, actionWithID("foo|||bar", resourceSecretDelete))
	assert.NoError(t, err, err)
	assert.Equal(t, "foo|||bar", d.Id())
}

func TestResourceSecretDelete_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/secrets/delete",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceSecret, nil, actionWithID("foo|||bar", resourceSecretDelete))
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "foo|||bar", d.Id())
}
