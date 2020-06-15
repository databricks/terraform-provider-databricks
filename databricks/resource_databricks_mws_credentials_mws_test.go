package databricks

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccMWSCredentials(t *testing.T) {
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
