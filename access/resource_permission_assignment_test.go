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
		Resource:  ResourcePermissionAssignment(),
		Create:    true,
		AccountID: "abc",
		HCL: `
		principal_id = 345
		permissions  = ["USER"]
		`,
	}.ApplyNoError(t)
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
