package scim

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceGroupMemberCreate(t *testing.T) {
	globalGroupsCache = newGroupCache()

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: PatchRequestWithValue("add", "members", "bcd"),
				Response: Group{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=members",
				Response: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					Members: []ComplexValue{
						{
							Value: "bcd",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceGroupMember(),
		State: map[string]any{
			"group_id":  "abc",
			"member_id": "bcd",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc|bcd", d.Id())

	assert.Equal(t, 1, len(globalGroupsCache.cache), "Cache should contain one entry after create")
	groupInfo, exists := globalGroupsCache.cache["abc"]
	assert.True(t, exists, "Cache should contain group 'abc'")

	assert.True(t, groupInfo.initialized, "Group info should be initialized")
	assert.Equal(t, 1, len(groupInfo.members), "Should have one member cached")
	_, memberExists := groupInfo.members["bcd"]
	assert.True(t, memberExists, "Member 'bcd' should be in cache")
}

func TestResourceGroupMemberCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceGroupMember(),
		State: map[string]any{
			"group_id":  "abc",
			"member_id": "bcd",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceGroupMemberRead(t *testing.T) {
	globalGroupsCache = newGroupCache()

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=members",
				Response: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					Members: []ComplexValue{
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
	assert.NoError(t, err)
	assert.Equal(t, "abc|bcd", d.Id(), "Id should not be empty")

	assert.Equal(t, 1, len(globalGroupsCache.cache), "Cache should contain one entry after read")
	groupInfo, exists := globalGroupsCache.cache["abc"]
	assert.True(t, exists, "Cache should contain group 'abc'")

	assert.True(t, groupInfo.initialized, "Group info should be initialized")
	assert.Equal(t, 1, len(groupInfo.members), "Should have one member cached")
	_, memberExists := groupInfo.members["bcd"]
	assert.True(t, memberExists, "Member 'bcd' should be in cache")
}

func TestResourceGroupMemberRead_NoMember(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=members",
				Response: Group{
					Schemas:     []URN{"urn:ietf:params:scim:schemas:core:2.0:Group"},
					DisplayName: "Data Scientists",
					ID:          "abc",
				},
			},
		},
		Resource: ResourceGroupMember(),
		Read:     true,
		Removed:  true,
		ID:       "abc|bcd",
	}.ApplyNoError(t)
}

func TestResourceGroupMemberRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=members",
				Response: apierr.APIError{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceGroupMember(),
		Read:     true,
		Removed:  true,
		ID:       "abc|bcd",
	}.ApplyNoError(t)
}

func TestResourceGroupMemberRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=members",
				Response: apierr.APIError{
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
	globalGroupsCache = newGroupCache()

	groupInfo := globalGroupsCache.getOrCreateGroupInfo("abc")
	groupInfo.initialized = true
	groupInfo.members["bcd"] = struct{}{}

	assert.Equal(t, 1, len(globalGroupsCache.cache), "Cache should contain one entry initially")
	groupInfo, exists := globalGroupsCache.cache["abc"]
	assert.True(t, exists, "Cache should contain group 'abc'")

	assert.Equal(t, 1, len(groupInfo.members), "Should have one member initially")
	_, memberExists := groupInfo.members["bcd"]
	assert.True(t, memberExists, "Member 'bcd' should be in cache initially")

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				ExpectedRequest: PatchRequest(
					"remove",
					`members[value eq "bcd"]`),
			},
		},
		Resource: ResourceGroupMember(),
		Delete:   true,
		ID:       "abc|bcd",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc|bcd", d.Id())

	groupInfo, exists = globalGroupsCache.cache["abc"]
	assert.True(t, exists, "Cache should still contain group 'abc'")

	assert.Equal(t, 0, len(groupInfo.members), "Should have no members after delete")
	_, memberExists = groupInfo.members["bcd"]
	assert.False(t, memberExists, "Member 'bcd' should not be in cache after delete")
}

func TestResourceGroupMemberDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: apierr.APIError{
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
