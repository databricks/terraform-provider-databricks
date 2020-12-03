package identity

// URN is a custom type for the SCIM spec for the schema
type URN string

// Possible schema URNs for the Databricks SCIM api
const (
	UserSchema             URN = "urn:ietf:params:scim:schemas:core:2.0:User"
	ServicePrincipalSchema URN = "urn:ietf:params:scim:schemas:core:2.0:ServicePrincipal"
	WorkspaceUserSchema    URN = "urn:ietf:params:scim:schemas:extension:workspace:2.0:User"
	PatchOp                URN = "urn:ietf:params:scim:api:messages:2.0:PatchOp"
	GroupSchema            URN = "urn:ietf:params:scim:schemas:core:2.0:Group"
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

// ScimGroup contains information about the SCIM group
type ScimGroup struct {
	ID           string                 `json:"id,omitempty"`
	Schemas      []URN                  `json:"schemas,omitempty"`
	DisplayName  string                 `json:"displayName,omitempty"`
	Members      []GroupMember          `json:"members,omitempty"`
	Groups       []GroupMember          `json:"groups,omitempty"`
	Roles        []roleListItem         `json:"roles,omitempty"`
	Entitlements []entitlementsListItem `json:"entitlements,omitempty"`
}

// HasMember returns true if group has given user or another group id as member
func (g ScimGroup) HasMember(memberID string) bool {
	for _, member := range g.Members {
		if member.Value == memberID {
			return true
		}
	}
	return false
}

// HasRole returns true if group has a role
func (g ScimGroup) HasRole(role string) bool {
	for _, groupRole := range g.Roles {
		if groupRole.Value == role {
			return true
		}
	}
	return false
}

// GroupList contains a list of groups fetched from a list api call from SCIM api
type GroupList struct {
	TotalResults int32       `json:"totalResults,omitempty"`
	StartIndex   int32       `json:"startIndex,omitempty"`
	ItemsPerPage int32       `json:"itemsPerPage,omitempty"`
	Schemas      []URN       `json:"schemas,omitempty"`
	Resources    []ScimGroup `json:"resources,omitempty"`
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

type groupsListItem struct {
	// TODO: combine entitlementsListItem & roleListItem into this one
	Display string `json:"display,omitempty"`
	Value   string `json:"value,omitempty"`
}

type entitlementsListItem struct {
	Value Entitlement `json:"value,omitempty"`
}

type roleListItem struct {
	Value string `json:"value,omitempty"`
}

type email struct {
	Type    interface{} `json:"type,omitempty"`
	Value   string      `json:"value,omitempty"`
	Primary interface{} `json:"primary,omitempty"`
}

// ScimUser is a struct that contains all the information about a SCIM user
type ScimUser struct {
	ID            string                 `json:"id,omitempty"`
	Emails        []email                `json:"emails,omitempty"`
	DisplayName   string                 `json:"displayName,omitempty"`
	Active        bool                   `json:"active,omitempty"`
	Schemas       []URN                  `json:"schemas,omitempty"`
	UserName      string                 `json:"userName,omitempty"`
	ApplicationID string                 `json:"application_id,omitempty"`
	Groups        []groupsListItem       `json:"groups,omitempty"`
	Name          map[string]string      `json:"name,omitempty"`
	Roles         []roleListItem         `json:"roles,omitempty"`
	Entitlements  []entitlementsListItem `json:"entitlements,omitempty"`
}

// HasRole returns true if group has a role
func (u ScimUser) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r.Value == role {
			return true
		}
	}
	return false
}

// UserList contains a list of Users fetched from a list api call from SCIM api
type UserList struct {
	TotalResults int32      `json:"totalResults,omitempty"`
	StartIndex   int32      `json:"startIndex,omitempty"`
	ItemsPerPage int32      `json:"itemsPerPage,omitempty"`
	Schemas      []URN      `json:"schemas,omitempty"`
	Resources    []ScimUser `json:"resources,omitempty"`
}

type patchOperation struct {
	Op    string      `json:"op,omitempty"`
	Path  string      `json:"path,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type patchRequest struct {
	Schemas    []URN            `json:"schemas,omitempty"`
	Operations []patchOperation `json:"Operations,omitempty"`
}

func scimPatchRequest(op, path, value string) patchRequest {
	o := patchOperation{
		Op:   op,
		Path: path,
	}
	if value != "" {
		o.Value = []ValueListItem{{value}}
	}
	return patchRequest{
		Schemas:    []URN{PatchOp},
		Operations: []patchOperation{o},
	}
}
