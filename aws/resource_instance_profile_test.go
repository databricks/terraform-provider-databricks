package aws

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
)

func TestResourceInstanceProfileCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/add",
				ExpectedRequest: InstanceProfileInfo{
					InstanceProfileArn: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
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
		State: map[string]any{
			"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceInstanceProfileWithRoleCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/add",
				ExpectedRequest: InstanceProfileInfo{
					InstanceProfileArn: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
					IamRoleArn:         "arn:aws:iam::999999999999:role/my-fake-instance-profile-role",
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
		State: map[string]any{
			"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
			"iam_role_arn":         "arn:aws:iam::999999999999:role/my-fake-instance-profile-role",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceInstanceProfileWithEmptyRoleCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/add",
				ExpectedRequest: InstanceProfileInfo{
					InstanceProfileArn: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
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
		State: map[string]any{
			"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
			"iam_role_arn":         "",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceInstanceProfileCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/add",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceInstanceProfile(),
		State: map[string]any{
			"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceInstanceProfileValidate_Error_InvalidInstanceProfileARN(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceInstanceProfile(),
		State: map[string]any{
			"instance_profile_arn": "abc",
		},
		Create: true,
	}.Apply(t)
	assert.EqualError(t, err, "invalid config supplied. [instance_profile_arn] Invalid ARN")
}

func TestResourceInstanceProfileValidate_Error_InvalidRoleARN(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceInstanceProfile(),
		State: map[string]any{
			"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
			"iam_role_arn":         "abc",
		},
		Create: true,
	}.Apply(t)
	assert.EqualError(t, err, "invalid config supplied. [iam_role_arn] Invalid ARN")
}

func TestResourceInstanceProfileValidate_Error_MalformedARN(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceInstanceProfile(),
		State: map[string]any{
			"instance_profile_arn": "arn:aws:iam::instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.Apply(t)
	assert.EqualError(t, err, "invalid config supplied. [instance_profile_arn] Invalid ARN")
}

func TestResourceInstanceProfileValidate_Error_WrongTypeProfileARN(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceInstanceProfile(),
		State: map[string]any{
			"instance_profile_arn": "arn:aws:iam::999999999999:failure/my-fake-instance-profile",
		},
		Create: true,
	}.Apply(t)
	assert.EqualError(t, err, "invalid config supplied. [instance_profile_arn] Invalid ARN")
}

func TestResourceInstanceProfileValidate_Error_WrongTypeRoleARN(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceInstanceProfile(),
		State: map[string]any{
			"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
			"iam_role_arn":         "arn:aws:iam::999999999999:failure/my-fake-instance-profile-role",
		},
		Create: true,
	}.Apply(t)
	assert.EqualError(t, err, "invalid config supplied. [iam_role_arn] Invalid ARN")
}

func TestResourceInstanceProfileValidate_Error_EmptyInstanceProfileARN(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceInstanceProfile(),
		State: map[string]any{
			"instance_profile_arn": "",
			"iam_role_arn":         "arn:aws:iam::999999999999:role/my-fake-instance-profile-role",
		},
		Create: true,
	}.Apply(t)
	assert.EqualError(t, err, "invalid config supplied. [instance_profile_arn] Invalid ARN")
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
		New:      true,
		ID:       "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err)
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
				Response: apierr.APIError{
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
	assert.NoError(t, err)
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceInstanceProfileDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/remove",
				Response: apierr.APIError{
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

func TestResourceInstanceProfileUpdate(t *testing.T) {
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
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/edit",
				ExpectedRequest: InstanceProfileInfo{
					InstanceProfileArn: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
					IamRoleArn:         "arn:aws:iam::999999999999:role/my-fake-instance-profile-role",
				},
			},
		},
		Resource: ResourceInstanceProfile(),
		State: map[string]any{
			"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
			"iam_role_arn":         "arn:aws:iam::999999999999:role/my-fake-instance-profile-role",
		},
		Update: true,
		ID:     "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceInstanceProfileUpdate_Error(t *testing.T) {
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
			{
				Method:   "POST",
				Resource: "/api/2.0/instance-profiles/edit",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceInstanceProfile(),
		State: map[string]any{
			"instance_profile_arn": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
			"iam_role_arn":         "",
		},
		Update: true,
		ID:     "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}
