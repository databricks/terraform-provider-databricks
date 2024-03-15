package aws

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/scim"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceUserInstanceProfileCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				ExpectedRequest: scim.PatchRequestWithValue(
					"add",
					"roles",
					"arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile"),
				Response: scim.User{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=roles",
				Response: scim.User{
					Schemas:     []scim.URN{"urn:ietf:params:scim:schemas:core:2.0:User"},
					DisplayName: "Data Scientists",
					Roles: []scim.ComplexValue{
						{
							Value: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceUserInstanceProfile(),
		State: map[string]any{
			"user_id":             "abc",
			"instance_profile_id": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceUserInstanceProfileCreate_Error_BadARN(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceUserInstanceProfile(),
		State: map[string]any{
			"user_id":             "abc",
			"instance_profile_id": "fake",
		},
		Create: true,
	}.Apply(t)
	assert.EqualError(t, err, "invalid config supplied. [instance_profile_id] Invalid ARN. Deprecated Resource")
}

func TestResourceUserInstanceProfileCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceUserInstanceProfile(),
		State: map[string]any{
			"user_id":             "abc",
			"instance_profile_id": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceUserInstanceProfileRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=roles",
				Response: scim.User{
					Schemas:     []scim.URN{"urn:ietf:params:scim:schemas:core:2.0:User"},
					DisplayName: "Data Scientists",
					Roles: []scim.ComplexValue{
						{
							Value: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceUserInstanceProfile(),
		Read:     true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id(), "Id should not be empty")
}

func TestResourceUserInstanceProfileRead_NoRole(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=roles",
				Response: scim.User{
					Schemas:     []scim.URN{"urn:ietf:params:scim:schemas:core:2.0:User"},
					DisplayName: "Data Scientists",
					ID:          "abc",
				},
			},
		},
		Resource: ResourceUserInstanceProfile(),
		Read:     true,
		Removed:  true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.ApplyNoError(t)
}

func TestResourceUserInstanceProfileRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=roles",
				Response: apierr.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceUserInstanceProfile(),
		Read:     true,
		Removed:  true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.ApplyNoError(t)
}

func TestResourceUserInstanceProfileRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/abc?attributes=roles",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceUserInstanceProfile(),
		Read:     true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id(), "Id should not be empty for error reads")
}

func TestResourceUserInstanceProfileDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				ExpectedRequest: scim.PatchRequest(
					"remove",
					`roles[value eq "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile"]`),
			},
		},
		Resource: ResourceUserInstanceProfile(),
		Delete:   true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceUserInstanceProfileDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceUserInstanceProfile(),
		Delete:   true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}
