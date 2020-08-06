package databricks

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccTokenResource(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	var tokenInfo model.TokenInfo

	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	rComment := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
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
					client := testAccProvider.Meta().(*service.DatabricksClient)
					err := client.Tokens().Delete(tokenInfo.TokenID)
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
	conn := testAccProvider.Meta().(*service.DatabricksClient)
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

func testCheckTokenValues(tokenInfo *model.TokenInfo, comment string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if tokenInfo.Comment != comment {
			return errors.New("the comment for the token created does not equal the value passed in")
		}
		return nil
	}
}

// testCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testCheckTokenResourceExists(n string, tokenInfo *model.TokenInfo, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DatabricksClient)
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

// testTokenResource returns an configuration for an Example Widget with the provided name
func testTokenResource(comment string) string {
	return fmt.Sprintf(`
		resource "databricks_token" "my-token" {
			lifetime_seconds = 6000
			comment = "%v"
		}
		`, comment)
}

func TestResourceTokenRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/token/list",
			Response: model.TokenList{
				TokenInfos: []model.TokenInfo{
					{
						Comment:      "Hello, world!",
						CreationTime: 10,
						ExpiryTime:   20,
						TokenID:      "abc",
					},
				},
			},
		},
	}, resourceToken, nil, actionWithID("abc", resourceTokenRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "Hello, world!", d.Get("comment"))
	assert.Equal(t, 10, d.Get("creation_time"))
	assert.Equal(t, 20, d.Get("expiry_time"))
	assert.Equal(t, "", d.Get("token_value"))
}

func TestResourceTokenRead_NotFound(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/token/list",
			Response: model.TokenList{
				TokenInfos: []model.TokenInfo{
					{
						Comment:      "Hello, world!",
						CreationTime: 10,
						ExpiryTime:   20,
						TokenID:      "bcd",
					},
				},
			},
		},
	}, resourceToken, nil, actionWithID("abc", resourceTokenRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceTokenRead_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/token/list",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceToken, nil, actionWithID("abc", resourceTokenRead))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceTokenCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/token/create",
			ExpectedRequest: model.TokenRequest{
				LifetimeSeconds: 300,
				Comment:         "Hello world!",
			},
			Response: model.TokenResponse{
				TokenValue: "dapi...",
				TokenInfo: &model.TokenInfo{
					TokenID: "abc",
					// other fields may be irrelevant...
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/token/list",
			Response: model.TokenList{
				TokenInfos: []model.TokenInfo{
					{
						Comment:      "Hello, world!",
						CreationTime: 10,
						ExpiryTime:   20,
						TokenID:      "abc",
					},
				},
			},
		},
	}, resourceToken, map[string]interface{}{
		"comment":          "Hello world!",
		"lifetime_seconds": 300,
	}, resourceTokenCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "dapi...", d.Get("token_value"))
}

func TestResourceTokenCreate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/token/create",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceToken, map[string]interface{}{
		"comment":          "Hello world!",
		"lifetime_seconds": 300,
	}, resourceTokenCreate)
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceTokenDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{ // read log output for better stub url...
			Method:   "POST",
			Resource: "/api/2.0/token/delete",
			ExpectedRequest: map[string]string{
				"token_id": "abc",
			},
		},
	}, resourceToken, nil, actionWithID("abc", resourceTokenDelete))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceTokenDelete_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/token/delete",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceToken, nil, actionWithID("abc", resourceTokenDelete))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
