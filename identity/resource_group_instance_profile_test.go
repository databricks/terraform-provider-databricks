package identity

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestValidateInstanceProfileARN(t *testing.T) {
	testCases := []struct {
		instanceProfileARN string
		errorCount         int
	}{
		{"arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", 0},
		{"arn:aws:iam::999999999999:role/not-an-instance-profile", 1},
		{"", 1},
		{"invalid-profile", 1},
	}
	for _, tc := range testCases {
		_, errs := ValidateInstanceProfileARN(tc.instanceProfileARN, "key")

		assert.Lenf(t, errs, tc.errorCount, "directory '%s' does not generate the expected error count", tc.instanceProfileARN)
	}
}

func TestResourceGroupInstanceProfileCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: scimPatchRequest(
					"add",
					"roles",
					"arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile"),
				Response: ScimGroup{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: ScimGroup{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					Roles: []RoleListItem{
						{"arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile"},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceGroupInstanceProfile(),
		State: map[string]interface{}{
			"group_id":            "abc",
			"instance_profile_id": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceGroupInstanceProfileCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroupInstanceProfile(),
		State: map[string]interface{}{
			"group_id":            "abc",
			"instance_profile_id": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceGroupInstanceProfileRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: ScimGroup{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					Roles: []RoleListItem{
						{"arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile"},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceGroupInstanceProfile(),
		Read:     true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id(), "Id should not be empty")
}

func TestResourceGroupInstanceProfileRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceGroupInstanceProfile(),
		Read:     true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceGroupInstanceProfileRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroupInstanceProfile(),
		Read:     true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id(), "Id should not be empty for error reads")
}

func TestResourceGroupInstanceProfileDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: scimPatchRequest(
					"remove",
					`roles[value eq "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile"]`,
					""),
			},
		},
		Resource: ResourceGroupInstanceProfile(),
		Delete:   true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceGroupInstanceProfileDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroupInstanceProfile(),
		Delete:   true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}
