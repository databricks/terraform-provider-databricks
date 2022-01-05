package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestPermissionsCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceGrants(), qa.CornerCaseID("schema/sandbox"))
}

type mvm map[string][]string

func (a mvm) toRaw() *schema.Set {
	grant := ResourceGrants().Schema["grant"]
	items := []interface{}{}
	for k, v := range a {
		privs := []interface{}{}
		for _, priv := range v {
			privs = append(privs, priv)
		}
		items = append(items, map[string]interface{}{
			"principal":  k,
			"privileges": schema.NewSet(schema.HashString, privs),
		})
	}
	result := schema.NewSet(schema.HashResource(grant.Elem.(*schema.Resource)), items)
	return result
}

func (a mvm) diff(t *testing.T, b mvm, changes []PermissionsChange) {
	old := a.toRaw()
	new := b.toRaw()
	diff := permissionDiffFromRaw(old, new)

	assert.Equal(t, PermissionsDiff{
		Changes: changes,
	}, diff)
}

func TestPermissionsDiff_OnlyAdd(t *testing.T) {
	mvm{
		// nothing
	}.diff(t, mvm{
		"a": {"b", "c"},
	}, []PermissionsChange{
		{
			Principal: "a",
			Add:       []string{"c", "b"},
		},
	})
}

func TestPermissionsDiff_OnlyRemove(t *testing.T) {
	mvm{
		"a": {"b", "c"},
	}.diff(t, mvm{
		// nothing
	}, []PermissionsChange{
		{
			Principal: "a",
			Remove:    []string{"c", "b"},
		},
	})
}

func TestPermissionsDiff_RemovePriv(t *testing.T) {
	mvm{
		"a": {"b", "c"},
	}.diff(t, mvm{
		"a": {"b"},
	}, []PermissionsChange{
		{
			Principal: "a",
			Remove:    []string{"c"},
		},
	})
}

func TestPermissionsDiff_AddPriv(t *testing.T) {
	mvm{
		"a": {"b", "c"},
	}.diff(t, mvm{
		"a": {"b", "c", "d"},
	}, []PermissionsChange{
		{
			Principal: "a",
			Add:       []string{"d"},
		},
	})
}

func TestPermissionsDiff_AddAndRemovePriv(t *testing.T) {
	mvm{
		"a": {"b", "c"},
	}.diff(t, mvm{
		"a": {"c", "d"},
	}, []PermissionsChange{
		{
			Principal: "a",
			Remove:    []string{"b"},
			Add:       []string{"d"},
		},
	})
}

func TestPermissionsDiff_RemovePrinc(t *testing.T) {
	mvm{
		"a": {"b", "c"},
		"z": {"x", "y"},
	}.diff(t, mvm{
		"a": {"c", "d"},
	}, []PermissionsChange{
		{
			Principal: "a",
			Remove:    []string{"b"},
			Add:       []string{"d"},
		},
		{
			Principal: "z",
			Remove:    []string{"x", "y"},
		},
	})
}

func TestPermissionsDiff_RemoveAndAddPrinc(t *testing.T) {
	mvm{
		"a": {"b", "c"},
		"z": {"x", "y"},
	}.diff(t, mvm{
		"a": {"c", "d"},
		"o": {"p", "r"},
	}, []PermissionsChange{
		{
			Principal: "o",
			Add:       []string{"r", "p"},
		},
		{
			Principal: "a",
			Remove:    []string{"b"},
			Add:       []string{"d"},
		},
		{
			Principal: "z",
			Remove:    []string{"x", "y"},
		},
	})
}

func TestPermissionsDiff_What(t *testing.T) {
	mvm{
		"a": {"b", "c"},
		"z": {"x", "y"},
	}.diff(t, mvm{
		"a": {"b"},
		"z": {"x"},
	}, []PermissionsChange{
		{
			Principal: "z",
			Remove:    []string{"y"},
		},
		{
			Principal: "a",
			Remove:    []string{"c"},
		},
	})
}

func TestGrantCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/unity-catalog/permissions/table/foo.bar.baz",
				ExpectedRequest: PermissionsDiff{
					Changes: []PermissionsChange{
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
		Create:   true,
		HCL: `
		table = "foo.bar.baz"

		grant {
			principal = "me"
			privileges = ["MODIFY", "SELECT"]
		}
		`,
	}.ApplyNoError(t)
}

func TestGrantUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/unity-catalog/permissions/table/foo.bar.baz",
				ExpectedRequest: PermissionsDiff{
					Changes: []PermissionsChange{
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

func TestPermissionsListToDelete(t *testing.T) {
	assert.Equal(t, PermissionsDiff{
		Changes: []PermissionsChange{
			{
				Principal: "me",
				Remove:    []string{"x"},
			},
		},
	}, PermissionsList{
		Assignments: []PrivilegeAssignment{
			{
				Principal:  "me",
				Privileges: []string{"x"},
			},
		},
	}.toDelete())
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
