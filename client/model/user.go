package model

// Entitlement is a custom type that contains a set of entitlements for a user/group
type Entitlement string

// List of possible entitlement constants on Databricks
const (
	AllowClusterCreateEntitlement      Entitlement = "allow-cluster-create"
	AllowInstancePoolCreateEntitlement Entitlement = "allow-instance-pool-create"
)

// GroupsListItem is a struct that contains a value of group id
type GroupsListItem struct {
	Value string `json:"value,omitempty"`
}

// EntitlementsListItem is a struct that contains a value of entitlement
type EntitlementsListItem struct {
	Value Entitlement `json:"value,omitempty"`
}

// RoleListItem is a struct that contains a value of role
type RoleListItem struct {
	Value string `json:"value,omitempty"`
}

// Email is a struct that contains information about a user's email
type Email struct {
	Type    interface{} `json:"type,omitempty"`
	Value   string      `json:"value,omitempty"`
	Primary interface{} `json:"primary,omitempty"`
}

// User is a struct that contains all the information about a SCIM user
type User struct {
	ID               string                 `json:"id,omitempty"`
	Emails           []Email                `json:"emails,omitempty"`
	DisplayName      string                 `json:"displayName,omitempty"`
	Active           bool                   `json:"active,omitempty"`
	Schemas          []URN                  `json:"schemas,omitempty"`
	UserName         string                 `json:"userName,omitempty"`
	Groups           []GroupsListItem       `json:"groups,omitempty"`
	Name             map[string]string      `json:"name,omitempty"`
	Roles            []RoleListItem         `json:"roles,omitempty"`
	Entitlements     []EntitlementsListItem `json:"entitlements,omitempty"`
	UnInheritedRoles []RoleListItem         `json:"uninherited_roles,omitempty"`
	InheritedRoles   []RoleListItem         `json:"inherited_roles,omitempty"`
}

// UserList contains a list of Users fetched from a list api call from SCIM api
type UserList struct {
	TotalResults int32  `json:"totalResults,omitempty"`
	StartIndex   int32  `json:"startIndex,omitempty"`
	ItemsPerPage int32  `json:"itemsPerPage,omitempty"`
	Schemas      []URN  `json:"schemas,omitempty"`
	Resources    []User `json:"resources,omitempty"`
}

// UserPatchRequest is a struct that contains all the information for a PATCH request to the SCIM users api
type UserPatchRequest struct {
	Schemas    []URN                 `json:"schemas,omitempty"`
	Operations []UserPatchOperations `json:"Operations,omitempty"`
}
