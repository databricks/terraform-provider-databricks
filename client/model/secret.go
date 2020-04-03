package model

type ScopeBackendType string

const (
	ScopeBackendTypeDatabricks ScopeBackendType = "DATABRICKS"
)

type SecretScope struct {
	Name        string           `json:"name,omitempty"`
	BackendType ScopeBackendType `json:"backend_type,omitempty"`
}

type SecretMetadata struct {
	Key                  string `json:"key,omitempty"`
	LastUpdatedTimestamp int64  `json:"last_updated_timestamp,omitempty"`
}

type AclPermission string

const (
	AclPermissionRead   AclPermission = "READ"
	AclPermissionWrite  AclPermission = "WRITE"
	AclPermissionManage AclPermission = "MANAGE"
)

func ValidSecretAclPermissions() []AclPermission {
	return []AclPermission{AclPermissionManage, AclPermissionRead, AclPermissionWrite}
}

type AclItem struct {
	Principal  string        `json:"principal,omitempty"`
	Permission AclPermission `json:"permission,omitempty"`
}
