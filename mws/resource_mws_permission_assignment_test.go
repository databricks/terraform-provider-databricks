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
		Resource:  ResourceMwsPermissionAssigntment(),
		Create:    true,
		AccountID: "abc",
		HCL: `
		workspace_id = 123
		principal_id = 345
		permissions  = ["USER"]
		`,
	}.ApplyNoError(t)
}

func TestPermssionAssignmentReadNotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
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
		Resource:  ResourceMwsPermissionAssigntment(),
		Read:      true,
		Removed:   true,
		AccountID: "abc",
		ID:        "123|456",
	}.ApplyNoError(t)
}

func TestPermssionAssignmentDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/accounts/abc/workspaces/123/permissionassignments/principals/456",
			},
		},
		Resource:  ResourceMwsPermissionAssigntment(),
		Delete:    true,
		ID:        "123|456",
		AccountID: "abc",
	}.ApplyNoError(t)
}

func TestPermssionAssignmentFuzz_NoAccountID(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMwsPermissionAssigntment(),
		qa.CornerCaseID("123|456"),
		qa.CornerCaseExpectError("must have `account_id` on provider"))
}

func TestPermssionAssignmentFuzz_InvalidID(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMwsPermissionAssigntment(),
		qa.CornerCaseExpectError("parse id: invalid ID: x"),
		qa.CornerCaseSkipCRUD("create"),
		qa.CornerCaseAccountID("abc"))
}

func TestPermssionAssignmentFuzz_ApiErrors(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMwsPermissionAssigntment(),
		qa.CornerCaseAccountID("abc"),
		qa.CornerCaseID("123|456"))
}
