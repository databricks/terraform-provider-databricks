package mws_permission_assignments

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/internal/service/iam_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlattenPermissionAssignments_Nil(t *testing.T) {
	result, diags := flattenPermissionAssignments(context.Background(), nil)
	require.False(t, diags.HasError())
	assert.Empty(t, result)
}

func TestFlattenPermissionAssignments_Empty(t *testing.T) {
	result, diags := flattenPermissionAssignments(context.Background(), &iam.PermissionAssignments{})
	require.False(t, diags.HasError())
	assert.Empty(t, result)
}

func TestFlattenPermissionAssignments(t *testing.T) {
	ctx := context.Background()
	sdkResponse := &iam.PermissionAssignments{
		PermissionAssignments: []iam.PermissionAssignment{
			{
				Permissions: []iam.WorkspacePermission{iam.WorkspacePermissionAdmin},
				Principal: &iam.PrincipalOutput{
					PrincipalId: 100,
					GroupName:   "data-engineers",
					DisplayName: "Data Engineers",
				},
			},
			{
				Permissions: []iam.WorkspacePermission{iam.WorkspacePermissionUser},
				Principal: &iam.PrincipalOutput{
					PrincipalId: 200,
					UserName:    "user@example.com",
					DisplayName: "Example User",
				},
			},
			{
				// Service-principal variant, plus a non-empty error, to cover the
				// service_principal_name and error attributes end-to-end.
				Error:       "principal not found",
				Permissions: []iam.WorkspacePermission{iam.WorkspacePermissionUser, iam.WorkspacePermissionAdmin},
				Principal: &iam.PrincipalOutput{
					PrincipalId:          300,
					ServicePrincipalName: "my-service-principal",
					DisplayName:          "My Service Principal",
				},
			},
		},
	}

	result, diags := flattenPermissionAssignments(ctx, sdkResponse)
	require.False(t, diags.HasError(), "flatten should not error: %v", diags)
	require.Len(t, result, 3)

	// Convert the first flattened element back into the TF model and assert its
	// fields survived the SDK -> TF conversion.
	first := attrValueToPermissionAssignment(t, ctx, result[0])
	assert.Equal(t, int64(100), first.principalID)
	assert.Equal(t, "data-engineers", first.groupName)
	assert.Equal(t, []string{"ADMIN"}, first.permissions)
	assert.Empty(t, first.errorMessage)

	second := attrValueToPermissionAssignment(t, ctx, result[1])
	assert.Equal(t, int64(200), second.principalID)
	assert.Equal(t, "user@example.com", second.userName)
	assert.Equal(t, []string{"USER"}, second.permissions)

	third := attrValueToPermissionAssignment(t, ctx, result[2])
	assert.Equal(t, int64(300), third.principalID)
	assert.Equal(t, "my-service-principal", third.servicePrincipalName)
	assert.Equal(t, []string{"USER", "ADMIN"}, third.permissions)
	assert.Equal(t, "principal not found", third.errorMessage)
}

type flattenedAssignment struct {
	principalID          int64
	groupName            string
	userName             string
	servicePrincipalName string
	errorMessage         string
	permissions          []string
}

// attrValueToPermissionAssignment converts a flattened attr.Value back into the
// iam_tf model and extracts the fields the test asserts on.
func attrValueToPermissionAssignment(t *testing.T, ctx context.Context, v attr.Value) flattenedAssignment {
	obj, ok := v.(types.Object)
	require.True(t, ok, "expected a types.Object, got %T", v)

	var pa iam_tf.PermissionAssignment
	diags := obj.As(ctx, &pa, basetypes.ObjectAsOptions{})
	require.False(t, diags.HasError(), "obj.As should not error: %v", diags)

	var principal iam_tf.PrincipalOutput
	principalDiags := pa.Principal.As(ctx, &principal, basetypes.ObjectAsOptions{})
	require.False(t, principalDiags.HasError(), "principal As should not error: %v", principalDiags)

	var perms []string
	permDiags := pa.Permissions.ElementsAs(ctx, &perms, false)
	require.False(t, permDiags.HasError(), "permissions ElementsAs should not error: %v", permDiags)

	return flattenedAssignment{
		principalID:          principal.PrincipalId.ValueInt64(),
		groupName:            principal.GroupName.ValueString(),
		userName:             principal.UserName.ValueString(),
		servicePrincipalName: principal.ServicePrincipalName.ValueString(),
		errorMessage:         pa.Error.ValueString(),
		permissions:          perms,
	}
}
