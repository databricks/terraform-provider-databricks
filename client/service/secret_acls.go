package service

import (
	"encoding/json"
	"github.com/stikkireddy/databricks-tf-provider/client/model"
	"net/http"
)

// SecretsAPI exposes the Secrets API
type SecretAclsAPI struct {
	Client DBApiClient
}

func (a SecretAclsAPI) init(client DBApiClient) SecretAclsAPI {
	a.Client = client
	return a
}

// Create creates or overwrites the ACL associated with the given principal (user or group) on the specified scope point
func (a SecretAclsAPI) Create(scope string, principal string, permission model.AclPermission) error {
	data := struct {
		Scope      string              `json:"scope,omitempty"`
		Principal  string              `json:"principal,omitempty"`
		Permission model.AclPermission `json:"permission,omitempty"`
	}{
		scope,
		principal,
		permission,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/acls/put", "2.0", nil, data)
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
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/acls/delete", "2.0", nil, data)
	return err
}

// Read describe the details about the given ACL, such as the group and permission
func (a SecretAclsAPI) Read(scope string, principal string) (model.AclItem, error) {
	var aclItem model.AclItem

	data := struct {
		Scope     string `json:"scope,omitempty" url:"scope,omitempty"`
		Principal string `json:"principal,omitempty" url:"principal,omitempty"`
	}{
		scope,
		principal,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/acls/get", "2.0", nil, data)
	if err != nil {
		return aclItem, err
	}

	err = json.Unmarshal(resp, &aclItem)
	return aclItem, err
}

// List lists the ACLs set on the given scope
func (a SecretAclsAPI) List(scope string) ([]model.AclItem, error) {
	var aclItem struct {
		Items []model.AclItem `json:"items,omitempty"`
	}

	data := struct {
		Scope string `json:"scope,omitempty" url:"scope,omitempty"`
	}{
		scope,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/acls/list", "2.0", nil, data)
	if err != nil {
		return aclItem.Items, err
	}

	err = json.Unmarshal(resp, &aclItem)
	return aclItem.Items, err
}
