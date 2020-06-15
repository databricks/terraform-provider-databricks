package databricks

import (
	"errors"
	"fmt"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccAwsTokenResource(t *testing.T) {
	var tokenInfo model.TokenInfo

	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	rComment := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccAwsCheckTokenResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testAccAwsTokenResource(rComment),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAccAwsCheckTokenResourceExists("databricks_token.my-token", &tokenInfo, t),
					// verify remote values
					testAccAwsCheckTokenValues(&tokenInfo, rComment),
					// verify local values
					resource.TestCheckResourceAttr("databricks_token.my-token", "lifetime_seconds", "6000"),
					resource.TestCheckResourceAttr("databricks_token.my-token", "comment", rComment),
				),
			},
			{
				//Deleting and recreating the token
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DBApiClient)
					err := client.Tokens().Delete(tokenInfo.TokenID)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testAccAwsTokenResource(rComment),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAccAwsCheckTokenResourceExists("databricks_token.my-token", &tokenInfo, t),
					// verify remote values
					testAccAwsCheckTokenValues(&tokenInfo, rComment),
					// verify local values
					resource.TestCheckResourceAttr("databricks_token.my-token", "lifetime_seconds", "6000"),
					resource.TestCheckResourceAttr("databricks_token.my-token", "comment", rComment),
				),
			},
		},
	})
}

func testAccAwsCheckTokenResourceDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*service.DBApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_token" {
			continue
		}
		_, err := conn.Tokens().Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource token is not cleaned up")
	}
	return nil
}

func testAccAwsCheckTokenValues(tokenInfo *model.TokenInfo, comment string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if tokenInfo.Comment != comment {
			return errors.New("the comment for the token created does not equal the value passed in")
		}
		return nil
	}
}

// testAccAwsCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testAccAwsCheckTokenResourceExists(n string, tokenInfo *model.TokenInfo, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DBApiClient)
		resp, err := conn.Tokens().Read(rs.Primary.ID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*tokenInfo = resp
		return nil
		//return fmt.Errorf("Token (%s) not found", rs.Primary.ID)
	}
}

// testAccAwsTokenResource returns an configuration for an Example Widget with the provided name
func testAccAwsTokenResource(comment string) string {
	return fmt.Sprintf(`
								resource "databricks_token" "my-token" {
								  lifetime_seconds = 6000
								  comment = "%v"
								}
								`, comment)
}
