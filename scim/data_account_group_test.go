package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceAccountGroup(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/account/scim/v2/Groups?filter=displayName%20eq%20%27ds%27",
				Response: GroupList{
					Resources: []Group{
						{
							DisplayName: "ds",
							ID:          "eerste",
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
				Resource: "/api/2.0/account/scim/v2/Groups/abc",
				Response: Group{
					DisplayName: "product",
					ID:          "abc",
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
		Resource:    DataSourceAccountGroup(),
		ID:          ".",
		State: map[string]any{
			"display_name": "ds",
		},
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "eerste", d.Id())
	assertContains(t, d.Get("roles"), "a")
	assertContains(t, d.Get("roles"), "b")
	assertContains(t, d.Get("members"), "1112")
	assertContains(t, d.Get("members"), "1113")
	assertContains(t, d.Get("groups"), "abc")

	assertContains(t, d.Get("users"), "1112")
	assertContains(t, d.Get("service_principals"), "1113")
	assertContains(t, d.Get("child_groups"), "1114")
}
