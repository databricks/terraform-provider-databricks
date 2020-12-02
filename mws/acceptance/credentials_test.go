package acceptance

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	. "github.com/databrickslabs/databricks-terraform/mws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"testing"
)

func TestMwsAccCredentials(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Cannot run test on non-MWS environment")
	}
	var creds Credentials
	config := qa.EnvironmentTemplate(t, `
	provider "databricks" {
		host     = "{env.DATABRICKS_HOST}"
		username = "{env.DATABRICKS_USERNAME}"
		password = "{env.DATABRICKS_PASSWORD}"
	}
	resource "databricks_mws_credentials" "my_e2_credentials" {
		account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
		credentials_name = "creds-test-{var.RANDOM}"
		role_arn         = "arn:aws:iam::999999999999:role/tf-test-{var.RANDOM}"
	}`)
	name := qa.FirstKeyValue(t, config, "credentials_name")
	arn := qa.FirstKeyValue(t, config, "role_arn")
	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testMWSCredentialsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testMWSCredentialsResourceExists("databricks_mws_credentials.my_e2_credentials", &creds, t),
					resource.TestCheckResourceAttr("databricks_mws_credentials.my_e2_credentials", "credentials_name", name),
					resource.TestCheckResourceAttr("databricks_mws_credentials.my_e2_credentials", "role_arn", arn),
				),
				Destroy: false,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testMWSCredentialsResourceExists("databricks_mws_credentials.my_e2_credentials", &creds, t),
					resource.TestCheckResourceAttr("databricks_mws_credentials.my_e2_credentials", "credentials_name", name),
					resource.TestCheckResourceAttr("databricks_mws_credentials.my_e2_credentials", "role_arn", arn),
				),
				ExpectNonEmptyPlan: false,
				Destroy:            false,
			},
			{
				PreConfig: func() {
					conn := common.CommonEnvironmentClient()
					err := NewCredentialsAPI(context.Background(), conn).Delete(creds.AccountID, creds.CredentialsID)
					if err != nil {
						panic(err)
					}
				},
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testMWSCredentialsResourceExists("databricks_mws_credentials.my_e2_credentials", &creds, t),
					resource.TestCheckResourceAttr("databricks_mws_credentials.my_e2_credentials", "credentials_name", name),
					resource.TestCheckResourceAttr("databricks_mws_credentials.my_e2_credentials", "role_arn", arn),
				),
				Destroy: false,
			},
		},
	})
}

func testMWSCredentialsResourceDestroy(s *terraform.State) error {
	client := common.CommonEnvironmentClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_mws_credentials" {
			continue
		}
		packagedMWSIds, err := UnpackMWSAccountID(rs.Primary.ID)
		if err != nil {
			return err
		}
		_, err = NewCredentialsAPI(context.Background(), client).Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		if err != nil {
			return nil
		}
		return errors.New("resource Scim Group is not cleaned up")
	}
	return nil
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testMWSCredentialsResourceExists(n string, mwsCreds *Credentials, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := common.CommonEnvironmentClient()
		packagedMWSIds, err := UnpackMWSAccountID(rs.Primary.ID)
		if err != nil {
			return err
		}
		resp, err := NewCredentialsAPI(context.Background(), conn).Read(packagedMWSIds.MwsAcctID, packagedMWSIds.ResourceID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*mwsCreds = resp
		return nil
	}
}
