package model

type GroupMember struct {
	Display string `json:"display,omitempty"`
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

type ValueListItem struct {
	Value string `json:"value,omitempty"`
}

type GroupPathType string

const (
	GroupMembersPath      GroupPathType = "members"
	GroupRolesPath        GroupPathType = "roles"
	GroupEntitlementsPath GroupPathType = "entitlements"
)

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

type GroupList struct {
	TotalResults int32   `json:"totalResults,omitempty"`
	StartIndex   int32   `json:"startIndex,omitempty"`
	ItemsPerPage int32   `json:"itemsPerPage,omitempty"`
	Schemas      []URN   `json:"schemas,omitempty"`
	Resources    []Group `json:"resources,omitempty"`
}

type GroupPatchRequest struct {
	Schemas    []URN                  `json:"schemas,omitempty"`
	Operations []GroupPatchOperations `json:"Operations,omitempty"`
}
