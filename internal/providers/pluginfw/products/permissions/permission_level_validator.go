package permissions

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/permissions"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// permissionLevelValidator validates that the permission_level is valid for the configured object type.
// It uses a hybrid approach:
// 1. If the object type is known in permission_definitions.go, validate against allowed levels
// 2. If the object type is unknown (new type), skip validation and let the API handle it
type permissionLevelValidator struct{}

func (v permissionLevelValidator) Description(ctx context.Context) string {
	return "validates that the permission level is valid for the configured object type when the object type is known"
}

func (v permissionLevelValidator) MarkdownDescription(ctx context.Context) string {
	return "validates that the permission level is valid for the configured object type when the object type is known"
}

func (v permissionLevelValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	permissionLevel := req.ConfigValue.ValueString()

	// Get the object_type from the configuration
	var objectType types.String
	diags := req.Config.GetAttribute(ctx, path.Root("object_type"), &objectType)
	if diags.HasError() || objectType.IsNull() || objectType.IsUnknown() {
		// Can't validate without object_type, let the API handle it
		return
	}

	objectTypeValue := objectType.ValueString()

	// Try to find the permission mapping for this object type
	allPermissions := permissions.AllResourcePermissions()
	var mapping permissions.WorkspaceObjectPermissions
	var found bool

	for _, m := range allPermissions {
		if m.GetRequestObjectType() == objectTypeValue {
			mapping = m
			found = true
			break
		}
	}

	if !found {
		// Object type not found in our definitions - this might be a new object type
		// Let the API handle validation
		return
	}

	// Get allowed permission levels for this object type
	allowedLevels := mapping.GetAllowedPermissionLevels(true) // true = include non-management permissions

	// Check if the configured permission level is allowed
	isValid := false
	for _, allowedLevel := range allowedLevels {
		if allowedLevel == permissionLevel {
			isValid = true
			break
		}
	}

	if !isValid {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Permission Level",
			fmt.Sprintf(
				"Permission level %q is not valid for object type %q. Allowed levels: %v",
				permissionLevel,
				objectTypeValue,
				allowedLevels,
			),
		)
	}
}

// ValidatePermissionLevel returns a validator that checks if the permission level is valid for the object type.
// Uses a hybrid approach: validates against known object types, lets API handle unknown ones.
func ValidatePermissionLevel() validator.String {
	return permissionLevelValidator{}
}
