package scim

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
				Method:       "GET",
				Resource:     "/api/2.0/preview/scim/v2/Groups?attributes=id%2Cmembers&count=10000&startIndex=1",
				ReuseRequest: true,
				Response: GroupList{
					TotalResults: 1,
					Resources: []Group{
						{
							ID:      "abc",
							Members: []ComplexValue{{Value: "bcd"}},
						},
					},
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

	// Verify the per-endpoint cache was populated by the bulk fetch.
	assert.Equal(t, 1, len(globalGroupsCache.endpoints), "Cache should have one endpoint after create")
	var ep *groupEndpointCache
	for _, v := range globalGroupsCache.endpoints {
		ep = v
		break
	}
	require.NotNil(t, ep, "endpoint cache must exist")
	assert.Equal(t, 1, len(ep.cache), "Endpoint cache should contain one group entry")
	groupInfo, exists := ep.cache["abc"]
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
				Method:       "GET",
				Resource:     "/api/2.0/preview/scim/v2/Groups?attributes=id%2Cmembers&count=10000&startIndex=1",
				ReuseRequest: true,
				Response: GroupList{
					TotalResults: 1,
					Resources: []Group{
						{
							ID:      "abc",
							Members: []ComplexValue{{Value: "bcd"}},
						},
					},
				},
			},
		},
		Resource: ResourceGroupMember(),
		Read:     true,
		ID:       "abc|bcd",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc|bcd", d.Id(), "Id should not be empty")

	// Verify the per-endpoint cache was populated by the bulk fetch.
	assert.Equal(t, 1, len(globalGroupsCache.endpoints), "Cache should have one endpoint after read")
	var ep *groupEndpointCache
	for _, v := range globalGroupsCache.endpoints {
		ep = v
		break
	}
	require.NotNil(t, ep, "endpoint cache must exist")
	assert.Equal(t, 1, len(ep.cache), "Endpoint cache should contain one group entry")
	groupInfo, exists := ep.cache["abc"]
	assert.True(t, exists, "Cache should contain group 'abc'")
	assert.True(t, groupInfo.initialized, "Group info should be initialized")
	assert.Equal(t, 1, len(groupInfo.members), "Should have one member cached")
	_, memberExists := groupInfo.members["bcd"]
	assert.True(t, memberExists, "Member 'bcd' should be in cache")
}

func TestResourceGroupMemberRead_NoMember(t *testing.T) {
	globalGroupsCache = newGroupCache()
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.0/preview/scim/v2/Groups?attributes=id%2Cmembers&count=10000&startIndex=1",
				ReuseRequest: true,
				Response: GroupList{
					TotalResults: 1,
					Resources: []Group{
						{ID: "abc"},
					},
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
	globalGroupsCache = newGroupCache()
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.0/preview/scim/v2/Groups?attributes=id%2Cmembers&count=10000&startIndex=1",
				ReuseRequest: true,
				Status:       404,
				Response: apierr.APIError{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
			},
		},
		Resource: ResourceGroupMember(),
		Read:     true,
		Removed:  true,
		ID:       "abc|bcd",
	}.ApplyNoError(t)
}

func TestResourceGroupMemberRead_Error(t *testing.T) {
	globalGroupsCache = newGroupCache()
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.0/preview/scim/v2/Groups?attributes=id%2Cmembers&count=10000&startIndex=1",
				ReuseRequest: true,
				Status:       400,
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
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
