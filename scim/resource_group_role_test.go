package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceGroupRoleCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: PatchRequest("add", "roles", "arn:aws:iam::000000000000:role/test-role"),
				Response: Group{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					Roles: []ComplexValue{
						{
							Value: "arn:aws:iam::000000000000:role/test-role",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceGroupRole(),
		State: map[string]any{
			"group_id": "abc",
			"role_arn": "arn:aws:iam::000000000000:role/test-role",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|arn:aws:iam::000000000000:role/test-role", d.Id())
}

func TestResourceGroupRoleCreate_Error(t *testing.T) {
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
		Resource: ResourceGroupRole(),
		State: map[string]any{
			"group_id": "abc",
			"role_arn": "arn:aws:iam::000000000000:role/test-role",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceGroupRoleRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					Roles: []ComplexValue{
						{
							Value: "arn:aws:iam::000000000000:role/test-role",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceGroupRole(),
		Read:     true,
		ID:       "abc|arn:aws:iam::000000000000:role/test-role",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|arn:aws:iam::000000000000:role/test-role", d.Id(), "Id should not be empty")
}

func TestResourceGroupRoleRead_NoRole(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					ID:          "abc",
				},
			},
		},
		Resource: ResourceGroupRole(),
		Read:     true,
		Removed:  true,
		ID:       "abc|arn:aws:iam::000000000000:role/test-role",
	}.ApplyNoError(t)
}

func TestResourceGroupRoleRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
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
		Resource: ResourceGroupRole(),
		Read:     true,
		Removed:  true,
		ID:       "abc|arn:aws:iam::000000000000:role/test-role",
	}.ApplyNoError(t)
}

func TestResourceGroupRoleRead_Error(t *testing.T) {
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
		Resource: ResourceGroupRole(),
		Read:     true,
		ID:       "abc|arn:aws:iam::000000000000:role/test-role",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|arn:aws:iam::000000000000:role/test-role", d.Id(), "Id should not be empty for error reads")
}

func TestResourceGroupRoleDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: PatchRequest(
					"remove",
					`roles[value eq "arn:aws:iam::000000000000:role/test-role"]`,
					""),
			},
		},
		Resource: ResourceGroupRole(),
		Delete:   true,
		ID:       "abc|arn:aws:iam::000000000000:role/test-role",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|arn:aws:iam::000000000000:role/test-role", d.Id())
}

func TestResourceGroupRoleDelete_Error(t *testing.T) {
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
		Resource: ResourceGroupRole(),
		Delete:   true,
		ID:       "abc|arn:aws:iam::000000000000:role/test-role",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|arn:aws:iam::000000000000:role/test-role", d.Id())
}
