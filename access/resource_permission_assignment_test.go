package access

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestPermissionAssignmentCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/permissionassignments/principals/345",
				ExpectedRequest: Permissions{
					Permissions: []string{"USER"},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/permissionassignments",
				Response: PermissionAssignmentList{
					PermissionAssignments: []PermissionAssignment{
						{
							Permissions: []string{"USER"},
							Principal: Principal{
								PrincipalID: 345,
							},
						},
					},
				},
			},
		},
		Resource: ResourcePermissionAssignment(),
		Create:   true,
		HCL: `
		principal_id = 345
		permissions  = ["USER"]
		`,
	}.ApplyNoError(t)
}

func TestPermissionAssignmentRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/permissionassignments",
				Response: PermissionAssignmentList{
					PermissionAssignments: []PermissionAssignment{
						{
							Permissions: []string{"USER"},
							Principal: Principal{
								PrincipalID: 345,
							},
						},
					},
				},
			},
		},
		Resource: ResourcePermissionAssignment(),
		Read:     true,
		New:      true,
		ID:       "345",
	}.ApplyAndExpectData(t, map[string]any{
		"principal_id": 345,
		"permissions":  []any{"USER"},
	})
}

func TestPermissionAssignmentReadNotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/permissionassignments",
				Response: PermissionAssignmentList{
					PermissionAssignments: []PermissionAssignment{
						{
							Permissions: []string{"USER"},
							Principal: Principal{
								PrincipalID: 345,
							},
						},
					},
				},
			},
		},
		Resource:  ResourcePermissionAssignment(),
		Read:      true,
		Removed:   true,
		AccountID: "abc",
		ID:        "123",
	}.ApplyNoError(t)
}

func TestPermissionAssignmentFuzz(t *testing.T) {
	qa.ResourceCornerCases(t, ResourcePermissionAssignment())
}
