package catalog

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/catalog/permissions"
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
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz",
				Response: permissions.UnityCatalogPermissionsList{
					Assignments: []permissions.UnityCatalogPrivilegeAssignment{
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
				ExpectedRequest: permissions.UnityCatalogPermissionsDiff{
					Changes: []permissions.UnityCatalogPermissionsChange{
						{
							Principal: "me",
							Add:       []string{"MODIFY"},
							Remove:    []string{"SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz",
				Response: permissions.UnityCatalogPermissionsList{
					Assignments: []permissions.UnityCatalogPrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []string{"MODIFY"},
						},
						{
							Principal:  "someone-else",
							Privileges: []string{"MODIFY", "SELECT"},
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
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz",
				Response: permissions.UnityCatalogPermissionsList{
					Assignments: []permissions.UnityCatalogPrivilegeAssignment{
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
				ExpectedRequest: permissions.UnityCatalogPermissionsDiff{
					Changes: []permissions.UnityCatalogPermissionsChange{
						{
							Principal: "me",
							Add:       []string{"MODIFY", "SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz",
				Response: permissions.UnityCatalogPermissionsList{
					Assignments: []permissions.UnityCatalogPrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []string{"MODIFY", "SELECT"},
						},
						{
							Principal:  "someone-else",
							Privileges: []string{"MODIFY", "SELECT"},
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
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz",
				Response: permissions.UnityCatalogPermissionsList{
					Assignments: []permissions.UnityCatalogPrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []string{"MODIFY", "SELECT"},
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
				ExpectedRequest: permissions.UnityCatalogPermissionsDiff{
					Changes: []permissions.UnityCatalogPermissionsChange{
						{
							Principal: "me",
							Remove:    []string{"MODIFY", "SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/table/foo.bar.baz",
				Response: permissions.UnityCatalogPermissionsList{
					Assignments: []permissions.UnityCatalogPrivilegeAssignment{
						{
							Principal:  "someone-else",
							Privileges: []string{"MODIFY", "SELECT"},
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
	}.ExpectError(t, "invalid config supplied. [catalog] Missing required argument. [external_location] Missing required argument. [foreign_connection] Missing required argument. [function] Missing required argument. [materialized_view] Missing required argument. [metastore] Missing required argument. [model] Missing required argument. [schema] Missing required argument. [share] Missing required argument. [storage_credential] Missing required argument. [table] Missing required argument. [view] Missing required argument. [volume] Missing required argument")
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
		permissions.UnityCatalogPermissionsList{ // config
			Assignments: []permissions.UnityCatalogPrivilegeAssignment{
				{
					Principal:  "a",
					Privileges: []string{"a"},
				},
				{
					Principal:  "c",
					Privileges: []string{"a"},
				},
			},
		},
		permissions.UnityCatalogPermissionsList{
			Assignments: []permissions.UnityCatalogPrivilegeAssignment{ // platform
				{
					Principal:  "a",
					Privileges: []string{"a"},
				},
				{
					Principal:  "b",
					Privileges: []string{"a"},
				},
			},
		},
	)
	assert.Len(t, diff.Changes, 0)
}

func TestResourceGrantPermissionsList_Diff_ExternallyAddedPriv(t *testing.T) {
	diff := diffPermissionsForPrincipal(
		"a",
		permissions.UnityCatalogPermissionsList{ // config
			Assignments: []permissions.UnityCatalogPrivilegeAssignment{
				{
					Principal:  "a",
					Privileges: []string{"a"},
				},
			},
		},
		permissions.UnityCatalogPermissionsList{
			Assignments: []permissions.UnityCatalogPrivilegeAssignment{ // platform
				{
					Principal:  "a",
					Privileges: []string{"a", "b"},
				},
			},
		},
	)
	assert.Len(t, diff.Changes, 1)
	assert.Len(t, diff.Changes[0].Add, 0)
	assert.Len(t, diff.Changes[0].Remove, 1)
	assert.Equal(t, "b", diff.Changes[0].Remove[0])
}

func TestResourceGrantPermissionsList_Diff_LocalRemoteDiff(t *testing.T) {
	diff := diffPermissionsForPrincipal(
		"a",
		permissions.UnityCatalogPermissionsList{ // config
			Assignments: []permissions.UnityCatalogPrivilegeAssignment{
				{
					Principal:  "a",
					Privileges: []string{"a", "b"},
				},
			},
		},
		permissions.UnityCatalogPermissionsList{
			Assignments: []permissions.UnityCatalogPrivilegeAssignment{ // platform
				{
					Principal:  "a",
					Privileges: []string{"b", "c"},
				},
			},
		},
	)
	assert.Len(t, diff.Changes, 1)
	assert.Len(t, diff.Changes[0].Add, 1)
	assert.Len(t, diff.Changes[0].Remove, 1)
	assert.Equal(t, "a", diff.Changes[0].Add[0])
	assert.Equal(t, "c", diff.Changes[0].Remove[0])
}

func TestResourceGrantShareGrantCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions",
				Response: permissions.UnityCatalogPermissionsList{
					Assignments: []permissions.UnityCatalogPrivilegeAssignment{},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions",
				ExpectedRequest: permissions.UnityCatalogPermissionsDiff{
					Changes: []permissions.UnityCatalogPermissionsChange{
						{
							Principal: "me",
							Add:       []string{"SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions",
				Response: permissions.UnityCatalogPermissionsList{
					Assignments: []permissions.UnityCatalogPrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []string{"SELECT"},
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
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions",
				Response: permissions.UnityCatalogPermissionsList{
					Assignments: []permissions.UnityCatalogPrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []string{"SELECT"},
						},
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions",
				ExpectedRequest: permissions.UnityCatalogPermissionsDiff{
					Changes: []permissions.UnityCatalogPermissionsChange{
						{
							Principal: "you",
							Add:       []string{"SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/myshare/permissions",
				Response: permissions.UnityCatalogPermissionsList{
					Assignments: []permissions.UnityCatalogPrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []string{"SELECT"},
						},
						{
							Principal:  "you",
							Privileges: []string{"SELECT"},
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
				Resource: "/api/2.1/unity-catalog/permissions/connection/myconn",
				Response: permissions.UnityCatalogPermissionsList{
					Assignments: []permissions.UnityCatalogPrivilegeAssignment{},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/permissions/connection/myconn",
				ExpectedRequest: permissions.UnityCatalogPermissionsDiff{
					Changes: []permissions.UnityCatalogPermissionsChange{
						{
							Principal: "me",
							Add:       []string{"USE_CONNECTION"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/permissions/connection/myconn",
				Response: permissions.UnityCatalogPermissionsList{
					Assignments: []permissions.UnityCatalogPrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []string{"USE_CONNECTION"},
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
