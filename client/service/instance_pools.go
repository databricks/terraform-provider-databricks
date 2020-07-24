package service

import (
	"github.com/databrickslabs/databricks-terraform/client/model"
)

// InstancePoolsAPI exposes the instance pools api
type InstancePoolsAPI struct {
	client *DatabricksClient
}

// Create creates the instance pool to given the instance pool configuration
func (a InstancePoolsAPI) Create(instancePool model.InstancePool) (model.InstancePoolAndStats, error) {
	var instancePoolInfo model.InstancePoolAndStats
	err := a.client.post("/instance-pools/create", instancePool, &instancePoolInfo)
	return instancePoolInfo, err
}

// Update edits the configuration of a instance pool to match the provided attributes and size
func (a InstancePoolsAPI) Update(instancePoolInfo model.InstancePoolAndStats) error {
	return a.client.post("/instance-pools/edit", instancePoolInfo, nil)
}

// Read retrieves the information for a instance pool given its identifier
func (a InstancePoolsAPI) Read(instancePoolID string) (model.InstancePoolAndStats, error) {
	var instancePoolInfo model.InstancePoolAndStats
	err := a.client.get("/instance-pools/get", map[string]string{
		"instance_pool_id": instancePoolID,
	}, &instancePoolInfo)
	return instancePoolInfo, err
}

// List retrieves the list of existing instance pools
func (a InstancePoolsAPI) List() (ipl model.InstancePoolList, err error) {
	err = a.client.get("/instance-pools/list", nil, &ipl)
	return
}

// Delete terminates a instance pool given its ID
func (a InstancePoolsAPI) Delete(instancePoolID string) error {
	return a.client.post("/instance-pools/delete", map[string]string{
		"instance_pool_id": instancePoolID,
	}, nil)
}
