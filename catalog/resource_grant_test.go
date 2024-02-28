package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestGrantPermissionsCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceGrant(), qa.CornerCaseID("schema/sandbox/me"))
}

func TestResourceGrantCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.PermissionsList{
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
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY"},
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
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY"},
						},
						{
							Principal:  "someone-else",
							Privileges: []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
		},
		Resource: ResourceGrant(),
		Create:   true,
		HCL: `
		table = "foo.bar.baz"

		principal = "me"
		privileges = ["MODIFY"]
		`,
	}.ApplyNoError(t)
}

func TestResourceGrantCreateMetastoreId(t *testing.T) {
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
				Response: catalog.PermissionsList{
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
				Resource: "/api/2.1/unity-catalog/permissions/metastore/metastore_id",
				ExpectedRequest: catalog.UpdatePermissions{
					Changes: []catalog.PermissionsChange{
						{
							Principal: "me",
							Add:       []catalog.Privilege{"MODIFY"},
							Remove:    []catalog.Privilege{"SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/metastore/metastore_id?",
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY"},
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
				Resource: "/api/2.1/unity-catalog/permissions/metastore/metastore_id?",
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY"},
						},
						{
							Principal:  "someone-else",
							Privileges: []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
		},
		Resource: ResourceGrant(),
		Create:   true,
		HCL: `
		metastore = "metastore_id"
		principal = "me"
		privileges = ["MODIFY"]
		`,
	}.ApplyNoError(t)
}

func TestFailsWhenWrongMetastoreId(t *testing.T) {
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
		Resource: ResourceGrant(),
		Create:   true,
		HCL: `
		metastore = "new_id"
		principal = "me"
		privileges = ["MODIFY"]
		`,
	}.ExpectError(t, "metastore_id must be empty or equal to the metastore id assigned to the workspace: old_id. "+
		"If the metastore assigned to the workspace has changed, the new metastore id must be explicitly set")
}

func TestResourceGrantWaitUntilReady(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.PermissionsList{
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
					},
				},
			},
			// This one is still the first one, to simulate a delay on updating the permissions
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.PermissionsList{
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
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY"},
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
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY"},
						},
						{
							Principal:  "someone-else",
							Privileges: []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
		},
		Resource: ResourceGrant(),
		Create:   true,
		HCL: `
		table = "foo.bar.baz"

		principal = "me"
		privileges = ["MODIFY"]
		`,
	}.ApplyNoError(t)
}

func TestResourceGrantUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
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
							Add:       []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY", "SELECT"},
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
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY", "SELECT"},
						},
						{
							Principal:  "someone-else",
							Privileges: []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
		},
		Resource:    ResourceGrant(),
		Update:      true,
		RequiresNew: false,
		ID:          "table/foo.bar.baz/me",
		InstanceState: map[string]string{
			"table":     "foo.bar.baz",
			"principal": "me",
		},
		HCL: `
		table = "foo.bar.baz"

		principal = "me"
		privileges = ["MODIFY", "SELECT"]
		`,
	}.ApplyNoError(t)
}

func TestResourceGrantUpdateWithChangedPrincipalForcesNewResource(t *testing.T) {
	qa.ResourceFixture{
		Resource:    ResourceGrant(),
		Update:      true,
		RequiresNew: false, // test the negative as the positive is not covered by qa.ResourceFixture
		ID:          "table/foo.bar.baz/me",
		InstanceState: map[string]string{
			"table":     "foo.bar.baz",
			"principal": "someone-else",
		},
		HCL: `
		table = "foo.bar.baz"

		principal = "me"
		privileges = ["MODIFY", "SELECT"]
		`,
	}.ExpectError(t, "changes require new: principal")
}

func TestResourceGrantDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"MODIFY", "SELECT"},
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
							Remove:    []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz?",
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "someone-else",
							Privileges: []catalog.Privilege{"MODIFY", "SELECT"},
						},
					},
				},
			},
		},
		Resource: ResourceGrant(),
		Delete:   true,
		ID:       "table/foo.bar.baz/me",
		InstanceState: map[string]string{
			"table":     "foo.bar.baz",
			"principal": "me",
		},
		HCL: `
		table = "foo.bar.baz"

		principal = "me"
		privileges = ["MODIFY", "SELECT"]
		`,
	}.ApplyNoError(t)
}

func TestResourceGrantReadMalformedId(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceGrant(),
		ID:       "foo.bar",
		Read:     true,
		HCL: `
		table = "foo"
		principal = "me"
		privileges = ["MODIFY", "SELECT"]
		`,
	}.ExpectError(t, "ID must be three elements split by `/`: foo.bar")
}

func TestResourceGrantCreateNoSecurable(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceGrant(),
		Create:   true,
		HCL: `
		principal = "me"
		privileges = ["MODIFY", "SELECT"]
		`,
	}.ExpectError(t, "invalid config supplied. [catalog] Missing required argument. [external_location] Missing required argument. [foreign_connection] Missing required argument. [function] Missing required argument. [metastore] Missing required argument. [model] Missing required argument. [pipeline] Missing required argument. [recipient] Missing required argument. [schema] Missing required argument. [share] Missing required argument. [storage_credential] Missing required argument. [table] Missing required argument. [volume] Missing required argument")
}

func TestResourceGrantCreateOneSecurableOnly(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceGrant(),
		Create:   true,
		HCL: `
		catalog = "foo"
		schema = "bar"
		table = "baz"
		principal = "me"
		privileges = ["MODIFY", "SELECT"]
		`,
	}.ExpectError(t, "invalid config supplied. [catalog] Conflicting configuration arguments. [schema] Conflicting configuration arguments. [table] Conflicting configuration arguments")
}

func TestResourceGrantCreatePrincipalRequired(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceGrant(),
		Create:   true,
		HCL: `
		table = "foo.bar.baz"
		privileges = ["MODIFY", "SELECT"]
		`,
	}.ExpectError(t, "invalid config supplied. [principal] Missing required argument")
}

func TestResourceGrantCreatePrivilegesRequired(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceGrant(),
		Create:   true,
		HCL: `
		table = "foo.bar.baz"
		principal = "me"
		`,
	}.ExpectError(t, "invalid config supplied. [privileges] Missing required argument")
}

func TestResourceGrantPermissionsList_Diff_ExternallyAddedPrincipal(t *testing.T) {
	diff := diffPermissionsForPrincipal(
		"a",
		catalog.PermissionsList{ // config
			PrivilegeAssignments: []catalog.PrivilegeAssignment{
				{
					Principal:  "a",
					Privileges: []catalog.Privilege{"a"},
				},
				{
					Principal:  "c",
					Privileges: []catalog.Privilege{"a"},
				},
			},
		},
		catalog.PermissionsList{
			PrivilegeAssignments: []catalog.PrivilegeAssignment{ // platform
				{
					Principal:  "a",
					Privileges: []catalog.Privilege{"a"},
				},
				{
					Principal:  "b",
					Privileges: []catalog.Privilege{"a"},
				},
			},
		},
	)
	assert.Len(t, diff, 0)
}

func TestResourceGrantPermissionsList_Diff_ExternallyAddedPriv(t *testing.T) {
	diff := diffPermissionsForPrincipal(
		"a",
		catalog.PermissionsList{ // config
			PrivilegeAssignments: []catalog.PrivilegeAssignment{
				{
					Principal:  "a",
					Privileges: []catalog.Privilege{"a"},
				},
			},
		},
		catalog.PermissionsList{
			PrivilegeAssignments: []catalog.PrivilegeAssignment{ // platform
				{
					Principal:  "a",
					Privileges: []catalog.Privilege{"a", "b"},
				},
			},
		},
	)
	assert.Len(t, diff, 1)
	assert.Len(t, diff[0].Add, 0)
	assert.Len(t, diff[0].Remove, 1)
	assert.Equal(t, catalog.Privilege("b"), diff[0].Remove[0])
}

func TestResourceGrantPermissionsList_Diff_LocalRemoteDiff(t *testing.T) {
	diff := diffPermissionsForPrincipal(
		"a",
		catalog.PermissionsList{ // config
			PrivilegeAssignments: []catalog.PrivilegeAssignment{
				{
					Principal:  "a",
					Privileges: []catalog.Privilege{"a", "b"},
				},
			},
		},
		catalog.PermissionsList{
			PrivilegeAssignments: []catalog.PrivilegeAssignment{ // platform
				{
					Principal:  "a",
					Privileges: []catalog.Privilege{"b", "c"},
				},
			},
		},
	)
	assert.Len(t, diff, 1)
	assert.Len(t, diff[0].Add, 1)
	assert.Len(t, diff[0].Remove, 1)
	assert.Equal(t, catalog.Privilege("a"), diff[0].Add[0])
	assert.Equal(t, catalog.Privilege("c"), diff[0].Remove[0])
}

func TestResourceGrantPermissionsList_Diff_Spaces(t *testing.T) {
	diff := diffPermissionsForPrincipal(
		"a",
		catalog.PermissionsList{ // config
			PrivilegeAssignments: []catalog.PrivilegeAssignment{
				{
					Principal:  "a",
					Privileges: []catalog.Privilege{"a b"},
				},
			},
		},
		catalog.PermissionsList{
			PrivilegeAssignments: []catalog.PrivilegeAssignment{ // platform
				{
					Principal:  "a",
					Privileges: []catalog.Privilege{"a_b"},
				},
			},
		},
	)
	assert.Len(t, diff, 0)
}

func TestResourceGrantShareGrantCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions?",
				Response: catalog.PermissionsList{
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
				Response: catalog.PermissionsList{
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
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"SELECT"},
						},
					},
				},
			},
		},
		Resource: ResourceGrant(),
		Create:   true,
		HCL: `
		share = "myshare"

		principal = "me"
		privileges = ["SELECT"]
		`,
	}.ApplyNoError(t)
}

func TestResourceGrantShareGrantUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions?",
				Response: catalog.PermissionsList{
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
							Principal: "you",
							Add:       []catalog.Privilege{"SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions?",
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"SELECT"},
						},
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
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"SELECT"},
						},
						{
							Principal:  "you",
							Privileges: []catalog.Privilege{"SELECT"},
						},
					},
				},
			},
		},
		Resource: ResourceGrant(),
		Update:   true,
		ID:       "share/myshare/you",
		InstanceState: map[string]string{
			"share":     "myshare",
			"principal": "you",
		},
		HCL: `
		share = "myshare"

		principal = "you"
		privileges = ["SELECT"]
		`,
	}.ApplyNoError(t)
}

func TestResourceGrantConnectionGrantCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/connection/myconn?",
				Response: catalog.PermissionsList{
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
				Response: catalog.PermissionsList{
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
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"USE_CONNECTION"},
						},
					},
				},
			},
		},
		Resource: ResourceGrant(),
		Create:   true,
		HCL: `
		foreign_connection = "myconn"

		principal = "me"
		privileges = ["USE_CONNECTION"]
		`,
	}.ApplyNoError(t)
}

func TestResourceGrantModelGrantCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/function/mymodel?",
				Response: catalog.PermissionsList{
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
				Response: catalog.PermissionsList{
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
				Response: catalog.PermissionsList{
					PrivilegeAssignments: []catalog.PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []catalog.Privilege{"EXECUTE"},
						},
					},
				},
			},
		},
		Resource: ResourceGrant(),
		Create:   true,
		HCL: `
		model = "mymodel"

		principal = "me"
		privileges = ["EXECUTE"]
		`,
	}.ApplyNoError(t)
}
