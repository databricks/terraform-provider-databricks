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
	// TODO: refactor for common instance pool & AZ CLI
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

func TestResourceInstanceProfileCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/instance-profiles/add",
			ExpectedRequest: map[string]interface{}{
				"instance_profile_arn": "abc",
				"skip_validation":      true,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-profiles/list?",
			Response: model.InstanceProfileList{
				InstanceProfiles: []model.InstanceProfileInfo{
					{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
	}, resourceInstanceProfile, map[string]interface{}{
		"instance_profile_arn": "abc",
		"skip_validation":      true,
	}, resourceInstanceProfileCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceInstanceProfileCreate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/instance-profiles/add",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceInstanceProfile, map[string]interface{}{
		"instance_profile_arn": "abc",
		"skip_validation":      true,
	}, resourceInstanceProfileCreate)
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceInstanceProfileRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-profiles/list?",
			Response: model.InstanceProfileList{
				InstanceProfiles: []model.InstanceProfileInfo{
					{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
	}, resourceInstanceProfile, nil, actionWithID("abc", resourceInstanceProfileRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "abc", d.Get("instance_profile_arn"))
	assert.Equal(t, false, d.Get("skip_validation"))
}

func TestResourceInstanceProfileRead_NotFound(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{ // read log output for correct url...
			Method:   "GET",
			Resource: "/api/2.0/instance-profiles/list?",
			Response: model.InstanceProfileList{
				InstanceProfiles: []model.InstanceProfileInfo{},
			},
		},
	}, resourceInstanceProfile, nil, actionWithID("abc", resourceInstanceProfileRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceInstanceProfileRead_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-profiles/list?",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceInstanceProfile, nil, actionWithID("abc", resourceInstanceProfileRead))
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceInstanceProfileDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/instance-profiles/remove",
			ExpectedRequest: model.InstanceProfileInfo{
				InstanceProfileArn: "abc",
			},
		},
	}, resourceInstanceProfile, nil, actionWithID("abc", resourceInstanceProfileDelete))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceInstanceProfileDelete_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/instance-profiles/remove",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceInstanceProfile, nil, actionWithID("abc", resourceInstanceProfileDelete))
	assert.EqualError(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
