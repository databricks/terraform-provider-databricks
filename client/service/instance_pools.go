package service

import (
	"encoding/json"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"log"
	"net/http"
)

type InstancePoolsAPI struct {
	Client DBApiClient
}

// Create creates the instance pool to given the instance pool configuration
func (a InstancePoolsAPI) Create(instancePool model.InstancePool) (model.InstancePoolInfo, error) {
	var instancePoolInfo model.InstancePoolInfo

	resp, err := a.Client.performQuery(http.MethodPost, "/instance-pools/create", "2.0", nil, instancePool)
	if err != nil {
		return instancePoolInfo, err
	}
	log.Println(resp)
	err = json.Unmarshal(resp, &instancePoolInfo)
	return instancePoolInfo, err
}

// Update edits the configuration of a instance pool to match the provided attributes and size
func (a InstancePoolsAPI) Update(instancePoolInfo model.InstancePoolInfo) error {
	_, err := a.Client.performQuery(http.MethodPost, "/instance-pools/edit", "2.0", nil, instancePoolInfo)
	return err
}

// Read retrieves the information for a instance pool given its identifier
func (a InstancePoolsAPI) Read(instancePoolId string) (model.InstancePoolInfo, error) {
	var instancePoolInfo model.InstancePoolInfo

	data := struct {
		InstancePoolId string `json:"instance_pool_id,omitempty" url:"instance_pool_id,omitempty"`
	}{
		instancePoolId,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/instance-pools/get", "2.0", nil, data)
	if err != nil {
		return instancePoolInfo, err
	}

	err = json.Unmarshal(resp, &instancePoolInfo)
	return instancePoolInfo, err
}

// Terminate terminates a instance pool given its ID
func (a InstancePoolsAPI) Delete(instancePoolId string) error {
	data := struct {
		InstancePoolId string `json:"instance_pool_id,omitempty" url:"instance_pool_id,omitempty"`
	}{
		instancePoolId,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/instance-pools/delete", "2.0", nil, data)
	return err
}
