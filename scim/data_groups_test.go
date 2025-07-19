package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceGroupsEmpty(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: `/api/2.0/preview/scim/v2/Groups?filter=displayName%20co%20%22nonexistent%22`,
				Response: GroupList{
					Resources: []Group{},
				},
			},
		},
		Resource:    DataSourceGroups(),
		HCL:         `filter = "displayName co \"nonexistent\""`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"groups": []any{},
	})
}


func TestDataSourceGroupsMultiple(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: `/api/2.0/preview/scim/v2/Groups?filter=displayName%20co%20%22team%22`,
				Response: GroupList{
					Resources: []Group{
						{
							DisplayName: "team-alpha",
							ID:          "g1",
							ExternalID:  "ext1",
							Members: []ComplexValue{
								{
									Ref:   "Users/u1",
									Value: "user1",
								},
								{
									Ref:   "ServicePrincipals/sp1",
									Value: "service-principal-1",
								},
								{
									Ref:   "Groups/g3",
									Value: "child-group-1",
								},
							},
							Groups: []ComplexValue{
								{
									Value: "parent-group-1",
								},
							},
							Roles: []ComplexValue{
								{
									Value: "instance-profile-1",
								},
							},
						},
						{
							DisplayName: "team-beta",
							ID:          "g2",
							ExternalID:  "ext2",
							Members: []ComplexValue{
								{
									Ref:   "Users/u2",
									Value: "user2",
								},
								{
									Ref:   "Users/u3",
									Value: "user3",
								},
							},
						},
					},
				},
			},
		},
		Resource:    DataSourceGroups(),
		HCL:         `filter = "displayName co \"team\""`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err)

	groups := d.Get("groups").([]any)
	require.Len(t, groups, 2)

	// Groups should be sorted alphabetically by display name
	group1 := groups[0].(map[string]any)
	group2 := groups[1].(map[string]any)

	// Verify team-alpha group
	assert.Equal(t, "team-alpha", group1["display_name"])
	assert.Equal(t, "ext1", group1["external_id"])
	assert.Equal(t, "groups/team-alpha", group1["acl_principal_id"])

	// Verify members are properly categorized in team-alpha
	assertContains(t, group1["users"], "user1")
	assertContains(t, group1["service_principals"], "service-principal-1")
	assertContains(t, group1["child_groups"], "child-group-1")
	assertContains(t, group1["groups"], "parent-group-1")
	assertContains(t, group1["instance_profiles"], "instance-profile-1")

	// Verify team-beta group
	assert.Equal(t, "team-beta", group2["display_name"])
	assert.Equal(t, "ext2", group2["external_id"])
	assert.Equal(t, "groups/team-beta", group2["acl_principal_id"])

	// Verify members are properly categorized in team-beta
	assertContains(t, group2["users"], "user2")
	assertContains(t, group2["users"], "user3")
}

func TestDataSourceGroupsSingleGroup(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: `/api/2.0/preview/scim/v2/Groups?filter=displayName%20eq%20%22admins%22`,
				Response: GroupList{
					Resources: []Group{
						{
							DisplayName: "admins",
							ID:          "admin-group",
							ExternalID:  "external-admin",
							Members: []ComplexValue{
								{
									Ref:   "Users/admin1",
									Value: "admin-user-1",
								},
								{
									Ref:   "ServicePrincipals/sp-admin",
									Value: "admin-service-principal",
								},
							},
							Roles: []ComplexValue{
								{
									Value: "admin-instance-profile",
								},
								{
									Value: "backup-instance-profile",
								},
							},
						},
					},
				},
			},
		},
		Resource:    DataSourceGroups(),
		HCL:         `filter = "displayName eq \"admins\""`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err)

	groups := d.Get("groups").([]any)
	require.Len(t, groups, 1)

	group := groups[0].(map[string]any)
	assert.Equal(t, "admins", group["display_name"])
	assert.Equal(t, "external-admin", group["external_id"])
	assert.Equal(t, "groups/admins", group["acl_principal_id"])

	// Verify members are properly categorized
	assertContains(t, group["users"], "admin-user-1")
	assertContains(t, group["service_principals"], "admin-service-principal")
	assertContains(t, group["instance_profiles"], "admin-instance-profile")
	assertContains(t, group["instance_profiles"], "backup-instance-profile")
}

func TestDataSourceGroupsError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: `/api/2.0/preview/scim/v2/Groups?filter=displayName%20co%20%22test%22`,
				Status:   500,
				Response: map[string]any{
					"error": "Internal Server Error",
				},
			},
		},
		Resource:    DataSourceGroups(),
		HCL:         `filter = "displayName co \"test\""`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.Error(t, err)
}

func TestDataSourceGroupsComplexMembers(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: `/api/2.0/preview/scim/v2/Groups?filter=displayName%20co%20%22complex%22`,
				Response: GroupList{
					Resources: []Group{
						{
							DisplayName: "complex-group",
							ID:          "complex-id",
							Members: []ComplexValue{
								{
									Ref:   "Users/u1",
									Value: "zebra-user",
								},
								{
									Ref:   "Users/u2",
									Value: "alpha-user",
								},
								{
									Ref:   "ServicePrincipals/sp1",
									Value: "zebra-sp",
								},
								{
									Ref:   "ServicePrincipals/sp2",
									Value: "alpha-sp",
								},
								{
									Ref:   "Groups/g1",
									Value: "zebra-group",
								},
								{
									Ref:   "Groups/g2",
									Value: "alpha-group",
								},
							},
							Groups: []ComplexValue{
								{
									Value: "zebra-parent",
								},
								{
									Value: "alpha-parent",
								},
							},
							Roles: []ComplexValue{
								{
									Value: "zebra-role",
								},
								{
									Value: "alpha-role",
								},
							},
						},
					},
				},
			},
		},
		Resource:    DataSourceGroups(),
		HCL:         `filter = "displayName co \"complex\""`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err)

	groups := d.Get("groups").([]any)
	require.Len(t, groups, 1)

	group := groups[0].(map[string]any)
	assert.Equal(t, "complex-group", group["display_name"])
	assert.Equal(t, "groups/complex-group", group["acl_principal_id"])

	// Verify all users are properly categorized and sorted
	assertContains(t, group["users"], "alpha-user")
	assertContains(t, group["users"], "zebra-user")

	// Verify all service principals are properly categorized and sorted
	assertContains(t, group["service_principals"], "alpha-sp")
	assertContains(t, group["service_principals"], "zebra-sp")

	// Verify all child groups are properly categorized and sorted
	assertContains(t, group["child_groups"], "alpha-group")
	assertContains(t, group["child_groups"], "zebra-group")

	// Verify all parent groups are properly categorized and sorted
	assertContains(t, group["groups"], "alpha-parent")
	assertContains(t, group["groups"], "zebra-parent")

	// Verify all instance profiles are properly categorized and sorted
	assertContains(t, group["instance_profiles"], "alpha-role")
	assertContains(t, group["instance_profiles"], "zebra-role")

	// Verify that all members are included in the members set
	assertContains(t, group["members"], "alpha-user")
	assertContains(t, group["members"], "zebra-user")
	assertContains(t, group["members"], "alpha-sp")
	assertContains(t, group["members"], "zebra-sp")
	assertContains(t, group["members"], "alpha-group")
	assertContains(t, group["members"], "zebra-group")
}

// Test that groups are sorted by display name
func TestDataSourceGroupsSorting(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: `/api/2.0/preview/scim/v2/Groups?filter=displayName%20co%20%22sort%22`,
				Response: GroupList{
					Resources: []Group{
						{
							DisplayName: "sort-zebra",
							ID:          "z-id",
						},
						{
							DisplayName: "sort-alpha",
							ID:          "a-id",
						},
						{
							DisplayName: "sort-beta",
							ID:          "b-id",
						},
					},
				},
			},
		},
		Resource:    DataSourceGroups(),
		HCL:         `filter = "displayName co \"sort\""`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err)

	groups := d.Get("groups").([]any)
	require.Len(t, groups, 3)

	// Verify sorting by display name
	group1 := groups[0].(map[string]any)
	group2 := groups[1].(map[string]any)
	group3 := groups[2].(map[string]any)

	assert.Equal(t, "sort-alpha", group1["display_name"])
	assert.Equal(t, "sort-beta", group2["display_name"])
	assert.Equal(t, "sort-zebra", group3["display_name"])
}
