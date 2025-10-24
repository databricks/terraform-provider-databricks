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
				Response: permissionAssignmentResponseItem{
					Permissions: []string{"USER"},
					Principal: principalInfo{
						PrincipalID: 345,
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/permissionassignments",
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"USER"},
							Principal: principalInfo{
								PrincipalID: 345,
							},
						},
					},
				},
			},
		},
		Resource: ResourcePermissionAssignment(),
		Create:   true,
		New:      true,
		HCL: `
		principal_id = 345
		permissions  = ["USER"]
		`,
	}.ApplyNoError(t)
}

func TestPermissionAssignmentCreateWithUserName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/permissionassignments",
				ExpectedRequest: permissionAssignmentRequest{
					PermissionAssignments: []permissionAssignmentRequestItem{
						{
							principalInfo: principalInfo{
								UserName: "test.user@databricks.com",
							},
							Permissions: []string{"USER"},
						},
					},
				},
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"USER"},
							Principal: principalInfo{
								PrincipalID: 123,
								UserName:    "test.user@databricks.com",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/permissionassignments",
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"USER"},
							Principal: principalInfo{
								PrincipalID: 123,
								UserName:    "test.user@databricks.com",
							},
						},
					},
				},
			},
		},
		Resource: ResourcePermissionAssignment(),
		Create:   true,
		New:      true,
		HCL: `
		user_name   = "test.user@databricks.com"
		permissions = ["USER"]
		`,
	}.ApplyNoError(t)
}

func TestPermissionAssignmentCreateWithUserNameAndAdminPermissions(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/permissionassignments",
				ExpectedRequest: permissionAssignmentRequest{
					PermissionAssignments: []permissionAssignmentRequestItem{
						{
							principalInfo: principalInfo{
								UserName: "test.user@databricks.com",
							},
							Permissions: []string{"USER"},
						},
					},
				},
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"USER"},
							Principal: principalInfo{
								PrincipalID: 123,
								UserName:    "test.user@databricks.com",
							},
						},
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/permissionassignments/principals/123",
				ExpectedRequest: Permissions{
					Permissions: []string{"ADMIN"},
				},
				Response: permissionAssignmentResponseItem{
					Permissions: []string{"ADMIN"},
					Principal: principalInfo{
						PrincipalID: 123,
						UserName:    "test.user@databricks.com",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/permissionassignments",
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"ADMIN"},
							Principal: principalInfo{
								PrincipalID: 123,
								UserName:    "test.user@databricks.com",
							},
						},
					},
				},
			},
		},
		Resource: ResourcePermissionAssignment(),
		Create:   true,
		New:      true,
		HCL: `
		user_name   = "test.user@databricks.com"
		permissions = ["ADMIN"]
		`,
	}.ApplyNoError(t)
}

func TestPermissionAssignmentCreateWithUserNameAndAdminPermissionsError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/permissionassignments",
				ExpectedRequest: permissionAssignmentRequest{
					PermissionAssignments: []permissionAssignmentRequestItem{
						{
							principalInfo: principalInfo{
								UserName: "test.user@databricks.com",
							},
							Permissions: []string{"USER"},
						},
					},
				},
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"USER"},
							Principal: principalInfo{
								PrincipalID: 123,
								UserName:    "test.user@databricks.com",
							},
						},
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/permissionassignments/principals/123",
				ExpectedRequest: Permissions{
					Permissions: []string{"ADMIN"},
				},
				Status: 500,
				Response: permissionAssignmentResponseItem{
					Error: "Internal error",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/permissionassignments/principals/123",
			},
		},
		Resource: ResourcePermissionAssignment(),
		Create:   true,
		New:      true,
		HCL: `
		user_name   = "test.user@databricks.com"
		permissions = ["ADMIN"]
		`,
	}.ExpectError(t, "Internal error")
}

func TestPermissionAssignmentCreateWithServicePrincipalName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/permissionassignments",
				ExpectedRequest: permissionAssignmentRequest{
					PermissionAssignments: []permissionAssignmentRequestItem{
						{
							principalInfo: principalInfo{
								ServicePrincipalName: "spn-123",
							},
							Permissions: []string{"USER"},
						},
					},
				},
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"USER"},
							Principal: principalInfo{
								PrincipalID:          456,
								ServicePrincipalName: "spn-123",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/permissionassignments",
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"USER"},
							Principal: principalInfo{
								PrincipalID:          456,
								ServicePrincipalName: "spn-123",
							},
						},
					},
				},
			},
		},
		Resource: ResourcePermissionAssignment(),
		Create:   true,
		New:      true,
		HCL: `
		service_principal_name = "spn-123"
		permissions            = ["USER"]
		`,
	}.ApplyNoError(t)
}

func TestPermissionAssignmentCreateWithGroupName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/permissionassignments",
				ExpectedRequest: permissionAssignmentRequest{
					PermissionAssignments: []permissionAssignmentRequestItem{
						{
							principalInfo: principalInfo{
								GroupName: "admins",
							},
							Permissions: []string{"USER"},
						},
					},
				},
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"USER"},
							Principal: principalInfo{
								PrincipalID: 789,
								GroupName:   "admins",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/permissionassignments",
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"USER"},
							Principal: principalInfo{
								PrincipalID: 789,
								GroupName:   "admins",
							},
						},
					},
				},
			},
		},
		Resource: ResourcePermissionAssignment(),
		Create:   true,
		New:      true,
		HCL: `
		group_name  = "admins"
		permissions = ["USER"]
		`,
	}.ApplyNoError(t)
}

func TestPermissionAssignmentCreateWithNonExistingGroup(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/permissionassignments",
				ExpectedRequest: permissionAssignmentRequest{
					PermissionAssignments: []permissionAssignmentRequestItem{
						{
							principalInfo: principalInfo{
								GroupName: "nonexistent-group",
							},
							Permissions: []string{"USER"},
						},
					},
				},
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Principal: principalInfo{
								GroupName: "nonexistent-group",
							},
							Error: "RESOURCE_DOES_NOT_EXIST: Principal not found in account.",
						},
					},
				},
			},
		},
		Resource: ResourcePermissionAssignment(),
		Create:   true,
		New:      true,
		HCL: `
		group_name  = "nonexistent-group"
		permissions = ["USER"]
		`,
	}.ExpectError(t, "RESOURCE_DOES_NOT_EXIST: Principal not found in account.")
}

func TestPermissionAssignmentUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				// Simulate updating permissions for a principal by principal_id
				Method:   "PUT",
				Resource: "/api/2.0/preview/permissionassignments/principals/123",
				ExpectedRequest: Permissions{
					Permissions: []string{"ADMIN"},
				},
				Response: permissionAssignmentResponseItem{
					Principal: principalInfo{
						PrincipalID: 123,
					},
					Permissions: []string{"ADMIN"},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/permissionassignments",
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"ADMIN"},
							Principal: principalInfo{
								PrincipalID: 123,
							},
						},
					},
				},
			},
		},
		Resource: ResourcePermissionAssignment(),
		Update:   true,
		ID:       "123",
		InstanceState: map[string]string{
			"principal_id": "123",
			"permissions":  "[\"USER\"]",
		},
		HCL: `
		principal_id = 123
		permissions  = ["ADMIN"]
		`,
	}.ApplyNoError(t)
}

func TestPermissionAssignmentUpdateWithError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.0/preview/permissionassignments/principals/123",
				ExpectedRequest: Permissions{
					Permissions: []string{"ADMIN"},
				},
				Status: 200,
				Response: permissionAssignmentResponseItem{
					Principal: principalInfo{
						PrincipalID: 123,
					},
					Permissions: []string{"ADMIN"},
					Error:       "Invalid permission",
				},
			},
		},
		Resource: ResourcePermissionAssignment(),
		Update:   true,
		ID:       "123",
		InstanceState: map[string]string{
			"principal_id": "123",
			"permissions":  "USER",
		},
		HCL: `
		principal_id = 123
		permissions  = ["ADMIN"]
		`,
	}.ExpectError(t, "Invalid permission")
}

func TestPermissionAssignmentRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/permissionassignments",
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"USER"},
							Principal: principalInfo{
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
				Response: permissionAssignmentResponse{
					PermissionAssignments: []permissionAssignmentResponseItem{
						{
							Permissions: []string{"USER"},
							Principal: principalInfo{
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
