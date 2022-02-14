package scim

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertContains(t *testing.T, s interface{}, e string) bool {
	return assert.True(t, s.(*schema.Set).Contains(e), "%#v doesn't contain %s", s, e)
}

func TestDataSourceGroup(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups?filter=displayName%20eq%20%27ds%27",
				Response: GroupList{
					Resources: []Group{
						{
							DisplayName: "ds",
							ID:          "eerste",
							Entitlements: []ComplexValue{
								{
									Value: "allow-cluster-create",
								},
							},
							Roles: []ComplexValue{
								{
									Value: "a",
								},
							},
							Members: []ComplexValue{
								{
									Ref:   "Users/1112",
									Value: "1112",
								},
								{
									Ref:   "ServicePrincipals/1113",
									Value: "1113",
								},
								{
									Ref:   "Groups/1114",
									Value: "1114",
								},
							},
							Groups: []ComplexValue{
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
				Response: Group{
					DisplayName: "product",
					ID:          "abc",
					Entitlements: []ComplexValue{
						{
							Value: "allow-instance-pool-create",
						},
					},
					Roles: []ComplexValue{
						{
							Value: "b",
						},
					},
					Members: []ComplexValue{
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
	assertContains(t, d.Get("instance_profiles"), "a")
	assertContains(t, d.Get("instance_profiles"), "b")
	assertContains(t, d.Get("members"), "1112")
	assertContains(t, d.Get("members"), "1113")
	assertContains(t, d.Get("groups"), "abc")
	assert.Equal(t, true, d.Get("allow_instance_pool_create"))
	assert.Equal(t, true, d.Get("allow_cluster_create"))

	assertContains(t, d.Get("users"), "1112")
	assertContains(t, d.Get("service_principals"), "1113")
	assertContains(t, d.Get("child_groups"), "1114")
}
