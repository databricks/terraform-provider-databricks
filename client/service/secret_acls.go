package service

import (
	"github.com/databrickslabs/databricks-terraform/client/model"
)

// SecretAclsAPI exposes the Secret ACL API
type SecretAclsAPI struct {
	client *DatabricksClient
}

type secretAclRequest struct {
	Scope      string              `json:"scope,omitempty" url:"scope,omitempty"`
	Principal  string              `json:"principal,omitempty" url:"principal,omitempty"`
	Permission model.ACLPermission `json:"permission,omitempty"`
}

// Create creates or overwrites the ACL associated with the given principal (user or group) on the specified scope point
func (a SecretAclsAPI) Create(scope string, principal string, permission model.ACLPermission) error {
	return a.client.post("/secrets/acls/put", secretAclRequest{
		Scope:      scope,
		Principal:  principal,
		Permission: permission,
	}, nil)
}

// Delete deletes the given ACL on the given scope
func (a SecretAclsAPI) Delete(scope string, principal string) error {
	return a.client.post("/secrets/acls/delete", secretAclRequest{
		Scope:     scope,
		Principal: principal,
	}, nil)
}

// Read describe the details about the given ACL, such as the group and permission
func (a SecretAclsAPI) Read(scope string, principal string) (model.ACLItem, error) {
	var aclItem model.ACLItem
	err := a.client.get("/secrets/acls/get", secretAclRequest{
		Scope:     scope,
		Principal: principal,
	}, &aclItem)
	return aclItem, err
}

// List lists the ACLs set on the given scope
func (a SecretAclsAPI) List(scope string) ([]model.ACLItem, error) {
	var aclItem struct {
		Items []model.ACLItem `json:"items,omitempty"`
	}
	err := a.client.get("/secrets/acls/list", map[string]string{
		"scope": scope,
	}, &aclItem)
	return aclItem.Items, err
}
