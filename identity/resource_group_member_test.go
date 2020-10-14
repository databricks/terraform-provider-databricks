package identity

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceGroupMemberCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: scimPatchRequest("add", "members", "bcd"),
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
					Members: []GroupMember{
						{
							Value: "bcd",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceGroupMember(),
		State: map[string]interface{}{
			"group_id":  "abc",
			"member_id": "bcd",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|bcd", d.Id())
}

func TestResourceGroupMemberCreate_Error(t *testing.T) {
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
		Resource: ResourceGroupMember(),
		State: map[string]interface{}{
			"group_id":  "abc",
			"member_id": "bcd",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceGroupMemberRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: ScimGroup{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					Members: []GroupMember{
						{
							Value: "bcd",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceGroupMember(),
		Read:     true,
		ID:       "abc|bcd",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|bcd", d.Id(), "Id should not be empty")
}

func TestResourceGroupMemberRead_NoMember(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: ScimGroup{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					ID:          "abc",
				},
			},
		},
		Resource: ResourceGroupMember(),
		Read:     true,
		ID:       "abc|bcd",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceGroupMemberRead_NotFound(t *testing.T) {
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
		Resource: ResourceGroupMember(),
		Read:     true,
		ID:       "abc|bcd",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceGroupMemberRead_Error(t *testing.T) {
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
		Resource: ResourceGroupMember(),
		Read:     true,
		ID:       "abc|bcd",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|bcd", d.Id(), "Id should not be empty for error reads")
}

func TestResourceGroupMemberDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: scimPatchRequest(
					"remove",
					`members[value eq "bcd"]`,
					""),
			},
		},
		Resource: ResourceGroupMember(),
		Delete:   true,
		ID:       "abc|bcd",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|bcd", d.Id())
}

func TestResourceGroupMemberDelete_Error(t *testing.T) {
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
		Resource: ResourceGroupMember(),
		Delete:   true,
		ID:       "abc|bcd",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|bcd", d.Id())
}
