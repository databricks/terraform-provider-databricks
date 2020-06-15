package service

import (
	"encoding/json"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// InstancePoolsAPI exposes the instance pools api
type InstancePoolsAPI struct {
	Client *DBApiClient
}

// Create creates the instance pool to given the instance pool configuration
func (a InstancePoolsAPI) Create(instancePool model.InstancePool) (model.InstancePoolInfo, error) {
	var instancePoolInfo model.InstancePoolInfo

	resp, err := a.Client.performQuery(http.MethodPost, "/instance-pools/create", "2.0", nil, instancePool, nil)
	if err != nil {
		return instancePoolInfo, err
	}
	err = json.Unmarshal(resp, &instancePoolInfo)
	return instancePoolInfo, err
}

// Update edits the configuration of a instance pool to match the provided attributes and size
func (a InstancePoolsAPI) Update(instancePoolInfo model.InstancePoolInfo) error {
	_, err := a.Client.performQuery(http.MethodPost, "/instance-pools/edit", "2.0", nil, instancePoolInfo, nil)
	return err
}

// Read retrieves the information for a instance pool given its identifier
func (a InstancePoolsAPI) Read(instancePoolID string) (model.InstancePoolInfo, error) {
	var instancePoolInfo model.InstancePoolInfo

	data := struct {
		InstancePoolID string `json:"instance_pool_id,omitempty" url:"instance_pool_id,omitempty"`
	}{
		instancePoolID,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/instance-pools/get", "2.0", nil, data, nil)
	if err != nil {
		return instancePoolInfo, err
	}

	err = json.Unmarshal(resp, &instancePoolInfo)
	return instancePoolInfo, err
}

// Delete terminates a instance pool given its ID
func (a InstancePoolsAPI) Delete(instancePoolID string) error {
	data := struct {
		InstancePoolID string `json:"instance_pool_id,omitempty" url:"instance_pool_id,omitempty"`
	}{
		instancePoolID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/instance-pools/delete", "2.0", nil, data, nil)
	return err
}
