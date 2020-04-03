package model

type Entitlement string

const (
	AllowClusterCreateEntitlement      Entitlement = "allow-cluster-create"
	AllowInstancePoolCreateEntitlement Entitlement = "allow-instance-pool-create"
)

type GroupsListItem struct {
	Value string `json:"value,omitempty"`
}

type EntitlementsListItem struct {
	Value Entitlement `json:"value,omitempty"`
}

type RoleListItem struct {
	Value string `json:"value,omitempty"`
}

type Email struct {
	Type    interface{} `json:"type,omitempty"`
	Value   string      `json:"value,omitempty"`
	Primary interface{} `json:"primary,omitempty"`
}

type User struct {
	ID           string                 `json:"id,omitempty"`
	Emails       []Email                `json:"emails,omitempty"`
	DisplayName  string                 `json:"displayName,omitempty"`
	Active       bool                   `json:"active,omitempty"`
	Schemas      []URN                  `json:"schemas,omitempty"`
	UserName     string                 `json:"userName,omitempty"`
	Groups       []GroupsListItem       `json:"groups,omitempty"`
	Name         map[string]string      `json:"name,omitempty"`
	Roles        []RoleListItem         `json:"roles,omitempty"`
	Entitlements []EntitlementsListItem `json:"entitlements,omitempty"`
}
