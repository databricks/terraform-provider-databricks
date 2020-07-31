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

func TestMwsAccCredentials(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Cannot run test on non-MWS environment")
	}
	var MWSCredentials model.MWSCredentials
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	mwsAcctID := os.Getenv("DATABRICKS_MWS_ACCT_ID")
	mwsHost := os.Getenv("DATABRICKS_MWS_HOST")
	awsAcctID := "999999999999"
	credentialsName := "test-mws-credentials-tf"
	roleName := "terraform-creds-role"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testMWSCredentialsResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testMWSCredentialsCreate(mwsAcctID, mwsHost, awsAcctID, roleName, credentialsName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testMWSCredentialsResourceExists("databricks_mws_credentials.my_e2_credentials", &MWSCredentials, t),
					// verify local values
					resource.TestCheckResourceAttr("databricks_mws_credentials.my_e2_credentials", "account_id", mwsAcctID),
					resource.TestCheckResourceAttr("databricks_mws_credentials.my_e2_credentials", "credentials_name", credentialsName),
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testMWSCredentialsCreate(mwsAcctID, mwsHost, awsAcctID, roleName, credentialsName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testMWSCredentialsResourceExists("databricks_mws_credentials.my_e2_credentials", &MWSCredentials, t),
					// verify local values
					resource.TestCheckResourceAttr("databricks_mws_credentials.my_e2_credentials", "account_id", mwsAcctID),
					resource.TestCheckResourceAttr("databricks_mws_credentials.my_e2_credentials", "credentials_name", credentialsName),
				),
				ExpectNonEmptyPlan: false,
				Destroy:            false,
			},
			{
				PreConfig: func() {
					conn := getMWSClient()
					err := conn.MWSCredentials().Delete(MWSCredentials.AccountID, MWSCredentials.CredentialsID)
					if err != nil {
						panic(err)
					}
				},
				// use a dynamic configuration with the random name from above
				Config: testMWSCredentialsCreate(mwsAcctID, mwsHost, awsAcctID, roleName, credentialsName),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// verify local values
					resource.TestCheckResourceAttr("databricks_mws_credentials.my_e2_credentials", "account_id", mwsAcctID),
					resource.TestCheckResourceAttr("databricks_mws_credentials.my_e2_credentials", "credentials_name", credentialsName),
				),
				Destroy: false,
			},
		},
	})
}

func testMWSCredentialsResourceDestroy(s *terraform.State) error {
	client := getMWSClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_mws_credentials" {
			continue
		}
		packagedMWSIds, err := unpackMWSAccountID(rs.Primary.ID)
		if err != nil {
			return err
		}
		_, err = client.MWSCredentials().Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		if err != nil {
			return nil
		}
		return errors.New("resource Scim Group is not cleaned up")
	}
	return nil
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testMWSCredentialsResourceExists(n string, mwsCreds *model.MWSCredentials, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := getMWSClient()
		packagedMWSIds, err := unpackMWSAccountID(rs.Primary.ID)
		if err != nil {
			return err
		}
		resp, err := conn.MWSCredentials().Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*mwsCreds = resp
		return nil
	}
}

func testMWSCredentialsCreate(mwsAcctID, mwsHost, awsAcctID, roleName, credentialsName string) string {
	return fmt.Sprintf(`
								provider "databricks" {
								  host = "%s"
								  basic_auth {}
								}
								resource "databricks_mws_credentials" "my_e2_credentials" {
 								  account_id       = "%s"
								  credentials_name = "%s"
								  role_arn         = "arn:aws:iam::%s:role/%s"
								}
								`, mwsHost, mwsAcctID, credentialsName, awsAcctID, roleName)
}

func TestResourceMWSCredentialsCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/abc/credentials",
			ExpectedRequest: model.MWSCredentials{
				CredentialsName: "Cross-account ARN",
				AwsCredentials: &model.AwsCredentials{
					StsRole: &model.StsRole{
						RoleArn: "arn:aws:iam::098765:role/cross-account",
					},
				},
			},
			Response: model.MWSCredentials{
				CredentialsID: "cid",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/credentials/cid?",
			Response: model.MWSCredentials{
				CredentialsID:   "cid",
				CredentialsName: "Cross-account ARN",
				AwsCredentials: &model.AwsCredentials{
					StsRole: &model.StsRole{
						RoleArn: "arn:aws:iam::098765:role/cross-account",
					},
				},
			},
		},
	}, resourceMWSCredentials, map[string]interface{}{
		"account_id":       "abc",
		"credentials_name": "Cross-account ARN",
		"role_arn":         "arn:aws:iam::098765:role/cross-account",
	}, resourceMWSCredentialsCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/cid", d.Id())
}

func TestResourceMWSCredentialsCreate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/abc/credentials",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSCredentials, map[string]interface{}{
		"account_id":       "abc",
		"credentials_name": "Cross-account ARN",
		"role_arn":         "arn:aws:iam::098765:role/cross-account",
	}, resourceMWSCredentialsCreate)
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceMWSCredentialsRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/credentials/cid?",
			Response: model.MWSCredentials{
				CredentialsID:   "cid",
				CredentialsName: "Cross-account ARN",
				AwsCredentials: &model.AwsCredentials{
					StsRole: &model.StsRole{
						RoleArn: "arn:aws:iam::098765:role/cross-account",
					},
				},
			},
		},
	}, resourceMWSCredentials, nil, actionWithID("abc/cid", resourceMWSCredentialsRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/cid", d.Id(), "Id should not be empty")
	assert.Equal(t, 0, d.Get("creation_time"))
	assert.Equal(t, "cid", d.Get("credentials_id"))
	assert.Equal(t, "Cross-account ARN", d.Get("credentials_name"))
	assert.Equal(t, "", d.Get("external_id"))
	assert.Equal(t, "arn:aws:iam::098765:role/cross-account", d.Get("role_arn"))
}

func TestResourceMWSCredentialsRead_NotFound(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/credentials/cid?",
			Response: service.APIErrorBody{
				ErrorCode: "NOT_FOUND",
				Message:   "Item not found",
			},
			Status: 404,
		},
	}, resourceMWSCredentials, nil, actionWithID("abc/cid", resourceMWSCredentialsRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceMWSCredentialsRead_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/credentials/cid?",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSCredentials, nil, actionWithID("abc/cid", resourceMWSCredentialsRead))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/cid", d.Id(), "Id should not be empty for error reads")
}

func TestResourceMWSCredentialsDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "DELETE",
			Resource: "/api/2.0/accounts/abc/credentials/cid",
		},
	}, resourceMWSCredentials, nil, actionWithID("abc/cid", resourceMWSCredentialsDelete))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/cid", d.Id())
}

func TestResourceMWSCredentialsDelete_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "DELETE",
			Resource: "/api/2.0/accounts/abc/credentials/cid",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceMWSCredentials, nil, actionWithID("abc/cid", resourceMWSCredentialsDelete))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/cid", d.Id())
}
