package acceptance

import (
	"context"
	"errors"
	"fmt"

	"github.com/databrickslabs/databricks-terraform/common"
	. "github.com/databrickslabs/databricks-terraform/identity"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"

	"os"
	"testing"
)

func TestAwsAccGroupInstanceProfileResource(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	var group ScimGroup
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest

	groupName := "group test"
	role := "arn:aws:iam::999999999999:instance-profile/terraform-group-test"

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testGroupInstanceProfileResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testGroupInstanceProfileResourceCreate(role, groupName),

				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testGroupInstanceProfileResourceExists("databricks_group.my_group", &group, t),
					// verify remote values
					testGroupInstanceProfileValues(t, &group, groupName, role),
					// verify local values
					resource.TestCheckResourceAttr("databricks_group.my_group", "display_name", groupName),
					resource.TestCheckResourceAttr("databricks_group_instance_profile.my_group_instance_profile", "instance_profile_id", role),
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testGroupInstanceProfileResourceCreate(role, groupName),

				// Test behavior to expect to attempt to create new role mapping because role is gone
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewInstanceProfilesAPI(client).Delete(role)
					assert.NoError(t, err, err)
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testGroupInstanceProfileResourceCreate(role, groupName),

				// Test behavior to expect to attempt to create new role mapping because role is gone
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewGroupsAPI(client).Delete(group.ID)
					assert.NoError(t, err, err)
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
		},
	})
}

func testGroupInstanceProfileResourceDestroy(s *terraform.State) error {
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_group" {
			continue
		}
		ctx := context.Background()
		usersAPI := NewUsersAPI(ctx, client)
		_, err := usersAPI.Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource Group is not cleaned up")
	}
	return nil
}

func testGroupInstanceProfileValues(t *testing.T, group *ScimGroup, displayName, role string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, group.DisplayName == displayName)
		assert.True(t, group.HasRole(role), "role is not in group")
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testGroupInstanceProfileResourceExists(n string, group *ScimGroup, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := common.CommonEnvironmentClient()
		resp, err := NewGroupsAPI(conn).Read(rs.Primary.ID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*group = resp
		return nil
	}
}

func testGroupInstanceProfileResourceCreate(roleArn, groupName string) string {
	return fmt.Sprintf(`
		resource "databricks_instance_profile" "instance_profile" {
			instance_profile_arn = "%s"
			skip_validation = true
		}
		resource "databricks_group" "my_group" {
			display_name = "%s"
		}
		resource "databricks_group_instance_profile" "my_group_instance_profile" {
			group_id = databricks_group.my_group.id
			instance_profile_id = databricks_instance_profile.instance_profile.id
		}
		`, roleArn, groupName)
}

func TestAwsAccInstanceProfileResource(t *testing.T) {
	var InstanceProfile InstanceProfileInfo
	// generate a random name for each tokenInfo test run, to avoid
	// collisions from multiple concurrent tests.
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	randomStr := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	instanceProfile := fmt.Sprintf("arn:aws:iam::999999999999:instance-profile/%s", randomStr)
	acceptance.AccTest(t, resource.TestCase{
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
					client := common.CommonEnvironmentClient()
					err := NewInstanceProfilesAPI(client).Delete(instanceProfile)
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
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_instance_profile" {
			continue
		}
		_, err := NewInstanceProfilesAPI(client).Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource InstanceProfile is not cleaned up")
	}
	return nil
}

func testAWSInstanceProfileValues(t *testing.T, instanceProfileInfo *InstanceProfileInfo, instanceProfile string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		assert.True(t, instanceProfileInfo.InstanceProfileArn == instanceProfile)
		return nil
	}
}

// testAccCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testAWSInstanceProfileResourceExists(n string, instanceProfileInfo *InstanceProfileInfo, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := common.CommonEnvironmentClient()
		resp, err := NewInstanceProfilesAPI(conn).Read(rs.Primary.ID)
		if err != nil {
			return err
		}

		// If no error, assign the response Widget attribute to the widget pointer
		*instanceProfileInfo = InstanceProfileInfo{InstanceProfileArn: resp}
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
