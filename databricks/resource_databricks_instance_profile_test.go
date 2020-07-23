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

func TestAwsAccInstanceProfileResource(t *testing.T) {
	var InstanceProfile model.InstanceProfileInfo
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	randomStr := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	instanceProfile := fmt.Sprintf("arn:aws:iam::999999999999:instance-profile/%s", randomStr)
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAWSInstanceProfileResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testAWSDatabricksInstanceProfile(instanceProfile),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAWSInstanceProfileResourceExists("databricks_instance_profile.my_instance_profile", &InstanceProfile, t),
					// verify remote values
					testAWSInstanceProfileValues(t, &InstanceProfile, instanceProfile),
					// verify local values
					resource.TestCheckResourceAttr("databricks_instance_profile.my_instance_profile",
						"instance_profile_arn", instanceProfile),
				),
				Destroy: false,
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DatabricksClient)
					err := client.InstanceProfiles().Delete(instanceProfile)
					assert.NoError(t, err, err)
				},
				// use a dynamic configuration with the random name from above
				Config: testAWSDatabricksInstanceProfile(instanceProfile),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAWSInstanceProfileResourceExists("databricks_instance_profile.my_instance_profile", &InstanceProfile, t),
					// verify remote values
					testAWSInstanceProfileValues(t, &InstanceProfile, instanceProfile),
					// verify local values
					resource.TestCheckResourceAttr("databricks_instance_profile.my_instance_profile",
						"instance_profile_arn", instanceProfile),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testAWSDatabricksInstanceProfile(instanceProfile),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testAWSInstanceProfileResourceExists("databricks_instance_profile.my_instance_profile", &InstanceProfile, t),
					// verify remote values
					testAWSInstanceProfileValues(t, &InstanceProfile, instanceProfile),
					// verify local values
					resource.TestCheckResourceAttr("databricks_instance_profile.my_instance_profile",
						"instance_profile_arn", instanceProfile),
				),
				Destroy: false,
			},
		},
	})
}

func testAWSInstanceProfileResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*service.DatabricksClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_instance_profile" {
			continue
		}
		_, err := client.InstanceProfiles().Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource InstanceProfile is not cleaned up")
	}
	return nil
}

func testAWSInstanceProfileValues(t *testing.T, instanceProfileInfo *model.InstanceProfileInfo, instanceProfile string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, instanceProfileInfo.InstanceProfileArn == instanceProfile)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testAWSInstanceProfileResourceExists(n string, instanceProfileInfo *model.InstanceProfileInfo, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DatabricksClient)
		resp, err := conn.InstanceProfiles().Read(rs.Primary.ID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*instanceProfileInfo = model.InstanceProfileInfo{InstanceProfileArn: resp}
		return nil
	}
}

func testAWSDatabricksInstanceProfile(instanceProfile string) string {
	return fmt.Sprintf(`
								resource "databricks_instance_profile" "my_instance_profile" {
								  instance_profile_arn = "%s"
  								  skip_validation = true
								}
								`, instanceProfile)
}
