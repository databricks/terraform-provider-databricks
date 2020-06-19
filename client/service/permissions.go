package service

import (
	"encoding/json"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// PermissionsAPI ...
type PermissionsAPI struct {
	Client *DBApiClient
}

// AddOrModify ...
func (a PermissionsAPI) AddOrModify(objectID string, objectACL *model.AccessControlChangeList) error {
	_, err := a.Client.performQuery(http.MethodPatch,
		"/preview/permissions"+objectID,
		"2.0", nil, objectACL, nil)
	if err != nil {
		return err
	}

	return err
}

// SetOrDelete ...
func (a PermissionsAPI) SetOrDelete(objectID string, objectACL *model.AccessControlChangeList) error {
	_, err := a.Client.performQuery(http.MethodPut,
		"/preview/permissions"+objectID,
		"2.0", nil, objectACL, nil)
	if err != nil {
		return err
	}

	return err
}

// Read ...
func (a PermissionsAPI) Read(objectID string) (*model.ObjectACL, error) {
	resp, err := a.Client.performQuery(http.MethodGet,
		"/preview/permissions"+objectID,
		"2.0", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	var objectACL = new(model.ObjectACL)
	err = json.Unmarshal(resp, &objectACL)
	return objectACL, err
}
