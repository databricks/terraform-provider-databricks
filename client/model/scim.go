package model

type URN string

const (
	UserSchema          URN = "urn:ietf:params:scim:schemas:core:2.0:User"
	WorkspaceUserSchema URN = "urn:ietf:params:scim:schemas:extension:workspace:2.0:User"
	PatchOp             URN = "urn:ietf:params:scim:api:messages:2.0:PatchOp"
	GroupSchema         URN = "urn:ietf:params:scim:schemas:core:2.0:Group"
)

type MembersValue struct {
	Members []ValueListItem `json:"members,omitempty"`
}

type RolesValue struct {
	Roles []ValueListItem `json:"roles,omitempty"`
}

type ValueList struct {
	Value []ValueListItem `json:"value,omitempty"`
}

type GroupsValue struct {
	Groups []ValueListItem `json:"groups,omitempty"`
}

type GroupPatchOperations struct {
	Op    string          `json:"op,omitempty"`
	Path  GroupPathType   `json:"path,omitempty"`
	Value []ValueListItem `json:"value,omitempty"`
}

type UserPatchOperations struct {
	Op    string       `json:"op,omitempty"`
	Path  string       `json:"path,omitempty"`
	Value *GroupsValue `json:"value,omitempty"`
}
