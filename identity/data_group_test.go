package identity

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataAwsCrossAccountRolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups?filter=displayName%20eq%20ds",
				Response: GroupList{
					Resources: []ScimGroup{
						{
							DisplayName: "ds",
							ID:          "eerste",
							Entitlements: []EntitlementsListItem{
								{
									Value: "allow-cluster-create",
								},
							},
							Roles: []RoleListItem{
								{
									Value: "a",
								},
							},
							Members: []GroupMember{
								{
									Value: "1112",
								},
							},
							Groups: []GroupMember{
								{
									Value: "abc",
								},
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
				Response: ScimGroup{
					DisplayName: "product",
					ID:          "abc",
					Entitlements: []EntitlementsListItem{
						{
							Value: "allow-instance-pool-create",
						},
					},
					Roles: []RoleListItem{
						{
							Value: "b",
						},
					},
					Members: []GroupMember{
						{
							Value: "1113",
						},
					},
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceGroup(),
		ID:          ".",
		State: map[string]interface{}{
			"display_name": "ds",
		},
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "eerste", d.Id())
	assert.Contains(t, d.Get("instance_profiles"), "a")
	assert.Contains(t, d.Get("instance_profiles"), "b")
	assert.Contains(t, d.Get("members"), "1112")
	assert.Contains(t, d.Get("members"), "1113")
	assert.Contains(t, d.Get("groups"), "abc")
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("allow_cluster_create"))
}
