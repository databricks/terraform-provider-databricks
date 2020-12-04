package identity

import (
	"context"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceInstanceProfileCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/add",
				ExpectedRequest: map[string]interface{}{
					"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
					"skip_validation":      false,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: InstanceProfileList{
					InstanceProfiles: []InstanceProfileInfo{
						{
							InstanceProfileArn: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
						},
					},
				},
			},
		},
		Resource: ResourceInstanceProfile(),
		State: map[string]interface{}{
			"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceInstanceProfileCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/add",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceInstanceProfile(),
		State: map[string]interface{}{
			"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceInstanceProfileCreate_Error_InvalidARN(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceInstanceProfile(),
		State: map[string]interface{}{
			"instance_profile_arn": "abc",
		},
		Create: true,
	}.Apply(t)
	assert.EqualError(t, err, "Invalid config supplied. [instance_profile_arn] Invalid ARN")
}

func TestResourceInstanceProfileRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: InstanceProfileList{
					InstanceProfiles: []InstanceProfileInfo{
						{
							InstanceProfileArn: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
						},
					},
				},
			},
		},
		Resource: ResourceInstanceProfile(),
		Read:     true,
		ID:       "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id(), "Id should not be empty")
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Get("instance_profile_arn"))
}

func TestResourceInstanceProfileRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: InstanceProfileList{
					InstanceProfiles: []InstanceProfileInfo{},
				},
			},
		},
		Resource: ResourceInstanceProfile(),
		Read:     true,
		Removed:  true,
		ID:       "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.ApplyNoError(t)
}

func TestResourceInstanceProfileRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceInstanceProfile(),
		Read:     true,
		ID:       "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id(), "Id should not be empty for error reads")
}

func TestResourceInstanceProfileDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/remove",
				ExpectedRequest: InstanceProfileInfo{
					InstanceProfileArn: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
				},
			},
		},
		Resource: ResourceInstanceProfile(),
		Delete:   true,
		ID:       "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceInstanceProfileDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/remove",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceInstanceProfile(),
		Delete:   true,
		ID:       "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestAwsAccInstanceProfiles(t *testing.T) {
	arn := qa.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	client := common.NewClientFromEnvironment()
	ctx := context.Background()
	instanceProfilesAPI := NewInstanceProfilesAPI(ctx, client)
	instanceProfilesAPI.Synchronized(arn, func() {
		err := instanceProfilesAPI.Create(arn, true)
		assert.NoError(t, err, err)
		defer func() {
			err := instanceProfilesAPI.Delete(arn)
			assert.NoError(t, err, err)
		}()

		arnSearch, err := instanceProfilesAPI.Read(arn)
		assert.NoError(t, err, err)
		assert.True(t, len(arnSearch) > 0)
	})
}
