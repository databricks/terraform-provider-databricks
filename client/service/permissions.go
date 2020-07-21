package service

import (
	"github.com/databrickslabs/databricks-terraform/client/model"
)

// PermissionsAPI exposes general permission related methods
type PermissionsAPI struct {
	client *DatabricksClient
}

// AddOrModify works with permissions change list
func (a PermissionsAPI) AddOrModify(objectID string, objectACL *model.AccessControlChangeList) error {
	return a.client.patch("/preview/permissions"+objectID, objectACL)
}

// SetOrDelete updates object permissions
func (a PermissionsAPI) SetOrDelete(objectID string, objectACL *model.AccessControlChangeList) error {
	return a.client.put("/preview/permissions"+objectID, objectACL)
}

// Read gets all relevant permissions for the object, including inherited ones
func (a PermissionsAPI) Read(objectID string) (objectACL *model.ObjectACL, err error) {
	err = a.client.get("/preview/permissions"+objectID, nil, &objectACL)
	return
}
