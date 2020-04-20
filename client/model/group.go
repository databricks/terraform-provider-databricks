package model

// GroupMember contains information of a member in a scim group
type GroupMember struct {
	Display string `json:"display,omitempty"`
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

// ValueListItem is a struct that contains a field Value.
// This is for the scim api.
type ValueListItem struct {
	Value string `json:"value,omitempty"`
}

// GroupPathType describes the possible paths in the SCIM RFC for patch operations
type GroupPathType string

const (
	// GroupMembersPath is the members path for SCIM patch operation.
	GroupMembersPath GroupPathType = "members"

	// GroupRolesPath is the roles path for SCIM patch operation.
	GroupRolesPath GroupPathType = "roles"

	// GroupEntitlementsPath is the entitlements path for SCIM patch operation.
	GroupEntitlementsPath GroupPathType = "entitlements"
)

// Group contains information about the SCIM group
type Group struct {
	ID               string                 `json:"id,omitempty"`
	Schemas          []URN                  `json:"schemas,omitempty"`
	DisplayName      string                 `json:"displayName,omitempty"`
	Members          []GroupMember          `json:"members,omitempty"`
	Groups           []GroupMember          `json:"groups,omitempty"`
	Roles            []RoleListItem         `json:"roles,omitempty"`
	Entitlements     []EntitlementsListItem `json:"entitlements,omitempty"`
	UnInheritedRoles []RoleListItem         `json:"uninherited_roles,omitempty"`
	InheritedRoles   []RoleListItem         `json:"inherited_roles,omitempty"`
}

// GroupList contains a list of groups fetched from a list api call from SCIM api
type GroupList struct {
	TotalResults int32   `json:"totalResults,omitempty"`
	StartIndex   int32   `json:"startIndex,omitempty"`
	ItemsPerPage int32   `json:"itemsPerPage,omitempty"`
	Schemas      []URN   `json:"schemas,omitempty"`
	Resources    []Group `json:"resources,omitempty"`
}

// GroupPatchRequest contains a request structure to make a patch op against SCIM api
type GroupPatchRequest struct {
	Schemas    []URN                  `json:"schemas,omitempty"`
	Operations []GroupPatchOperations `json:"Operations,omitempty"`
}
