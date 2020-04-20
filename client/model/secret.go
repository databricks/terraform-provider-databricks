package model

// ScopeBackendType is a custom type for the backend type for secret scopes
type ScopeBackendType string

// List of constants of ScopeBackendType
const (
	ScopeBackendTypeDatabricks ScopeBackendType = "DATABRICKS"
)

// SecretScope is a struct that encapsulates the secret scope
type SecretScope struct {
	Name        string           `json:"name,omitempty"`
	BackendType ScopeBackendType `json:"backend_type,omitempty"`
}

// SecretMetadata is a struct that encapsulates the metadata for a secret object in a scope
type SecretMetadata struct {
	Key                  string `json:"key,omitempty"`
	LastUpdatedTimestamp int64  `json:"last_updated_timestamp,omitempty"`
}

// ACLPermission is a custom type for acl permissions
type ACLPermission string

// List of possible ACL Permissions on Databricks
const (
	ACLPermissionRead   ACLPermission = "READ"
	ACLPermissionWrite  ACLPermission = "WRITE"
	ACLPermissionManage ACLPermission = "MANAGE"
)

// ACLItem is a struct that contains information about a secret scope acl
type ACLItem struct {
	Principal  string        `json:"principal,omitempty"`
	Permission ACLPermission `json:"permission,omitempty"`
}
