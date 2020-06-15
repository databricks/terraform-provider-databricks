package service

import (
	"encoding/json"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// SecretAclsAPI exposes the Secret ACL API
type SecretAclsAPI struct {
	Client *DBApiClient
}

// Create creates or overwrites the ACL associated with the given principal (user or group) on the specified scope point
func (a SecretAclsAPI) Create(scope string, principal string, permission model.ACLPermission) error {
	data := struct {
		Scope      string              `json:"scope,omitempty"`
		Principal  string              `json:"principal,omitempty"`
		Permission model.ACLPermission `json:"permission,omitempty"`
	}{
		scope,
		principal,
		permission,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/acls/put", "2.0", nil, data, nil)
	return err
}

// Delete deletes the given ACL on the given scope
func (a SecretAclsAPI) Delete(scope string, principal string) error {
	data := struct {
		Scope     string `json:"scope,omitempty"`
		Principal string `json:"principal,omitempty"`
	}{
		scope,
		principal,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/acls/delete", "2.0", nil, data, nil)
	return err
}

// Read describe the details about the given ACL, such as the group and permission
func (a SecretAclsAPI) Read(scope string, principal string) (model.ACLItem, error) {
	var aclItem model.ACLItem

	data := struct {
		Scope     string `json:"scope,omitempty" url:"scope,omitempty"`
		Principal string `json:"principal,omitempty" url:"principal,omitempty"`
	}{
		scope,
		principal,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/acls/get", "2.0", nil, data, nil)
	if err != nil {
		return aclItem, err
	}

	err = json.Unmarshal(resp, &aclItem)
	return aclItem, err
}

// List lists the ACLs set on the given scope
func (a SecretAclsAPI) List(scope string) ([]model.ACLItem, error) {
	var aclItem struct {
		Items []model.ACLItem `json:"items,omitempty"`
	}

	data := struct {
		Scope string `json:"scope,omitempty" url:"scope,omitempty"`
	}{
		scope,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/acls/list", "2.0", nil, data, nil)
	if err != nil {
		return aclItem.Items, err
	}

	err = json.Unmarshal(resp, &aclItem)
	return aclItem.Items, err
}
