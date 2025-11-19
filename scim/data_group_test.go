package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceGroup(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: `/api/2.0/preview/scim/v2/Groups?filter=displayName%20eq%20%22ds%22`,
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
							Roles: []ComplexValue{
								{
									Value: "a",
								},
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=displayName,members,roles,entitlements,externalId,groups",
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
		New:         true,
		Resource:    DataSourceGroup(),
		ID:          ".",
		State: map[string]any{
			"display_name": "ds",
		},
	}.ApplyAndExpectData(t, map[string]any{
		"acl_principal_id":           "groups/ds",
		"instance_profiles":          []string{"a", "b"},
		"members":                    []string{"1112", "1113", "1114"},
		"groups":                     []string{"abc"},
		"allow_instance_pool_create": true,
		"allow_cluster_create":       true,
		"users":                      []string{"1112"},
		"service_principals":         []string{"1113"},
		"child_groups":               []string{"1114"},
		"id":                         "eerste",
	})
}

func TestDataSourceGroupAccountClient(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: `/api/2.0/accounts/1234567890/scim/v2/Groups?attributes=id&filter=displayName%20eq%20%22ds%22`,
				Response: GroupList{
					Resources: []Group{
						{
							DisplayName: "ds",
							ID:          "eerste",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/1234567890/scim/v2/Groups/eerste?attributes=displayName,members,roles,entitlements,externalId,groups",
				Response: Group{
					DisplayName: "ds",
					ID:          "eerste",
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
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/1234567890/scim/v2/Groups/abc?attributes=displayName,members,roles,entitlements,externalId,groups",
				Response: Group{
					DisplayName: "product",
					ID:          "abc",
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
		AccountID:   "1234567890",
		ID:          ".",
		State: map[string]any{
			"display_name": "ds",
		},
	}.ApplyAndExpectData(t, map[string]any{
		"acl_principal_id":   "groups/ds",
		"members":            []string{"1112", "1113", "1114"},
		"groups":             []string{"abc"},
		"users":              []string{"1112"},
		"service_principals": []string{"1113"},
		"child_groups":       []string{"1114"},
		"id":                 "eerste",
	})

}
