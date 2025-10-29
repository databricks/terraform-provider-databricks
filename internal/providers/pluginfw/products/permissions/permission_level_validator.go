package permissions

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/permissions"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// permissionLevelValidator validates that the permission_level is valid for the configured object type
type permissionLevelValidator struct{}

func (v permissionLevelValidator) Description(ctx context.Context) string {
	return "validates that the permission level is valid for the configured object type"
}

func (v permissionLevelValidator) MarkdownDescription(ctx context.Context) string {
	return "validates that the permission level is valid for the configured object type"
}

func (v permissionLevelValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	permissionLevel := req.ConfigValue.ValueString()

	// Dynamically iterate through all permission definitions to find which object ID is set
	allPermissions := permissions.AllResourcePermissions()
	var mapping permissions.WorkspaceObjectPermissions
	var found bool

	for _, m := range allPermissions {
		var attrValue types.String
		diags := req.Config.GetAttribute(ctx, path.Root(m.GetField()), &attrValue)
		if diags.HasError() {
			continue // Attribute doesn't exist or has errors, try next
		}

		if !attrValue.IsNull() && !attrValue.IsUnknown() && attrValue.ValueString() != "" {
			mapping = m
			found = true
			break
		}
	}

	if !found {
		// If we can't determine the object type, let the ConflictsWith validators handle it
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
				mapping.GetObjectType(),
				allowedLevels,
			),
		)
	}
}

// ValidatePermissionLevel returns a validator that checks if the permission level is valid for the object type
func ValidatePermissionLevel() validator.String {
	return permissionLevelValidator{}
}
