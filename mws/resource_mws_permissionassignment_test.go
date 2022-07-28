package mws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestPermssionAssignmentCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/accounts/abc/workspaces/123/permissionassignments/principals/345",
				ExpectedRequest: Permissions{
					Permissions: []string{"USER"},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/accounts/abc/workspaces/123/permissionassignments",
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
		Resource: ResourceMwsPermissionassigntment(),
		Create: true,
		AccountID: "abc",
		HCL: `
		workspace_id = 123
		principal_id = 345
		permissions  = ["USER"]
		`,
	}.ApplyNoError(t)
}