package identity

// URN is a custom type for the SCIM spec for the schema
type URN string

// Possible schema URNs for the Databricks SCIM api
const (
	UserSchema          URN = "urn:ietf:params:scim:schemas:core:2.0:User"
	WorkspaceUserSchema URN = "urn:ietf:params:scim:schemas:extension:workspace:2.0:User"
	PatchOp             URN = "urn:ietf:params:scim:api:messages:2.0:PatchOp"
	GroupSchema         URN = "urn:ietf:params:scim:schemas:core:2.0:Group"
)

// MembersValue is a list of value items for the members path
type MembersValue struct {
	Members []ValueListItem `json:"members,omitempty"`
}

// RolesValue is a list of value items for the roles path
type RolesValue struct {
	Roles []ValueListItem `json:"roles,omitempty"`
}

// ValueList is a generic list of value items for any path
type ValueList struct {
	Value []ValueListItem `json:"value,omitempty"`
}

// GroupsValue is a list of value items for the groups path
type GroupsValue struct {
	Groups []ValueListItem `json:"groups,omitempty"`
}

// GroupPatchOperations is a list of path operations for add or removing group attributes
type GroupPatchOperations struct {
	Op    string          `json:"op,omitempty"`
	Path  GroupPathType   `json:"path,omitempty"`
	Value []ValueListItem `json:"value,omitempty"`
}

// UserPatchOperations is a list of path operations for add or removing user attributes
type UserPatchOperations struct {
	Op    string       `json:"op,omitempty"`
	Path  string       `json:"path,omitempty"`
	Value *GroupsValue `json:"value,omitempty"`
}

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