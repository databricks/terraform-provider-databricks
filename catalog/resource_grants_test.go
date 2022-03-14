package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
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
				Resource: "/api/2.0/unity-catalog/permissions/table/foo.bar.baz",
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
				Resource: "/api/2.0/unity-catalog/permissions/table/foo.bar.baz",
				ExpectedRequest: permissionsDiff{
					Changes: []permissionsChange{
						{
							Principal: "me",
							Add:       []string{"MODIFY"},
							Remove:    []string{"SELECT"},
						},
						{
							Principal: "someone-else",
							Remove:    []string{"MODIFY", "SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/permissions/table/foo.bar.baz",
				Response: PermissionsList{
					Assignments: []PrivilegeAssignment{
						{
							Principal:  "me",
							Privileges: []string{"MODIFY"},
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
				Resource: "/api/2.0/unity-catalog/permissions/table/foo.bar.baz",
				Response: PermissionsList{
					Assignments: []PrivilegeAssignment{},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.0/unity-catalog/permissions/table/foo.bar.baz",
				ExpectedRequest: permissionsDiff{
					Changes: []permissionsChange{
						{
							Principal: "me",
							Add:       []string{"MODIFY", "SELECT"},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/permissions/table/foo.bar.baz",
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

type data map[string]string

func (a data) Get(k string) interface{} {
	return a[k]
}

func TestMappingUnsupported(t *testing.T) {
	d := data{"nothing": "here"}
	err := mapping.validate(d, PermissionsList{})
	assert.EqualError(t, err, "unknown is not fully supported yet")
}

func TestInvalidPrivilege(t *testing.T) {
	d := data{"table": "me"}
	err := mapping.validate(d, PermissionsList{
		Assignments: []PrivilegeAssignment{
			{
				Principal:  "me",
				Privileges: []string{"EVERYTHING"},
			},
		},
	})
	assert.EqualError(t, err, "EVERYTHING is not allowed on table")
}

func TestPermissionsList_Diff_ExternallyAddedPrincipal(t *testing.T) {
	diff := PermissionsList{ // config
		Assignments: []PrivilegeAssignment{
			{
				Principal: "a",
				Privileges: []string{"a"},
			},
			{
				Principal: "c",
				Privileges: []string{"a"},
			},
		},
	}.diff(PermissionsList{
		Assignments: []PrivilegeAssignment{ // platform
			{
				Principal: "a",
				Privileges: []string{"a"},
			},
			{
				Principal: "b",
				Privileges: []string{"a"},
			},
		},
	})
	assert.Len(t, diff.Changes, 2)
	assert.Len(t, diff.Changes[0].Add, 0)
	assert.Len(t, diff.Changes[0].Remove, 1)
	assert.Equal(t, "b", diff.Changes[0].Principal)
	assert.Equal(t, "a", diff.Changes[0].Remove[0])
	assert.Equal(t, "c", diff.Changes[1].Principal)
}

func TestPermissionsList_Diff_ExternallyAddedPriv(t *testing.T) {
	diff := PermissionsList{ // config
		Assignments: []PrivilegeAssignment{
			{
				Principal: "a",
				Privileges: []string{"a"},
			},
		},
	}.diff(PermissionsList{
		Assignments: []PrivilegeAssignment{ // platform
			{
				Principal: "a",
				Privileges: []string{"a", "b"},
			},
		},
	})
	assert.Len(t, diff.Changes, 1)
	assert.Len(t, diff.Changes[0].Add, 0)
	assert.Len(t, diff.Changes[0].Remove, 1)
	assert.Equal(t, "b", diff.Changes[0].Remove[0])
}

func TestPermissionsList_Diff_LocalRemoteDiff(t *testing.T) {
	diff := PermissionsList{ // config
		Assignments: []PrivilegeAssignment{
			{
				Principal: "a",
				Privileges: []string{"a", "b"},
			},
		},
	}.diff(PermissionsList{
		Assignments: []PrivilegeAssignment{ // platform
			{
				Principal: "a",
				Privileges: []string{"b", "c"},
			},
		},
	})
	assert.Len(t, diff.Changes, 1)
	assert.Len(t, diff.Changes[0].Add, 1)
	assert.Len(t, diff.Changes[0].Remove, 1)
	assert.Equal(t, "a", diff.Changes[0].Add[0])
	assert.Equal(t, "c", diff.Changes[0].Remove[0])
}