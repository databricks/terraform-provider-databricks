package permissions

import "github.com/databricks/databricks-sdk-go/service/iam"

// PermissionEntity is the entity used for singular databricks_permission resource
// It represents permissions for a single principal on a single object
// Note: Currently not used, the resource uses PermissionResourceModel directly
type PermissionEntity struct {
	// Object type - computed
	ObjectType string `json:"object_type,omitempty" tf:"computed"`

	// Principal identifiers - exactly one required
	UserName             string `json:"user_name,omitempty"`
	GroupName            string `json:"group_name,omitempty"`
	ServicePrincipalName string `json:"service_principal_name,omitempty"`

	// Permission level for this principal
	PermissionLevel iam.PermissionLevel `json:"permission_level"`
}

// GetPrincipalName returns the principal identifier from the entity
func (p PermissionEntity) GetPrincipalName() string {
	if p.UserName != "" {
		return p.UserName
	}
	if p.GroupName != "" {
		return p.GroupName
	}
	if p.ServicePrincipalName != "" {
		return p.ServicePrincipalName
	}
	return ""
}

// ToAccessControlRequest converts PermissionEntity to AccessControlRequest for API calls
func (p PermissionEntity) ToAccessControlRequest() iam.AccessControlRequest {
	return iam.AccessControlRequest{
		UserName:             p.UserName,
		GroupName:            p.GroupName,
		ServicePrincipalName: p.ServicePrincipalName,
		PermissionLevel:      p.PermissionLevel,
	}
}

// FromAccessControlResponse creates a PermissionEntity from an AccessControlResponse
func FromAccessControlResponse(acr iam.AccessControlResponse) PermissionEntity {
	// Get the highest permission level from AllPermissions
	var permissionLevel iam.PermissionLevel
	if len(acr.AllPermissions) > 0 {
		permissionLevel = acr.AllPermissions[0].PermissionLevel
	}

	return PermissionEntity{
		UserName:             acr.UserName,
		GroupName:            acr.GroupName,
		ServicePrincipalName: acr.ServicePrincipalName,
		PermissionLevel:      permissionLevel,
	}
}
