package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestPermissionsCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceGrants(), qa.CornerCaseID("schema/sandbox"))
}

func TestGrantCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"SELECT"},
						},
						{
							Principal:  "someone-else",
							Privileges: []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz",
				ExpectedRequest: catalog.UpdatePermissions{
					Changes: []catalog.PermissionsChange{
						{
							Principal: "me",
							Add:       []catalog.Privilege{"MODIFY"},
							Remove:    []catalog.Privilege{"SELECT"},
						},
						{
							Principal: "someone-else",
							Remove:    []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY"},
						},
					},
				},
			},
		},
		Resource: ResourceGrants(),
		Create:   true,
		HCL: `
		table = "foo.bar.baz"

		grant {
			principal = "me"
			privileges = ["MODIFY"]
		}`,
	}.ApplyNoError(t)
}

func TestGrantCreateMetastoreId(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/current-metastore-assignment",
				Response: catalog.MetastoreAssignment{
					MetastoreId: "metastore_id",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/metastore/metastore_id?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"SELECT"},
						},
						{
							Principal:  "someone-else",
							Privileges: []catalog.Privilege{"SELECT", "USE_SHARE"},
						},
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/permissions/metastore/metastore_id",
				ExpectedRequest: catalog.UpdatePermissions{
					Changes: []catalog.PermissionsChange{
						{
							Principal: "me",
							Add:       []catalog.Privilege{"USE_SHARE"},
							Remove:    []catalog.Privilege{"SELECT"},
						},
						{
							Principal: "someone-else",
							Remove:    []catalog.Privilege{"SELECT", "USE_SHARE"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/metastore/metastore_id?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"USE_SHARE"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/metastore/metastore_id?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"USE_SHARE"},
						},
					},
				},
			},
		},
		Resource: ResourceGrants(),
		Create:   true,
		HCL: `
		metastore = "metastore_id"

		grant {
			principal = "me"
			privileges = ["USE_SHARE"]
		}`,
	}.ApplyNoError(t)
}

func TestErrorIfWrongMetastoreId(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/current-metastore-assignment",
				Response: catalog.MetastoreAssignment{
					MetastoreId: "old_id",
				},
			},
		},
		Resource: ResourceGrants(),
		Create:   true,
		HCL: `
		metastore = "new_id"

		grant {
			principal = "me"
			privileges = ["MODIFY"]
		}`,
	}.ExpectError(t, "metastore_id must be empty or equal to the metastore id assigned to the workspace: old_id. "+
		"If the metastore assigned to the workspace has changed, the new metastore id must be explicitly set")
}

func TestWaitUntilReady(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: PermissionsList{
					Assignments: []PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []string{"SELECT"},
						},
						{
							Principal:  "someone-else",
							Privileges: []string{"MODIFY", "SELECT"},
						},
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz",
				ExpectedRequest: catalog.UpdatePermissions{
					Changes: []catalog.PermissionsChange{
						{
							Principal: "me",
							Add:       []catalog.Privilege{"MODIFY"},
							Remove:    []catalog.Privilege{"SELECT"},
						},
						{
							Principal: "someone-else",
							Remove:    []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
			// This one is still the first one, to simulate a delay on updating the permissions
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"SELECT"},
						},
						{
							Principal:  "someone-else",
							Privileges: []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY"},
						},
					},
				},
			},
		},
		Resource: ResourceGrants(),
		Create:   true,
		HCL: `
		table = "foo.bar.baz"

		grant {
			principal = "me"
			privileges = ["MODIFY"]
		}`,
	}.ApplyNoError(t)
}

func TestGrantUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz",
				ExpectedRequest: catalog.UpdatePermissions{
					Changes: []catalog.PermissionsChange{
						{
							Principal: "me",
							Add:       []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: PermissionsList{
					Assignments: []PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []string{"MODIFY", "SELECT"},
						},
					},
				},
			},
		},
		Resource: ResourceGrants(),
		Update:   true,
		ID:       "table/foo.bar.baz",
		InstanceState: map[string]string{
			"table": "foo.bar.baz",
		},
		HCL: `
		table = "foo.bar.baz"

		grant {
			principal = "me"
			privileges = ["MODIFY", "SELECT"]
		}
		`,
	}.ApplyNoError(t)
}

func TestGrantReadMalformedId(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceGrants(),
		ID:       "foo.bar",
		Read:     true,
		HCL: `
		table = "foo"
		grant {
			principal = "me"
			privileges = ["MODIFY", "SELECT"]
		}
		`,
	}.ExpectError(t, "ID must be two elements split by `/`: foo.bar")
}

func TestPermissionsList_Diff_ExternallyAddedPrincipal(t *testing.T) {
	diff := diffPermissions(
		[]catalog.PrivilegeAssignment{ // config
			{
				Principal:  "a",
				Privileges: []catalog.Privilege{"a"},
			},
			{
				Principal:  "c",
				Privileges: []catalog.Privilege{"a"},
			},
		},
		[]catalog.PrivilegeAssignment{ // platform
			{
				Principal:  "a",
				Privileges: []catalog.Privilege{"a"},
			},
			{
				Principal:  "b",
				Privileges: []catalog.Privilege{"a"},
			},
		},
	)
	assert.Len(t, diff, 2)
	assert.Len(t, diff[0].Add, 0)
	assert.Len(t, diff[0].Remove, 1)
	assert.Equal(t, "b", diff[0].Principal)
	assert.Equal(t, catalog.Privilege("A"), diff[0].Remove[0])
	assert.Equal(t, "c", diff[1].Principal)
}

func TestPermissionsList_Diff_ExternallyAddedPriv(t *testing.T) {
	diff := diffPermissions(
		[]catalog.PrivilegeAssignment{ // config
			{
				Principal:  "a",
				Privileges: []catalog.Privilege{"a"},
			},
		},
		[]catalog.PrivilegeAssignment{ // platform
			{
				Principal:  "a",
				Privileges: []catalog.Privilege{"a", "b"},
			},
		},
	)
	assert.Len(t, diff, 1)
	assert.Len(t, diff[0].Add, 0)
	assert.Len(t, diff[0].Remove, 1)
	assert.Equal(t, catalog.Privilege("B"), diff[0].Remove[0])
}

func TestPermissionsList_Diff_CaseSensitivePrincipal(t *testing.T) {
	diff := diffPermissions(
		[]catalog.PrivilegeAssignment{ // config
			{
				Principal:  "A",
				Privileges: []catalog.Privilege{"a"},
			},
		},
		[]catalog.PrivilegeAssignment{ // platform
			{
				Principal:  "a",
				Privileges: []catalog.Privilege{"a"},
			},
		},
	)
	assert.Len(t, diff, 2)
}

func TestPermissionsList_Diff_LocalRemoteDiff(t *testing.T) {
	diff := diffPermissions(
		[]catalog.PrivilegeAssignment{ // config
			{
				Principal:  "a",
				Privileges: []catalog.Privilege{"a", "b"},
			},
		},
		[]catalog.PrivilegeAssignment{ // platform
			{
				Principal:  "a",
				Privileges: []catalog.Privilege{"b", "c"},
			},
		},
	)
	assert.Len(t, diff, 1)
	assert.Len(t, diff[0].Add, 1)
	assert.Len(t, diff[0].Remove, 1)
	assert.Equal(t, catalog.Privilege("A"), diff[0].Add[0])
	assert.Equal(t, catalog.Privilege("C"), diff[0].Remove[0])
}

func TestShareGrantCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions",
				ExpectedRequest: catalog.UpdatePermissions{
					Changes: []catalog.PermissionsChange{
						{
							Principal: "me",
							Add:       []catalog.Privilege{"SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"SELECT"},
						},
					},
				},
			},
		},
		Resource: ResourceGrants(),
		Create:   true,
		HCL: `
		share = "myshare"

		grant {
			principal = "me"
			privileges = ["SELECT"]
		}`,
	}.ApplyNoError(t)
}

func TestShareGrantUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"SELECT"},
						},
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions",
				ExpectedRequest: catalog.UpdatePermissions{
					Changes: []catalog.PermissionsChange{
						{
							Principal: "me",
							Remove:    []catalog.Privilege{"SELECT"},
						},
						{
							Principal: "you",
							Add:       []catalog.Privilege{"SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "you",
							Privileges: []catalog.Privilege{"SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "you",
							Privileges: []catalog.Privilege{"SELECT"},
						},
					},
				},
			},
		},
		Resource: ResourceGrants(),
		Update:   true,
		ID:       "share/myshare",
		InstanceState: map[string]string{
			"share": "myshare",
		},
		HCL: `
		share = "myshare"

		grant {
			principal = "you"
			privileges = ["SELECT"]
		}`,
	}.ApplyNoError(t)
}

func TestConnectionGrantCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/connection/myconn?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/permissions/connection/myconn",
				ExpectedRequest: catalog.UpdatePermissions{
					Changes: []catalog.PermissionsChange{
						{
							Principal: "me",
							Add:       []catalog.Privilege{"USE_CONNECTION"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/connection/myconn?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"USE_CONNECTION"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/connection/myconn?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"USE_CONNECTION"},
						},
					},
				},
			},
		},
		Resource: ResourceGrants(),
		Create:   true,
		HCL: `
		foreign_connection = "myconn"

		grant {
			principal = "me"
			privileges = ["USE_CONNECTION"]
		}`,
	}.ApplyNoError(t)
}

func TestModelGrantCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/function/mymodel?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/permissions/function/mymodel",
				ExpectedRequest: catalog.UpdatePermissions{
					Changes: []catalog.PermissionsChange{
						{
							Principal: "me",
							Add:       []catalog.Privilege{"EXECUTE"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/function/mymodel?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"EXECUTE"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/function/mymodel?",
				Response: catalog.GetPermissionsResponse{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"EXECUTE"},
						},
					},
				},
			},
		},
		Resource: ResourceGrants(),
		Create:   true,
		HCL: `
		model = "mymodel"

		grant {
			principal = "me"
			privileges = ["EXECUTE"]
		}`,
	}.ApplyNoError(t)
}
