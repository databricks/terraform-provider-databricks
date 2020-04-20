package model

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
