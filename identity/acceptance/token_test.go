package acceptance

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/databrickslabs/databricks-terraform/common"
	. "github.com/databrickslabs/databricks-terraform/identity"
	"github.com/stretchr/testify/assert"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"testing"
)

func TestAccTokenResource(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	var tokenInfo TokenInfo

	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	rComment := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testCheckTokenResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testTokenResource(rComment),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testCheckTokenResourceExists("databricks_token.my-token", &tokenInfo, t),
					// verify remote values
					testCheckTokenValues(&tokenInfo, rComment),
					// verify local values
					resource.TestCheckResourceAttr("databricks_token.my-token", "lifetime_seconds", "6000"),
					resource.TestCheckResourceAttr("databricks_token.my-token", "comment", rComment),
				),
			},
			{
				//Deleting and recreating the token
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewTokensAPI(context.Background(), client).Delete(tokenInfo.TokenID)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testTokenResource(rComment),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testCheckTokenResourceExists("databricks_token.my-token", &tokenInfo, t),
					// verify remote values
					testCheckTokenValues(&tokenInfo, rComment),
					// verify local values
					resource.TestCheckResourceAttr("databricks_token.my-token", "lifetime_seconds", "6000"),
					resource.TestCheckResourceAttr("databricks_token.my-token", "comment", rComment),
				),
			},
		},
	})
}

func testCheckTokenResourceDestroy(s *terraform.State) error {
	conn := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_token" {
			continue
		}
		_, err := NewTokensAPI(context.Background(), conn).Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource token is not cleaned up")
	}
	return nil
}

func testCheckTokenValues(tokenInfo *TokenInfo, comment string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if tokenInfo.Comment != comment {
			return errors.New("the comment for the token created does not equal the value passed in")
		}
		return nil
	}
}

// testCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testCheckTokenResourceExists(n string, tokenInfo *TokenInfo, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := common.CommonEnvironmentClient()
		resp, err := NewTokensAPI(context.Background(), conn).Read(rs.Primary.ID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*tokenInfo = resp
		return nil
		//return fmt.Errorf("Token (%s) not found", rs.Primary.ID)
	}
}

// testTokenResource returns an configuration for an Example Widget with the provided name
func testTokenResource(comment string) string {
	return fmt.Sprintf(`
		resource "databricks_token" "my-token" {
			lifetime_seconds = 6000
			comment = "%v"
		}
		`, comment)
}
