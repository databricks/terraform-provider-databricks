package mws

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestPermissionAssignmentCreate(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockWorkspaceAssignmentAPI().EXPECT()
			e.Update(mock.Anything, iam.UpdateWorkspaceAssignments{
				Permissions: []iam.WorkspacePermission{iam.WorkspacePermissionUser},
				PrincipalId: 345,
				WorkspaceId: 123,
			}).Return(&iam.PermissionAssignment{
				Permissions: []iam.WorkspacePermission{iam.WorkspacePermissionUser},
				Principal: &iam.PrincipalOutput{
					PrincipalId: 345,
				},
			}, nil)
			e.ListByWorkspaceId(mock.Anything, int64(123)).Return(&iam.PermissionAssignments{
				PermissionAssignments: []iam.PermissionAssignment{
					{
						Permissions: []iam.WorkspacePermission{iam.WorkspacePermissionUser},
						Principal: &iam.PrincipalOutput{
							PrincipalId: 345,
						},
					},
				},
			}, nil)
		},
		Resource:  ResourceMwsPermissionAssignment(),
		Create:    true,
		AccountID: "abc",
		HCL: `
		workspace_id = 123
		principal_id = 345
		permissions  = ["USER"]
		`,
	}.ApplyNoError(t)
}

func TestPermissionAssignmentReadNotFound(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockWorkspaceAssignmentAPI().EXPECT()
			e.ListByWorkspaceId(mock.Anything, int64(123)).Return(&iam.PermissionAssignments{
				PermissionAssignments: []iam.PermissionAssignment{
					{
						Permissions: []iam.WorkspacePermission{iam.WorkspacePermissionUser},
						Principal: &iam.PrincipalOutput{
							PrincipalId: 345,
						},
					},
				},
			}, nil)
		},
		Resource:  ResourceMwsPermissionAssignment(),
		Read:      true,
		Removed:   true,
		AccountID: "abc",
		ID:        "123|456",
	}.ApplyNoError(t)
}

func TestPermissionAssignmentDelete(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockWorkspaceAssignmentAPI().EXPECT()
			e.DeleteByWorkspaceIdAndPrincipalId(mock.Anything, int64(123), int64(456)).Return(nil)
		},
		Resource:  ResourceMwsPermissionAssignment(),
		Delete:    true,
		ID:        "123|456",
		AccountID: "abc",
	}.ApplyNoError(t)
}

func TestPermissionAssignmentFuzz_NoAccountID(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMwsPermissionAssignment(),
		qa.CornerCaseID("123|456"),
		qa.CornerCaseExpectError("must have `account_id` on provider"))
}

func TestPermissionAssignmentFuzz_InvalidID(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMwsPermissionAssignment(),
		qa.CornerCaseExpectError("parse id: invalid ID: x"),
		qa.CornerCaseSkipCRUD("create"),
		qa.CornerCaseAccountID("abc"))
}

func TestPermissionAssignmentFuzz_ApiErrors(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMwsPermissionAssignment(),
		qa.CornerCaseAccountID("abc"),
		qa.CornerCaseID("123|456"))
}
